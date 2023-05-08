package gift_application_detail

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	gift_application_detailModel "eirc.app/internal/v1/structure/gift_application_details"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *gift_application_detailModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	gift_application_detail, err := r.GiftApplicationDetailService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, gift_application_detail.GdID)
}

func (r *resolver) List(input *gift_application_detailModel.Fields) interface{} {
	output := &gift_application_detailModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, gift_application_detail, err := r.GiftApplicationDetailService.List(input)
	output.Total = quantity
	gift_application_detailByte, err := json.Marshal(gift_application_detail)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(gift_application_detailByte, &output.GiftApplicationDetails)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *gift_application_detailModel.Field) interface{} {
	gift_application_detail, err := r.GiftApplicationDetailService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &gift_application_detailModel.Single{}
	gift_application_detailByte, _ := json.Marshal(gift_application_detail)
	err = json.Unmarshal(gift_application_detailByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *gift_application_detailModel.Updated) interface{} {
	_, err := r.GiftApplicationDetailService.GetByID(&gift_application_detailModel.Field{GdID: input.GdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GiftApplicationDetailService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *gift_application_detailModel.Updated) interface{} {
	gift_application_detail, err := r.GiftApplicationDetailService.GetByID(&gift_application_detailModel.Field{GdID: input.GdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GiftApplicationDetailService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, gift_application_detail.GdID)
}
