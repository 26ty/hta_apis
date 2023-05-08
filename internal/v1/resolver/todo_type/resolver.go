package todo_type

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/todo_type"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/department"
	model "eirc.app/internal/v1/structure/todo_type"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByUserID(input *model.Field) interface{}
	Delete(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	Account account.Service
	Company company.Service
	Department department.Service
	Todo_typeService todo_type.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		Account: account.New(db),
		Company: company.New(db),
		Department: department.New(db),
		Todo_typeService: todo_type.New(db),
	}
}
