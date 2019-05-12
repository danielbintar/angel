package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/router"

	"github.com/stretchr/testify/assert"
)

type route struct {
	method string
	uri    string
	code   int
}

func TestNewRouter(t *testing.T) {
	assert.NotNil(t, router.NewRouter())
}

func TestPublic(t *testing.T) {
	r := router.NewRouter()
	m := users.Instance()
	router.Public(r, m)
	rr := httptest.NewRecorder()

	tests := []route{
		{ "GET", "/healthz", http.StatusOK },
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, test.uri, nil)
		if err != nil { t.Fatal(err) }

		r.ServeHTTP(rr, req)
		assert.Equal(t, test.code, rr.Code)
	}
}
