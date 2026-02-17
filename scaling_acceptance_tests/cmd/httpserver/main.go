package main

import (
	"log"
	"net/http"

	"github.com/hrutvikyadav/go-specs-greet/adapters/httpserver"
)

func main() {
	http.HandleFunc("/greet", httpserver.GreetHandler)
	http.HandleFunc("/curse", httpserver.CurseHandler)
	// TODO: use NewServeMux instead of the Default
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
