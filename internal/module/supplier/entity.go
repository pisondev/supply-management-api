package supplier

import "time"

type Supplier struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	ContactPerson string    `gorm:"type:varchar(255)" json:"contact_person"`
	Phone         string    `gorm:"type:varchar(50)" json:"phone"`
	Address       string    `gorm:"type:text" json:"address"`
	Status        string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt     time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamptz;not null;default:now()" json:"updated_at"`
}
