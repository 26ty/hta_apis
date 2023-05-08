package manufacture_order

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/manufacture_order"
	"eirc.app/internal/v1/service/manufacture_user"
	model "eirc.app/internal/v1/structure/manufacture_order"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	// ManufactureOrderCdListUser(input *model.Fields) interface{}
	ManufactureOrderProjectListUser(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByIDOne(input *model.Field) interface{}
	GetByPIDList(input *model.Fields) interface{}
	GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoManufacture(input []Model.GetCaseListOutput) interface{} 
	GetByCaseIDtoTop(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoConfirm(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoSave(input []Model.GetCaseListOutput) interface{} 
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	Updated_Bonita(input *model.Updated_Bonita) interface{}
}

type resolver struct {
	AccountService          account.Service
	CompanyService          company.Service
	FileService             file.Service
	ManufactureOrderService manufacture_order.Service
	ManufactureUserService manufacture_user.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:          account.New(db),
		CompanyService:          company.New(db),
		FileService:             file.New(db),
		ManufactureOrderService: manufacture_order.New(db),
		ManufactureUserService: manufacture_user.New(db),
	}
}
