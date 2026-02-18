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
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processPost(w, r)
	case http.MethodGet:
		p.processGet(w, r)
	}
}

func (p *PlayerServer) processGet(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processPost(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.Store.PostWin(player)
	w.WriteHeader(http.StatusAccepted)
}
