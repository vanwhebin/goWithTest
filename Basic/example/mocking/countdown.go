package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func CountConsistentDown(out io.Writer) {
	count := countDownStart
	for i := count; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func CountDownConfigurable(out io.Writer, sleeper Sleeper) {
	count := countDownStart
	for i := count; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	// CountConsistentDown(os.Stdout)
	CountDownConfigurable(os.Stdout, sleeper)
}
