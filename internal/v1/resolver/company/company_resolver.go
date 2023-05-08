package company

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *companyModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	company, err := r.Company.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, company.CompanyID)
}

func (r *resolver) List(input *companyModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &companyModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, companies, err := r.Company.List(input)
	companiesByte, err := json.Marshal(companies)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(companiesByte, &output.Companies)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *companyModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	base, err := r.Company.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontCompany := &companyModel.Single{}
	companyByte, _ := json.Marshal(base)
	err = json.Unmarshal(companyByte, &frontCompany)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontCompany)
}

func (r *resolver) Deleted(input *companyModel.Updated) interface{} {
	_, err := r.Company.GetByID(&companyModel.Field{CompanyID: input.CompanyID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.Company.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *companyModel.Updated) interface{} {
	company, err := r.Company.GetByID(&companyModel.Field{CompanyID: input.CompanyID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.Company.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, company.CompanyID)
}
