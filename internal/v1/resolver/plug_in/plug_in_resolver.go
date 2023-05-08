package plug_in

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	plug_inModel "eirc.app/internal/v1/structure/plug_in"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *plug_inModel.Created) interface{} {
	defer trx.Rollback()

	plug_in, err := r.PlugInService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, plug_in.PiID)
}

func (r *resolver) List(input *plug_inModel.Fields) interface{} {
	output := &plug_inModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, plug_in, err := r.PlugInService.List(input)
	output.Total = quantity
	plug_inByte, err := json.Marshal(plug_in)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(plug_inByte, &output.PlugIn)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByPIDList(input *plug_inModel.Fields) interface{} {
	output := &plug_inModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, plug_in, err := r.PlugInService.GetByPIDList(input)
	output.Total = quantity
	plug_inByte, err := json.Marshal(plug_in)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(plug_inByte, &output.PlugIn)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *plug_inModel.Field) interface{} {
	plug_in, err := r.PlugInService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &plug_inModel.Single{}
	plug_inByte, _ := json.Marshal(plug_in)
	err = json.Unmarshal(plug_inByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *plug_inModel.Updated) interface{} {
	_, err := r.PlugInService.GetByID(&plug_inModel.Field{PiID: input.PiID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PlugInService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *plug_inModel.Updated) interface{} {
	plug_in, err := r.PlugInService.GetByID(&plug_inModel.Field{PiID: input.PiID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PlugInService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, plug_in.PiID)
}
