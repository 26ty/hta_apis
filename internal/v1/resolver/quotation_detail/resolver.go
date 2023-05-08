package quotation_detail

import (
	"eirc.app/internal/v1/service/quotation_detail"
	model "eirc.app/internal/v1/structure/quotation_details"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	QuotationDetailService quotation_detail.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		QuotationDetailService: quotation_detail.New(db),
	}
}
