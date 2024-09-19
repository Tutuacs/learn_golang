package auth

import (
	"fmt"
	"net/http"

	"todo-app/internal/auth"
	"todo-app/internal/common"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var body auth.RegisterUserDTO

	err := common.GetBody(r, body)
	if err != nil {
		resp := map[string]string{"Error": "Unexpected error getting body"}
		common.WriteResponse(w, http.StatusBadRequest, resp)
		return
	}

	if err := common.Validate.Struct(body); err != nil {
		resp := map[string]string{"Validate error": fmt.Sprintf("Unexpected validate error: %v", err)}
		common.WriteResponse(w, http.StatusBadRequest, resp)
		return
	}

	body.Password, err = auth.HashPassword(body.Password)
	if err != nil {
		resp := map[string]string{"Encrypt error": fmt.Sprintf("Unexpected error while encrypting pass: %v", err)}
		common.WriteResponse(w, http.StatusInternalServerError, resp)
		return
	}

	common.WriteResponse(w, http.StatusOK, map[string]string{})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	common.WriteResponse(w, http.StatusOK, map[string]string{})
}
