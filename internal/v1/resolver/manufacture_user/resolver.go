package manufacture_user

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/manufacture_user"
	"eirc.app/internal/v1/service/personnel_affiliation"

	model "eirc.app/internal/v1/structure/manufacture_user"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByManufactureID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService account.Service
	CompanyService company.Service
	FileService file.Service
	ManufactureUserService manufacture_user.Service
	PersonnelAffiliationService personnel_affiliation.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService: account.New(db),
		CompanyService: company.New(db),
		FileService: file.New(db),
		ManufactureUserService: manufacture_user.New(db),
		PersonnelAffiliationService: personnel_affiliation.New(db),
	}
}
