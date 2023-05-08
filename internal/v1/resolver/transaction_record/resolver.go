package transaction_record

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/transaction_record"
	model "eirc.app/internal/v1/structure/transaction_record"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByDocumentIDUserList(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService           account.Service
	CompanyService           company.Service
	TransactionRecordService transaction_record.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService:           account.New(db),
		CompanyService:           company.New(db),
		TransactionRecordService: transaction_record.New(db),
	}
}
