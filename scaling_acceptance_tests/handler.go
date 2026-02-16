package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler (writer http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// Isolate domain logic (Greet) from handler (which only responsible for http stuff)
	// This means we can test domain logic in isolation from the http specification
	fmt.Fprint(writer, Greet(name))
}
