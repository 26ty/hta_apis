package customer

import (
	"eirc.app/internal/v1/entity/customer"
	model "eirc.app/internal/v1/structure/customers"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	GetByCIDAccount(input *model.Field) (output *model.Customer_Account, err error)
	AccountList(input *model.Fields) (quantity int64, output []*model.Customer_Account, err error)
}

type service struct {
	Entity customer.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: customer.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
