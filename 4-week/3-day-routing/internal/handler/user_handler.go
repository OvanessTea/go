package handler

import (
	"encoding/json"
	"errors"
	"myapp/internal/dto"
	"myapp/internal/user"
	"net/http"
	"strconv"
	"strings"
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
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) == 1 {
			h.GetUsers(w, r)
		}
		if len(parts) == 2 {
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{
					Error: "invalid ID",
				})
				return
			}
			h.GetUser(id, w, r)
		}
		return

	case http.MethodPost:
		h.CreateUser(w, r)
		return

	default:
		w.Header().Set("Allow", http.MethodGet+", "+http.MethodPost)

		writeJSON(w, http.StatusMethodNotAllowed, dto.ErrorResponse{
			Error: "method not allowed",
		})
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid JSON",
		})
		return
	}

	err = h.service.CreateUser(req.Name)

	if errors.Is(err, user.ErrInvalidName) {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid name",
		})
		return
	}

	if errors.Is(err, user.ErrUserExists) {
		writeJSON(w, http.StatusConflict, dto.ErrorResponse{
			Error: "user already exists",
		})
		return
	}

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, dto.ErrorResponse{
			Error: "internal error",
		})
		return
	}

	writeJSON(w, http.StatusCreated, dto.CreateUserResponse{
		Message: "user created",
		Name:    req.Name,
	})
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetUsers()
	res := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		res = append(res, dto.UserResponse{ID: user.ID, Name: user.Name})
	}

	writeJSON(w, http.StatusOK, res)
}

func (h UserHandler) GetUser(id int, w http.ResponseWriter, r *http.Request) {
	user, err := h.service.GetUser(id)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, dto.ErrorResponse{
			Error: "user not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, dto.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	})
}
