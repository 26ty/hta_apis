package transaction_record

import (
	"eirc.app/internal/v1/entity/transaction_record"
	model "eirc.app/internal/v1/structure/transaction_record"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByDocumentIDUserList(input *model.Fields) (quantity int64, output []*model.Record_user_list, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity transaction_record.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: transaction_record.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
