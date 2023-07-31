package products

import (
	"github.com/lumiano/go-expert/database/domain/products"
	"github.com/lumiano/go-expert/database/services/products/dtos"
	"github.com/lumiano/go-expert/database/usecases/products"
)

type ProductService interface {
	Create(createProductDto dtos.CreateProductsDto) (products.Model, error)
}

type Services struct {
	NewProductsUseCases usecases.ProductsUseCases
}

func NewProductsService(NewProductsUseCases usecases.ProductsUseCases) ProductService {
	return &Services{
		NewProductsUseCases,
	}
}
