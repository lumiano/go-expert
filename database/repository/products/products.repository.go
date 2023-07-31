package products

import (
	"github.com/lumiano/go-expert/database/domain/products"
	"github.com/lumiano/go-expert/database/infra/database"
)

type Products interface {
	Create(model products.Model) (products.Model, error)
}

type Repository struct {
	databaseService *database.Database
}

func NewProductsRepository(databaseService *database.Database) *Repository {
	return &Repository{
		databaseService: databaseService,
	}
}
