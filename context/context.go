package contezt

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	// store might have more dependencies that need to derive from parent context.
	// instead of exposing the Cancel() api and having the caller cancel only the stores context, we handle it inside store itself;
	// passing along the context to any other functions between the req res pipeline.
}
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // TODO: iterate over error later
		}
		fmt.Fprint(w, data)
	}
}
