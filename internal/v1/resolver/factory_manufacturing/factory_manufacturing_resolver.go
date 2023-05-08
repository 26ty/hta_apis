package factory_manufacturing

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	factory_manufacturingModel "eirc.app/internal/v1/structure/factory_manufacturings"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *factory_manufacturingModel.Created_List) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	// factory_manufacturing, err := r.FactoryManufacturingService.WithTrx(trx).Created(input)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	arr := []string{}
	for _, value := range input.Manufacturing {
		manufacturing, err := r.FactoryManufacturingService.WithTrx(trx).Created(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		arr = append(arr, manufacturing.FmID)
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) List(input *factory_manufacturingModel.Fields) interface{} {
	output := &factory_manufacturingModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, factory_manufacturing, err := r.FactoryManufacturingService.List(input)
	output.Total = quantity
	factory_manufacturingByte, err := json.Marshal(factory_manufacturing)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(factory_manufacturingByte, &output.FactoryManufacturings)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *factory_manufacturingModel.Field) interface{} {
	factory_manufacturing, err := r.FactoryManufacturingService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &factory_manufacturingModel.Single{}
	factory_manufacturingByte, _ := json.Marshal(factory_manufacturing)
	err = json.Unmarshal(factory_manufacturingByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *factory_manufacturingModel.Updated) interface{} {
	_, err := r.FactoryManufacturingService.GetByID(&factory_manufacturingModel.Field{FmID: input.FmID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FactoryManufacturingService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *factory_manufacturingModel.Updated) interface{} {
	factory_manufacturing, err := r.FactoryManufacturingService.GetByID(&factory_manufacturingModel.Field{FmID: input.FmID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FactoryManufacturingService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, factory_manufacturing.FmID)
}
