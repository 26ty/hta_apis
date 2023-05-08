package jig_demand

import (
	"eirc.app/internal/v1/service/jig_demand"
	model "eirc.app/internal/v1/structure/jig_demands"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	Updated_Bonita(input *model.Updated_Bonita) interface{}
	GetByJIDJigDetailListUser(input *model.Field) interface{}
	JigDetailListUser(input *model.Fields) interface{}
	SearchJigDemand(input *model.Fields) interface{}
	GetByUserIDListJD(input *model.Users) interface{}
}

type resolver struct {
	JigDemandService jig_demand.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		JigDemandService: jig_demand.New(db),
	}
}
