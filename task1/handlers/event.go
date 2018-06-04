package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	types "github.com/Abdujabbor/atuko/task1/types"
	"github.com/julienschmidt/httprouter"
)

//dummy event endpoint for local checking
func event(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		statItems := types.Stats{}
		err = json.Unmarshal(requestBody, &statItems)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(statItems)
		}
	}
}
