package service_test

import (
	"testing"

	"github.com/danielbintar/angel/server-library/service"

	"github.com/danielbintar/angel/server-library/slice"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	t.Run("error validation", func(t *testing.T) {
		form := ServiceForm{Options: []string{"error_validate"}}
		obj, err := service.Start(&form)
		assert.Nil(t, obj)
		assert.NotNil(t, err)
	})

	t.Run("success validation", func(t *testing.T) {
		form := ServiceForm{}
		_, err := service.Start(&form)
		assert.Nil(t, err)
	})
}

type ServiceForm struct {
	Options []string
}

func (self *ServiceForm) Validate() *service.Error {
	if slice.InStrings("error_validate", self.Options) {
		return &service.Error { Error: "error" }
	}

	return nil
}

func (self *ServiceForm) Perform() (interface{}, *service.Error) {
	return nil, nil
}
