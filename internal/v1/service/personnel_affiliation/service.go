package personnel_affiliation

import (
	"eirc.app/internal/v1/entity/personnel_affiliation"
	model "eirc.app/internal/v1/structure/personnel_affiliation"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByUserID(input *model.Field) (output []*model.Affiliation_Account, err error) 
	GetByDepartmentID(input *model.Field) (output []*model.Deparment_User, err error)
	GetByParentDepartmentID(bonita_group_id string) (output []*model.Deparment_User, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity personnel_affiliation.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: personnel_affiliation.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
