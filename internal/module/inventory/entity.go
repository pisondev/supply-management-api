package inventory

import "time"

type InventoryMovement struct {
	ID            string
	WarehouseID   int
	IngredientID  string
	SupplierID    *string
	MovementType  string
	Quantity      float64
	BalanceAfter  float64
	ReferenceCode string
	Notes         string
	MovementDate  time.Time
	CreatedAt     time.Time
}

type Inventory struct {
	WarehouseID  int
	IngredientID string
	StockLevel   float64
	UpdatedAt    time.Time
}
