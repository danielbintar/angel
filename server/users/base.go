package users

import (
	"github.com/danielbintar/angel/server-library/pubsub"
	"github.com/danielbintar/angel/server/users/db"
)

type UserManager struct {
	DatabaseManager db.DatabaseManagerInterface
	Publisher       pubsub.AsyncPublisher
}
