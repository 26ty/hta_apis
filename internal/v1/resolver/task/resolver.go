package task

import (
	//"eirc.app/internal/v1/entity/project"
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/project"
	"eirc.app/internal/v1/service/task"
	model "eirc.app/internal/v1/structure/task"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created_List) interface{}
	List(input *model.Fields) interface{}
	TaskListUser(input *model.Users) interface{}
	GetByIDListTaskHour(input *model.Users) interface{}
	GetByTaskListUser(input *model.Users) interface{}
	GetByID(input *model.Field) interface{}
	GetByIDTaskBonitaUserList(input *model.Users) interface{}
	GetByOriginIDAndUserID(input *model.Users) interface{}
	GetTaskListHourByUserID(input *model.Field) interface{}
	GetByTaskListHourDocumentsAndUserID(input *model.Field) interface{}
	GetByTIDTaskListUser(input *model.Fields) interface{}
	GetByDocumentIDTaskListLast(input *model.Fields) interface{}
	GetByDocumentIDTaskList(input *model.Field) interface{}
	DeleteList(input *model.Updated_List) interface{}
	Delete(input *model.Updated) interface{}
	Updated(input *model.Updated_List) interface{}
}

type resolver struct {
	ProjectService project.Service
	AccountService account.Service
	CompanyService company.Service
	TaskService    task.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		ProjectService: project.New(db),
		AccountService: account.New(db),
		CompanyService: company.New(db),
		TaskService:    task.New(db),
	}
}
