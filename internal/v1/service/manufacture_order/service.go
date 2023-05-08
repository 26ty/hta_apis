package manufacture_order

import (
	"eirc.app/internal/v1/entity/manufacture_order"
	model "eirc.app/internal/v1/structure/manufacture_order"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	// ManufactureOrderCdListUser(input *model.Fields) (quantity int64, output []*model.ManufactureOrder_Cd_Account, err error)
	ManufactureOrderProjectListUser(input *model.Fields) (quantity int64, output []*model.ManufactureOrder_Project_Account, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByIDOne(input *model.Field) (output *model.One, err error) 
	GetByPIDList(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByCaseID(input string) (output *model.Review, err error) 
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	Updated_Bonita(input *model.Updated_Bonita) (err error)
}

type service struct {
	Entity manufacture_order.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: manufacture_order.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
