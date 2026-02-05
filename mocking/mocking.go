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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}


func Countdown(w io.Writer, sleeper Sleeper) {
	for i:= countDownN; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	configurableSleeper := &ConfigurableSleeper{
		duration: 2 * time.Second,
		sleep: time.Sleep,
	}
	Countdown(os.Stdout, configurableSleeper)
}
