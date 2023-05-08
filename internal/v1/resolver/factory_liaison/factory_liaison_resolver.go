package factory_liaison

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	factory_liaisonModel "eirc.app/internal/v1/structure/factory_liaisons"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *factory_liaisonModel.Created_List) interface{} {
	// defer trx.Rollback()
	// // Todo 角色名稱

	// factory_liaison, err := r.FactoryLiaisonService.WithTrx(trx).Created(input)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	arr := []string{}
	for _, value := range input.Liaison {
		liaison, err := r.FactoryLiaisonService.WithTrx(trx).Created(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		arr = append(arr, liaison.FlID)
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) List(input *factory_liaisonModel.Fields) interface{} {
	output := &factory_liaisonModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, factory_liaison, err := r.FactoryLiaisonService.List(input)
	output.Total = quantity
	factory_liaisonByte, err := json.Marshal(factory_liaison)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(factory_liaisonByte, &output.FactoryLiaison)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *factory_liaisonModel.Field) interface{} {
	factory_liaison, err := r.FactoryLiaisonService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &factory_liaisonModel.Single{}
	factory_liaisonByte, _ := json.Marshal(factory_liaison)
	err = json.Unmarshal(factory_liaisonByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *factory_liaisonModel.Updated) interface{} {
	_, err := r.FactoryLiaisonService.GetByID(&factory_liaisonModel.Field{FlID: input.FlID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FactoryLiaisonService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *factory_liaisonModel.Updated) interface{} {
	factory_liaison, err := r.FactoryLiaisonService.GetByID(&factory_liaisonModel.Field{FlID: input.FlID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FactoryLiaisonService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, factory_liaison.FlID)
}
