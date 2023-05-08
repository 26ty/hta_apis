package antivirus_software

import (
	"eirc.app/internal/v1/entity/antivirus_software"
	model "eirc.app/internal/v1/structure/antivirus_software"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByCaseID(input string) (output *model.Review, err error) 
	GetByPIDList(input *model.Fields) (quantity int64, output []*model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity antivirus_software.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: antivirus_software.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
