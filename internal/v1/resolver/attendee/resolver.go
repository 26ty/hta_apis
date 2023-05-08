package attendee

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/attendee"
	"eirc.app/internal/v1/service/company"
	model "eirc.app/internal/v1/structure/attendee"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated_List) interface{}
}

type resolver struct {
	AttendeeService attendee.Service
	AccountService  account.Service
	CompanyService  company.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AttendeeService: attendee.New(db),
		AccountService:  account.New(db),
		CompanyService:  company.New(db),
	}
}
