package handler

import (
	"encoding/json"
	"errors"
	"myapp/internal/user"
	"net/http"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{
			Error: "method not allowed",
		})
		return
	}
	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "invalid JSON",
		})
		return
	}

	err = h.service.CreateUser(req.Name)

	if errors.Is(err, user.ErrInvalidName) {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "invalid name",
		})
		return
	}

	if errors.Is(err, user.ErrUserExists) {
		writeJSON(w, http.StatusConflict, ErrorResponse{
			Error: "user already exists",
		})
		return
	}

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{
			Error: "internal error",
		})
		return
	}

	writeJSON(w, http.StatusCreated, CreateUserResponse{
		Message: "user created",
		Name:    req.Name,
	})
}
