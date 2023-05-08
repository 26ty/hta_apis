package labor_hour

import (
	model "eirc.app/internal/v1/structure/labor_hour"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByUserIdCategoryList(input *model.Field) (amount int64, output []*model.GetUserCategoryLabor, err error) 
	GetByUserIdCategory(input *model.Field) (amount int64, output []*model.GetUserCategoryLabor, err error) 
	GetByUserIdMonthList(input *model.Field_Month) (amount int64, output []*model.GetUserAllLabor, err error)
	GetByUserIdOneLaborhour(input *model.Field_Month) (output []*model.GetUserOneLabor, err error)
	GetByUserIdOneSumLaborhour(input *model.Field_Month) (output *model.GetUserOneSumLabor, err error)
	GetByUserIdMonthSumList(input *model.Field_Month) (output []*model.GetUserAllSumLabor, err error)
	GetByUserIdList(input *model.Field) (amount int64, output []*model.LaborHour, err error)
	GetByUserIdLaborHourList(input *model.Field) (amount int64, output []*model.Table, err error)
	GetByCuIdLaborHourList(input *model.Field) (amount int64, output []*model.Table, err error)
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
