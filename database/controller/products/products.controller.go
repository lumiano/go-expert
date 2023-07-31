package products

import (
	"encoding/json"
	"github.com/lumiano/go-expert/database/infra/database"
	repository "github.com/lumiano/go-expert/database/repository/products"
	services "github.com/lumiano/go-expert/database/services/products"
	"github.com/lumiano/go-expert/database/services/products/dtos"
	usecases "github.com/lumiano/go-expert/database/usecases/products"

	"io"
	"log"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {

	log.Printf("Request received: %v from Product Controller", r)

	body := r.Body

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(body)

	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		res := make(map[string]string)

		res["error"] = "Method not allowed"
		res["status"] = "405"

		err := json.NewEncoder(w).Encode(res)

		if err != nil {
			log.Println("Error on encode json:", err)
		}

		w.WriteHeader(http.StatusMethodNotAllowed)

	}

	if r.Method == "POST" {

		newProduct := dtos.CreateProductsDto{}

		decoder := json.NewDecoder(body)

		if err := decoder.Decode(&newProduct); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		db := database.NewDatabaseService()

		productRepository := repository.NewProductsRepository(db)

		productsUseCases := usecases.NewProductsUseCases(productRepository)

		productsService := services.NewProductsService(productsUseCases)

		productCreated, err := productsService.Create(newProduct)

		if err != nil {
			log.Println("Error on create new product:", err)
			w.WriteHeader(http.StatusBadRequest)

			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusOK)

		res := make(map[string]interface{})

		res["data"] = productCreated
		res["status"] = "200"

		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			log.Println("Error on encode json:", err)
		}

	}

}
