package sales_call_record

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	sales_call_recordModel "eirc.app/internal/v1/structure/sales_call_records"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *sales_call_recordModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	sales_call_record, err := r.SalesCallRecordService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, sales_call_record.SID)
}

func (r *resolver) List(input *sales_call_recordModel.Fields) interface{} {
	output := &sales_call_recordModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, sales_call_record, err := r.SalesCallRecordService.AccountList(input)
	output.Total = quantity
	sales_call_recordByte, err := json.Marshal(sales_call_record)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(sales_call_recordByte, &output.SalesCallRecords)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *sales_call_recordModel.Field) interface{} {
	sales_call_record, err := r.SalesCallRecordService.GetBySIDAccount(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &sales_call_recordModel.SalesCallRecord_Account{}
	sales_call_recordByte, _ := json.Marshal(sales_call_record)
	err = json.Unmarshal(sales_call_recordByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *sales_call_recordModel.Updated) interface{} {
	_, err := r.SalesCallRecordService.GetByID(&sales_call_recordModel.Field{SID: input.SID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.SalesCallRecordService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *sales_call_recordModel.Updated) interface{} {
	sales_call_record, err := r.SalesCallRecordService.GetByID(&sales_call_recordModel.Field{SID: input.SID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.SalesCallRecordService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, sales_call_record.SID)
}
