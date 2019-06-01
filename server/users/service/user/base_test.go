package user_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users/factory"
	"github.com/danielbintar/angel/server/users/service/user"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	manager := factory.MockBase("broken_find_user_by_username")
	form := user.CreateForm {
		Username: "a",
		Password: "a",
		Manager: manager,
	}

	u, err := user.Create(form)
	assert.Nil(t, u)
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	manager := factory.MockBase("broken_find_user_by_username")
	form := user.LoginForm {
		Username: "a",
		Password: "a",
		Manager: manager,
	}

	u, err := user.Login(form)
	assert.Nil(t, u)
	assert.NotNil(t, err)
}