package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello mito"
	svr := Server(&SpyStore{data})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	svr.ServeHTTP(res, req)

	if res.Body.String() != data {
		t.Error("err")
	}
}
