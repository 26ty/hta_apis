package jig_demand_detail

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	jig_demand_detailModel "eirc.app/internal/v1/structure/jig_demand_details"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *jig_demand_detailModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	jig_demand_detail, err := r.JigDemandDetailService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, jig_demand_detail.JdID)

	// arr := []string{}
	// for _, value := range input.Detail{
	// 	jig_demand_detail, err := r.JigDemandDetailService.WithTrx(trx).Created(value)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err.Error())
	// 	}
	// 	arr = append(arr, jig_demand_detail.JdID)
	// }

	// trx.Commit()
	// return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) List(input *jig_demand_detailModel.Fields) interface{} {
	output := &jig_demand_detailModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, jig_demand_detail, err := r.JigDemandDetailService.List(input)
	output.Total = quantity
	jig_demand_detailByte, err := json.Marshal(jig_demand_detail)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(jig_demand_detailByte, &output.JigDemandDetails)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *jig_demand_detailModel.Field) interface{} {
	jig_demand_detail, err := r.JigDemandDetailService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &jig_demand_detailModel.Single{}
	jig_demand_detailByte, _ := json.Marshal(jig_demand_detail)
	err = json.Unmarshal(jig_demand_detailByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *jig_demand_detailModel.Updated) interface{} {
	_, err := r.JigDemandDetailService.GetByID(&jig_demand_detailModel.Field{JdID: input.JdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JigDemandDetailService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *jig_demand_detailModel.Updated) interface{} {
	jig_demand_detail, err := r.JigDemandDetailService.GetByID(&jig_demand_detailModel.Field{JdID: input.JdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JigDemandDetailService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, jig_demand_detail.JdID)
}

func (r *resolver) UpdatedByJigID(input *jig_demand_detailModel.Updated) interface{} {
	_, err := r.JigDemandDetailService.GetByJigID(&jig_demand_detailModel.Field{JigID: input.JigID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JigDemandDetailService.UpdatedByJigID(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "update success!")
}
