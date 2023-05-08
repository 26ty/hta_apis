package project_template

import (
	"eirc.app/internal/v1/entity/project_template"
	model "eirc.app/internal/v1/structure/project_template"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity project_template.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: project_template.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
