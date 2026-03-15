package supplier

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(req *CreateSupplierRequest) (*Supplier, error)
	GetAll() ([]Supplier, error)
	GetByID(id string) (*Supplier, error)
	Update(id string, req *UpdateSupplierRequest) (*Supplier, error)
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

func (s *service) Create(req *CreateSupplierRequest) (*Supplier, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	status := "active"
	if req.Status != "" {
		status = req.Status
	}

	newSupplier := &Supplier{
		Name:          req.Name,
		ContactPerson: req.ContactPerson,
		Phone:         req.Phone,
		Address:       req.Address,
		Status:        status,
	}

	if err := s.repo.Create(newSupplier); err != nil {
		return nil, err
	}
	return newSupplier, nil
}

func (s *service) GetAll() ([]Supplier, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id string) (*Supplier, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(id string, req *UpdateSupplierRequest) (*Supplier, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Name = req.Name
	existing.ContactPerson = req.ContactPerson
	existing.Phone = req.Phone
	existing.Address = req.Address
	existing.Status = req.Status

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
