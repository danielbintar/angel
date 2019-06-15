package service

// simple operation with validation
// if it not pass validation, it will not perform the operation
type Service interface {
	Validate() *Error
	Perform() (interface{}, *Error)
}

// Private for error privacy
// ex if there is database error and 
// the application doesn't need to tell the end user about it
// default is false
type Error struct {
	Error   string
	Private bool
}

// start the service
// will check the validation and do perform operation
func Start(svc Service) (interface{}, *Error) {
	var object interface{}

	err := svc.Validate()
	if err != nil {
		return object, err
	}

	return svc.Perform()
}
