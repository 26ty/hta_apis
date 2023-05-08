package quotation

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	quotationModel "eirc.app/internal/v1/structure/quotations"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *quotationModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	quotation, err := r.QuotationService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, quotation.QID)
}

func (r *resolver) List(input *quotationModel.Fields) interface{} {
	output := &quotationModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, quotation, err := r.QuotationService.List(input)
	output.Total = quantity
	quotationByte, err := json.Marshal(quotation)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(quotationByte, &output.Quotations)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) QuotationDetailListUser(input *quotationModel.Fields) interface{} {
	quantity, quotation, err := r.QuotationService.QuotationDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &quotationModel.AllQuotationDetail{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)

	quotationByte, _ := json.Marshal(quotation)
	err = json.Unmarshal(quotationByte, &output.Quotation)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	// for _, value := range output.Quotation {
	// 	One_input := &quotationModel.Field{}
	// 	One_input.QID = value.QID

	// 	detial, err := r.QuotationService.GetByQIDDetail(One_input)
	// 	if err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			return code.GetCodeMessage(code.DoesNotExist, err)
	// 		}

	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// 	detialByte, _ := json.Marshal(detial)
	// 	err = json.Unmarshal(detialByte, &value.Detail)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// }

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByQIDQuotationDetailListUser(input *quotationModel.Field) interface{} {
	quotation, err := r.QuotationService.GetByQIDQuotationDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &quotationModel.QuotationDetail{}
	quotationByte, _ := json.Marshal(quotation)
	err = json.Unmarshal(quotationByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *quotationModel.Field) interface{} {
	quotation, err := r.QuotationService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &quotationModel.Single{}
	quotationByte, _ := json.Marshal(quotation)
	err = json.Unmarshal(quotationByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *quotationModel.Updated) interface{} {
	_, err := r.QuotationService.GetByID(&quotationModel.Field{QID: input.QID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.QuotationService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *quotationModel.Updated) interface{} {
	quotation, err := r.QuotationService.GetByID(&quotationModel.Field{QID: input.QID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.QuotationService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, quotation.QID)
}

func (r *resolver) Updated_Bonita(input *quotationModel.Updated_Bonita) interface{} {
	_, err := r.QuotationService.GetByID(&quotationModel.Field{QID: input.QID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.QuotationService.Updated_Bonita(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, input.BonitaCaseID)
}
