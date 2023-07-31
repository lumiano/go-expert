package products

import (
	"database/sql"
	"github.com/lumiano/go-expert/database/domain/products"
	"log"
)

func (r Repository) Create(model products.Model) (products.Model, error) {

	log.Printf("Create product from %v", model)

	query := "INSERT INTO products (id, name, price) VALUES ($1, $2, $3)"

	stmt, err := r.databaseService.Prepare(query)

	if err != nil {
		return products.Model{}, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(model.ID, model.Name, model.Price)

	if err != nil {
		return products.Model{}, err
	}

	return model, nil
}
