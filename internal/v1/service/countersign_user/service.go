package countersign_user

import (
	"eirc.app/internal/v1/entity/countersign_user"
	model "eirc.app/internal/v1/structure/countersign_user"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByIDCountersignUserListUser(input *model.Documents) (quantity int64, output []*model.CountersignUser_Account, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	Updated_Bonita(ID string,ParentCaseID string) (err error)
}

type service struct {
	Entity countersign_user.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: countersign_user.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
