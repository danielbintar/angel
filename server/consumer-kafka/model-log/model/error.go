package model

type ErrInvalidMessage struct {
	message string
}

func NewErrInvalidMessage(message string) *ErrInvalidMessage {
	return &ErrInvalidMessage{
		message: message,
	}
}

func (self *ErrInvalidMessage) Error() string {
    return self.message
}
