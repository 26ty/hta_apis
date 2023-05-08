package project

import (
	"eirc.app/internal/v1/entity/project"
	model "eirc.app/internal/v1/structure/project"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByProjectBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error)
	ProjectListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error)
	ProduceSalesListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) 
	ProjectTemplateListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) 
	ProjectAuthorizationListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) 
	GetByProjectListUser(input *model.Field) (output *model.Project_Account, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByCaseID(input string) (output *model.Project_Account, err error)
	GetByCaseIDBonitaUserID(caseId string,userId string) (output []*model.Tm_Return, err error)
	GetByCaseIDTaskUserParentcaseID(caseId string,bonita_parentcase_id string) (output *model.Tm_Return, err error) 
	GetByCaseIDTaskUserStatus2(caseId string,status_type_id string,status_type_id2 string) (output []*model.Tm_Return, err error)
	GetByCaseIDTaskUserDepartment(caseId string,status_type_id string,dep string) (output []*model.Tm_Return, err error) 
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	Updated_Bonita(input *model.Updated_Bonita) (err error)
}

type service struct {
	Entity project.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: project.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
