package gift_application_detail

import (
	"eirc.app/internal/v1/service/gift_application_detail"
	model "eirc.app/internal/v1/structure/gift_application_details"
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
	GiftApplicationDetailService gift_application_detail.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		GiftApplicationDetailService: gift_application_detail.New(db),
	}
}
