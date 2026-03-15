package supplier

import "gorm.io/gorm"

type Repository interface {
	Create(supplier *Supplier) error
	FindAll() ([]Supplier, error)
	FindByID(id string) (*Supplier, error)
	Update(supplier *Supplier) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(supplier *Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *repository) FindAll() ([]Supplier, error) {
	var suppliers []Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *repository) FindByID(id string) (*Supplier, error) {
	var supplier Supplier
	err := r.db.First(&supplier, "id = ?", id).Error
	return &supplier, err
}

func (r *repository) Update(supplier *Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Supplier{}, "id = ?", id).Error
}
