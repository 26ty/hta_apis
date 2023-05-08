package transaction_record

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	transactionRecordModel "eirc.app/internal/v1/structure/transaction_record"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *transactionRecordModel.Created) interface{} {
	defer trx.Rollback()

	transaction_record, err := r.TransactionRecordService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, transaction_record.TrID)
}

func (r *resolver) List(input *transactionRecordModel.Fields) interface{} {
	output := &transactionRecordModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, transaction_record, err := r.TransactionRecordService.List(input)
	output.Total = quantity
	transaction_recordByte, err := json.Marshal(transaction_record)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(transaction_recordByte, &output.TransactionRecord)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDocumentIDUserList(input *transactionRecordModel.Fields) interface{} {
	output := &transactionRecordModel.Record_user_lists{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, transaction_record, err := r.TransactionRecordService.GetByDocumentIDUserList(input)
	output.Total = quantity
	transaction_recordByte, err := json.Marshal(transaction_record)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(transaction_recordByte, &output.TransactionRecord)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *transactionRecordModel.Field) interface{} {
	transaction_record, err := r.TransactionRecordService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &transactionRecordModel.Single{}
	transaction_recordByte, _ := json.Marshal(transaction_record)
	err = json.Unmarshal(transaction_recordByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *transactionRecordModel.Updated) interface{} {
	_, err := r.TransactionRecordService.GetByID(&transactionRecordModel.Field{TrID: input.TrID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.TransactionRecordService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *transactionRecordModel.Updated) interface{} {
	transaction_record, err := r.TransactionRecordService.GetByID(&transactionRecordModel.Field{TrID: input.TrID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.TransactionRecordService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, transaction_record.TrID)
}
