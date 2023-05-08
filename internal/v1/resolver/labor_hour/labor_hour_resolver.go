package labor_hour

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	labor_hourModel "eirc.app/internal/v1/structure/labor_hour"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *labor_hourModel.Created) interface{} {
	defer trx.Rollback()

	labor_hour, err := r.LaborHourService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, labor_hour.HID)
}

func (r *resolver) List(input *labor_hourModel.Fields) interface{} {
	output := &labor_hourModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, labor_hour, err := r.LaborHourService.List(input)
	output.Total = quantity
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output.LaborHour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *labor_hourModel.Field) interface{} {
	labor_hour, err := r.LaborHourService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &labor_hourModel.Single{}
	labor_hourByte, _ := json.Marshal(labor_hour)
	err = json.Unmarshal(labor_hourByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdCategoryList(input *labor_hourModel.Field) interface{} {
	output := &[]labor_hourModel.GetUserCategoryLabor{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	_, labor_hour, err := r.LaborHourService.GetByUserIdCategoryList(input)
	//output.Total = quantity
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdCategory(input *labor_hourModel.Field) interface{} {
	output := &[]labor_hourModel.GetUserCategoryLabor{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	_, labor_hour, err := r.LaborHourService.GetByUserIdCategory(input)
	//output.Total = quantity
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdMonthList(input *labor_hourModel.Field_Month) interface{} {
	output := &labor_hourModel.GetUserAllLabors{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	_, labor_hour, err := r.LaborHourService.GetByUserIdMonthList(input)
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output.LaborHour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for _,value := range output.LaborHour{
		One_input := input
		One_input.Category = value.Category

		labor_hour_one, err := r.LaborHourService.GetByUserIdOneLaborhour(One_input)
		labor_hour_oneByte, err := json.Marshal(labor_hour_one)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(labor_hour_oneByte, &value.DateOfLaborhour)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		labor_hour_one_sum, err := r.LaborHourService.GetByUserIdOneSumLaborhour(One_input)
		value.SumOfLaborhour = labor_hour_one_sum.SumOfLaborhour
	}

	labor_hour_sum, err := r.LaborHourService.GetByUserIdMonthSumList(input)
	labor_hour_sumByte, err := json.Marshal(labor_hour_sum)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(labor_hour_sumByte, &output.DateOfSum)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdList(input *labor_hourModel.Field) interface{} {
	output := &labor_hourModel.LaborHours{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, labor_hour, err := r.LaborHourService.GetByUserIdList(input)
	output.Total = quantity
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output.LaborHour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdLaborHourList(input *labor_hourModel.Field) interface{} {
	output := &labor_hourModel.List{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, labor_hour, err := r.LaborHourService.GetByUserIdLaborHourList(input)
	output.Total = quantity
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output.LaborHour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCuIdLaborHourList(input *labor_hourModel.Field) interface{} {
	output := &labor_hourModel.List{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, labor_hour, err := r.LaborHourService.GetByCuIdLaborHourList(input)
	output.Total = quantity
	labor_hourByte, err := json.Marshal(labor_hour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hourByte, &output.LaborHour)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *labor_hourModel.Updated) interface{} {
	_, err := r.LaborHourService.GetByID(&labor_hourModel.Field{HID: input.HID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *labor_hourModel.Updated) interface{} {
	labor_hour, err := r.LaborHourService.GetByID(&labor_hourModel.Field{HID: input.HID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, labor_hour.HID)
}
