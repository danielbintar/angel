package users

import (
	"github.com/danielbintar/angel/server/users/db"

	"github.com/jinzhu/gorm"
)

type UserManager struct {
	DB *gorm.DB
}

func Instance() *UserManager {
	m := &UserManager {
		DB: db.NewDB(),
	}

	return m
}
