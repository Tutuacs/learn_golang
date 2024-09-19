package todo

import (
	"todo-app/internal/common"
)

func (h *Handler) StartRoutes(baseRoute *common.Router) {
	router := baseRoute.BaseRoute("/todo")
	router.HandleFunc("", h.CreateTodo).Methods("POST")
	router.HandleFunc("", h.ListTodos).Methods("GET")
	router.HandleFunc("/{id}", h.FindTodo).Methods("GET")
	router.HandleFunc("/{id}", h.UpdateTodo).Methods("PUT")
	router.HandleFunc("/{id}", h.DeleteTodo).Methods("DELETE")
}
