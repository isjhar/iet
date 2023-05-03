package usecases

import "isjhar/template/echo-golang/domain/repositories"

type TransactionalUseCase struct {
	Repository repositories.Repository
}
