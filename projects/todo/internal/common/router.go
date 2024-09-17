package common

import "github.com/gorilla/mux"

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{router: new(mux.Router)}
}

func (r *Router) BaseRoute(route string) *mux.Router {
	router := r.router.PathPrefix(route).Subrouter()
	return router
}

func (r *Router) GetRouter() *mux.Router {
	return r.router
}
