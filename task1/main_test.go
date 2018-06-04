package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	handlers "github.com/Abdujabbor/atuko/task1/handlers"
	messagequeue "github.com/Abdujabbor/atuko/task1/messagequeue"
	thirdparty "github.com/Abdujabbor/atuko/task1/thirdparty"
	"github.com/Abdujabbor/atuko/task1/transport"
	"github.com/julienschmidt/httprouter"
)

func TestTransportNotFound(t *testing.T) {
	router := httprouter.New()
	ts := httptest.NewServer(router)
	defer ts.Close()
	thirdPartyService := thirdparty.NewService(ts.URL, "random", 90)
	messageQueueService := messagequeue.NewService(ts.URL, "event")
	transportService := transport.NewService(thirdPartyService, messageQueueService)
	handlers.Init(router, transportService)
	err := transportService.Run("somestring")
	if !strings.HasPrefix(err.Error(), fmt.Sprintf("Wrong response code from third party service: %v", http.StatusNotFound)) {
		t.Errorf("Failed on testing not found")
	}
}

func TestEmptyJsonThirdParty(t *testing.T) {
	router := httprouter.New()
	ts := httptest.NewServer(router)
	defer ts.Close()
	thirdPartyService := thirdparty.NewService(ts.URL, "/fake-third-party-empty-stats", 90)
	messageQueueService := messagequeue.NewService(ts.URL, "event")
	transportService := transport.NewService(thirdPartyService, messageQueueService)
	handlers.Init(router, transportService)
	err := transportService.Run("somestring")

	if err.Error() != "unexpected end of JSON input" {
		t.Errorf("Failed on testing json response")
	}
}

func TestWrongStructThirdParty(t *testing.T) {
	router := httprouter.New()
	ts := httptest.NewServer(router)
	defer ts.Close()
	thirdPartyService := thirdparty.NewService(ts.URL, "/fake-third-party-wrong-stats-struct", 90)
	messageQueueService := messagequeue.NewService(ts.URL, "event")
	transportService := transport.NewService(thirdPartyService, messageQueueService)
	handlers.Init(router, transportService)
	err := transportService.Run("somestring")
	if !strings.HasPrefix(err.Error(), "Wrong data items from third party") {
		t.Errorf("Failed on validate third party response struct")
	}
}
