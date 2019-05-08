package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/handler"

	"github.com/julienschmidt/httprouter"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	m := users.Instance()
	h := handler.NewBaseHandler(m)
	assert.NotNil(t, h)
}

func TestHealthz(t *testing.T) {
	uri := "/healthz"
	m := users.Instance()
	h := handler.NewBaseHandler(m)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil { t.Fatal(err) }

	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET(uri, h.Healthz)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `ok`, rr.Body.String())
}
