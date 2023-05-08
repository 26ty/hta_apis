package factory

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	factoryModel "eirc.app/internal/v1/structure/factorys"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *factoryModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	factory, err := r.FactoryService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, factory.FID)
}

func (r *resolver) List(input *factoryModel.Fields) interface{} {
	output := &factoryModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, factory, err := r.FactoryService.List(input)
	output.Total = quantity
	factoryByte, err := json.Marshal(factory)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(factoryByte, &output.Factorys)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) SearchFactory(input *factoryModel.Fields) interface{} {
	quantity, factory, err := r.FactoryService.FLMListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &factoryModel.SearchFactory{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity

	factoryByte, _ := json.Marshal(factory)
	err = json.Unmarshal(factoryByte, &output.SearchFactory)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	// for _, value := range output.SearchFactory {
	// 	One_input := &factoryModel.Field{}
	// 	One_input.FID = value.FID

	// 	factoryManufacturing, err := r.FactoryService.GetByFIDManufacturing(One_input)
	// 	if err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			return code.GetCodeMessage(code.DoesNotExist, err)
	// 		}

	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// 	factoryManufacturingByte, _ := json.Marshal(factoryManufacturing)
	// 	err = json.Unmarshal(factoryManufacturingByte, &value.Manufacturing)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// }

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) FLMListUser(input *factoryModel.Fields) interface{} {
	quantity, factory, err := r.FactoryService.FLMListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &factoryModel.AllFLM{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)

	factoryByte, _ := json.Marshal(factory)
	err = json.Unmarshal(factoryByte, &output.Factory)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	// for _, value := range output.Factory {
	// 	One_input := &factoryModel.Field{}
	// 	One_input.FID = value.FID

	// factoryLiaison, err := r.FactoryService.GetByFIDLiaison(One_input)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }
	// factoryLiaisonByte, _ := json.Marshal(factoryLiaison)
	// err = json.Unmarshal(factoryLiaisonByte, &value.Liaison)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// factoryManufacturing, err := r.FactoryService.GetByFIDManufacturing(One_input)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// factoryManufacturingByte, _ := json.Marshal(factoryManufacturing)
	// err = json.Unmarshal(factoryManufacturingByte, &value.Manufacturing)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// }

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByFIDFLMListUser(input *factoryModel.Field) interface{} {
	factory, err := r.FactoryService.GetByFIDFLMListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &factoryModel.FLM{}
	factoryByte, _ := json.Marshal(factory)
	err = json.Unmarshal(factoryByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *factoryModel.Field) interface{} {
	factory, err := r.FactoryService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &factoryModel.Single{}
	factoryByte, _ := json.Marshal(factory)
	err = json.Unmarshal(factoryByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *factoryModel.Updated) interface{} {
	_, err := r.FactoryService.GetByID(&factoryModel.Field{FID: input.FID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FactoryService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *factoryModel.Updated) interface{} {
	factory, err := r.FactoryService.GetByID(&factoryModel.Field{FID: input.FID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FactoryService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, factory.FID)
}
