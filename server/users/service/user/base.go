package user

import "github.com/danielbintar/angel/server/users/service"

var Create = func(form CreateForm) (interface{}, *service.Error) {
	return service.Start(&form)
}

var Login = func(form LoginForm) (interface{}, *service.Error) {
	return service.Start(&form)
}
