package user

import (
	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/users/service"

	"gopkg.in/validator.v2"

	"golang.org/x/crypto/bcrypt"
)

type CreateForm struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Manager  *users.UserManager `    validate:"nonzero"`
}

func (self *CreateForm) Validate() *service.Error {
	if err := validator.Validate(self); err != nil {
		return &service.Error { Error: err.Error() }
	}

	return nil
}

func (self *CreateForm) Perform() (interface{}, *service.Error) {
	user, err := self.Manager.DatabaseManager.FindUserByUsername(self.Username)

	if err != nil {
		return nil, &service.Error { Error: err.Error(), Private: true }
	}
	if user != nil {
		return user, &service.Error { Error: "username already used" }
	}

	user = &model.User{Username: self.Username}

	// 0 for using default cost
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(self.Password), 0)
	user.Username = self.Username
	user.Password = string(encryptedPassword)

	if err := self.Manager.DatabaseManager.InsertUser(user); err != nil {
		return nil, &service.Error { Error: err.Error(), Private: true }
	}

	return user, nil
}
