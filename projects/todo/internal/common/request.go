package common

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"todo-app/pkg/utils"
)

func GetBody(r *http.Request, response interface{}) {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(response)
	utils.ErrorPanic(err)
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
