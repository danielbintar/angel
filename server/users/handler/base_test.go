package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/danielbintar/angel/server/users/factory"
	"github.com/danielbintar/angel/server/users/handler"
	"github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/users/router"

	"github.com/julienschmidt/httprouter"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	m := factory.MockBase()
	h := handler.NewBaseHandler(m)
	assert.NotNil(t, h)
}

func TestHealthz(t *testing.T) {
	m := factory.MockBase()
	h := handler.NewBaseHandler(m)

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/healthz", h.Healthz)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `ok`, rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	t.Run("no username", func(t *testing.T) {
		m := factory.MockBase()

		body := []byte(`{"password":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("no password", func(t *testing.T) {
		m := factory.MockBase()

		body := []byte(`{"username":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("mysql problem", func(t *testing.T) {
		m := factory.MockBase("broken_find_user_by_username")

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusServiceUnavailable, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		m := factory.MockBase("find_user_by_username_404")

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("no username", func(t *testing.T) {
		m := factory.MockBase()

		body := []byte(`{"password":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("no password", func(t *testing.T) {
		m := factory.MockBase()

		body := []byte(`{"username":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("mysql problem", func(t *testing.T) {
		m := factory.MockBase("broken_find_user_by_username")

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusServiceUnavailable, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		m := factory.MockBase("real_database")

		plainPass := "123456"
		pass, _ := bcrypt.GenerateFromPassword([]byte(plainPass), 0)
		m.DatabaseManager.InsertUser(&model.User{Username: "123456", Password: string(pass)})

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		r := httprouter.New()
		router.Public(r, m)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
