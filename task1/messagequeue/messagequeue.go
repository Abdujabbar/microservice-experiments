package messagequeue

import (
	"bytes"
	"net/http"
)

//Farwarder interface
type Farwarder interface {
	Farward([]byte) error
}

//Service struct
type Service struct {
	host     string
	endpoint string
}

//NewService initilizer
func NewService(host, endpoint string) *Service {
	return &Service{
		host:     host,
		endpoint: endpoint,
	}
}

//Farward received data
func (s *Service) Farward(body []byte) error {
	_, err := http.Post(s.host+"/"+s.endpoint, "application/json", bytes.NewReader(body))
	return err
}
