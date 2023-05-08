package customer

import (
	"eirc.app/internal/v1/service/customer"
	model "eirc.app/internal/v1/structure/customers"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	CustomerService customer.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		CustomerService: customer.New(db),
	}
}
