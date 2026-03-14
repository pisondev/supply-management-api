package inventory

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface
type Repository interface {
	GetStock(warehouseID int, ingredientID string) (*Inventory, error)
	InsertMovementWithTx(tx *gorm.DB, movement *InventoryMovement) error
	UpsertInventoryWithTx(tx *gorm.DB, inv *Inventory) error
}

type repository struct {
	db *gorm.DB
}

// constructor
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
