package quotation

import (
	"eirc.app/internal/v1/entity/quotation"
	model "eirc.app/internal/v1/structure/quotations"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	Updated_Bonita(input *model.Updated_Bonita) (err error)
	QuotationDetailListUser(input *model.Fields) (quantity int64, output []*model.QuotationDetail, err error)
	GetByQIDQuotationDetailListUser(input *model.Field) (output *model.QuotationDetail, err error)
}

type service struct {
	Entity quotation.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: quotation.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
