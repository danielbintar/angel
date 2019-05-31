package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/users/service"
	"github.com/danielbintar/angel/server/users/service/user"

	"github.com/julienschmidt/httprouter"
)

type baseHandler struct {
	manager *users.UserManager
}

func NewBaseHandler(m *users.UserManager) *baseHandler {
	return &baseHandler {
		manager: m,
	}
}

type Response struct {
	Data interface{} `json:"data"`
}

func WriteSuccess(w http.ResponseWriter, data interface{}) {
	resp := Response { Data: data }
	encoded, _ := json.Marshal(&resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(encoded))
}

func WriteServiceError(w http.ResponseWriter, err *service.Error) {
	if err.Private {
		http.Error(w, "Something has gone wrong, please try again in a few moment", http.StatusServiceUnavailable)
	} else {
		http.Error(w, err.Error, http.StatusUnprocessableEntity)
	}
}

func (self *baseHandler) Healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (self *baseHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var form user.CreateForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	form.Manager = self.manager
	userI, serviceErr := user.Create(form)
	if serviceErr != nil {
		WriteServiceError(w, serviceErr)
		return
	}

	byteData, _ := json.Marshal(userI)
	var user model.User
	json.Unmarshal(byteData, &user)

	WriteSuccess(w, &user)
}

func (self *baseHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var form user.LoginForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	form.Manager = self.manager
	userI, serviceErr := user.Login(form)
	if serviceErr != nil {
		WriteServiceError(w, serviceErr)
		return
	}

	byteData, _ := json.Marshal(userI)
	var user model.User
	json.Unmarshal(byteData, &user)

	WriteSuccess(w, &user)
}
