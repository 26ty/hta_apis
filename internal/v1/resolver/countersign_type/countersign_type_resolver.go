package countersign_type

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	countersign_typeModel "eirc.app/internal/v1/structure/countersign_type"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *countersign_typeModel.Created) interface{} {
	defer trx.Rollback()

	countersign_type, err := r.CountersignTypeService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, countersign_type.CtID)
}

func (r *resolver) List(input *countersign_typeModel.Fields) interface{} {
	output := &countersign_typeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, countersign_type, err := r.CountersignTypeService.List(input)
	output.Total = quantity
	countersign_typeByte, err := json.Marshal(countersign_type)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(countersign_typeByte, &output.CountersignType)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *countersign_typeModel.Field) interface{} {
	countersign_type, err := r.CountersignTypeService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &countersign_typeModel.Single{}
	countersign_typeByte, _ := json.Marshal(countersign_type)
	err = json.Unmarshal(countersign_typeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *countersign_typeModel.Updated) interface{} {
	_, err := r.CountersignTypeService.GetByID(&countersign_typeModel.Field{CtID: input.CtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CountersignTypeService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *countersign_typeModel.Updated) interface{} {
	countersign_type, err := r.CountersignTypeService.GetByID(&countersign_typeModel.Field{CtID: input.CtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CountersignTypeService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, countersign_type.CtID)
}
