package task_user

import (
	"eirc.app/internal/v1/entity/task_user"
	model "eirc.app/internal/v1/structure/task_user"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Task_user_Account, err error)
	GetByDocumnetIDListHour(input *model.Field) (quantity int64, output []*model.Task_user_Labor_Hour, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetName(input *model.Field) (output *model.Task_user_Account, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	UpdatedStatus(input *model.Updated_Review) (err error)
	Updated_Bonita(ID string,ParentCaseID string) (err error)
}

type service struct {
	Entity task_user.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: task_user.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
