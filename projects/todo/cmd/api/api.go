package api

import (
	"fmt"
	"net/http"

	"todo-app/internal/common"
	todoInternal "todo-app/internal/todo/handler"
	userInternal "todo-app/internal/user/handler"
	"todo-app/pkg/logs"
)

type APIServer struct {
	addr string
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (srv *APIServer) Start() error {
	router := common.NewRouter()

	todoHandler := todoInternal.NewHandler()
	todoHandler.StartRoutes(router)

	userHandler := userInternal.NewHandler()
	userHandler.StartRoutes(router)

	logs.OkLog(fmt.Sprintf("Running server on port %v", srv.addr))

	return http.ListenAndServe(srv.addr, router.GetRouter())
}
