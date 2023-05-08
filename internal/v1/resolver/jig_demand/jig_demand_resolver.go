package jig_demand

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	jig_demandModel "eirc.app/internal/v1/structure/jig_demands"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *jig_demandModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	jig_demand, err := r.JigDemandService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, jig_demand.JID)
}

func (r *resolver) List(input *jig_demandModel.Fields) interface{} {
	output := &jig_demandModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, jig_demand, err := r.JigDemandService.List(input)
	output.Total = quantity
	jig_demandByte, err := json.Marshal(jig_demand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(jig_demandByte, &output.JigDemands)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) JigDetailListUser(input *jig_demandModel.Fields) interface{} {
	quantity, jig_demand, err := r.JigDemandService.JigDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &jig_demandModel.AllJigDetail{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)

	jig_demandByte, _ := json.Marshal(jig_demand)
	err = json.Unmarshal(jig_demandByte, &output.JigDemand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	// for _, value := range output.JigDemand {
	// 	One_input := &jig_demandModel.Field{}
	// 	One_input.JID = value.JID

	// 	detial, err := r.JigDemandService.GetByJIDDetail(One_input)
	// 	if err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			return code.GetCodeMessage(code.DoesNotExist, err)
	// 		}

	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// 	detialByte, _ := json.Marshal(detial)
	// 	err = json.Unmarshal(detialByte, &value.Detail)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// }

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) SearchJigDemand(input *jig_demandModel.Fields) interface{} {
	quantity, jig_demand, err := r.JigDemandService.JigDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &jig_demandModel.SearchJigDemand{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)

	jig_demandByte, _ := json.Marshal(jig_demand)
	err = json.Unmarshal(jig_demandByte, &output.SearchJigDemand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByJIDJigDetailListUser(input *jig_demandModel.Field) interface{} {
	jig_demand, err := r.JigDemandService.GetByJIDJigDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &jig_demandModel.JigDetail{}
	jig_demandByte, _ := json.Marshal(jig_demand)
	err = json.Unmarshal(jig_demandByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIDListJD(input *jig_demandModel.Users) interface{} {
	output := &jig_demandModel.JDs{}
	quantity, jig_demand, err := r.JigDemandService.GetByUserIDListJD(input)
	output.Total = quantity
	jig_demandByte, err := json.Marshal(jig_demand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(jig_demandByte, &output.JigDemand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *jig_demandModel.Field) interface{} {
	jig_demand, err := r.JigDemandService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &jig_demandModel.Single{}
	jig_demandByte, _ := json.Marshal(jig_demand)
	err = json.Unmarshal(jig_demandByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *jig_demandModel.Updated) interface{} {
	_, err := r.JigDemandService.GetByID(&jig_demandModel.Field{JID: input.JID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JigDemandService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *jig_demandModel.Updated) interface{} {
	jig_demand, err := r.JigDemandService.GetByID(&jig_demandModel.Field{JID: input.JID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JigDemandService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, jig_demand.JID)
}

func (r *resolver) Updated_Bonita(input *jig_demandModel.Updated_Bonita) interface{} {
	_, err := r.JigDemandService.GetByID(&jig_demandModel.Field{JID: input.JID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.JigDemandService.Updated_Bonita(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, input.BonitaCaseID)
}
