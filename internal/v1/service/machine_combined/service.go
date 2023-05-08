package machine_combined

import (
	"eirc.app/internal/v1/entity/machine_combined"
	model "eirc.app/internal/v1/structure/machine_combined"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	MachineCombinedListLast(input *model.Fields) (quantity int64, output []*model.Machine_Combined_Last, err error)
	GetByPIDMachineCombinedListLast(input *model.Fields) (quantity int64, output []*model.Machine_Combined_Last, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity machine_combined.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: machine_combined.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
