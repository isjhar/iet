package entities

import "gopkg.in/guregu/null.v4"

type User struct {
	ID       int64
	Username string
	Password string
	Name     null.String
}
