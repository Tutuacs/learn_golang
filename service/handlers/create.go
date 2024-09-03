package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"service/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	// Fazendo decode do payload
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Erro ao fazer decode do json:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
