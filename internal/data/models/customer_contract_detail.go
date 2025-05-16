
package models

type CustomerContractDetail struct {
	ID int64
}

func (CustomerContractDetail) TableName() string {
	return "customer_contract_detail_detail"
}

type CustomerContractDetails []CustomerContractDetail
