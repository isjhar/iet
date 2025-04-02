package repositories

import (
	"context"

	"github.com/isjhar/iet/domain/entities"
)

type UserRepository interface {
	Find(ctx context.Context, username string) (entities.User, error)
}
