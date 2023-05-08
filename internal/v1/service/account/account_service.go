package account

import (
	"encoding/json"

	"errors"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/accounts"
)

func (s *service) Created(input *model.Created) (output *model.Base, err error) {
	key := "423CD5C09F7DD58950F1E494099EB075"
	fields := &model.Users{}
	// fields.Limit = 1
	// fields.Page = 1
	fields.Account = util.PointerString(input.Account)
	fields.CompanyID = util.PointerString(input.CompanyID)
	fields.IsDeleted = util.PointerBool(false)
	amount, _, err := s.Entity.List(fields)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if amount > 0 {
		log.Info("Account already exists. Account: ", input.Account, ",CompanyID:", input.CompanyID)
		return nil, errors.New("account already exists")
	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	input.Password = util.HmacSha512(input.Password, key)
	log.Debug(input.Password)
	password, err := util.AesEncryptOFB([]byte(input.Password), []byte(key))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	output.Password = util.Base64BydEncode(password)
	output.AccountID = util.GenerateUUID()
	output.CreatedAt = util.NowToUTC()
	output.UpdatedAt = util.PointerTime(util.NowToUTC())
	output.UpdatedBy = &input.CreatedBy
	output.IsDeleted = false
	output.Status = true

	marshal, err = json.Marshal(output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	table := &model.Table{}
	err = json.Unmarshal(marshal, &table)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = s.Entity.Created(table)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) List(input *model.Users) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.List(input)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	return amount, output, err
}

func (s *service) AccountNameList(input *model.Users) (quantity int64, output []*model.Account_Name, err error) {
	amount, fields, err := s.Entity.AccountNameList(input)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	return amount, output, err
}

func (s *service) AccountNameDepartmentList(input *model.Users) (quantity int64, output []*model.Account_Name, err error) {
	amount, fields, err := s.Entity.AccountNameDepartmentList(input)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	return amount, output, err
}

func (s *service) GetByID(input *model.Field) (output *model.Base, err error) {
	field, err := s.Entity.GetByID(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) GetByAccount(input *model.Field) (output *model.Base, err error) {
	field, err := s.Entity.GetByAccount(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) GetByBonitaUserID(userID string) (output *model.Base, err error) {
	field, err := s.Entity.GetByBonitaUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)
		return err
	}

	field.UpdatedBy = input.UpdatedBy
	field.UpdatedAt = util.PointerTime(util.NowToUTC())
	field.IsDeleted = true
	field.Status = false
	err = s.Entity.Updated(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)
		return err
	}

	key := "423CD5C09F7DD58950F1E494099EB075"

	if input.Password !=""{
		input.Password = util.HmacSha512(input.Password, key)
		log.Debug(input.Password)
		password, err := util.AesEncryptOFB([]byte(input.Password), []byte(key))
		if err != nil {
			log.Error(err)
			return err
		}
		input.Password = util.Base64BydEncode(password)

	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Entity.Updated(field)

	return err
}

//為了導入EMAIL用的(暫時寫死)
func (s *service) UpdatedCsv(input *model.UpdatedCsv) (err error) {
	field, err := s.Entity.GetByAccount(&model.Field{Account: input.Account,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)
		return err
	}

	key := "423CD5C09F7DD58950F1E494099EB075"

	if input.Password !=""{
		input.Password = util.HmacSha512(input.Password, key)
		log.Debug(input.Password)
		password, err := util.AesEncryptOFB([]byte(input.Password), []byte(key))
		if err != nil {
			log.Error(err)
			return err
		}
		input.Password = util.Base64BydEncode(password)

	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Entity.UpdatedCsv(field)

	return err
}


func (s *service) AcknowledgeAccount(input *model.Field) (acknowledge bool, output []*model.Base, err error) {
	key := "423CD5C09F7DD58950F1E494099EB075"
	input.Password = util.PointerString(util.HmacSha512(*input.Password, key))
	fields := &model.Users{}
	// fields.Limit = 1
	// fields.Page = 1
	fields.Account = input.Account
	fields.CompanyID = input.CompanyID
	fields.IsDeleted = util.PointerBool(false)
	amount, accounts, err := s.Entity.List(fields)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	if amount == 0 {
		return false, nil, errors.New("account error")
	}

	password, err := util.AesDecryptOFB(util.Base64StdDecode(accounts[0].Password), []byte(key))
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	if string(password) != *input.Password {
		return false, nil, errors.New("incorrect password")
	}

	marshal, err := json.Marshal(accounts)
	if err != nil {
		log.Error(err)
		return false, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return false, output, err
	}

	return true, output, nil
}
