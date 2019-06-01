package users_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/factory"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	database := factory.MockDatabase()
	assert.NotNil(t, users.Instance(database))
}
