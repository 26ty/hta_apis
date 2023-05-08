package machine_combined

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/machine_combined"
	model "eirc.app/internal/v1/structure/machine_combined"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	MachineCombinedListLast(input *model.Fields) interface{}
	GetByPIDMachineCombinedListLast(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService account.Service
	CompanyService company.Service
	FileService file.Service
	MachineCombinedService machine_combined.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService: account.New(db),
		CompanyService: company.New(db),
		FileService: file.New(db),
		MachineCombinedService: machine_combined.New(db),
	}
}
