package contezt

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t *testing.T
}

// make use of context so that in case of cancellation this information can be consistently conveyed throughout the pipeline
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func(){
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
				log.Print("store got cancelled")
				return
			default:
				time.Sleep(100*time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case d := <- data:
		return d, nil
	case <- ctx.Done():
		return "", ctx.Err()
	}
}

func TestServer(t *testing.T) {
	t.Run("return data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	// t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	// 	data := "hello, world"
	// 	store := &SpyStore{ data, false, t}
	// 	svr := Server(store)
	//
	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
	//
	// 	cancellingCtx, cancel := context.WithCancel(request.Context())
	// 	time.AfterFunc(10*time.Millisecond, cancel)
	// 	request = request.WithContext(cancellingCtx)
	//
	// 	response := httptest.NewRecorder()
	//
	// 	svr.ServeHTTP(response, request)
	// 	store.assertCancel()
	// })
}
