package usecases

import "github.com/isjhar/iet/domain/repositories"

type TransactionalUseCase struct {
	Repository repositories.Repository
}
