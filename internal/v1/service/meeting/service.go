package meeting

import (
	"eirc.app/internal/v1/entity/meeting"
	model "eirc.app/internal/v1/structure/meeting"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByUserIDMeetingListUser(input *model.Users) (quantity int64, output []*model.MeetingListUser, err error) 
	GetByMIDMeetingListUser(input *model.Fields) (quantity int64, output []*model.MeetingListUser, err error)
	GetByDIDMeetingListUser(input *model.Fields) (quantity int64, output []*model.MeetingListUser, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByMIDMeetingUser(input *model.Field) (output *model.MeetingListUser, err error)
	GetByMIDMeetingUserParticipant(input *model.Field) (output []*model.MeetingParticipant, err error)
	MeetingUser(input *model.Fields) (quantity int64, output []*model.MeetingListUser, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity meeting.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: meeting.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
