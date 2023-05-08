package project

import (
	//"eirc.app/internal/v1/entity/project"
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/task_user"
	"eirc.app/internal/v1/service/project"
	model "eirc.app/internal/v1/structure/project"
	p_model "eirc.app/internal/v1/structure"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	GetByProjectBonitaUserList(input *model.Users) interface{}
	List(input *model.Fields) interface{}
	ProjectListUser(input *model.Fields) interface{}
	ProduceSalesListUser(input *model.Fields) interface{}
	ProjectTemplateListUser(input *model.Fields) interface{} 
	ProjectAuthorizationListUser(input *model.Fields) interface{}
	GetByProjectListUser(input *model.Field) interface{}
	GetByID(input *model.Field) interface{}
	GetByCaseIDtoPM(input []p_model.GetCaseListOutput) interface{}
	GetBonitaCaseListPMCompleted(input []p_model.GetCaseListOutput) interface{}
	GetB2BonitaCaseListTM(input []p_model.GetCaseListOutput,userId string) interface{}
	GetB2BonitaCaseListCountersign(input []p_model.GetCaseListOutput) interface{}
	GetB2BonitaCaseListConfirm(input []p_model.GetCaseListOutput) interface{}
	GetB2BonitaCaseListDepartment(input []p_model.GetCaseListOutput,userId string) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
	Updated_Bonita(input *model.Updated_Bonita) interface{}
}

type resolver struct {
	ProjectService project.Service
	AccountService account.Service
	TaskUserService task_user.Service
	CompanyService company.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		ProjectService: project.New(db),
		AccountService: account.New(db),
		TaskUserService: task_user.New(db),
		CompanyService: company.New(db),
	}
}
