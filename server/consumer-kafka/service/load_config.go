package service

import (
	serviceLib "github.com/danielbintar/angel/server-library/service"

	"gopkg.in/validator.v2"
)

type LoadConfigForm struct {
	MicroName    string `validate:"nonzero"`
	ConsumerName string `validate:"nonzero"`
}

func (self *LoadConfigForm) Validate() *serviceLib.Error {
	if err := validator.Validate(self); err != nil {
		return &serviceLib.Error { Error: err.Error() }
	}

	return nil
}
