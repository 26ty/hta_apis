package customer_demand

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/customer_demand"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/task_user"
	"eirc.app/internal/v1/service/countersign_user"
	"eirc.app/internal/v1/service/personnel_affiliation"

	model "eirc.app/internal/v1/structure/customer_demand"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	CustomerDemandListUser(input *model.Field) interface{}
	GetByCuIDCustomerDemandListUser(input *model.Field) interface{}
	GetByUserIDListCR(input *model.Users) interface{}
	GetByUserIDListHCR(input *model.Users) interface{} 
	GetByID(input *model.Field) interface{}
	GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoDirector(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoTop(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoDispatch(input []Model.GetCaseListOutput) interface{} 
	GetByCaseIDtoEvaluation(input []Model.GetCaseListOutput,userId string) interface{}
	GetByCaseIDtoCountersign(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoPMEvaluation(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoBusiness(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoBusinessManager(input []Model.GetCaseListOutput) interface{} 
	GetByCaseIDtoBusinessDirector(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoTaskFinish(input []Model.GetCaseListOutput,userId string) interface{} 
	GetByCaseIDtoTaskFinishManager(input []Model.GetCaseListOutput) interface{} 
	GetByCaseIDtoBusinessClose(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoDepartmentClose(input []Model.GetCaseListOutput) interface{} 
	GetByCaseIDtoDirectorClose(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoTopClose(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoProductionClose(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoCountersignClose(input []Model.GetCaseListOutput) interface{}
	GetByCaseIDtoPMClose(input []Model.GetCaseListOutput) interface{} 
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	Updated_Bonita(input *model.Updated_Bonita) interface{}
}

type resolver struct {
	AccountService            account.Service
	CompanyService            company.Service
	FileService               file.Service
	TaskUserService               task_user.Service
	CountersignUserService               countersign_user.Service
	CustomerDemandService customer_demand.Service
	PersonnelAffiliationService personnel_affiliation.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:            account.New(db),
		CompanyService:            company.New(db),
		FileService:               file.New(db),
		TaskUserService:               task_user.New(db),
		CountersignUserService:               countersign_user.New(db),
		CustomerDemandService: customer_demand.New(db),
		PersonnelAffiliationService: personnel_affiliation.New(db),
	}
}
