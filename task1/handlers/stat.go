package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func stat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	account := params.ByName("account")
	response := Response{
		Success: true,
		Error:   "",
	}
	err := transportService.Run(account)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
	}
	jsonResponseBody, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonResponseBody)
}
