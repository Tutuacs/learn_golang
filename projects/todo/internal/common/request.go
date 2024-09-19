package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func GetBody(r *http.Request, response interface{}) error {

	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(response)
	return err
}

func GetParam(r *http.Request, paramName string) (param string) {
	param = mux.Vars(r)[paramName]
	return
}

func WriteResponse(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(w)
	encode.Encode(result)
}

var Validate = validator.New()
