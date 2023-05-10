package usecases

import (
	"context"
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/domain/repositories"
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
	accessToken, err := r.JwtRepository.GenerateToken(accessTokenPayload)
	if err != nil {
		return result, err
	}

	refresToken, err := r.JwtRepository.GenerateToken(user.Username)
	if err != nil {
		return result, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refresToken
	return result, err
}
