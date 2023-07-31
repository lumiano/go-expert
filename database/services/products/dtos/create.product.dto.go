package dtos

type CreateProductDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}
