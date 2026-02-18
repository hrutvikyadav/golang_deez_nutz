package go_httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Score int
}

type PlayerStore interface {
	GetScore(name string) int
	PostWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler // PlayerServer now exposes all apis of this embedding
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.Store = store

	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	p.Handler = router

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
	lTable := p.Store.GetLeague()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&lTable)
	// WARN: does not work w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
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
