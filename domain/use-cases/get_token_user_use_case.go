package usecases

import (
	"context"
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/domain/repositories"

	"gopkg.in/guregu/null.v4"
)

type GetTokenUserUseCase struct {
	JwtRepository repositories.JwtRepository
}

func (r *GetTokenUserUseCase) Execute(ctx context.Context, token string) (entities.User, error) {
	var user entities.User

	raw, err := r.JwtRepository.GetData(token)
	if err != nil {
		return user, nil
	}
	userRaw := raw.(map[string]interface{})
	user = entities.User{
		ID:       int64(userRaw["id"].(float64)),
		Username: userRaw["username"].(string),
		Name:     null.StringFrom(userRaw["name"].(string)),
	}
	return user, nil
}
