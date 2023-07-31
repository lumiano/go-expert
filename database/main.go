package main

import (
	"fmt"
	"github.com/lumiano/go-expert/database/controller/products"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/products", products.Controller)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		return
	}

	fmt.Println("Server running on port 8080")

}
