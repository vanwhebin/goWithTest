package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("expected %q, while got %q", want, got)
	}

}

func TestConsistentCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountConsistentDown(buffer)

	got := buffer.String()
	want := "3\n2\n1\nGO!"

	if got != want {
		t.Errorf("expected %q, while got %q", want, got)
	}

}

func TestMonitorSleepCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	CountDownMonitor(buffer, spySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGO!"
	assertSleepTimes := 3

	if got != want {
		t.Errorf("expected %q, while got %q", want, got)
	}

	if spySleeper.Calls != assertSleepTimes {
		t.Errorf("not enough calls to sleeper, expected %d, while got %d", assertSleepTimes, spySleeper.Calls)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeper := &CountdownOperationSpy{}
		buffer := &bytes.Buffer{}

		CountDownMonitor(buffer, spySleeper)

		want := []string{
			sleep,
			sleep,
			sleep,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("want calls %v got %v", want, spySleeper.Calls)
		}

	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
