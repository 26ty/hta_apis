package sales_call_record

import (
	"eirc.app/internal/v1/entity/sales_call_record"
	model "eirc.app/internal/v1/structure/sales_call_records"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	GetBySIDAccount(input *model.Field) (output *model.SalesCallRecord_Account, err error)
	AccountList(input *model.Fields) (quantity int64, output []*model.SalesCallRecord_Account, err error)
}

type service struct {
	Entity sales_call_record.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: sales_call_record.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
