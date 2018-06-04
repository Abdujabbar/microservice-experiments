ATUKO third party service
=============================

All envirement values stored in .env file listing is here:

- THIRD_PARTY_SITE: third party site address, example: http://localhost:8000
- THIRD_PARTY_ENDPOINT: third party site endpoint name: fake-third-party-stats
- THIRD_PARTY_REQUEST_LIMIT_PER_MINUTE: third party site requests limit per minute, example: 90
- MESSAGE_QUEUE_SITE: message queue site address, example: http://localhost:8000
- MESSAGE_QUEUE_ENDPOINT: message queue site endpoint, example: fake-event
- LOCAL_PORT: port for running on your machine, for example: 8000


for running service you can just run with current command:

```
    go run main.go
```

Available endpoint is here: 
===============================
- /stat/{anyString} - it runs the transport fetches the data from third party service, and farwards to message queue if it would have errors, on response in this endpoint you can see json data of information


Running tests
===============================
For running all test cases run this:

```
    go test -v ./...
```






