package factory_manufacturing

import (
	"eirc.app/internal/v1/service/factory_manufacturing"
	model "eirc.app/internal/v1/structure/factory_manufacturings"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created_List) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	FactoryManufacturingService factory_manufacturing.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		FactoryManufacturingService: factory_manufacturing.New(db),
	}
}
