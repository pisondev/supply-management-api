package ingredient

import "time"

type Ingredient struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	SKU          string    `gorm:"type:varchar(50);not null;unique" json:"sku"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	UnitID       int       `gorm:"not null" json:"unit_id"`
	IsPerishable bool      `gorm:"default:false" json:"is_perishable"`
	CreatedAt    time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamptz;not null;default:now()" json:"updated_at"`
}
