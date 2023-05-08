package manufacture_user

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	personnel_affiliationModel "eirc.app/internal/v1/structure/personnel_affiliation"
	manufacture_userModel "eirc.app/internal/v1/structure/manufacture_user"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *manufacture_userModel.Created) interface{} {
	defer trx.Rollback()

	manufacture_user, err := r.ManufactureUserService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, manufacture_user.MuID)
}

func (r *resolver) List(input *manufacture_userModel.Fields) interface{} {
	output := &manufacture_userModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, manufacture_user, err := r.ManufactureUserService.List(input)
	output.Total = quantity
	manufacture_userByte, err := json.Marshal(manufacture_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(manufacture_userByte, &output.ManufactureUser)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *manufacture_userModel.Field) interface{} {
	manufacture_user, err := r.ManufactureUserService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &manufacture_userModel.Single{}
	manufacture_userByte, _ := json.Marshal(manufacture_user)
	err = json.Unmarshal(manufacture_userByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByManufactureID(input *manufacture_userModel.Field) interface{} {
	manufacture_user, err := r.ManufactureUserService.GetByManufactureID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &manufacture_userModel.ManufactureAccounts{}
	manufacture_userByte, _ := json.Marshal(manufacture_user)
	err = json.Unmarshal(manufacture_userByte, &output.ManufactureAccount)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	for _,value := range output.ManufactureAccount{
		personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: value.UserID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
		err = json.Unmarshal(personnel_affiliationByte, &value.Deps)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *manufacture_userModel.Updated) interface{} {
	_, err := r.ManufactureUserService.GetByID(&manufacture_userModel.Field{MuID: input.MuID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureUserService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *manufacture_userModel.Updated) interface{} {
	manufacture_user, err := r.ManufactureUserService.GetByID(&manufacture_userModel.Field{MuID: input.MuID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureUserService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, manufacture_user.MuID)
}
