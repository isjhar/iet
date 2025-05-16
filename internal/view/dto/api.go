package dto

import (
	"gopkg.in/guregu/null.v4"
)

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetParams struct {
	Sort   null.String `query:"sort"`
	Order  null.String `query:"order"`
	Limit  null.Int    `query:"limit"`
	Offset null.Int    `query:"offset"`
	FilterParams
}

type FilterParams struct {
	Search null.String `query:"search"`
	ID     null.Int    `query:"id" param:"id"`
}

type FindParams struct {
	ID int64 `param:"id"`
}

type GetData struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

type GetItems struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}

type Response struct {
	Message string `json:"message"`
}

type CreateResponse struct {
	Response
	Data int64 `json:"data"`
}

type FloatResponse struct {
	Response
	Data float64 `json:"data"`
}

type IntResponse struct {
	Response
	Data int64 `json:"data"`
}
