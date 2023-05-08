package labor_hour_modify

import (
	"eirc.app/internal/v1/entity/labor_hour_modify"
	model "eirc.app/internal/v1/structure/labor_hour_modify"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByUserIdLaborHourModifyList(input *model.Field) (quantity int64, output []*model.Base, err error)
	GetByCuIdLaborHourModifyList(input *model.Field) (quantity int64, output []*model.Base, err error) 
	GetByUserIdList(input *model.Field) (quantity int64, output []*model.LaborHourModify, err error) 
	GetByCaseID(input string) (output *model.ReviewByDepartment, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	UpdatedStatus(input *model.Updated_Review) (err error)
}

type service struct {
	Entity labor_hour_modify.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: labor_hour_modify.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
