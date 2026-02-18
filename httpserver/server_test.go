package go_httpserver_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
	league []go_httpserver.Player
}

func (sp StubPlayerStore) GetScore(name string) int {
	return sp.scores[name]
}

func (sp *StubPlayerStore) PostWin(name string) {
	sp.winCalls = append(sp.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []go_httpserver.Player {
	return s.league
}

func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{ "Oggy": 69, "Cockroach": 420, },
		[]string{},
		nil,
	}
	server := go_httpserver.NewPlayerServer(&store)

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
	server := go_httpserver.NewPlayerServer(&store)

	t.Run("records Wins in store on POST/players/name", func(t *testing.T) {
		const PLAYAH = "Bheem"
		request := newPostWinReq(t, PLAYAH)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("expected Postwin to be called %d times but was called %d times", 1, len( store.winCalls ))
		}
		if store.winCalls[0] != PLAYAH {
			t.Errorf("expected to invoke store with %s got %s instead", PLAYAH, store.winCalls[0])
		}
	})
}

func TestGetLeague(t *testing.T) {
	wantedLeague := []go_httpserver.Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{nil, nil, wantedLeague}
	server := go_httpserver.NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := getLeagueReq()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeagues(t, got, wantedLeague)

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

func getLeagueFromResponse(t testing.TB, body io.Reader) []go_httpserver.Player {
	t.Helper()
	var got []go_httpserver.Player
	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("unable to parse response %q into Player slice, %v", body, err)
	}

	return got
}

func assertLeagues(t testing.TB, got, want []go_httpserver.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func getLeagueReq() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
