package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler (writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Hello, %s", r.URL.Query().Get("name"))
}
