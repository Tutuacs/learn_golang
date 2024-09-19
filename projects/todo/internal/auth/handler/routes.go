package auth

import "todo-app/internal/common"

func (h *Handler) StartRoutes(baseRouter *common.Router) {
	router := baseRouter.BaseRoute("/auth")
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/register", h.Register).Methods("POST")
}
