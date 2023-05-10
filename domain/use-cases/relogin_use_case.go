package usecases

import (
	"context"
	"isjhar/template/echo-golang/domain/repositories"
)

type ReloginUseCase struct {
	UserRepository repositories.UserRepository
	JwtRepository  repositories.JwtRepository
}

func (r *ReloginUseCase) Execute(ctx context.Context, refreshToken string) (LoginJwtUseCaseResult, error) {
	var result LoginJwtUseCaseResult
	refereshTokenData, err := r.JwtRepository.GetData(refreshToken)
	if err != nil {
		return result, err
	}
	username := refereshTokenData.(string)
	user, err := r.UserRepository.Find(ctx, username)
	if err != nil {
		return result, err
	}

	generatePairTokenUseCase := GeneratePairTokenUseCase{
		JwtRepository: r.JwtRepository,
	}
	pairToken, err := generatePairTokenUseCase.Execute(ctx, user)
	if err != nil {
		return result, err
	}
	result.AccessToken = pairToken.AccessToken
	result.RefreshToken = pairToken.RefreshToken
	return result, nil
}
