package sales_call_record

import (
	"eirc.app/internal/v1/service/sales_call_record"
	model "eirc.app/internal/v1/structure/sales_call_records"
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
	SalesCallRecordService sales_call_record.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		SalesCallRecordService: sales_call_record.New(db),
	}
}
