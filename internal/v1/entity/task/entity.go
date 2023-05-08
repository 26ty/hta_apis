package task

import (
	model "eirc.app/internal/v1/structure/task"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Create_Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	TaskListUser(input *model.Users) (amount int64, output []*model.Task_Account, err error)
	GetByIDListTaskHour(input *model.Users) (amount int64, output []*model.Task_Account_Labor_Hour, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByOriginIDAndUserID(input *model.Users) (amount int64, output []*model.Task_OriginId, err error)
	GetTaskListHourByUserID(input *model.Field) (amount int64, output []*model.Task_Hour_User, err error)
	GetByTaskListHourDocumentsAndUserID(input *model.Field) (amount int64, output []*model.Task_Hour_User, err error)
	GetByTaskListUser(input *model.Users) (amount int64, output []*model.Task_Account, err error)
	GetByTIDTaskListUser(input *model.Fields) (amount int64, output []*model.Task_User_Account, err error)
	GetByDocumentIDTaskListLast(input *model.Fields) (amount int64, output []*model.Task_Template_Last, err error)
	GetByDocumentIDTaskList(input *model.Field) (amount int64, output []*model.Task_Template, err error)
	GetByIDTaskBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error)
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
