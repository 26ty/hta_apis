package gift_application

import (
	"eirc.app/internal/v1/entity/gift_application"
	model "eirc.app/internal/v1/structure/gift_applications"
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
	GetByGIDGiftDetailListUser(input *model.Field) (output *model.Base, err error)
	GiftDetailListUser(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByCaseID(input string) (output *model.Review, err error)
}

type service struct {
	Entity gift_application.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: gift_application.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
