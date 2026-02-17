package go_httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/players/")
	var score string

	if p == "Oggy" {
		score = "69"
	}
	if p == "Cockroach" {
		score = "420"
	}

	fmt.Fprint(w, score)
}

