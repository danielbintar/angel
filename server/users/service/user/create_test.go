package user_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users"
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

func generateValidateTestCase() []ValidationCaseTest {
	manager := users.Instance()
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
