package usecases

import (
	"github.com/lumiano/go-expert/database/domain/products"
	"github.com/lumiano/go-expert/database/shared/helpers/uuid"
	"log"
)

func (r Repository) Create(model products.Model) (products.Model, error) {

	log.Printf("Create product from usecases")

	uuid := uuid.UUID{}

	newProductId := uuid.Create()

	model.ID = newProductId

	model, err := r.productsRepository.Create(model)

	if err != nil {
		return products.Model{}, err
	}

	return model, nil
}
