package supplier

type CreateSupplierRequest struct {
	Name          string `json:"name" validate:"required"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Status        string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateSupplierRequest struct {
	Name          string `json:"name" validate:"required"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Status        string `json:"status" validate:"required,oneof=active inactive"`
}
