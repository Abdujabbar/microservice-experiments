package thirdparty

import (
	"sync"
	"testing"
	"time"
)

func TestRequestLimit(t *testing.T) {
	requestCnt := requestCounter{
		limit:   9,
		mutex:   &sync.Mutex{},
		counter: make(map[string]int),
	}
	tm := time.Now().Format("02.01.2006 15:04")
	for i := 0; i < 10; i++ {
		requestCnt.increment(tm)
	}

	if requestCnt.isValid(tm) {
		t.Errorf("Request counter failed on checking minute requests")
	}

}
