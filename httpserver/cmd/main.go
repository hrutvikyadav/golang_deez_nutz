package main

import (
	"log"
	"net/http"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

type InMemoryStore struct {}

func (ims *InMemoryStore) GetScore(name string) int {
	return 69420
}

func (ims *InMemoryStore) PostWin(name string) {}

func main () {
	server := &go_httpserver.PlayerServer{&InMemoryStore{}}
	log.Fatal(http.ListenAndServe(":8099", server))
}
