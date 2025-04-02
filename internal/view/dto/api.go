package dto

import "gopkg.in/guregu/null.v4"

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetParams struct {
	Sort   null.String `query:"sort"`
	Order  null.String `query:"order"`
	Limit  null.Int    `query:"limit"`
	Offset null.Int    `query:"offset"`
	CountParams
}

type CountParams struct {
	Search null.String `query:"search"`
}

type FindParams struct {
	ID int64 `param:"id"`
}

type GetData struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}
