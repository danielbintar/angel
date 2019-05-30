package user_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users/factory"
	"github.com/danielbintar/angel/server/users/service/user"

	"github.com/stretchr/testify/assert"
)

type ValidationCaseTest struct {
	Form      user.CreateForm
	NilResult bool
}

func TestValidate(t *testing.T) {
	cases := generateValidateTestCase()

	for _, c := range cases {
		r := c.Form.Validate()
		if c.NilResult {
			assert.Nil(t, r)
		} else {
			assert.NotNil(t, r)
		}
	}
}

func TestPerform(t *testing.T) {
	t.Run("fail find user", func(t *testing.T) {
		manager := factory.MockBase("broken_find_user_by_username")
		form := user.CreateForm {
			Username: "a",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})

	t.Run("user already exists", func(t *testing.T) {
		manager := factory.MockBase()
		form := user.CreateForm {
			Username: "lala",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})

	t.Run("fail insert user", func(t *testing.T) {
		manager := factory.MockBase("find_user_by_username_404", "broken_insert_user")
		form := user.CreateForm {
			Username: "a",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		manager := factory.MockBase("find_user_by_username_404")
		form := user.CreateForm {
			Username: "a",
			Password: "a",
			Manager: manager,
		}
		u, err := form.Perform()
		assert.NotNil(t, u)
		assert.Nil(t, err)
	})
}

func generateValidateTestCase() []ValidationCaseTest {
	manager := factory.MockBase()
	cases := []ValidationCaseTest {
		ValidationCaseTest {
			Form: user.CreateForm {},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Username: "a",
			},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Password: "a",
			},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Manager: manager,
			},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Username: "a",
				Password: "a",
			},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Username: "a",
				Manager: manager,
			},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Password: "a",
				Manager: manager,
			},
			NilResult: false,
		},
		ValidationCaseTest {
			Form: user.CreateForm {
				Username: "a",
				Password: "a",
				Manager: manager,
			},
			NilResult: true,
		},
	}
	return cases
}
