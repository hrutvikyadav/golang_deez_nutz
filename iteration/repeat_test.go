package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat (t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("sixnine", 15)
	}
}

func ExampleRepeat() {
	repeated := Repeat("signal_", 4)
	fmt.Println(repeated)
	// Output: signal_signal_signal_signal_
}
