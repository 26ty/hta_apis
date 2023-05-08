package manufacture_order

import (
	model "eirc.app/internal/v1/structure/manufacture_order"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Create_Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	// ManufactureOrderCdListUser(input *model.Fields) (amount int64, output []*model.ManufactureOrder_Cd_Account, err error)
	ManufactureOrderProjectListUser(input *model.Fields) (amount int64, output []*model.ManufactureOrder_Project_Account, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByIDOne(input *model.Field) (output *model.One, err error) 
	GetByPIDList(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByCaseID(input string) (output *model.Review, err error) 
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
