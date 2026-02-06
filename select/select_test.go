package racer

import (
	"fmt"
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
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return error if a server takes more than 10 seconds to respond", func(t *testing.T){
		server := makeDelayedServer(10 * time.Second)

		// defer server.Close() // WARN: neeed to remove this to actually make timeout work otherwise it keeps waiting to close

		_, err := ConfigurableRacer(server.URL, server.URL, 2 * time.Second)

		if err == nil {
			t.Error("expected error but did not get one")
		}
		fmt.Printf("t: %v\n", err)
	})
}

func makeDelayedServer(delay time.Duration) (*httptest.Server) {
	return httptest.NewServer(http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
