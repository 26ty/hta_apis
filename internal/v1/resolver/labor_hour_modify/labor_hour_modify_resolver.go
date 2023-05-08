package labor_hour_modify

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"
	"strconv"
	Model "eirc.app/internal/v1/structure"
	labor_hour_modifyModel "eirc.app/internal/v1/structure/labor_hour_modify"
	labor_hourModel "eirc.app/internal/v1/structure/labor_hour"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *labor_hour_modifyModel.Created) interface{} {
	defer trx.Rollback()

	labor_hour_modify, err := r.LaborHourModifyService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, labor_hour_modify.HmID)
}

func (r *resolver) List(input *labor_hour_modifyModel.Fields) interface{} {
	output := &labor_hour_modifyModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, labor_hour_modify, err := r.LaborHourModifyService.List(input)
	output.Total = quantity
	labor_hour_modifyByte, err := json.Marshal(labor_hour_modify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hour_modifyByte, &output.LaborHourModify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *labor_hour_modifyModel.Field) interface{} {
	labor_hour_modify, err := r.LaborHourModifyService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &labor_hour_modifyModel.Single{}
	labor_hour_modifyByte, _ := json.Marshal(labor_hour_modify)
	err = json.Unmarshal(labor_hour_modifyByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdList(input *labor_hour_modifyModel.Field) interface{} {
	output := &labor_hour_modifyModel.LaborHourModifys{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, labor_hour_modify, err := r.LaborHourModifyService.GetByUserIdList(input)
	output.Total = quantity
	labor_hour_modifyByte, err := json.Marshal(labor_hour_modify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hour_modifyByte, &output.LaborHourModify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIdLaborHourModifyList(input *labor_hour_modifyModel.Field) interface{} {
	output := &labor_hour_modifyModel.List{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, labor_hour_modify, err := r.LaborHourModifyService.GetByUserIdLaborHourModifyList(input)
	output.Total = quantity
	labor_hour_modifyByte, err := json.Marshal(labor_hour_modify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hour_modifyByte, &output.LaborHourModify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCuIdLaborHourModifyList(input *labor_hour_modifyModel.Field) interface{} {
	output := &labor_hour_modifyModel.List{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, labor_hour_modify, err := r.LaborHourModifyService.GetByCuIdLaborHourModifyList(input)
	output.Total = quantity
	labor_hour_modifyByte, err := json.Marshal(labor_hour_modify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(labor_hour_modifyByte, &output.LaborHourModify)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []labor_hour_modifyModel.ReviewByDepartment{}

	for _, value := range input {
		if value.Name == "主管審核"{
			labor_hour_modify, err := r.LaborHourModifyService.GetByCaseID(value.CaseID)
			if err == nil {
				output := labor_hour_modifyModel.ReviewByDepartment{}
				labor_hour_modifyByte, _ := json.Marshal(labor_hour_modify)
				err = json.Unmarshal(labor_hour_modifyByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
				output.BonitaCaseID = float32(caseid)
				output.BonitaTaskID = value.ID
				output.BonitaTaskName = value.Name
				arr = append(arr, output)
			}
		}	
	}

	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) Deleted(input *labor_hour_modifyModel.Updated) interface{} {
	_, err := r.LaborHourModifyService.GetByID(&labor_hour_modifyModel.Field{HmID: input.HmID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourModifyService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *labor_hour_modifyModel.Updated) interface{} {
	labor_hour_modify, err := r.LaborHourModifyService.GetByID(&labor_hour_modifyModel.Field{HmID: input.HmID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourModifyService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, labor_hour_modify.HmID)
}

func (r *resolver) UpdatedStatus(input *labor_hour_modifyModel.Updated_Review) interface{} {
	labor_hour_modify, err := r.LaborHourModifyService.GetByID(&labor_hour_modifyModel.Field{HmID: input.HmID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourModifyService.UpdatedStatus(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, labor_hour_modify.HmID)
}

func (r *resolver) Replace(input *labor_hour_modifyModel.Field) interface{} {
	labor_hour_modify, err := r.LaborHourModifyService.GetByID(&labor_hour_modifyModel.Field{HmID: input.HmID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourModifyService.UpdatedStatus(&labor_hour_modifyModel.Updated_Review{HmID: input.HmID,StatusTypeID: "fc758b55-e02d-462e-9ab8-734417bcbcae"})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	update := &labor_hourModel.Updated{}
	labor_hour_modifyByte, _ := json.Marshal(labor_hour_modify)
	err = json.Unmarshal(labor_hour_modifyByte, &update)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	update.HID = labor_hour_modify.HourID

	labor_hour, err := r.LaborHourService.GetByID(&labor_hourModel.Field{HID: update.HID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.LaborHourService.Updated(update)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, labor_hour.HID)

	// err = r.LaborHourService.Updated(update)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// return code.GetCodeMessage(code.Successful, update.HID)
}
