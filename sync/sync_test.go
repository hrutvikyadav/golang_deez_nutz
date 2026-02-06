package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increementing counter 3 times should track the count as 3", func(t *testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})
}

func assertCount(t testing.TB, counter Counter, want int) {
	got := counter.CountValue()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
