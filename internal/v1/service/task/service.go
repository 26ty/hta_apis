package task

import (
	"eirc.app/internal/v1/entity/task"
	model "eirc.app/internal/v1/structure/task"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	TaskListUser(input *model.Users) (quantity int64, output []*model.Task_Account, err error)
	GetByIDListTaskHour(input *model.Users) (quantity int64, output []*model.Task_Account_Labor_Hour, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByIDTaskBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error)
	GetByTaskListUser(input *model.Users) (quantity int64, output []*model.Task_Account, err error)
	GetByOriginIDAndUserID(input *model.Users) (quantity int64, output []*model.Task_OriginId, err error)
	GetTaskListHourByUserID(input *model.Field) (quantity int64, output []*model.Task_Hour_User, err error)
	GetByTaskListHourDocumentsAndUserID(input *model.Field) (quantity int64, output []*model.Task_Hour_User, err error)
	GetByTIDTaskListUser(input *model.Fields) (quantity int64, output []*model.Task_User_Account, err error)
	GetByDocumentIDTaskListLast(input *model.Fields) (quantity int64, output []*model.Task_Template_Last, err error)
	GetByDocumentIDTaskList(input *model.Field) (quantity int64, output []*model.Task_Template, err error) 
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity task.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: task.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
