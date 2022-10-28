package selectIO

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowUrl := "http://github.com"
	fastUrl := "http://baidu.com"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if want != got {
		t.Errorf("expected %v, while got %v", want, got)
	}
}

func TestHttpRacer(t *testing.T) {
	slowServer := getTestHttpServer(20 * time.Millisecond)
	fastServer := getTestHttpServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if want != got {
		t.Errorf("expected %v, while got %v", want, got)
	}
}

func getTestHttpServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestHttpSelectRacer(t *testing.T) {
	slowServer := getTestHttpServer(20 * time.Millisecond)
	fastServer := getTestHttpServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := RacerWithSelect(slowUrl, fastUrl)

	if want != got {
		t.Errorf("expected %v, while got %v", want, got)
	}
}

func TestTimeoutRaceWithSelect(t *testing.T) {
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := getTestHttpServer(11 * time.Second)
		fastServer := getTestHttpServer(10 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		maxDuration := 10 * time.Second
		_, err := RacerWithSelectTimeOut(slowUrl, fastUrl, maxDuration)

		if err == nil {
			t.Error("expected an error while get none")
		}

		want := fastUrl
		got := RacerWithSelect(slowUrl, fastUrl)

		if want != got {
			t.Errorf("expected %v, while got %v", want, got)
		}
	})
}
