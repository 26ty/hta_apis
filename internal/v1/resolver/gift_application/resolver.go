package gift_application

import (
	"eirc.app/internal/v1/service/gift_application"
	Model "eirc.app/internal/v1/structure"
	model "eirc.app/internal/v1/structure/gift_applications"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	Updated_Bonita(input *model.Updated_Bonita) interface{}
	GetByGIDGiftDetailListUser(input *model.Field) interface{}
	GiftDetailListUser(input *model.Fields) interface{}
	GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoViceTop(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoTop(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoAttm(input []Model.GetCaseListOutput) interface{}
}

type resolver struct {
	GiftApplicationService gift_application.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		GiftApplicationService: gift_application.New(db),
	}
}
