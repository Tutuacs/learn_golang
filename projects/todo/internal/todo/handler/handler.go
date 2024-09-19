package todo

import (
	"fmt"
	"net/http"

	"todo-app/internal/common"
	"todo-app/internal/todo"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	var body todo.NewTodoDTO

	err := common.GetBody(r, body)
	if err != nil {
		resp := map[string]string{"Error": "Error getting the body request"}
		common.WriteResponse(w, http.StatusBadRequest, resp)
		return
	}

	if err := common.Validate.Struct(body); err != nil {
		common.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	store, err := todo.NewStore()
	if err != nil {
		resp := map[string]string{"Fail": fmt.Sprintf("Server cant connect with DB: %v", err)}
		common.WriteResponse(w, http.StatusPreconditionFailed, resp)
		return
	}
	defer store.CloseStore()

	dbResp, err := store.CreateTodo(body)
	if err != nil {
		resp := map[string]string{"Error": fmt.Sprintf("Unexpected error on creating: %v", err)}
		common.WriteResponse(w, http.StatusInternalServerError, resp)
		return
	}

	resp := map[string]todo.Todo{"Created": dbResp}

	common.WriteResponse(w, http.StatusCreated, resp)
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	common.WriteResponse(w, http.StatusOK, map[string]string{})
}

func (h *Handler) FindTodo(w http.ResponseWriter, r *http.Request) {
	common.WriteResponse(w, http.StatusCreated, map[string]string{})
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	common.WriteResponse(w, http.StatusCreated, map[string]string{})
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	common.WriteResponse(w, http.StatusCreated, map[string]string{})
}
