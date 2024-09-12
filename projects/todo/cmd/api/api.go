package api

import (
	"database/sql"

	"github.com/gorilla/mux"
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

func (srv *APIServer) Start() {
	router := mux.NewRouter()

	router.PathPrefix("/")

}
