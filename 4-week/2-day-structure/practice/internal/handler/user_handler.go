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

func (h UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUsers(w, r)
		return

	case http.MethodPost:
		h.CreateUser(w, r)
		return

	default:
		w.Header().Set("Allow", http.MethodGet+", "+http.MethodPost)

		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{
			Error: "method not allowed",
		})
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

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

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetUsers()
	res := make([]UserResponse, 0, len(users))
	for _, name := range users {
		res = append(res, UserResponse{Name: name})
	}

	writeJSON(w, http.StatusOK, res)
}
