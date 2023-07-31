package products

type Model struct {
	ID    string  `json:"id" validate:"uuid"`
	Name  string  `json:"name"`
	Price float64 `json:"price" validate:"min=0"`
}
