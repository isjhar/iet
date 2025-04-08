package usecases

import (
	"context"

	"github.com/isjhar/iet/internal/domain/entities"
	"github.com/isjhar/iet/internal/domain/repositories"
)

type GeneratePairTokenUseCase struct {
	JwtRepository      repositories.JwtRepository
	JwtStoreRepository repositories.JwtStoreRepository
}

type GeneratePairTokenUseCaseResult struct {
	AccessToken  string
	RefreshToken string
}

func (r *GeneratePairTokenUseCase) Execute(ctx context.Context, user entities.User) (GeneratePairTokenUseCaseResult, error) {
	var result GeneratePairTokenUseCaseResult
	accessTokenPayload := r.createPayload(user)
	accessToken, err := r.JwtRepository.GenerateToken(accessTokenPayload)
	if err != nil {
		return result, err
	}

	refresToken, err := r.JwtRepository.GenerateRefreshToken()
	if err != nil {
		return result, err
	}

	err = r.JwtStoreRepository.StoreRefreshToken(ctx, repositories.StoreRefreshTokenParams{
		Token:     refresToken,
		SessionID: user.Username,
	})
	if err != nil {
		return result, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refresToken
	return result, err
}

func (r *GeneratePairTokenUseCase) createPayload(user entities.User) map[string]any {
	accessTokenPayload := make(map[string]any)
	accessTokenPayload["id"] = user.ID
	accessTokenPayload["username"] = user.Username
	accessTokenPayload["name"] = user.Name
	return accessTokenPayload
}
