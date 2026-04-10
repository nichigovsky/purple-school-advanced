package main

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/internal/users"
	"5-order-api-auth/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	userRepo := users.NewUserRepository(database)

	users.NewUserHandler(router, users.UserHandlerDeps{
		UserRepository: userRepo,
		Config: conf,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}