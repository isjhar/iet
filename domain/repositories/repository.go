package repositories

import (
	"context"

	"gopkg.in/guregu/null.v4"
)

type Repository interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type GetParams struct {
	Limit  null.Int
	Offset null.Int
	Sort   null.String
	Order  null.String
	Search null.String
}
