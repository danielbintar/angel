package users

import (
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server-library/pubsub"
)

type UserManager struct {
	DatabaseManager db.DatabaseManagerInterface
	Publisher       pubsub.AsyncPublisher
}

func Instance(db db.DatabaseManagerInterface, pb pubsub.AsyncPublisher) *UserManager {
	m := &UserManager {
		DatabaseManager: db,
		Publisher: pb,
	}

	return m
}
