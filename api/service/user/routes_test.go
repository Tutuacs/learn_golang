package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/Tutuacs/learn_golang/api.git/types"
)

func TestUser(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// ! Fail if user payload is invalid
	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserDto{
			FirstName: "user",
			LastName:  "123",
			Email:     "invalid",
			Password:  "123456789",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRecorder()

		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(r, req)

		if r.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, r.Code)
		}
	})

	// * Complete register a new User
	t.Run("Should correctly register a new user", func(t *testing.T) {
		payload := types.RegisterUserDto{
			FirstName: "user",
			LastName:  "123",
			Email:     "valid@gmail.com",
			Password:  "123456789",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRecorder()

		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(r, req)

		if r.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, r.Code)
		}
	})

}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
