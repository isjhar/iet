
package repositories

import (
	"context"

	"github.com/isjhar/iet/internal/domain/entities"
)

type CustomerContractDetailRepository interface {
	Count(ctx context.Context, arg CountCustomerContractDetailsParams) (int64, error)
	Get(ctx context.Context, arg GetCustomerContractDetailsParams) (entities.CustomerContractDetails, error)
}

type GetCustomerContractDetailsParams struct {
	GetParams
	FilterCustomerContractDetailParams
}

type CountCustomerContractDetailsParams struct {
	FilterParams
	FilterCustomerContractDetailParams
}

type FilterCustomerContractDetailParams struct {
}
