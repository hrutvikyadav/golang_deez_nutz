package contezt

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		// race both cases, so we take an action based on whichever happens first
		select {
		case d:= <- data:
			fmt.Fprint(w, d)
		case <- ctx.Done():
			store.Cancel()
		}
	}
}
