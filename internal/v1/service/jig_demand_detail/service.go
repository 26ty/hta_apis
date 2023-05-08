package jig_demand_detail

import (
	"eirc.app/internal/v1/entity/jig_demand_detail"
	model "eirc.app/internal/v1/structure/jig_demand_details"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	UpdatedByJigID(input *model.Updated) (err error)
	GetByJigID(input *model.Field) (output *model.Base, err error)
}

type service struct {
	Entity jig_demand_detail.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: jig_demand_detail.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
