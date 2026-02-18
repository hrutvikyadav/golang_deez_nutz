package main

import go_httpserver "github.com/hrutvikyadav/go-httpserver"

type InMemoryStore struct {
	// the store will be accessed by http handlers.
	// in go for each request, new goroutine is spinned up
	// data race for 👇
	scores map[string]int
	// TODO: put it under a mutex, use locking to avoid race conditions
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

func (ims *InMemoryStore) GetLeague() []go_httpserver.Player{
	return nil
}
