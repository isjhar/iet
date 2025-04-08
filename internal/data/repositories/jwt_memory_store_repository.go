package repositories

import (
	"context"
	"time"

	"github.com/isjhar/iet/internal/config"
	"github.com/isjhar/iet/internal/data/models"
	"github.com/isjhar/iet/internal/domain/entities"
	"github.com/isjhar/iet/internal/domain/repositories"
	"gopkg.in/guregu/null.v4"
)

var refreshTokens map[string]models.MemoryRefreshToken

type JwtDbStoreRepository struct {
}

func (r JwtDbStoreRepository) StoreRefreshToken(ctx context.Context, arg repositories.StoreRefreshTokenParams) error {
	expiresAt := time.Now().Add(time.Second * time.Duration(config.Jwt.AccessTokenExpiresIn.Int64))
	refreshToken := models.MemoryRefreshToken{
		SessionID: null.StringFrom(arg.SessionID),
		Token:     null.StringFrom(arg.Token),
		ExpiresAt: null.TimeFrom(expiresAt),
	}

	refreshTokens[refreshToken.Token.String] = refreshToken

	return nil
}

func (r JwtDbStoreRepository) ValidateRefreshToken(ctx context.Context, token string) (string, error) {
	var result string

	if _, ok := refreshTokens[token]; !ok {
		return result, entities.ErrorUnauthorized
	}
	refreshToken := refreshTokens[token]
	result = refreshToken.SessionID.String

	return result, nil
}

func (r JwtDbStoreRepository) RevokeToken(ctx context.Context, token string) error {
	if _, ok := refreshTokens[token]; !ok {
		return nil
	}

	delete(refreshTokens, token)
	return nil
}
