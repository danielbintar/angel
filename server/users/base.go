package users

import (
	"github.com/danielbintar/angel/server/users/db"
)

type UserManager struct {
	DatabaseManager db.DatabaseManagerInterface
}
