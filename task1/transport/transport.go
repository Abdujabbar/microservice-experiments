package transport

import thirdparty "github.com/Abdujabbor/atuko/task1/thirdparty"
import messagequeue "github.com/Abdujabbor/atuko/task1/messagequeue"

//Service transport
type Service struct {
	fetcher   thirdparty.Fetcher
	farwarder messagequeue.Farwarder
}

//NewService service initialize
func NewService(fetcher thirdparty.Fetcher, farwarder messagequeue.Farwarder) *Service {
	return &Service{
		fetcher:   fetcher,
		farwarder: farwarder,
	}
}

//Run runs full process fetching and farwarding
func (s *Service) Run(account string) error {
	response, err := s.fetcher.Fetch(account)
	if err != nil {
		return err
	}
	err = s.farwarder.Farward(response)
	return err
}
