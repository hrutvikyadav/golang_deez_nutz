package go_httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

func TestGETPlayer(t *testing.T) {
	t.Run("should return scores of player", func(t *testing.T) {
		request := newGetScoreReq(t, "Oggy")
		response := httptest.NewRecorder()

		go_httpserver.PlayerServer(response, request)

		assertPlayerScore(t, response.Body.String(), "69")

	})

	t.Run("should return scores of player2", func(t *testing.T) {
		request := newGetScoreReq(t, "Cockroach")
		response := httptest.NewRecorder()

		go_httpserver.PlayerServer(response, request)

		assertPlayerScore(t, response.Body.String(), "420")
	})
}

func newGetScoreReq(t testing.TB, player string) *http.Request {
	t.Helper()
	request, err := http.NewRequest(http.MethodGet, "/players/" + player, nil)
	if err != nil { t.Fatal(err) }
	return request
}

func assertPlayerScore(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
