package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrUserExists = errors.New("user already exists")

type UserRepository interface {
	Save(name string) error
}

type InMemoryUserRepository struct {
	users map[string]bool
}

type UserService struct {
	repo UserRepository
}

type UserHandler struct {
	service UserService
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Message string `json:message`
	Name    string `json:name`
}

type CreateUserError struct {
	Error string `json:error`
}

func main() {
	repo := &InMemoryUserRepository{
		users: make(map[string]bool),
	}
	service := NewUserService(repo)

	handler := NewUserHandler(service)

	http.HandleFunc("/users", handler.CreateUser)

	http.ListenAndServe(":8080", nil)
}

func (r *InMemoryUserRepository) Save(name string) error {
	if r.users[name] {
		return ErrUserExists
	}

	r.users[name] = true
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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateUserError{
			Error: "Method not allowed",
		})
		return
	}

	var req CreateUserRequest
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateUserError{
			Error: "invalid JSON",
		})
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateUserError{
			Error: "invalid name",
		})
		return
	}

	err = h.service.CreateUser(req.Name)

	if errors.Is(err, ErrUserExists) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(CreateUserError{
			Error: "user already exists",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CreateUserError{
			Error: "internal error",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{
		Message: "user created",
		Name:    req.Name,
	})
}
