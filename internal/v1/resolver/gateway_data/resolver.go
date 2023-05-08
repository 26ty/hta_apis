package gateway_data

import (
	"eirc.app/internal/v1/service/gateway_data"
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/personnel_affiliation"

	model "eirc.app/internal/v1/structure/gateway_data"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByClassificationTitle(input *model.Field) interface{}
	GetByDataDemand(input []Model.GetCaseListOutput,GdID string,userId string) interface{} 
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	GatewayDataService gateway_data.Service
	AccountService          account.Service
	PersonnelAffiliationService personnel_affiliation.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		GatewayDataService: gateway_data.New(db),
		AccountService:          account.New(db),
		PersonnelAffiliationService: personnel_affiliation.New(db),
	}
}
