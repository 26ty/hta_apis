package project_template

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/project_template"
	model "eirc.app/internal/v1/structure/project_template"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService         account.Service
	CompanyService         company.Service
	ProjectTemplateService project_template.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:         account.New(db),
		CompanyService:         company.New(db),
		ProjectTemplateService: project_template.New(db),
	}
}
