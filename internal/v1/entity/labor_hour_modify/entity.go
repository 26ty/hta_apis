package labor_hour_modify

import (
	model "eirc.app/internal/v1/structure/labor_hour_modify"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByUserIdList(input *model.Field) (amount int64, output []*model.LaborHourModify, err error)
	GetByUserIdLaborHourModifyList(input *model.Field) (amount int64, output []*model.Table, err error)
	GetByCuIdLaborHourModifyList(input *model.Field) (amount int64, output []*model.Table, err error)
	GetByCaseID(input string) (output *model.ReviewByDepartment, err error)
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
