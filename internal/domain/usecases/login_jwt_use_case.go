package usecases

import (
	"context"

	"github.com/isjhar/iet/internal/domain/repositories"
)

type LoginJwtUseCase struct {
	UserRepository repositories.UserRepository
	JwtRepository  repositories.JwtRepository
}

type LoginJwtUseCaseParams struct {
	Username string
	Password string
}

type LoginJwtUseCaseResult struct {
	AccessToken  string
	RefreshToken string
}

func (r *LoginJwtUseCase) Execute(ctx context.Context, arg LoginJwtUseCaseParams) (LoginJwtUseCaseResult, error) {
	var result LoginJwtUseCaseResult
	loginUseCase := LoginUseCase{
		UserRepository: r.UserRepository,
	}
	user, err := loginUseCase.Execute(ctx, LoginParams(arg))
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
