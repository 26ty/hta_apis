package customer

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	customerModel "eirc.app/internal/v1/structure/customers"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *customerModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	customer, err := r.CustomerService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, customer.CID)
}

func (r *resolver) List(input *customerModel.Fields) interface{} {
	output := &customerModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, customer, err := r.CustomerService.AccountList(input)
	output.Total = quantity
	customerByte, err := json.Marshal(customer)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(customerByte, &output.Customers)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *customerModel.Field) interface{} {
	customer, err := r.CustomerService.GetByCIDAccount(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &customerModel.Single{}
	customerByte, _ := json.Marshal(customer)
	err = json.Unmarshal(customerByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *customerModel.Updated) interface{} {
	_, err := r.CustomerService.GetByID(&customerModel.Field{CID: input.CID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CustomerService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *customerModel.Updated) interface{} {
	customer, err := r.CustomerService.GetByID(&customerModel.Field{CID: input.CID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CustomerService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, customer.CID)
}
