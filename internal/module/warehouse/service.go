package warehouse

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(req *CreateWarehouseRequest) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetByID(id int) (*Warehouse, error)
	Update(id int, req *UpdateWarehouseRequest) (*Warehouse, error)
	Delete(id int) error
}

type service struct {
	repo     Repository
	validate *validator.Validate
}

func NewService(repo Repository) Service {
	return &service{repo, validator.New()}
}

func (s *service) Create(req *CreateWarehouseRequest) (*Warehouse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	w := &Warehouse{Name: req.Name, Location: req.Location}
	if err := s.repo.Create(w); err != nil {
		return nil, err
	}
	return w, nil
}
func (s *service) GetAll() ([]Warehouse, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) (*Warehouse, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(id int, req *UpdateWarehouseRequest) (*Warehouse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	existing.Name = req.Name
	existing.Location = req.Location
	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *service) Delete(id int) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
