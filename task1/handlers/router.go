package handlers

import (
	transport "github.com/Abdujabbor/atuko/task1/transport"
	"github.com/julienschmidt/httprouter"
)

var transportService *transport.Service

//Init method
func Init(router *httprouter.Router, transport *transport.Service) {
	transportService = transport
	router.GET("/", home)
	router.GET("/stat/:account", stat)

	//Fake endpoints for testing
	router.POST("/fake-event", event)
	router.GET("/fake-third-party-stats/:account", stats)
	router.GET("/fake-third-party-empty-stats/:account", emptystats)
	router.GET("/fake-third-party-wrong-stats-struct/:account", wrongstats)
}

//Response struct
type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
