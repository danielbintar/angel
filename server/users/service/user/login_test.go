package user_test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/danielbintar/angel/server/users/factory"
	"github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/users/service/user"

	"github.com/stretchr/testify/assert"

	"github.com/subosito/gotenv"
)

type LoginFormValidationCaseTest struct {
	Form      user.LoginForm
	NilResult bool
}

func TestLoginFormValidate(t *testing.T) {
	cases := generateLoginFormValidateTestCase()

	for _, c := range cases {
		r := c.Form.Validate()
		if c.NilResult {
			assert.Nil(t, r)
		} else {
			assert.NotNil(t, r)
		}
	}
}

func TestLoginFormPerform(t *testing.T) {
	t.Run("fail find user", func(t *testing.T) {
		manager := factory.MockBase("broken_find_user_by_username")
		form := user.LoginForm {
			Username: "a",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})

	t.Run("user not exists", func(t *testing.T) {
		manager := factory.MockBase("find_user_by_username_404")
		form := user.LoginForm {
			Username: "lala",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})

	t.Run("wrong password", func(t *testing.T) {
		manager := factory.MockBase()
		form := user.LoginForm {
			Username: "a",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		gotenv.Load("../../.env")
		manager := factory.MockBase("real_database")
		form := user.LoginForm {
			Username: "a",
			Password: "a",
			Manager: manager,
		}

		pass, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 0)
		manager.DatabaseManager.InsertUser(&model.User{Username: form.Username, Password: string(pass)})

		u, err := form.Perform()
		assert.NotNil(t, u)
		assert.Nil(t, err)
	})
}

func generateLoginFormValidateTestCase() []LoginFormValidationCaseTest {
	manager := factory.MockBase()
	cases := []LoginFormValidationCaseTest {
		LoginFormValidationCaseTest {
			Form: user.LoginForm {},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Username: "a",
			},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Password: "a",
			},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Manager: manager,
			},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Username: "a",
				Password: "a",
			},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Username: "a",
				Manager: manager,
			},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Password: "a",
				Manager: manager,
			},
			NilResult: false,
		},
		LoginFormValidationCaseTest {
			Form: user.LoginForm {
				Username: "a",
				Password: "a",
				Manager: manager,
			},
			NilResult: true,
		},
	}
	return cases
}
