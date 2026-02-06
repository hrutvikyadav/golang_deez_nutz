package syncz

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increementing counter 3 times should track the count as 3", func(t *testing.T){
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

// dont pass mutex (Counter) by value, pass by reference instead
// run `go vet` for more
func assertCount(t testing.TB, counter *Counter, want int) {
	got := counter.CountValue()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
