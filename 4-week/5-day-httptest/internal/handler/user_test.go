package handler

import (
	"encoding/json"
	"myapp/internal/dto"
	"myapp/internal/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)

	repo.Save("Alex")

	h := NewUserHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", h.GetUser)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users/0",
		nil,
	)

	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var response dto.UserResponse

	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.ID != 0 {
		t.Errorf("expected ID 0, got %d", response.ID)
	}

	if response.Name != "Alex" {
		t.Errorf("expected name Alex, got %q", response.Name)
	}
}
