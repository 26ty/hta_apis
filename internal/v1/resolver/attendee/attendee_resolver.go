package attendee

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	attendeeModel "eirc.app/internal/v1/structure/attendee"

	//accountModel "eirc.app/internal/v1/structure/accounts"
	//companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *attendeeModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱
	// _, err := r.CompanyService.WithTrx(trx).GetByID(&companyModel.Field{CompanyID: input.CompanyID,
	// 	IsDeleted: util.PointerBool(false)})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	attendee, err := r.AttendeeService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, attendee.AID)
}

func (r *resolver) List(input *attendeeModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &attendeeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, attendee, err := r.AttendeeService.List(input)
	output.Total = quantity
	attendeeByte, err := json.Marshal(attendee)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(attendeeByte, &output.Attendee)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *attendeeModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	attendee, err := r.AttendeeService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &attendeeModel.Single{}
	attendeeByte, _ := json.Marshal(attendee)
	err = json.Unmarshal(attendeeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *attendeeModel.Updated) interface{} {
	_, err := r.AttendeeService.GetByID(&attendeeModel.Field{AID: input.AID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.AttendeeService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *attendeeModel.Updated_List) interface{} {
	arr := []string{}
	for _, value := range input.Attendee {
		attendee, err := r.AttendeeService.GetByID(&attendeeModel.Field{AID: value.AID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		err = r.AttendeeService.Updated(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
		arr = append(arr, attendee.AID)
	}

	return code.GetCodeMessage(code.Successful, arr)
}
