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
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantUser   *dto.UserResponse
		wantError  string
	}{
		{
			name:       "user exists",
			path:       "/users/0",
			wantStatus: http.StatusOK,
			wantUser: &dto.UserResponse{
				ID:   0,
				Name: "Alex",
			},
		},
		{
			name:       "user not found",
			path:       "/users/999",
			wantStatus: http.StatusNotFound,
			wantError:  "user not found",
		},
		{
			name:       "invalid id",
			path:       "/users/abc",
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid ID",
		},
	}

	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)

	repo.Save("Alex")

	h := NewUserHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", h.GetUser)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodGet,
				tt.path,
				nil,
			)

			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if got := rec.Header().Get("Content-Type"); got != "application/json" {
				t.Errorf("expected Content-Type %q, got %q", "application/json", got)
			}

			if rec.Code != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			if tt.wantUser != nil {
				var got dto.UserResponse

				err := json.NewDecoder(rec.Body).Decode(&got)
				if err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if got.ID != tt.wantUser.ID {
					t.Errorf("expected ID %d, got %d", tt.wantUser.ID, got.ID)
				}

				if got.Name != tt.wantUser.Name {
					t.Errorf("expected name %q, got %q", tt.wantUser.Name, got.Name)
				}
			}

			if tt.wantError != "" {
				var response dto.ErrorResponse

				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}

				if response.Error != tt.wantError {
					t.Errorf("expected error %q, got %q", tt.wantError, response.Error)
				}
			}
		})
	}
}
