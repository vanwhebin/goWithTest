package sync

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func (counter *Counter) Inc() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.value++
}

func (counter *Counter) Value() int {
	return counter.value
}
