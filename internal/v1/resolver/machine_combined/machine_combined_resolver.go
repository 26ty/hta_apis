package machine_combined

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	machine_combinedModel "eirc.app/internal/v1/structure/machine_combined"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *machine_combinedModel.Created) interface{} {
	defer trx.Rollback()

	machine_combined, err := r.MachineCombinedService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, machine_combined.McID)
}

func (r *resolver) List(input *machine_combinedModel.Fields) interface{} {
	output := &machine_combinedModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, machine_combined, err := r.MachineCombinedService.List(input)
	output.Total = quantity
	machine_combinedByte, err := json.Marshal(machine_combined)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(machine_combinedByte, &output.MachineCombined)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) MachineCombinedListLast(input *machine_combinedModel.Fields) interface{} {
	output := &machine_combinedModel.Machine_Combined_Lasts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, machine_combined, err := r.MachineCombinedService.MachineCombinedListLast(input)
	output.Total = quantity
	machine_combinedByte, err := json.Marshal(machine_combined)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(machine_combinedByte, &output.MachineCombined)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByPIDMachineCombinedListLast(input *machine_combinedModel.Fields) interface{} {
	output := &machine_combinedModel.Machine_Combined_Lasts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, machine_combined, err := r.MachineCombinedService.GetByPIDMachineCombinedListLast(input)
	output.Total = quantity
	machine_combinedByte, err := json.Marshal(machine_combined)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(machine_combinedByte, &output.MachineCombined)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *machine_combinedModel.Field) interface{} {
	machine_combined, err := r.MachineCombinedService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &machine_combinedModel.Single{}
	machine_combinedByte, _ := json.Marshal(machine_combined)
	err = json.Unmarshal(machine_combinedByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *machine_combinedModel.Updated) interface{} {
	_, err := r.MachineCombinedService.GetByID(&machine_combinedModel.Field{McID: input.McID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.MachineCombinedService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *machine_combinedModel.Updated) interface{} {
	machine_combined, err := r.MachineCombinedService.GetByID(&machine_combinedModel.Field{McID: input.McID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.MachineCombinedService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, machine_combined.McID)
}
