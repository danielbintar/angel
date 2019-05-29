package router

import (
	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/handler"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	return httprouter.New()
}

func Public(r *httprouter.Router, m *users.UserManager) {
	h := handler.NewBaseHandler(m)

	r.GET("/healthz", h.Healthz)
	r.POST("/users", h.CreateUser)
	r.POST("/users/my-session", h.Login)
}
