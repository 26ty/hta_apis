package labor_hour_modify

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/labor_hour_modify"
	"eirc.app/internal/v1/service/labor_hour"
	model "eirc.app/internal/v1/structure/labor_hour_modify"
	Model "eirc.app/internal/v1/structure"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{}
	GetByUserIdLaborHourModifyList(input *model.Field) interface{}
	GetByCuIdLaborHourModifyList(input *model.Field) interface{}
	GetByUserIdList(input *model.Field) interface{} 
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	UpdatedStatus(input *model.Updated_Review) interface{}
	Replace(input *model.Field) interface{}
}

type resolver struct {
	AccountService   account.Service
	CompanyService   company.Service
	FileService      file.Service
	LaborHourModifyService labor_hour_modify.Service
	LaborHourService labor_hour.Service

}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:   account.New(db),
		CompanyService:   company.New(db),
		FileService:      file.New(db),
		LaborHourModifyService: labor_hour_modify.New(db),
		LaborHourService: labor_hour.New(db),
	}
}
