package labor_hour

import (
	"eirc.app/internal/v1/entity/labor_hour"
	model "eirc.app/internal/v1/structure/labor_hour"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByUserIdLaborHourList(input *model.Field) (quantity int64, output []*model.Base, err error)
	GetByCuIdLaborHourList(input *model.Field) (quantity int64, output []*model.Base, err error) 
	GetByUserIdList(input *model.Field) (quantity int64, output []*model.LaborHour, err error)
	GetByUserIdCategoryList(input *model.Field) (quantity int64, output []*model.GetUserCategoryLabor, err error) 
	GetByUserIdCategory(input *model.Field) (quantity int64, output []*model.GetUserCategoryLabor, err error)
	GetByUserIdMonthList(input *model.Field_Month) (quantity int64, output []*model.GetUserAllLabor, err error) 
	GetByUserIdOneLaborhour(input *model.Field_Month) (output []*model.GetUserOneLabor, err error) 
	GetByUserIdOneSumLaborhour(input *model.Field_Month) (output *model.GetUserOneSumLabor, err error) 
	GetByUserIdMonthSumList(input *model.Field_Month) (output []*model.GetUserAllSumLabor, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity labor_hour.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: labor_hour.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
