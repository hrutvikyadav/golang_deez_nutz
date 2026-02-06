package contezt

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	cancelled bool
	t *testing.T
}

// make use of context so that in case of cancellation this information can be consistently conveyed throughout the application
func (s *SpyStore) Fetch() string {
	time.Sleep(100*time.Millisecond)
	return s.response
}
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertCancel() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("expected to cancel but store was not told to cancel")
	}
}

func (s *SpyStore) assertNotCancel() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("should not have cancelled")
	}
}

func TestServer(t *testing.T) {
	t.Run("return data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data, false, t} // TODO: contrib upstream
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
		store.assertNotCancel()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{ data, false, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(10*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		store.assertCancel()
	})
}
