package repositories

import (
	"context"
)

type JwtStoreRepository interface {
	StoreRefreshToken(ctx context.Context, arg StoreRefreshTokenParams) error
	ValidateRefreshToken(ctx context.Context, token string) (string, error)
	RevokeToken(ctx context.Context, token string) error
}

type StoreRefreshTokenParams struct {
	Token     string
	SessionID string
}
