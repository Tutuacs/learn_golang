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

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
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
			t.Error(string(Red), "expected status code", http.StatusBadRequest, ", got", r.Code, Reset)
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
			t.Error(string(Red), "expected status code", http.StatusCreated, ", got", r.Code, Reset)
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
