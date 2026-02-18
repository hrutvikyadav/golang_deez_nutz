package go_httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (sp StubPlayerStore) GetScore(name string) int {
	return sp.scores[name]
}

func (sp *StubPlayerStore) PostWin(name string) {
	sp.winCalls = append(sp.winCalls, name)
}


func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{ "Oggy": 69, "Cockroach": 420, },
		[]string{},
	}
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
	store := StubPlayerStore{
		scores:   map[string]int{},
		winCalls: nil,
	}
	server := &go_httpserver.PlayerServer{&store}

	t.Run("records Wins in store on POST/players/name", func(t *testing.T) {
		request := newPostWinReq(t, "Bheem")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("expected Postwin to be called %d times but was called %d times", 1, len( store.winCalls ))
		}
	})
}

func newGetScoreReq(t testing.TB, player string) *http.Request {
	t.Helper()
	request, err := http.NewRequest(http.MethodGet, "/players/" + player, nil)
	if err != nil { t.Fatal(err) }
	return request
}

func newPostWinReq(t testing.TB, player string) (*http.Request) {
	request, err := http.NewRequest(http.MethodPost, "/players/" + player, nil)
	if err != nil {
		t.Fatal(err)
	}
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
