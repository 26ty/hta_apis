package factory

import (
	"eirc.app/internal/v1/service/factory"
	model "eirc.app/internal/v1/structure/factorys"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	GetByFIDFLMListUser(input *model.Field) interface{}
	FLMListUser(input *model.Fields) interface{}
	SearchFactory(input *model.Fields) interface{}
}

type resolver struct {
	FactoryService factory.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		FactoryService: factory.New(db),
	}
}
