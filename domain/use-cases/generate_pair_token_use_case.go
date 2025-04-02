package usecases

import (
	"context"

	"github.com/isjhar/iet/domain/entities"
	"github.com/isjhar/iet/domain/repositories"

	"gopkg.in/guregu/null.v4"
)

type GeneratePairTokenUseCase struct {
	JwtRepository repositories.JwtRepository
}

type GeneratePairTokenUseCaseResult struct {
	AccessToken  string
	RefreshToken string
}

func (r *GeneratePairTokenUseCase) Execute(ctx context.Context, user entities.User) (GeneratePairTokenUseCaseResult, error) {
	var result GeneratePairTokenUseCaseResult
	accessTokenPayload := make(map[string]interface{})
	accessTokenPayload["id"] = user.ID
	accessTokenPayload["username"] = user.Username
	accessTokenPayload["name"] = user.Name
	accessToken, err := r.JwtRepository.GenerateToken(accessTokenPayload, null.IntFrom(60))
	if err != nil {
		return result, err
	}

	refresToken, err := r.JwtRepository.GenerateToken(user.Username, null.NewInt(0, false))
	if err != nil {
		return result, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refresToken
	return result, err
}
