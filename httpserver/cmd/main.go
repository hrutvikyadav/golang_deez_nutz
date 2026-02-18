package main

import (
	"log"
	"net/http"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

type InMemoryStore struct {
	scores map[string]int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{make(map[string]int)}
}

func (ims *InMemoryStore) GetScore(name string) int {
	return ims.scores[name]
}

func (ims *InMemoryStore) PostWin(name string) {
	ims.scores[name]++ // adding entry to nil map? solve by constructor
}

func main () {
	server := &go_httpserver.PlayerServer{&InMemoryStore{}}
	log.Fatal(http.ListenAndServe(":8099", server))
}
