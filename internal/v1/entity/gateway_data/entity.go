package gateway_data

import (
	model "eirc.app/internal/v1/structure/gateway_data"
	gg_data_demand "eirc.app/internal/v1/structure/gg_data_demand"

	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByClassificationTitle(input *model.Field) (output *model.Table, err error) 
	GetByDataDemand(input string) (output *[]gg_data_demand.Review, err error)
	Deleted(input *model.Table) (err error)
	Updated(input *model.Table) (err error)
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
