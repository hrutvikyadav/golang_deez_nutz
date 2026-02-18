package go_httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (sp StubPlayerStore) GetScore(name string) int {
	return sp.scores[name]
}

func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{map[string]int{
		"Oggy": 69,
		"Cockroach": 420,
	}}
	server := &go_httpserver.PlayerServer{&store}

	t.Run("should return scores of player", func(t *testing.T) {
		request := newGetScoreReq(t, "Oggy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertPlayerScore(t, response.Body.String(), "69")

	})

	t.Run("should return scores of player2", func(t *testing.T) {
		request := newGetScoreReq(t, "Cockroach")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertPlayerScore(t, response.Body.String(), "420")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreReq(t, "Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestPOSTWins(t *testing.T) {
	store := StubPlayerStore{map[string]int{
	}}
	server := &go_httpserver.PlayerServer{&store}

	t.Run("returns accepted on POST/players/name", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/players/Bheem", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted
		assertStatus(t, got, want)
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

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
