package repositories

import (
	"context"
	"isjhar/template/echo-golang/domain/entities"
)

type UserRepository interface {
	Find(ctx context.Context, username string) (entities.User, error)
}
