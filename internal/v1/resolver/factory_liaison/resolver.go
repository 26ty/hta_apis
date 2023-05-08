package factory_liaison

import (
	"eirc.app/internal/v1/service/factory_liaison"
	model "eirc.app/internal/v1/structure/factory_liaisons"
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
	FactoryLiaisonService factory_liaison.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		FactoryLiaisonService: factory_liaison.New(db),
	}
}
