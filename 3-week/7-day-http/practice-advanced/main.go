package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrUserExists = errors.New("user already exists")

type UserRepository interface {
	Save(name string) error
}

type PostgresUserRepository struct{}

type UserService struct {
	repo UserRepository
}

type UserHandler struct {
	service UserService
}

func main() {
	repo := PostgresUserRepository{}
	service := NewUserService(repo)

	handler := NewUserHandler(service)

	http.HandleFunc("/users", handler.CreateUser)

	http.ListenAndServe(":8080", nil)
}

func (r PostgresUserRepository) Save(name string) error {
	dbErr := errors.New("duplicate key violates unique constraint")

	if dbErr != nil {
		return ErrUserExists
	}

	return nil
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(name string) error {
	if err := s.repo.Save(name); err != nil {
		return fmt.Errorf("create user %q: %w", name, err)
	}

	return nil
}

func NewUserHandler(service UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	err := h.service.CreateUser(name)

	if errors.Is(err, ErrUserExists) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user created")
}
