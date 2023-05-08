package personnel_affiliation

import (
	model "eirc.app/internal/v1/structure/personnel_affiliation"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByUserID(input *model.Field) (output []*model.Affiliation_Account, err error) 
	GetByDepartmentID(input *model.Field) (output []*model.Deparment_User, err error)
	GetByParentDepartmentID(bonita_group_id string) (output []*model.Deparment_User, err error)
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
