package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("HelloWorld"))
}
