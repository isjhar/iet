package entities

import "errors"

var InternalServerError = errors.New("Internal server error")
var EntityNotFound = errors.New("Entity not found")
var WrongPassword = errors.New("Wrong password")
