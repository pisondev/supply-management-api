package inventory

import "time"

type InventoryMovement struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	WarehouseID   int       `gorm:"not null" json:"warehouse_id"`
	IngredientID  string    `gorm:"type:uuid;not null" json:"ingredient_id"`
	SupplierID    *string   `gorm:"type:uuid" json:"supplier_id"`
	MovementType  string    `gorm:"type:varchar(20);not null" json:"movement_type"`
	Quantity      float64   `gorm:"type:numeric(12,2);not null" json:"quantity"`
	BalanceAfter  float64   `gorm:"type:numeric(12,2);not null" json:"balance_after"`
	ReferenceCode string    `gorm:"type:varchar(100)" json:"reference_code"`
	Notes         string    `gorm:"type:text" json:"notes"`
	MovementDate  time.Time `gorm:"type:timestamptz;not null;default:now()" json:"movement_date"`
	CreatedAt     time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
}

type Inventory struct {
	WarehouseID  int       `gorm:"primaryKey;autoIncrement:false" json:"warehouse_id"`
	IngredientID string    `gorm:"primaryKey;type:uuid" json:"ingredient_id"`
	StockLevel   float64   `gorm:"type:numeric(12,2);not null;default:0" json:"stock_level"`
	UpdatedAt    time.Time `gorm:"type:timestamptz;not null;default:now()" json:"updated_at"`
}
