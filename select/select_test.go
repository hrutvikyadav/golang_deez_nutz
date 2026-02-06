package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare response speed of 2 servers and return winner", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return error if a server takes more than 10 seconds to respond", func(t *testing.T){
		serverOne := makeDelayedServer(17 * time.Second)
		serverTwo := makeDelayedServer(10 * time.Second)

		defer serverOne.Close()
		defer serverTwo.Close()

		_, err := Racer(serverOne.URL, serverTwo.URL)

		if err == nil {
			t.Error("expected error but did not get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) (*httptest.Server) {
	return httptest.NewServer(http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
