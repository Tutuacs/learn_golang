package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"service/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.List()
	if err != nil {
		log.Printf("Erro ao listar os todos: %v", err)
	}

	if len(todos) == 0 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.Todo{})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
