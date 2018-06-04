package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/Abdujabbor/atuko/task1/types"
	"github.com/julienschmidt/httprouter"
)

//dummy stats endpoint for third party testing
func stats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	statItems := types.Stats{}
	t := time.Now()
	for i := 0; i < 10; i++ {
		statItems = append(statItems, types.Stat{
			ID:     rand.Intn(100000),
			Date:   t.Format("2006-02-01T15:04:00Z"),
			Shows:  rand.Intn(10000),
			Clicks: rand.Intn(10000),
			Costs:  rand.Float32(),
		})
	}
	jsonBody, err := json.Marshal(statItems)
	if err != nil {
		w.Write([]byte(""))
	} else {
		w.Write(jsonBody)
	}
}
