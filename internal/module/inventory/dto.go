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

type PaginationParam struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type StockFilterParam struct {
	PaginationParam
	WarehouseID int `query:"warehouse_id"`
}

type MovementFilterParam struct {
	PaginationParam
	WarehouseID  int    `query:"warehouse_id"`
	MovementType string `query:"movement_type"`
}

type PaginatedResult struct {
	Items      interface{} `json:"items"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
}
