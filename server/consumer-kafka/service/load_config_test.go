package service_test

import (
	"encoding/json"
	"testing"

	// "golang.org/x/crypto/bcrypt"

	// "github.com/danielbintar/angel/server/users/factory"
	// "github.com/danielbintar/angel/server/users/model"
	"github.com/danielbintar/angel/server/consumer-kafka/model"
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

func TestLoadConfigFormPerform(t *testing.T) {
	t.Run("config not exists", func(t *testing.T) {
		form := service.LoadConfigForm {
			MicroName: "a",
			ConsumerName: "a",
		}

		assert.Panics(t, func() { form.Perform() })
	})

	t.Run("wrong configured config file", func(t *testing.T) {
		form := service.LoadConfigForm {
			MicroName: "dummy-micro",
			ConsumerName: "wrong-configured-consumer",
		}

		assert.Panics(t, func() { form.Perform() })
	})

	t.Run("config exists", func(t *testing.T) {
		form := service.LoadConfigForm {
			MicroName: "dummy-micro",
			ConsumerName: "success-consumer",
		}

		assert.NotPanics(t, func() { form.Perform() })

		configI, err := form.Perform()
		assert.Nil(t, err)
		assert.NotNil(t, configI)

		byteData, _ := json.Marshal(configI)
		var config model.Config
		json.Unmarshal(byteData, &config)

		assert.Equal(t, 2, len(config.Topics))
		if config.Topics[0] == "users-model-log" {
			assert.Equal(t, "items-model-log", config.Topics[1])
		} else {
			assert.Equal(t, "items-model-log", config.Topics[0])
			assert.Equal(t, "users-model-log", config.Topics[1])
		}
	})
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
