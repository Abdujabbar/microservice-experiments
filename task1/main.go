package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	handlers "github.com/Abdujabbor/atuko/task1/handlers"
	messagequeue "github.com/Abdujabbor/atuko/task1/messagequeue"
	thirdparty "github.com/Abdujabbor/atuko/task1/thirdparty"
	"github.com/Abdujabbor/atuko/task1/transport"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	localPort := os.Getenv("LOCAL_PORT")
	if localPort == "" {
		localPort = "8000"
	}
	thirdPartySite := os.Getenv("THIRD_PARTY_SITE")
	thirdPartyEndpoint := os.Getenv("THIRD_PARTY_ENDPOINT")
	thirdPartyRequestLimit := os.Getenv("THIRD_PARTY_REQUEST_LIMIT")
	requestLimit, err := strconv.Atoi(thirdPartyRequestLimit)
	if err != nil {
		requestLimit = 90
	}
	thirdpartyService := thirdparty.NewService(thirdPartySite, thirdPartyEndpoint, requestLimit)

	messageQueueSite := os.Getenv("MESSAGE_QUEUE_SITE")
	messageQueueEndpoint := os.Getenv("MESSAGE_QUEUE_ENDPOINT")
	messageQueueService := messagequeue.NewService(messageQueueSite, messageQueueEndpoint)
	transportService := transport.NewService(thirdpartyService, messageQueueService)
	router := httprouter.New()
	handlers.Init(router, transportService)
	log.Fatal(http.ListenAndServe(":"+localPort, router))
}
