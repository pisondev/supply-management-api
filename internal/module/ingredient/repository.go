package ingredient

import "gorm.io/gorm"

type Repository interface {
	Create(ingredient *Ingredient) error
	FindAll() ([]Ingredient, error)
	FindByID(id string) (*Ingredient, error)
	Update(ingredient *Ingredient) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(ingredient *Ingredient) error {
	return r.db.Create(ingredient).Error
}

func (r *repository) FindAll() ([]Ingredient, error) {
	var ingredients []Ingredient
	err := r.db.Find(&ingredients).Error
	return ingredients, err
}

func (r *repository) FindByID(id string) (*Ingredient, error) {
	var ingredient Ingredient
	err := r.db.First(&ingredient, "id = ?", id).Error
	return &ingredient, err
}

func (r *repository) Update(ingredient *Ingredient) error {
	return r.db.Save(ingredient).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Ingredient{}, "id = ?", id).Error
}
