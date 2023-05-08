package file

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	model "eirc.app/internal/v1/structure/file"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *model.Created) interface{} {
	defer trx.Rollback()
	//Todo 檢查重複
	file, err := r.FileService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// aa := New()

	trx.Commit()
	return code.GetCodeMessage(code.Successful, file.FID)
}

func (r *resolver) List(input *model.Fields) interface{} {
	output := &model.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, file, err := r.FileService.List(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total = quantity
	fileByte, err := json.Marshal(file)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(fileByte, &output.File)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDocumentID(input *model.Field) interface{} {
	output := &model.FilebydocumentIds{}
	quantity, file, err := r.FileService.GetByDocumentID(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total = quantity
	fileByte, err := json.Marshal(file)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(fileByte, &output.File)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDocumentIDUserID(input *model.Users) interface{} {
	output := &model.FilebydocumentIds{}
	quantity, file, err := r.FileService.GetByDocumentIDUserID(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total = quantity
	fileByte, err := json.Marshal(file)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(fileByte, &output.File)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *model.Field) interface{} {
	file, err := r.FileService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontMessage := &model.Single{}
	messageByte, _ := json.Marshal(file)
	err = json.Unmarshal(messageByte, &frontMessage)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontMessage)
}

func (r *resolver) Deleted(input *model.Updated) interface{} {
	_, err := r.FileService.GetByID(&model.Field{FID: input.FID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FileService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *model.Updated) interface{} {
	file, err := r.FileService.GetByID(&model.Field{FID: input.FID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FileService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, file.FID)
}
