package router

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) PrepareRoutes() {

}

func (r *Router) BuildRoute() (route string) {
	return
}
