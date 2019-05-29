package user

import (
	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/model"
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
	user := &model.User{Username: self.Username}

	if err := self.Manager.DB.Where(&user).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, &service.Error { Error: "wrong username or password" }
		} else {
			return nil, &service.Error { Error: err.Error(), Private: true }
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(self.Password)); err != nil {
		return nil, &service.Error { Error: "wrong username or password" }
	}

	return user, nil
}
