package inventory

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetStock(warehouseID int, ingredientID string) (*Inventory, error)
	InsertMovementWithTx(tx *gorm.DB, movement *InventoryMovement) error
	UpsertInventoryWithTx(tx *gorm.DB, inv *Inventory) error

	FindStocks(filter *StockFilterParam) ([]Inventory, int64, error)
	FindMovements(filter *MovementFilterParam) ([]InventoryMovement, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetStock(warehouseID int, ingredientID string) (*Inventory, error) {
	var inv Inventory
	err := r.db.Where("warehouse_id = ? AND ingredient_id = ?", warehouseID, ingredientID).First(&inv).Error
	return &inv, err
}

func (r *repository) InsertMovementWithTx(tx *gorm.DB, movement *InventoryMovement) error {
	return tx.Create(movement).Error
}

func (r *repository) UpsertInventoryWithTx(tx *gorm.DB, inv *Inventory) error {
	return tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "warehouse_id"}, {Name: "ingredient_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"stock_level", "updated_at"}),
	}).Create(inv).Error
}

func (r *repository) FindStocks(filter *StockFilterParam) ([]Inventory, int64, error) {
	var stocks []Inventory
	var total int64

	query := r.db.Model(&Inventory{})

	if filter.WarehouseID != 0 {
		query = query.Where("warehouse_id = ?", filter.WarehouseID)
	}

	query.Count(&total)

	offset := (filter.Page - 1) * filter.Limit
	err := query.Offset(offset).Limit(filter.Limit).Order("updated_at desc").Find(&stocks).Error

	return stocks, total, err
}

func (r *repository) FindMovements(filter *MovementFilterParam) ([]InventoryMovement, int64, error) {
	var movements []InventoryMovement
	var total int64

	query := r.db.Model(&InventoryMovement{})

	if filter.WarehouseID != 0 {
		query = query.Where("warehouse_id = ?", filter.WarehouseID)
	}
	if filter.MovementType != "" {
		query = query.Where("movement_type = ?", filter.MovementType)
	}

	query.Count(&total)

	offset := (filter.Page - 1) * filter.Limit
	err := query.Offset(offset).Limit(filter.Limit).Order("movement_date desc").Find(&movements).Error

	return movements, total, err
}
