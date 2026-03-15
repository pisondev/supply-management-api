package ingredient

type CreateIngredientRequest struct {
	SKU          string `json:"sku" validate:"required,min=3"`
	Name         string `json:"name" validate:"required"`
	UnitID       int    `json:"unit_id" validate:"required,gt=0"`
	IsPerishable bool   `json:"is_perishable"`
}

type UpdateIngredientRequest struct {
	Name         string `json:"name" validate:"required"`
	UnitID       int    `json:"unit_id" validate:"required,gt=0"`
	IsPerishable bool   `json:"is_perishable"`
}
