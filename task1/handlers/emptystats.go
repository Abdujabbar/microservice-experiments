package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//dummy emptystats endpoint for third party testing
func emptystats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}
