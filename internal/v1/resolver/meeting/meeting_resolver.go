package meeting

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	meetingModel "eirc.app/internal/v1/structure/meeting"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *meetingModel.Created) interface{} {
	defer trx.Rollback()

	meeting, err := r.MeetingService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, meeting.MID)
}

func (r *resolver) List(input *meetingModel.Fields) interface{} {
	output := &meetingModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, meeting, err := r.MeetingService.List(input)
	output.Total = quantity
	meetingByte, err := json.Marshal(meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(meetingByte, &output.Meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIDMeetingListUser(input *meetingModel.Users) interface{} {
	output := &meetingModel.AllMeetingListUserParticipant{}
	output.Limit = input.Limit
	output.Page = input.Page
	//output.Total = input.Total
	quantity, meeting, err := r.MeetingService.GetByUserIDMeetingListUser(input)
	output.Total = quantity
	meetingByte, err := json.Marshal(meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(meetingByte, &output.Meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	
	for _,value := range output.Meeting{
		One_input := &meetingModel.Field{}
		One_input.MID = value.MID

		meetingParticipant, err := r.MeetingService.GetByMIDMeetingUserParticipant(One_input)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		meetingParticipantByte, _ := json.Marshal(meetingParticipant)
		err = json.Unmarshal(meetingParticipantByte, &value.Participant)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	}
	return code.GetCodeMessage(code.Successful, output)

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByMIDMeetingListUser(input *meetingModel.Fields) interface{} {
	output := &meetingModel.MeetingListUsers{}
	output.Limit = input.Limit
	output.Page = input.Page
	//output.Total = input.Total
	quantity, meeting, err := r.MeetingService.GetByMIDMeetingListUser(input)
	output.Total = quantity
	meetingByte, err := json.Marshal(meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(meetingByte, &output.Meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDIDMeetingListUser(input *meetingModel.Fields) interface{} {
	output := &meetingModel.MeetingListUsers{}
	output.Limit = input.Limit
	output.Page = input.Page
	//output.Total = input.Total
	quantity, meeting, err := r.MeetingService.GetByDIDMeetingListUser(input)
	output.Total = quantity
	meetingByte, err := json.Marshal(meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(meetingByte, &output.Meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) MeetingUser(input *meetingModel.Fields) interface{} {
	quantity, meeting, err := r.MeetingService.MeetingUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &meetingModel.AllMeetingListUserParticipant{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity

	meetingByte, _ := json.Marshal(meeting)
	err = json.Unmarshal(meetingByte, &output.Meeting)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	for _,value := range output.Meeting{
		One_input := &meetingModel.Field{}
		One_input.MID = value.MID

		meetingParticipant, err := r.MeetingService.GetByMIDMeetingUserParticipant(One_input)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		meetingParticipantByte, _ := json.Marshal(meetingParticipant)
		err = json.Unmarshal(meetingParticipantByte, &value.Participant)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	}
	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByMIDMeetingUser(input *meetingModel.Field) interface{} {
	meeting, err := r.MeetingService.GetByMIDMeetingUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &meetingModel.MeetingListUserParticipant{}
	meetingByte, _ := json.Marshal(meeting)
	err = json.Unmarshal(meetingByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	meetingParticipant, err := r.MeetingService.GetByMIDMeetingUserParticipant(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	meetingParticipantByte, _ := json.Marshal(meetingParticipant)
	err = json.Unmarshal(meetingParticipantByte, &output.Participant)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *meetingModel.Field) interface{} {
	meeting, err := r.MeetingService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &meetingModel.Single{}
	meetingByte, _ := json.Marshal(meeting)
	err = json.Unmarshal(meetingByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *meetingModel.Updated) interface{} {
	_, err := r.MeetingService.GetByID(&meetingModel.Field{MID: input.MID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.MeetingService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *meetingModel.Updated) interface{} {
	meeting, err := r.MeetingService.GetByID(&meetingModel.Field{MID: input.MID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.MeetingService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, meeting.MID)
}
