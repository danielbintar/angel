package service

import "github.com/danielbintar/angel/server-library/service"

var LoadConfig = func(form LoadConfigForm) (interface{}, *service.Error) {
	return service.Start(&form)
}
