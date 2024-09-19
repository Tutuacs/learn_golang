package auth

import (
	"fmt"
	"net/http"

	"todo-app/internal/auth"
	"todo-app/internal/common"
	"todo-app/internal/user"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var body auth.RegisterUserDTO

	err := common.GetBody(r, &body)
	if err != nil {
		resp := map[string]string{"Error": fmt.Sprintf("Unexpected error getting body: %v", err)}
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

	store, err := auth.NewStore()
	if err != nil {
		resp := map[string]string{"Fail": fmt.Sprintf("Server cant connect with DB: %v", err)}
		common.WriteResponse(w, http.StatusPreconditionFailed, resp)
		return
	}
	defer store.CloseStore()

	userStore, _ := user.NewStoreInstance(store.GetConn())

	usr, err := userStore.FindUserByEmail(body.Email)
	if err != nil {
		resp := map[string]string{"Error": fmt.Sprintf("Error verifying if user exist: %v", err)}
		common.WriteResponse(w, http.StatusInternalServerError, resp)
		return
	}

	if usr.ID == 0 {
		resp := map[string]string{"Conflict": fmt.Sprintf("Cant use an existent user email: %v", err)}
		common.WriteResponse(w, http.StatusConflict, resp)
		return
	}

	resp, err := store.RegisterUser(body)
	if err != nil {
		resp := map[string]string{"Error": fmt.Sprintf("Error while register user: %v", err)}
		common.WriteResponse(w, http.StatusInternalServerError, resp)
		return
	}

	usr, err = userStore.FindUserByID(resp.User.ID)
	if err != nil {
		resp := map[string]string{"Created": fmt.Sprintf("Error while returning user: %v", err)}
		common.WriteResponse(w, http.StatusOK, resp)
		return
	}

	resp.User = usr

	common.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	common.WriteResponse(w, http.StatusOK, map[string]string{})
}
