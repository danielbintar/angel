package users_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/factory"

	libFactory "github.com/danielbintar/angel/server-library/factory"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	database := factory.MockDatabase()
	publisher := libFactory.MockAsyncPublisher()
	assert.NotNil(t, users.Instance(database, publisher))
}
