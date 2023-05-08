package antivirus_software

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/antivirus_software"
	model "eirc.app/internal/v1/structure/antivirus_software"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByCaseIDtoTop(input []Model.GetCaseListOutput) interface{} 
	GetByPIDList(input *model.Fields) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService account.Service
	CompanyService company.Service
	FileService file.Service
	AntivirusSoftwareService antivirus_software.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService: account.New(db),
		CompanyService: company.New(db),
		FileService: file.New(db),
		AntivirusSoftwareService: antivirus_software.New(db),
	}
}
