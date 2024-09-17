package user

import (
	"net/http"

	"todo-app/internal/user"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {

	db, err := user.NewStore()
	if err != nil {
		return
	}

	defer db.CloseStore()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
