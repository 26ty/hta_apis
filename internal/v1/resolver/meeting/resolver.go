package meeting

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/company"
	"eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/meeting"
	model "eirc.app/internal/v1/structure/meeting"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByUserIDMeetingListUser(input *model.Users) interface{} 
	GetByMIDMeetingListUser(input *model.Fields) interface{}
	GetByDIDMeetingListUser(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetByMIDMeetingUser(input *model.Field) interface{}
	MeetingUser(input *model.Fields) interface{} 
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	AccountService account.Service
	CompanyService company.Service
	FileService    file.Service
	MeetingService meeting.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		AccountService: account.New(db),
		CompanyService: company.New(db),
		FileService:    file.New(db),
		MeetingService: meeting.New(db),
	}
}
