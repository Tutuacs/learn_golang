package todo

import (
	"todo-app/internal/common"
)

func (h *Handler) StartRoutes(baseRoute *common.Router) {
	router := baseRoute.BaseRoute("/todo")
	router.HandleFunc("", h.CreateTodo).Methods("POST")
}
