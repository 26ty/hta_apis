package project

import (
	model "eirc.app/internal/v1/structure/project"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByProjectBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error)
	ProjectListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error)
	ProduceSalesListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) 
	ProjectTemplateListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) 
	ProjectAuthorizationListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) 
	GetByProjectListUser(input *model.Field) (output *model.Project_Account, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByCaseID(input string) (output *model.Project_Account, err error)
	GetByCaseIDBonitaUserID(caseId string,userId string) (output []*model.Tm_Return, err error)
	GetByCaseIDTaskUserParentcaseID(caseId string,bonita_parentcase_id string) (output *model.Tm_Return, err error) 
	GetByCaseIDTaskUserStatus2(caseId string,status_type_id string,status_type_id2 string) (output []*model.Tm_Return, err error)
	GetByCaseIDTaskUserDepartment(caseId string,status_type_id string,dep string) (output []*model.Tm_Return, err error)
	Deleted(input *model.Table) (err error)
	Updated(input *model.Table) (err error)
}

type entity struct {
	db *gorm.DB
}

func New(db *gorm.DB) Entity {
	return &entity{
		db: db,
	}
}

func (e *entity) WithTrx(tx *gorm.DB) Entity {
	return &entity{
		db: tx,
	}
}
