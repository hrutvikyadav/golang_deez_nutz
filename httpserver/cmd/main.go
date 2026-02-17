package main

import (
	"log"
	"net/http"

	go_httpserver "github.com/hrutvikyadav/go-httpserver"
)

func main () {
	handler := http.HandlerFunc(go_httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":8099", handler))
}
