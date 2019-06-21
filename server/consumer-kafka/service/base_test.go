package service_test

import (
	"encoding/json"
	"testing"

	"github.com/danielbintar/angel/server/consumer-kafka/model"
	"github.com/danielbintar/angel/server/consumer-kafka/service"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	base := "../"
	form := service.LoadConfigForm{
		MicroName:    "dummy-micro",
		ConsumerName: "valid-consumer",
		Base:         &base,
	}

	configI, err := service.LoadConfig(form)
	assert.Nil(t, err)
	assert.NotNil(t, configI)

	byteData, _ := json.Marshal(configI)
	var config model.Config
	json.Unmarshal(byteData, &config)

	assert.Equal(t, 2, len(config.Topics))
	if config.Topics[0] == "users_model-log" {
		assert.Equal(t, "items_model-log", config.Topics[1])
	} else {
		assert.Equal(t, "items_model-log", config.Topics[0])
		assert.Equal(t, "users_model-log", config.Topics[1])
	}
}
