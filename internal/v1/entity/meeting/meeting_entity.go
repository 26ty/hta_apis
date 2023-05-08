package meeting

import (
	model "eirc.app/internal/v1/structure/meeting"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.DocumentID != nil {
	// 	db.Where("document_id = ?", input.DocumentID)
	// }

	if input.MName != nil {
		db.Where("name like %?%", *input.MName)
	}

	if input.Room != nil {
		db.Where("room = ?", input.Room)
	}

	if input.TimeForStart != nil {
		db.Where("time_for_start = ?", input.TimeForStart)
	}

	if input.TimeForEnd != nil {
		db.Where("time_for_end = ?", input.TimeForEnd)
	}

	if input.DateForStart != nil {
		db.Where("date_for_start = ?", input.DateForStart)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

//判斷該會議的餐與者
func (e *entity) GetByMIDMeetingListUser(input *model.Fields) (amount int64, output []*model.MeetingListUser, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page-1)*input.Limit)).
		Limit(int(input.Limit)).
		Select("meeting.m_id,meeting.m_name,accounts.account_id,accounts.name,attendee.chairman,attendee.receive_email").
		Joins("left join attendee on attendee.meet_id = meeting.m_id").
		Joins("left join accounts on accounts.account_id = attendee.user_id").
		Where("attendee.chairman=false").
		Where("m_id = ?", input.MID).
		Order("meeting.create_time desc").Find(&output).Error

	return amount, output, err
}


func (e *entity) GetByUserIDMeetingListUser(input *model.Users) (amount int64, output []*model.MeetingListUser, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("attendee.a_id,accounts.email,meeting.m_id,meeting.m_name,meeting.room,meeting.documents_id,meeting.time_for_start,meeting.time_for_end,meeting.date_for_start,accounts.account_id,attendee.receive_email,accounts.name,meeting.create_time").
		Joins("left join attendee on attendee.meet_id = meeting.m_id").
		Joins("left join accounts on accounts.account_id = attendee.user_id").
		Where("attendee.user_id=?",input.UserID).
		Order("create_time desc").Count(&amount).Offset(int((input.Page-1)*input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

//判斷該單據下的會議
func (e *entity) GetByDIDMeetingListUser(input *model.Fields) (amount int64, output []*model.MeetingListUser, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("meeting.m_id,meeting.m_name,meeting.room,meeting.documents_id,meeting.time_for_start,meeting.time_for_end,meeting.date_for_start,accounts.account_id,attendee.receive_email,accounts.name,meeting.create_time").
		Joins("left join attendee on attendee.meet_id = meeting.m_id").
		Joins("left join accounts on accounts.account_id = attendee.user_id").
		Where("attendee.chairman=true").
		Where("documents_id = ?", input.DocumentsID).
		Order("create_time desc").Count(&amount).Offset(int((input.Page-1)*input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) MeetingUser(input *model.Fields) (amount int64, output []*model.MeetingListUser, err error) {
	db := e.db.Model(&model.Table{})

	err = db.
		Select("attendee.a_id,accounts.email,meeting.m_id,meeting.m_name,meeting.room,meeting.documents_id,meeting.time_for_start,meeting.time_for_end,meeting.date_for_start,accounts.account_id,attendee.receive_email,accounts.name,meeting.create_time").
		Joins("left join attendee on attendee.meet_id = meeting.m_id").
		Joins("left join accounts on accounts.account_id = attendee.user_id").
		Where("attendee.chairman=true").
		Order("meeting.create_time desc").
		Count(&amount).Offset(int((input.Page-1)*input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByMIDMeetingUser(input *model.Field) (output *model.MeetingListUser, err error) {
	db := e.db.Model(&model.Table{})

	err = db.
		Select("attendee.a_id,accounts.email,meeting.m_id,meeting.m_name,meeting.room,meeting.documents_id,meeting.time_for_start,meeting.time_for_end,meeting.date_for_start,accounts.account_id,accounts.name,attendee.receive_email,meeting.create_time").
		Joins("left join attendee on attendee.meet_id = meeting.m_id").
		Joins("left join accounts on accounts.account_id = attendee.user_id").
		Where("attendee.chairman=true").
		Where("m_id = ?", input.MID).
		Order("meeting.create_time desc").First(&output).Error

	return output, err
}

func (e *entity) GetByMIDMeetingUserParticipant(input *model.Field) (output []*model.MeetingParticipant, err error) {
	db := e.db.Model(&model.Table{})

	err = db.
		Select("attendee.a_id,accounts.email,attendee.receive_email,accounts.account_id as participant_id,accounts.name as participant_name").
		Joins("left join attendee on attendee.meet_id = meeting.m_id").
		Joins("left join accounts on accounts.account_id = attendee.user_id").
		Where("attendee.chairman=false").
		Where("m_id = ?", input.MID).
		Order("meeting.create_time desc").Find(&output).Error

	return output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("m_id = ?", input.MID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("m_id = ?", input.MID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("m_id = ?", input.MID).Save(&input).Error

	return err
}
