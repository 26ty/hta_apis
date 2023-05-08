package project_template

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	project_templateModel "eirc.app/internal/v1/structure/project_template"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *project_templateModel.Created) interface{} {
	defer trx.Rollback()

	project_template, err := r.ProjectTemplateService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, project_template.PtID)
}

func (r *resolver) List(input *project_templateModel.Fields) interface{} {
	output := &project_templateModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, project_template, err := r.ProjectTemplateService.List(input)
	output.Total = quantity
	project_templateByte, err := json.Marshal(project_template)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(project_templateByte, &output.ProjectTemplate)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *project_templateModel.Field) interface{} {
	project_template, err := r.ProjectTemplateService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &project_templateModel.Single{}
	project_templateByte, _ := json.Marshal(project_template)
	err = json.Unmarshal(project_templateByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *project_templateModel.Updated) interface{} {
	_, err := r.ProjectTemplateService.GetByID(&project_templateModel.Field{PtID: input.PtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProjectTemplateService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *project_templateModel.Updated) interface{} {
	project_template, err := r.ProjectTemplateService.GetByID(&project_templateModel.Field{PtID: input.PtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProjectTemplateService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, project_template.PtID)
}
