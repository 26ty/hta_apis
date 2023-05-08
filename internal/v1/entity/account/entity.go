package account

import (
	model "eirc.app/internal/v1/structure/accounts"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Users) (amount int64, output []*model.Table, err error)
	AccountNameList(input *model.Users) (amount int64, output []*model.Account_Name, err error)
	AccountNameDepartmentList(input *model.Users) (amount int64, output []*model.Account_Name, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByAccount(input *model.Field) (output *model.Table, err error) 
	GetByBonitaUserID(userID string) (output *model.Table, err error)
	Deleted(input *model.Field) (err error)
	Updated(input *model.Table) (err error)
	UpdatedCsv(input *model.Table) (err error) 
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