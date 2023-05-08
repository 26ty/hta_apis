package account

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/department"
	"eirc.app/internal/v1/service/jobtitle"
	"eirc.app/internal/v1/service/personnel_affiliation"

	model "eirc.app/internal/v1/structure/accounts"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Users) interface{}
	AccountNameList(input *model.Users) interface{}
	AccountNameDepartmentList(input *model.Users) interface{} 
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	UpdatedCsv(input *model.UpdatedCsvList) interface{} 
}

type resolver struct {
	Account account.Service
	Company company.Service
	Department department.Service
	Jobtitle jobtitle.Service
	PersonnelAffiliation personnel_affiliation.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		Account: account.New(db),
		Company: company.New(db),
		Department: department.New(db),
		Jobtitle: jobtitle.New(db),
		PersonnelAffiliation: personnel_affiliation.New(db),
	}
}
