package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func wrongstats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(`[{"ss": 123}, {"qq":12345}]`))
}
