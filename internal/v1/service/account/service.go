package account

import (
	"eirc.app/internal/v1/entity/account"
	model "eirc.app/internal/v1/structure/accounts"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Users) (quantity int64, output []*model.Base, err error)
	AccountNameList(input *model.Users) (quantity int64, output []*model.Account_Name, err error)
	AccountNameDepartmentList(input *model.Users) (quantity int64, output []*model.Account_Name, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByAccount(input *model.Field) (output *model.Base, err error)
	GetByBonitaUserID(userID string) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	UpdatedCsv(input *model.UpdatedCsv) (err error) 
	AcknowledgeAccount(input *model.Field) (acknowledge bool, output []*model.Base, err error)
}

type service struct {
	Entity account.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: account.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
