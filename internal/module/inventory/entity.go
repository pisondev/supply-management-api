package inventory

import "time"

type InventoryMovement struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WarehouseID   int       `gorm:"not null"`
	IngredientID  string    `gorm:"type:uuid;not null"`
	SupplierID    *string   `gorm:"type:uuid"`
	MovementType  string    `gorm:"type:varchar(20);not null"`
	Quantity      float64   `gorm:"type:numeric(12,2);not null"`
	BalanceAfter  float64   `gorm:"type:numeric(12,2);not null"`
	ReferenceCode string    `gorm:"type:varchar(100)"`
	Notes         string    `gorm:"type:text"`
	MovementDate  time.Time `gorm:"type:timestamptz;not null;default:now()"`
	CreatedAt     time.Time `gorm:"type:timestamptz;not null;default:now()"`
}

type Inventory struct {
	WarehouseID  int       `gorm:"primaryKey;autoIncrement:false"`
	IngredientID string    `gorm:"primaryKey;type:uuid"`
	StockLevel   float64   `gorm:"type:numeric(12,2);not null;default:0"`
	UpdatedAt    time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
