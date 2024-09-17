package user

import "todo-app/internal/common"

func (h *Handler) StartRoutes(baseRoute *common.Router) {
	router := baseRoute.BaseRoute("/user")
	router.HandleFunc("", h.FindUser).Methods("GET")
}
