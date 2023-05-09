package entities

type ServiceError struct {
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewServerError(message string) error {
	return &ServiceError{Message: message}
}

var InternalServerError = NewServerError("Internal server error")
var EntityNotFound = NewServerError("Entity not found")
var WrongPassword = NewServerError("Wrong password")
