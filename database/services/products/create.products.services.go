package products

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/lumiano/go-expert/database/domain/products"
	"github.com/lumiano/go-expert/database/services/products/dtos"
)

func (s Services) Create(createProductDto dtos.CreateProductDto) (products.Model, error) {

	validator := validator.New()

	if err := validator.Struct(createProductDto); err != nil {
		return products.Model{}, errors.New("name and price are required")
	}

	newProductModel := products.Model{
		Name:  createProductDto.Name,
		Price: createProductDto.Price,
	}

	model, err := s.NewProductsUseCases.Create(newProductModel)

	if err != nil {
		return products.Model{}, err
	}

	return model, nil
}
