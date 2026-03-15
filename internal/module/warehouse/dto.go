package warehouse

type CreateWarehouseRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type UpdateWarehouseRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}
