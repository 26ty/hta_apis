package countersign_user

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/countersign_user"
	model "eirc.app/internal/v1/structure/countersign_user"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByIDCountersignUserListUser(input *model.Documents) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService account.Service
	CompanyService company.Service
	FileService file.Service
	CountersignUserService countersign_user.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService: account.New(db),
		CompanyService: company.New(db),
		FileService: file.New(db),
		CountersignUserService: countersign_user.New(db),
	}
}
