package thirdparty

import "sync"

type requestCounter struct {
	limit   int
	mutex   *sync.Mutex
	counter map[string]int
}

func (r *requestCounter) increment(t string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, ok := r.counter[t]; !ok {
		r.counter[t] = 0
	}
	r.counter[t]++
}

func (r *requestCounter) get(t string) int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.counter[t]
}

func (r *requestCounter) isValid(t string) bool {
	currentValue := r.get(t)
	if currentValue < r.limit {
		return true
	}
	return false
}
