package usecases

import (
	"context"
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/domain/repositories"
)

type LoginUseCase struct {
	UserRepository repositories.UserRepository
}

type LoginParams struct {
	Username string
	Password string
}

func (i *LoginUseCase) Execute(ctx context.Context, arg LoginParams) (entities.User, error) {
	var user entities.User
	user, err := i.UserRepository.Find(ctx, arg.Username)
	if err != nil {
		return user, err
	}
	if user.Password == arg.Password {
		return user, entities.WrongPassword
	}
	return user, nil
}
