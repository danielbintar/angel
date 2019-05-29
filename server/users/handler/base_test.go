package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/handler"

	"github.com/julienschmidt/httprouter"

	"github.com/stretchr/testify/assert"

	"github.com/subosito/gotenv"
)

func TestNewHandler(t *testing.T) {
	m := users.Instance()
	h := handler.NewBaseHandler(m)
	assert.NotNil(t, h)
}

func TestHealthz(t *testing.T) {
	m := users.Instance()
	h := handler.NewBaseHandler(m)

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil { t.Fatal(err) }

	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/healthz", h.Healthz)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `ok`, rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	t.Run("no body", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		req, err := http.NewRequest("POST", "/users", nil)
		if err != nil { t.Fatal(err) }

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("no username", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"password":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("invalid type username", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":12}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("invalid type password", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"password":12}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("no password", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("mysql problem", func(t *testing.T) {
		os.Setenv("TEST_MYSQL_USER", "invalid_user")
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusServiceUnavailable, rr.Code)
		os.Unsetenv("TEST_MYSQL_USER")
	})

	t.Run("success", func(t *testing.T) {
		gotenv.Load("../.env")
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users", h.CreateUser)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("no body", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		req, err := http.NewRequest("POST", "/users/my-session", nil)
		if err != nil { t.Fatal(err) }

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("no username", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"password":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("invalid type username", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":12}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("invalid type password", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"password":12}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("no password", func(t *testing.T) {
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})

	t.Run("mysql problem", func(t *testing.T) {
		os.Setenv("TEST_MYSQL_USER", "invalid_user")
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusServiceUnavailable, rr.Code)
		os.Unsetenv("TEST_MYSQL_USER")
	})

	t.Run("success", func(t *testing.T) {
		gotenv.Load("../.env")
		m := users.Instance()
		h := handler.NewBaseHandler(m)

		body := []byte(`{"username":"123456","password":"123456"}`)
		req, err := http.NewRequest("POST", "/users/my-session", bytes.NewBuffer(body))
		if err != nil { t.Fatal(err) }
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.POST("/users/my-session", h.Login)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
