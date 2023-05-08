package meeting

import (
	model "eirc.app/internal/v1/structure/meeting"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByUserIDMeetingListUser(input *model.Users) (amount int64, output []*model.MeetingListUser, err error) 
	GetByMIDMeetingListUser(input *model.Fields) (amount int64, output []*model.MeetingListUser, err error)
	GetByDIDMeetingListUser(input *model.Fields) (amount int64, output []*model.MeetingListUser, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByMIDMeetingUser(input *model.Field) (output *model.MeetingListUser, err error)
	GetByMIDMeetingUserParticipant(input *model.Field) (output []*model.MeetingParticipant, err error)
	MeetingUser(input *model.Fields) (amount int64, output []*model.MeetingListUser, err error)
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
