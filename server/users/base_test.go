package users_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	assert.NotNil(t, users.Instance())
}
