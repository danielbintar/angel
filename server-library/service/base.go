package service

type Error struct {
	Error   string
	Private bool
}

type Service interface {
	Validate() *Error
	Perform() (interface{}, *Error)
}

func Start(svc Service) (interface{}, *Error) {
	var object interface{}

	err := svc.Validate()
	if err != nil {
		return object, err
	}

	return svc.Perform()
}
