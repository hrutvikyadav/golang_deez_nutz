package main

import (
	"log"
	"net/http"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

func main () {
	server := go_httpserver.NewPlayerServer(&InMemoryStore{})
	log.Fatal(http.ListenAndServe(":8099", server))
}
