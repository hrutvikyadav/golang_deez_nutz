package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/hrutvikyadav/go-specs-greet/domain/interactions"

)

func GreetHandler (writer http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// Isolate domain logic (Greet) from handler (which only responsible for http stuff)
	// This means we can test domain logic in isolation from the http specification
	fmt.Fprint(writer, go_specs_greet.Greet(name))
}

func CurseHandler (w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, go_specs_greet.Curse(name))
}
