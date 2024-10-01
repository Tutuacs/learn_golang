package api

import (
	"log"
	"net/http"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Run() error {

	router := http.NewServeMux()
	router.HandleFunc("GET /user", func(http.ResponseWriter, *http.Request) {

	})
	router.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {

	})

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Println("Server has started")

	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}

}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check auth
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unautorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Middleware func(next http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}
