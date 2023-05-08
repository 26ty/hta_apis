package department

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/personnel_affiliation"
	"eirc.app/internal/v1/service/department"
	model "eirc.app/internal/v1/structure/department"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	A1Department(input *model.Field) interface{}
	AllDepartment(input *model.Field) interface{}
	DepartmentAccountList(input *model.Users) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService    account.Service
	PersonnelAffiliationService personnel_affiliation.Service
	DepartmentService department.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:    account.New(db),
		PersonnelAffiliationService: personnel_affiliation.New(db),
		DepartmentService: department.New(db),
	}
}
