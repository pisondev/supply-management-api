package warehouse

import "gorm.io/gorm"

type Repository interface {
	Create(warehouse *Warehouse) error
	FindAll() ([]Warehouse, error)
	FindByID(id int) (*Warehouse, error)
	Update(warehouse *Warehouse) error
	Delete(id int) error
}

type repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(w *Warehouse) error {
	return r.db.Create(w).Error
}

func (r *repository) FindAll() ([]Warehouse, error) {
	var w []Warehouse
	err := r.db.Find(&w).Error
	return w, err
}

func (r *repository) FindByID(id int) (*Warehouse, error) {
	var w Warehouse
	err := r.db.First(&w, "id = ?", id).Error
	return &w, err
}
func (r *repository) Update(w *Warehouse) error {
	return r.db.Save(w).Error
}

func (r *repository) Delete(id int) error {
	return r.db.Delete(&Warehouse{}, "id = ?", id).Error
}
