package handler

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Message string `json:message`
	Name    string `json:name`
}

type ErrorResponse struct {
	Error string `json:error`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
