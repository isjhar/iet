package models

import "gopkg.in/guregu/null.v4"

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FindRowsRequestParams struct {
	Sort   null.String `query:"sort"`
	Order  null.String `query:"order"`
	Limit  null.Int    `query:"limit"`
	Offset null.Int    `query:"offset"`
}

type FindRowsResponseData struct {
	Total int64 `json:"total"`
}
