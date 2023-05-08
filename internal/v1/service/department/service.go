package department

import (
	"eirc.app/internal/v1/entity/department"
	model "eirc.app/internal/v1/structure/department"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	A1Department(input *model.Field) (output []*model.Table, err error)
	AllDepartment(input *model.Field) (output []*model.Table, err error)  
	DepartmentAccountList(input *model.Users) (quantity int64, output []*model.Deparment_Account, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity department.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: department.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
