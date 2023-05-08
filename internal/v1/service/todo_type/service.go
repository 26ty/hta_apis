package todo_type

import (
	"eirc.app/internal/v1/entity/todo_type"
	model "eirc.app/internal/v1/structure/todo_type"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByUserID(input *model.Field) (output []*model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity todo_type.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: todo_type.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
