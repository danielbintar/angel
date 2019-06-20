package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielbintar/angel/server-library/middleware"

	"github.com/julienschmidt/httprouter"

	"github.com/stretchr/testify/assert"
)

type Form struct {
	Name string `json:"name"`
}

func TestMustHaveForm(t *testing.T) {
	mid := middleware.MustHaveForm(&Form{})

	t.Run("no body", func(t *testing.T) {
		router := httprouter.New()
		router.POST("/success", middleware.Adapt(success, mid))
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/success", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with non hash body", func(t *testing.T) {
		router := httprouter.New()
		router.POST("/success", middleware.Adapt(success, mid))
		w := httptest.NewRecorder()
		body := []byte(``)
		req, err := http.NewRequest("POST", "/success", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with wrong type body", func(t *testing.T) {
		router := httprouter.New()
		router.POST("/success", middleware.Adapt(success, mid))
		w := httptest.NewRecorder()
		body := []byte(`{"name":123}`)
		req, err := http.NewRequest("POST", "/success", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with body", func(t *testing.T) {
		router := httprouter.New()
		router.POST("/success", middleware.Adapt(success, mid))
		w := httptest.NewRecorder()
		body := []byte(`{}`)
		req, err := http.NewRequest("POST", "/success", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func success(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
