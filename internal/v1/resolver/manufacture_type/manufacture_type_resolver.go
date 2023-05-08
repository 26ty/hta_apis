package manufacture_type

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	manufacture_typeModel "eirc.app/internal/v1/structure/manufacture_type"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *manufacture_typeModel.Created) interface{} {
	defer trx.Rollback()

	manufacture_type, err := r.ManufactureTypeService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, manufacture_type.MtID)
}

func (r *resolver) List(input *manufacture_typeModel.Fields) interface{} {
	output := &manufacture_typeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, manufacture_type, err := r.ManufactureTypeService.List(input)
	output.Total = quantity
	manufacture_typeByte, err := json.Marshal(manufacture_type)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(manufacture_typeByte, &output.ManufactureType)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *manufacture_typeModel.Field) interface{} {
	manufacture_type, err := r.ManufactureTypeService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &manufacture_typeModel.Single{}
	manufacture_typeByte, _ := json.Marshal(manufacture_type)
	err = json.Unmarshal(manufacture_typeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *manufacture_typeModel.Updated) interface{} {
	_, err := r.ManufactureTypeService.GetByID(&manufacture_typeModel.Field{MtID: input.MtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureTypeService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *manufacture_typeModel.Updated) interface{} {
	manufacture_type, err := r.ManufactureTypeService.GetByID(&manufacture_typeModel.Field{MtID: input.MtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureTypeService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, manufacture_type.MtID)
}
