package jobtitle

import (
	//bpm "eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	jobtitleModel "eirc.app/internal/v1/structure/jobtitle"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *jobtitleModel.Created) interface{} {
	defer trx.Rollback()

	jobtitle, err := r.JobtitleService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, jobtitle.JID)
}

func (r *resolver) List(input *jobtitleModel.Fields) interface{} {
	output := &jobtitleModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, jobtitle, err := r.JobtitleService.List(input)
	output.Total = quantity
	jobtitleByte, err := json.Marshal(jobtitle)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(jobtitleByte, &output.Jobtitle)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *jobtitleModel.Field) interface{} {
	jobtitle, err := r.JobtitleService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &jobtitleModel.Single{}
	jobtitleByte, _ := json.Marshal(jobtitle)
	err = json.Unmarshal(jobtitleByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *jobtitleModel.Updated) interface{} {
	_, err := r.JobtitleService.GetByID(&jobtitleModel.Field{JID: input.JID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JobtitleService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *jobtitleModel.Updated) interface{} {
	jobtitle, err := r.JobtitleService.GetByID(&jobtitleModel.Field{JID: input.JID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JobtitleService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, jobtitle.JID)
}
