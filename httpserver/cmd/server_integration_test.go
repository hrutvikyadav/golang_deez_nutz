package main

import (
	"testing"
	"net/http"
	"net/http/httptest"

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

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreReq(t, PLAYER))

	assertStatus(t, response.Code, http.StatusOK)
	assertPlayerScore(t, response.Body.String(), "3")

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
