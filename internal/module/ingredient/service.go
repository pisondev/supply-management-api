package ingredient

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(req *CreateIngredientRequest) (*Ingredient, error)
	GetAll() ([]Ingredient, error)
	GetByID(id string) (*Ingredient, error)
	Update(id string, req *UpdateIngredientRequest) (*Ingredient, error)
	Delete(id string) error
}

type service struct {
	repo     Repository
	validate *validator.Validate
}

func NewService(repo Repository) Service {
	return &service{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *service) Create(req *CreateIngredientRequest) (*Ingredient, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	newIngredient := &Ingredient{
		SKU:          req.SKU,
		Name:         req.Name,
		UnitID:       req.UnitID,
		IsPerishable: req.IsPerishable,
	}

	if err := s.repo.Create(newIngredient); err != nil {
		return nil, err
	}
	return newIngredient, nil
}

func (s *service) GetAll() ([]Ingredient, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id string) (*Ingredient, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(id string, req *UpdateIngredientRequest) (*Ingredient, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Name = req.Name
	existing.UnitID = req.UnitID
	existing.IsPerishable = req.IsPerishable

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *service) Delete(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
