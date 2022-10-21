package mocking

import (
	"bytes"
	"fmt"
	"time"
)

const finalWord = "GO!"
const countDownStart = 3
const sleep = "sleep"
const write = "write"

func CountDown(out *bytes.Buffer) {
	fmt.Fprintf(out, "3")
}

func CountConsistentDown(out *bytes.Buffer) {
	count := countDownStart
	for i := count; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func CountDownMonitor(out *bytes.Buffer, sleeper Sleeper) {
	count := countDownStart
	for i := count; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
