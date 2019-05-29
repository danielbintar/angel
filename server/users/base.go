package users

import (
	"github.com/danielbintar/angel/server/users/db"
)

type UserManager struct {
	DatabaseManager db.DatabaseManagerInterface
}

func Instance(db db.DatabaseManagerInterface) *UserManager {
	m := &UserManager {
		DatabaseManager: db,
	}

	return m
}
