package task_user

import (
	"eirc.app/internal/v1/service/task_user"
	model "eirc.app/internal/v1/structure/task_user"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created_List) interface{}
	List(input *model.Fields) interface{}
	GetByDocumnetIDListHour(input *model.Field) interface{}
	GetByID(input *model.Field) interface{}
	GetName(input *model.Field) interface{}
	Delete(input *model.Updated) interface{}
	DeleteList(input *model.Updated_List) interface{}
	Updated(input *model.Updated_List) interface{}
	UpdatedStatus(input *model.Updated_Review) interface{}
}

type resolver struct {
	TaskUserService task_user.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		TaskUserService: task_user.New(db),
	}
}
