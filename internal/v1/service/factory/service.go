package factory

import (
	"eirc.app/internal/v1/entity/factory"
	model "eirc.app/internal/v1/structure/factorys"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	GetByFIDFLMListUser(input *model.Field) (output *model.Base, err error)
	FLMListUser(input *model.Fields) (quantity int64, output []*model.Base, err error)
}

type service struct {
	Entity factory.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: factory.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
