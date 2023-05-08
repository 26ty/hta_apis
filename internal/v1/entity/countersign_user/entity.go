package countersign_user

import (
	model "eirc.app/internal/v1/structure/countersign_user"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByIDCountersignUserListUser(input *model.Documents) (amount int64, output []*model.CountersignUser_Account, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	Deleted(input *model.Table) (err error)
	Updated(input *model.Table) (err error)
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
