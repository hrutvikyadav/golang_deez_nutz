package go_httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

func TestGETPlayer(t *testing.T) {
	t.Run("should return scores of player", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/players/Oggy", nil)
		if err != nil { t.Fatal(err) }
		response := httptest.NewRecorder()

		go_httpserver.PlayerServer(response, request)

		got := response.Body.String()
		want := "69"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("should return scores of player2", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/players/Cockroach", nil)
		if err != nil { t.Fatal(err) }
		response := httptest.NewRecorder()

		go_httpserver.PlayerServer(response, request)

		got := response.Body.String()
		want := "420"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
