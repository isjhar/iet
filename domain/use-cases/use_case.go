package usecases

import "gopkg.in/guregu/null.v4"

type GetUseCaseParams struct {
	Limit  null.Int
	Offset null.Int
	Sort   null.String
	Order  null.String
	CountUseCaseParams
}

type CountUseCaseParams struct {
	Search null.String
}
