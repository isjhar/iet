
package repositories

import (
	"context"

	"github.com/isjhar/iet/internal/data/models"

	"github.com/isjhar/iet/pkg"

	"github.com/isjhar/iet/internal/domain/entities"
	"github.com/isjhar/iet/internal/domain/repositories"

	"gorm.io/gorm"
)

type CustomerContractDetailRepository struct {
}

func (r CustomerContractDetailRepository) Count(ctx context.Context, arg repositories.CountCustomerContractDetailsParams) (int64, error) {
	var results int64
	var query *gorm.DB

	if arg.ID.Valid {
		query = r.createFindFilterQuery(ctx, arg)
	} else {
		query = r.createGetFilterQuery(ctx, arg)
	}

	err := query.Count(&results).Error
	if err != nil {
		pkg.LogError("error count customer_contract_detail %v", err)
		return results, entities.InternalServerError
	}

	return results, nil
}

func (r CustomerContractDetailRepository) Get(ctx context.Context, arg repositories.GetCustomerContractDetailsParams) (entities.CustomerContractDetails, error) {
	results := entities.CustomerContractDetails{}
	var err error

	if arg.ID.Valid {
		item, err := r.find(ctx, arg)
		if err != nil {
			return results, err
		}
		results = append(results, item)
	} else {
		results, err = r.get(ctx, arg)
		if err != nil {
			return results, err
		}
	}
	return results, nil
}

func (r CustomerContractDetailRepository) get(ctx context.Context, arg repositories.GetCustomerContractDetailsParams) (entities.CustomerContractDetails, error) {
	var results entities.CustomerContractDetails
	var models models.CustomerContractDetails

	query := r.createGetFilterQuery(ctx, repositories.CountCustomerContractDetailsParams{
		FilterCustomerContractDetailParams: arg.FilterCustomerContractDetailParams,
		FilterParams:                       arg.FilterParams,
	})

	limit := int(-1)
	if arg.Limit.Valid {
		limit = int(arg.Limit.Int64)
	}
	offset := int(0)
	if arg.Offset.Valid {
		offset = int(arg.Offset.Int64)
	}
	orderBy := ""
	switch arg.Sort.String {
	default:
		orderBy += "customer_contract_detail.tanggal"
	}
	orderBy += " " + GetOrderQuery(arg.Order)

	query = query.Limit(limit).Offset(offset).Order(orderBy + `, customer_contract_detail."OID"`)
	err := query.Find(&models).Error
	if err != nil {
		pkg.LogError("error get customer_contract_detail %v", err)
		return results, entities.InternalServerError
	}

	for _, model := range models {
		result := entities.CustomerContractDetail{
			ID: model.ID,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r CustomerContractDetailRepository) find(ctx context.Context, arg repositories.GetCustomerContractDetailsParams) (entities.CustomerContractDetail, error) {
	var model models.CustomerContractDetail
	var result entities.CustomerContractDetail

	query := r.createFindFilterQuery(ctx, repositories.CountCustomerContractDetailsParams{
		FilterParams:                       arg.FilterParams,
		FilterCustomerContractDetailParams: arg.FilterCustomerContractDetailParams,
	})

	err := query.First(&model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		pkg.LogError("error find customer_contract_detail %v", err)
		return result, entities.InternalServerError
	}

	result.ID = model.ID
	return result, nil
}

func (r CustomerContractDetailRepository) createGetFilterQuery(ctx context.Context, arg repositories.CountCustomerContractDetailsParams) *gorm.DB {
	query := ORM.WithContext(ctx).
		Model(&models.CustomerContractDetail{})

	query = r.applyDefaultFilterQuery(query)

	query = r.applySearchFilterQuery(query, arg)

	return query
}

func (r CustomerContractDetailRepository) createFindFilterQuery(ctx context.Context, arg repositories.CountCustomerContractDetailsParams) *gorm.DB {
	query := ORM.WithContext(ctx).
		Model(&models.CustomerContractDetail{})

	query = r.applyDefaultFilterQuery(query)

	query = query.Where(`customer_contract_detail."OID" = ?`, arg.ID)

	return query
}

func (r CustomerContractDetailRepository) applySearchFilterQuery(query *gorm.DB, arg repositories.CountCustomerContractDetailsParams) *gorm.DB {
	if arg.Search.Valid && arg.Search.String != "" {
		query = query.Where(
			ORM.Where("customer_contract_detail.nomor ilike ?", "%"+arg.Search.String+"%"),
		)
	}

	return query
}

func (r CustomerContractDetailRepository) applyDefaultFilterQuery(query *gorm.DB) *gorm.DB {
	query = query.Where(`customer_contract_detail."GCRecord" is null`)

	return query
}
