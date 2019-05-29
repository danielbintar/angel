package user

import (
	// "github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/users/service"

	"gopkg.in/validator.v2"

	"golang.org/x/crypto/bcrypt"
)

type CreateForm struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func (self *CreateForm) Validate() *service.Error {
	if err := validator.Validate(self); err != nil {
		return &service.Error { Error: err.Error() }
	}

	return nil
}

func (self *CreateForm) Perform() (interface{}, *service.Error) {
	user := &model.User{Username: self.Username}

	// if err := db.DB().Where(&user).First(&user).Error; err != nil {
	// 	if err.Error() != "record not found" {
	// 		return nil, &service.Error { Error: err.Error(), Private: true }
	// 	}
	// } else {
		return user, &service.Error { Error: "username already used" }
	// }

	// 0 for using default cost
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(self.Password), 0)
	user.Username = self.Username
	user.Password = string(encryptedPassword)

	// if err := db.DB().Create(&user).Error; err != nil {
	// 	return nil, &service.Error { Error: err.Error(), Private: true }
	// }

	return user, nil
}
