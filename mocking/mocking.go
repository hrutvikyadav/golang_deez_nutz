package main

import (
	"fmt"
	"io"
	"os"
)

const countDownN = 3
const finalWord = "Go!"

func Countdown(w io.Writer) {
	for i:= countDownN; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
