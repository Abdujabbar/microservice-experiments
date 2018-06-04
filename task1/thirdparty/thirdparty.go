package thirdparty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	types "github.com/Abdujabbor/atuko/task1/types"
	"gopkg.in/validator.v2"
)

//Fetcher interface
type Fetcher interface {
	Fetch(account string) ([]byte, error)
}

//Service struct
type Service struct {
	host           string
	endpoint       string
	requestCounter requestCounter
}

//NewService initializer
func NewService(host, endpoint string, requestLimit int) *Service {
	return &Service{
		host:     host,
		endpoint: endpoint,
		requestCounter: requestCounter{
			limit:   requestLimit,
			mutex:   &sync.Mutex{},
			counter: make(map[string]int),
		},
	}
}

func (s *Service) getCacheKeyTime() string {
	t := time.Now()
	return t.Format("02.01.2006 15:04")
}

func (s *Service) isAvailable() bool {
	cacheKey := s.getCacheKeyTime()
	if s.requestCounter.isValid(cacheKey) {
		s.requestCounter.increment(cacheKey)
		return true
	}
	return false
}

//Fetch get response from third party
func (s *Service) Fetch(account string) ([]byte, error) {
	if !s.isAvailable() {
		return nil, fmt.Errorf("Fetching temporary not available")
	}

	client := http.Client{}
	address := s.host + "/" + s.endpoint + "/" + account
	resp, err := client.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong response code from third party service: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	stats := types.Stats{}
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return nil, err
	}
	for _, st := range stats {
		if err := validator.Validate(st); err != nil {
			return nil, fmt.Errorf("Wrong data items from third party: %v", st)
		}
	}
	return body, nil
}
