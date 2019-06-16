package service_test

import (
	"testing"

	// "golang.org/x/crypto/bcrypt"

	// "github.com/danielbintar/angel/server/users/factory"
	// "github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/consumer-kafka/service"

	"github.com/stretchr/testify/assert"
)

type LoadConfigValidationTestCase struct {
	Form      service.LoadConfigForm
	NilResult bool
}

func TestLoadConfigFormValidate(t *testing.T) {
	cases := generateLoadConfigFormValidateTestCase()

	for _, c := range cases {
		r := c.Form.Validate()
		if c.NilResult {
			assert.Nil(t, r)
		} else {
			assert.NotNil(t, r)
		}
	}
}

func generateLoadConfigFormValidateTestCase() []LoadConfigValidationTestCase {
	cases := []LoadConfigValidationTestCase {
		LoadConfigValidationTestCase {
			Form: service.LoadConfigForm {},
			NilResult: false,
		},
		LoadConfigValidationTestCase {
			Form: service.LoadConfigForm {
				MicroName: "a",
			},
			NilResult: false,
		},
		LoadConfigValidationTestCase {
			Form: service.LoadConfigForm {
				ConsumerName: "a",
			},
			NilResult: false,
		},
		LoadConfigValidationTestCase {
			Form: service.LoadConfigForm {
				MicroName: "a",
				ConsumerName: "a",
			},
			NilResult: true,
		},
	}
	return cases
}
