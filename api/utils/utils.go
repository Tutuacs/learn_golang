package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

func ParseJSON(r *http.Request, request interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	decode := json.NewDecoder(r.Body)
	return decode.Decode(request)
}

func WriteJSON(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	return encoder.Encode(response)
}

func WriteError(w http.ResponseWriter, status int, response error) {
	WriteJSON(w, status, map[string]string{"Error": response.Error()})
}

var Validate = validator.New()
