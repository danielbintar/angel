package user

import (
	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/service"

	"gopkg.in/validator.v2"

	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Manager  *users.UserManager `    validate:"nonzero"`
}

func (self *LoginForm) Validate() *service.Error {
	if err := validator.Validate(self); err != nil {
		return &service.Error { Error: err.Error() }
	}

	return nil
}

func (self *LoginForm) Perform() (interface{}, *service.Error) {
	user, err := self.Manager.DatabaseManager.FindUserByUsername(self.Username)

	if err != nil {
		return nil, &service.Error { Error: err.Error(), Private: true }
	}
	if user == nil {
		return user, &service.Error { Error: "wrong username or password" }
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(self.Password)); err != nil {
		return nil, &service.Error { Error: "wrong username or password" }
	}

	return user, nil
}
