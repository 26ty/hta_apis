package department

import (
	model "eirc.app/internal/v1/structure/department"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	A1Department(input *model.Field) (output []*model.Table, err error) 
	AllDepartment(input *model.Field) (output []*model.Table, err error)
	DepartmentAccountList(input *model.Users) (amount int64, output []*model.Deparment_Account, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	Deleted(input *model.Table) (err error)
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
