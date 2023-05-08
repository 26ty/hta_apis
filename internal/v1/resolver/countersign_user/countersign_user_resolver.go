package countersign_user

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	countersign_userModel "eirc.app/internal/v1/structure/countersign_user"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *countersign_userModel.Created) interface{} {
	defer trx.Rollback()

	countersign_user, err := r.CountersignUserService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, countersign_user.CuID)
}

func (r *resolver) List(input *countersign_userModel.Fields) interface{} {
	output := &countersign_userModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, countersign_user, err := r.CountersignUserService.List(input)
	output.Total = quantity
	countersign_userByte, err := json.Marshal(countersign_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(countersign_userByte, &output.CountersignUser)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByIDCountersignUserListUser(input *countersign_userModel.Documents) interface{} {
	output := &countersign_userModel.CountersignUser_Accounts{}
	quantity, countersign_user, err := r.CountersignUserService.GetByIDCountersignUserListUser(input)
	output.Total = quantity
	countersign_userByte, err := json.Marshal(countersign_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(countersign_userByte, &output.CountersignUser)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *countersign_userModel.Field) interface{} {
	countersign_user, err := r.CountersignUserService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &countersign_userModel.Single{}
	countersign_userByte, _ := json.Marshal(countersign_user)
	err = json.Unmarshal(countersign_userByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *countersign_userModel.Updated) interface{} {
	_, err := r.CountersignUserService.GetByID(&countersign_userModel.Field{CuID: input.CuID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CountersignUserService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *countersign_userModel.Updated) interface{} {
	countersign_user, err := r.CountersignUserService.GetByID(&countersign_userModel.Field{CuID: input.CuID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CountersignUserService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, countersign_user.CuID)
}
