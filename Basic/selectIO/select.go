package selectIO

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := getRequestDuration(a)
	bDuration := getRequestDuration(b)
	if aDuration > bDuration {
		return b
	}
	buffer := &bytes.Buffer{}
	fmt.Fprint(buffer, a, aDuration)
	fmt.Fprint(buffer, b, bDuration)
	return a
}

func getRequestDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func RacerWithSelect(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func RacerWithSelectTimeOut(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}
