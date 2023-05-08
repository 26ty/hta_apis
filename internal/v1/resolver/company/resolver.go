package company

import (
	"eirc.app/internal/v1/service/company"
	model "eirc.app/internal/v1/structure/companies"
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
	Company company.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		Company: company.New(db),
	}
}
