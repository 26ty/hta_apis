package personnel_affiliation

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/personnel_affiliation"
	"eirc.app/internal/v1/service/department"
	"eirc.app/internal/v1/service/jobtitle"

	model "eirc.app/internal/v1/structure/personnel_affiliation"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByUserID(input *model.Field) interface{}
	GetByDepartmentID(input *model.Field) interface{}
	GetByParentDepartmentID(input *model.Field) interface{} 
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService    account.Service
	CompanyService    company.Service
	PersonnelAffiliationService personnel_affiliation.Service
	DepartmentService department.Service
	JobtitleService jobtitle.Service

}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:    account.New(db),
		CompanyService:    company.New(db),
		PersonnelAffiliationService: personnel_affiliation.New(db),
		DepartmentService: department.New(db),
		JobtitleService: jobtitle.New(db),

	}
}
