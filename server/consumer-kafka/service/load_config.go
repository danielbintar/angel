package service

import (
	"fmt"
	"io/ioutil"
	"os"

	serviceLib "github.com/danielbintar/angel/server-library/service"
	"github.com/danielbintar/angel/server/consumer-kafka/model"

	"gopkg.in/validator.v2"

	"gopkg.in/yaml.v2"
)

type LoadConfigForm struct {
	MicroName    string `validate:"nonzero"`
	ConsumerName string `validate:"nonzero"`
}

func (self *LoadConfigForm) Validate() *serviceLib.Error {
	if err := validator.Validate(self); err != nil {
		return &serviceLib.Error{Error: err.Error()}
	}

	return nil
}

func (self *LoadConfigForm) Perform() (interface{}, *serviceLib.Error) {
	fmt.Println(os.Getwd())
	yamlFile, err := ioutil.ReadFile("consumers/" + self.MicroName + "/" + self.ConsumerName + "/config.yaml")
	if err != nil {
		panic(self.ConsumerName + " not found in " + self.MicroName)
	}

	var config model.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic("wrong configuration file on " + self.ConsumerName + " in " + self.MicroName)
	}

	return config, nil
}
