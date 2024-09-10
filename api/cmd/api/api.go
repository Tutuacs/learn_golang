package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Tutuacs/learn_golang/api.git/service/product"
	"github.com/Tutuacs/learn_golang/api.git/service/user"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subRouter)

	log.Println(string(Green), "Listening on", s.addr, Reset)

	return http.ListenAndServe(s.addr, router)
}
