package context

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type StubStore struct {
	response string
}

type SpyStore struct {
	response  string
	t         *testing.T
	cancelled bool
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// store.Cancel()
		// fmt.Fprint(w, store.Fetch())
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprintf(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}
