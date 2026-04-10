package main

import (
	"4-order-api/configs"
	"4-order-api/internal/products"
	"4-order-api/pkg/db"
	"fmt"
	"net/http"
)

func main(){
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	productRepository := products.NewProductRepository(database)

	// Handlers
	products.NewProductHandler(router, products.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}