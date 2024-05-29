package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func TestServer(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		data := "hello mito"
		svr := Server(&SpyStore{data, false})

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Error("err")
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {

	})
}
