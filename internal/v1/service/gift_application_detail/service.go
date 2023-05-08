package gift_application_detail

import (
	"eirc.app/internal/v1/entity/gift_application_detail"
	model "eirc.app/internal/v1/structure/gift_application_details"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity gift_application_detail.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: gift_application_detail.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
