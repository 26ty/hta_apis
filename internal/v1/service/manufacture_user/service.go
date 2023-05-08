package manufacture_user

import (
	"eirc.app/internal/v1/entity/manufacture_user"
	model "eirc.app/internal/v1/structure/manufacture_user"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByManufactureID(input *model.Field) (output []*model.ManufactureAccount, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity manufacture_user.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: manufacture_user.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
