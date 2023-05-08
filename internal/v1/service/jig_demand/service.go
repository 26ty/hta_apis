package jig_demand

import (
	"eirc.app/internal/v1/entity/jig_demand"
	model "eirc.app/internal/v1/structure/jig_demands"
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
	GetByJIDJigDetailListUser(input *model.Field) (output *model.Base, err error)
	JigDetailListUser(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByUserIDListJD(input *model.Users) (quantity int64, output []*model.JD, err error)
}

type service struct {
	Entity jig_demand.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: jig_demand.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
