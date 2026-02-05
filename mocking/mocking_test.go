package main

import (
	"bytes"
	"testing"
)

// If we can mock time.Sleep we can use DI to use the mock instead of a "real" time.Sleep;
// then we can spy on the calls to make assertions on them.
// type Sleeper interface {
// 	Sleep()
// }
type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{Calls: 0}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to Sleeper want 3 got %d", spySleeper.Calls)
	}
}
