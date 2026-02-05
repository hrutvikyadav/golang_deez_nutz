package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownN = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i:= countDownN; i > 0; i-- {
		fmt.Fprintln(w, i)
		// time.Sleep(1 * time.Second) // WARN: we now have a dependency on `Sleep`ingwhich we need to extrace out of this function so that we have more control
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
