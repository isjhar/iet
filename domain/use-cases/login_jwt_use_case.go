package usecases

import (
	"context"
	"isjhar/template/echo-golang/domain/repositories"
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
	user, err := loginUseCase.Execute(ctx, LoginParams{
		Username: arg.Username,
		Password: arg.Password,
	})
	if err != nil {
		return result, err
	}

	userPayload := make(map[string]interface{})

	accessToken, err := r.JwtRepository.GenerateToken(userPayload)
	if err != nil {
		return result, err
	}
	result.AccessToken = accessToken
	return result, nil
}
