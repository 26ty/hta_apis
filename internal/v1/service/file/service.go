package file

import (
	"eirc.app/internal/v1/entity/file"
	model "eirc.app/internal/v1/structure/file"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByDocumentID(input *model.Field) (quantity int64, output []*model.FilebydocumentId, err error)
	GetByDocumentIDUserID(input *model.Users) (quantity int64, output []*model.FilebydocumentId, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity file.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: file.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
