package go_httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetScore(name string) int
	PostWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{store, http.NewServeMux()}

	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	return p
}

func (p *PlayerServer) playersHandler (w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processPost(w, player)
	case http.MethodGet:
		p.processGet(w, player)
	}
}

func (p *PlayerServer) leagueHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *PlayerServer) processGet(w http.ResponseWriter, player string) {
	score := p.Store.GetScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processPost(w http.ResponseWriter, player string) {
	p.Store.PostWin(player)
	w.WriteHeader(http.StatusAccepted)
}
