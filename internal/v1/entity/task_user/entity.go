package task_user

import (
	model "eirc.app/internal/v1/structure/task_user"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Task_user_Account, err error)
	GetByDocumnetIDListHour(input *model.Field) (amount int64, output []*model.Task_user_Labor_Hour, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetName(input *model.Field) (output *model.Task_user_Account, err error)
	Deleted(input *model.Table) (err error)
	Updated_Bonita(input *model.Updated_Bonita) (err error) 
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
