package quotation

import (
	"eirc.app/internal/v1/service/quotation"
	model "eirc.app/internal/v1/structure/quotations"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	Updated_Bonita(input *model.Updated_Bonita) interface{}
	GetByQIDQuotationDetailListUser(input *model.Field) interface{}
	QuotationDetailListUser(input *model.Fields) interface{}
}

type resolver struct {
	QuotationService quotation.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		QuotationService: quotation.New(db),
	}
}
