package jig_demand_detail

import (
	"eirc.app/internal/v1/service/jig_demand_detail"
	model "eirc.app/internal/v1/structure/jig_demand_details"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	UpdatedByJigID(input *model.Updated) interface{}
}

type resolver struct {
	JigDemandDetailService jig_demand_detail.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		JigDemandDetailService: jig_demand_detail.New(db),
	}
}
