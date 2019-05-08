package handler

import (
	"net/http"

	"github.com/danielbintar/angel/server/users"

	"github.com/julienschmidt/httprouter"
)

type baseHandler struct {
	m *users.UserManager
}

func NewBaseHandler(m *users.UserManager) *baseHandler {
	return &baseHandler{
		m: m,
	}
}

func (self *baseHandler) Healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
