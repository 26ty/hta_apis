package quotation_detail

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	quotation_detailModel "eirc.app/internal/v1/structure/quotation_details"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *quotation_detailModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	quotation_detail, err := r.QuotationDetailService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, quotation_detail.QdID)
}

func (r *resolver) List(input *quotation_detailModel.Fields) interface{} {
	output := &quotation_detailModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, quotation_detail, err := r.QuotationDetailService.List(input)
	output.Total = quantity
	quotation_detailByte, err := json.Marshal(quotation_detail)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(quotation_detailByte, &output.QuotationDetails)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *quotation_detailModel.Field) interface{} {
	quotation_detail, err := r.QuotationDetailService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &quotation_detailModel.Single{}
	quotation_detailByte, _ := json.Marshal(quotation_detail)
	err = json.Unmarshal(quotation_detailByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *quotation_detailModel.Updated) interface{} {
	_, err := r.QuotationDetailService.GetByID(&quotation_detailModel.Field{QdID: input.QdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.QuotationDetailService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *quotation_detailModel.Updated) interface{} {
	quotation_detail, err := r.QuotationDetailService.GetByID(&quotation_detailModel.Field{QdID: input.QdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.QuotationDetailService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, quotation_detail.QdID)
}
