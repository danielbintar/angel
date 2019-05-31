package router

import (
	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/handler"
	userService "github.com/danielbintar/angel/server/users/service/user"

	"github.com/danielbintar/angel/server-library/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	return httprouter.New()
}

func Public(r *httprouter.Router, m *users.UserManager) {
	h := handler.NewBaseHandler(m)

	r.GET("/healthz", h.Healthz)
	r.POST("/users", middleware.Adapt(h.CreateUser, middleware.MustHaveForm(&userService.CreateForm{})))
	r.POST("/users/my-session", middleware.Adapt(h.Login, middleware.MustHaveForm(&userService.LoginForm{})))
}
