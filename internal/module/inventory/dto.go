package inventory

type RecordMovementRequest struct {
	WarehouseID   int     `json:"warehouse_id" validate:"required"`
	IngredientID  string  `json:"ingredient_id" validate:"required,uuid"`
	SupplierID    *string `json:"supplier_id,omitempty" validate:"omitempty,uuid"`
	MovementType  string  `json:"movement_type" validate:"required,oneof=IN OUT ADJUSTMENT"`
	Quantity      float64 `json:"quantity" validate:"required,gt=0"`
	ReferenceCode string  `json:"reference_code,omitempty"`
	Notes         string  `json:"notes,omitempty"`
}
