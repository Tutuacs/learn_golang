package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"service/configs"
	"service/handlers"
)

func main() {
	err := configs.LoadConfigs()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/", handlers.Create)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.Get)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), router)
}
