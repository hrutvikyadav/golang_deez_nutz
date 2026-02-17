package go_httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetScore(p))
}


func GetScore(pname string) (score string ) {
	if pname == "Oggy" {
		score = "69"
	}
	if pname == "Cockroach" {
		score = "420"
	}

	return
}
