package usecases

import (
	"github.com/lumiano/go-expert/database/domain/products"
	repository "github.com/lumiano/go-expert/database/repository/products"
)

type ProductsUseCases interface {
	Create(model products.Model) (products.Model, error)
}

type Repository struct {
	productsRepository *repository.Repository
}

func NewProductsUseCases(productRepository *repository.Repository) *Repository {
	return &Repository{
		productRepository,
	}
}
