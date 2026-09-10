package main

import (
	"net/http"

	"myapp/internal/handler"
	"myapp/internal/user"
)

func main() {
	mux := http.NewServeMux()

	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)
	userHandler := handler.NewUserHandler(service)

	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)

	http.ListenAndServe(":8080", mux)
}
