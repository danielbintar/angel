package gateway_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielbintar/angel/server/gateway"

	"github.com/stretchr/testify/assert"

	"github.com/subosito/gotenv"
)

func TestNewServer(t *testing.T) {
	assert.NotNil(t, gateway.NewServer())
}

func TestServeHTTP(t *testing.T) {
	gotenv.Load()
	s := gateway.NewServer()

	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/healthz", nil)
		if err != nil { t.Fatal(err) }

		rr := httptest.NewRecorder()
		s.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("exists", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/users", nil)
		if err != nil { t.Fatal(err) }

		rr := httptest.NewRecorder()
		s.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})
}
