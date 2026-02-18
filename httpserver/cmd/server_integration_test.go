package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

// Test recording wins and retrieving them; to and from the store
func TestServerStoreIntegration(t *testing.T) {
	const PLAYER = "Ben10"
	store := NewInMemoryStore()
	server := go_httpserver.NewPlayerServer(store)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinReq(t, PLAYER))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinReq(t, PLAYER))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinReq(t, PLAYER))

	t.Run("Get posted score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreReq(t, PLAYER))

		assertStatus(t, response.Code, http.StatusOK)
		assertPlayerScore(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getLeagueReq())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []go_httpserver.Player{
			{PLAYER, 3},
		}
		assertLeagues(t, got, want)
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
