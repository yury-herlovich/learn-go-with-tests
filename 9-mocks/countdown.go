package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var COUNT_START = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := COUNT_START; i >= 0; i-- {
		sleeper.Sleep()

		if i > 0 {
			fmt.Fprintln(w, i)
		} else {
			fmt.Fprintf(w, "Go!")
		}
	}
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
