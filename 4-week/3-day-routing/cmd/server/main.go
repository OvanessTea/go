package main

import (
	"net/http"

	"myapp/internal/handler"
	"myapp/internal/user"
)

func main() {
	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)
	userHandler := handler.NewUserHandler(service)

	http.HandleFunc("/users", userHandler.Handle)
	http.HandleFunc("/users/{id}", userHandler.Handle)

	http.ListenAndServe(":8080", nil)
}
