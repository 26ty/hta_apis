package quotation

import (
	model "eirc.app/internal/v1/structure/quotations"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	Deleted(input *model.Table) (err error)
	Updated(input *model.Table) (err error)
	QuotationDetailListUser(input *model.Fields) (amount int64, output []*model.QuotationDetail, err error)
	GetByQIDQuotationDetailListUser(input *model.Field) (output *model.QuotationDetail, err error)
}

type entity struct {
	db *gorm.DB
}

func New(db *gorm.DB) Entity {
	return &entity{
		db: db,
	}
}

func (e *entity) WithTrx(tx *gorm.DB) Entity {
	return &entity{
		db: tx,
	}
}
