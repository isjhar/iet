package usecases

import "github.com/isjhar/iet/internal/domain/repositories"

type TransactionalUseCase struct {
	Repository repositories.Repository
}
