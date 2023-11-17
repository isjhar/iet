package entities

import "fmt"

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
var InvalidParams = NewServerError("Invalid params")

func FileSizeReachLimit(maxSize int64) error {
	return NewServerError(fmt.Sprintf("File exceeds limit %d bytes", maxSize))
}
