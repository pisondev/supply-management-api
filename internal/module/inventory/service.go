package inventory

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// interface
type Service interface {
	RecordMovement(req *RecordMovementRequest) (*InventoryMovement, error)
}

type service struct {
	repo     Repository
	db       *gorm.DB
	validate *validator.Validate
}

// constructor
func NewService(repo Repository, db *gorm.DB) Service {
	return &service{
		repo:     repo,
		db:       db,
		validate: validator.New(),
	}
}

func (s *service) RecordMovement(req *RecordMovementRequest) (*InventoryMovement, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	if req.MovementType == "OUT" && req.SupplierID != nil {
		return nil, errors.New("supplier_id must be null for OUT movements")
	}
	if req.MovementType == "IN" && req.SupplierID == nil {
		return nil, errors.New("supplier_id is required for IN movements")
	}

	var newMovement *InventoryMovement

	err := s.db.Transaction(func(tx *gorm.DB) error {
		currentInv, err := s.repo.GetStock(req.WarehouseID, req.IngredientID)

		var currentStock float64 = 0
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if req.MovementType != "IN" {
					return errors.New("item not found in warehouse inventory")
				}
				currentStock = 0
			} else {
				return fmt.Errorf("failed to fetch current stock: %w", err)
			}
		} else {
			currentStock = currentInv.StockLevel
		}

		var balanceAfter float64

		switch req.MovementType {
		case "IN":
			balanceAfter = currentStock + req.Quantity
		case "OUT":
			if currentStock < req.Quantity {
				return fmt.Errorf("insufficient stock. current: %.2f, requested: %.2f", currentStock, req.Quantity)
			}
			balanceAfter = currentStock - req.Quantity
		case "ADJUSTMENT":
			if currentStock < req.Quantity {
				return fmt.Errorf("insufficient stock for adjustment. current: %.2f, deduct: %.2f", currentStock, req.Quantity)
			}
			balanceAfter = currentStock - req.Quantity
		default:
			return errors.New("invalid movement_type")
		}

		newMovement = &InventoryMovement{
			WarehouseID:   req.WarehouseID,
			IngredientID:  req.IngredientID,
			SupplierID:    req.SupplierID,
			MovementType:  req.MovementType,
			Quantity:      req.Quantity,
			BalanceAfter:  balanceAfter,
			ReferenceCode: req.ReferenceCode,
			Notes:         req.Notes,
			MovementDate:  time.Now(),
		}

		if err := s.repo.InsertMovementWithTx(tx, newMovement); err != nil {
			return fmt.Errorf("failed to insert movement log: %w", err)
		}

		newInvState := &Inventory{
			WarehouseID:  req.WarehouseID,
			IngredientID: req.IngredientID,
			StockLevel:   balanceAfter,
			UpdatedAt:    time.Now(),
		}

		if err := s.repo.UpsertInventoryWithTx(tx, newInvState); err != nil {
			return fmt.Errorf("failed to upsert inventory state: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newMovement, nil
}
