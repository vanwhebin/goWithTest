package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	data := "hello world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("got %s, want %s", response.Body.String(), data)
	}

}

func TestCancelHandle(t *testing.T) {
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}

func TestServer(t *testing.T) {
	data := "hello world"
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}

		store.assertWasCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())

		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		store.assertWasNotCancelled()
	})

}
