package gateway_data

import (
	"eirc.app/internal/v1/entity/gateway_data"
	model "eirc.app/internal/v1/structure/gateway_data"
	gg_data_demand "eirc.app/internal/v1/structure/gg_data_demand"

	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByClassificationTitle(input *model.Field) (output *model.Base, err error)
	GetByDataDemand(input string) (output *[]gg_data_demand.Review, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity gateway_data.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: gateway_data.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
