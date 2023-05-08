package account

import (
	"encoding/json"
	"errors"
	"strconv"

	bpm "eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	accountModel "eirc.app/internal/v1/structure/accounts"

	//personnel_affiliationModel "eirc.app/internal/v1/structure/personnel_affiliation"

	// companyModel "eirc.app/internal/v1/structure/companies"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *accountModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱
	// _, err := r.Company.WithTrx(trx).GetByID(&companyModel.Field{CompanyID: input.CompanyID,
	// 	IsDeleted: util.PointerBool(false)})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	account, err := r.Account.WithTrx(trx).Created(input)
	if err != nil {
		if err.Error() == "account already exists" {
			return code.GetCodeMessage(code.BadRequest, err.Error())
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, account.AccountID)

}

func (r *resolver) List(input *accountModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &accountModel.List{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, accounts, err := r.Account.List(input)
	output.Total = quantity
	accountsByte, err := json.Marshal(accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	//output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountsByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) AccountNameList(input *accountModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &accountModel.Account_Names{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, accounts, err := r.Account.AccountNameList(input)
	output.Total = quantity
	accountsByte, err := json.Marshal(accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	//output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountsByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) AccountNameDepartmentList(input *accountModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &accountModel.Account_Names{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, accounts, err := r.Account.AccountNameDepartmentList(input)
	output.Total = quantity
	accountsByte, err := json.Marshal(accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	//output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountsByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *accountModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	account, err := r.Account.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &accountModel.Single{}
	accountByte, _ := json.Marshal(account)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *accountModel.Updated) interface{} {
	accounts, err := r.Account.GetByID(&accountModel.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	result := bpm.DeleteUserGetcode(input, accounts.BonitaUserID)

	if result != 200 {
		return code.GetCodeMessage(result, "bonita delete error")
	}

	err = r.Account.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *accountModel.Updated) interface{} {
	account, err := r.Account.GetByID(&accountModel.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	if input.BonitaManagerID != "" {
		input.ManagerID = input.BonitaManagerID
	}

	input.Enabled = strconv.FormatBool(input.Status)

	result := bpm.UpdateUserGetcode(input, account.BonitaUserID)

	if result != 200 {
		return code.GetCodeMessage(result, "bonita update error")
	}

	err = r.Account.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, account.AccountID)
}

// 為了導入EMAIL用的(暫時寫死)
func (r *resolver) UpdatedCsv(input *accountModel.UpdatedCsvList) interface{} {
	arr := []string{}
	for _, value := range input.Account {
		account, err := r.Account.GetByAccount(&accountModel.Field{Account: value.Account,
			IsDeleted: util.PointerBool(false)})
		value.AccountID = account.AccountID
		//value.Dep = account.Dep
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		if value.BonitaManagerID != "" {
			value.ManagerID = value.BonitaManagerID
		}

		value.Enabled = strconv.FormatBool(value.Status)

		marshal, err := json.Marshal(value)
		if err != nil {
			log.Error(err)
			return err
		}
		update := &accountModel.Updated{}
		err = json.Unmarshal(marshal, &update)
		if err != nil {
			log.Error(err)
			return err
		}

		result := bpm.UpdateUserGetcode(update, account.BonitaUserID)

		if result != 200 {
			return code.GetCodeMessage(result, "bonita update error")
		}

		err = r.Account.UpdatedCsv(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		arr = append(arr, account.AccountID)
	}

	return code.GetCodeMessage(code.Successful, arr)
}
