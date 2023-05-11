package repositories

import "gopkg.in/guregu/null.v4"

type JwtRepository interface {
	GenerateToken(data interface{}, exp null.Int) (string, error)
	GetData(token string) (interface{}, error)
}
