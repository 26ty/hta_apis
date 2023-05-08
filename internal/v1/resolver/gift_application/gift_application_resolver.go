package gift_application

import (
	"encoding/json"
	"errors"
	"strconv"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	Model "eirc.app/internal/v1/structure"
	gift_applicationModel "eirc.app/internal/v1/structure/gift_applications"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *gift_applicationModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	gift_application, err := r.GiftApplicationService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, gift_application.GID)
}

func (r *resolver) List(input *gift_applicationModel.Fields) interface{} {
	output := &gift_applicationModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, gift_application, err := r.GiftApplicationService.List(input)
	output.Total = quantity
	gift_applicationByte, err := json.Marshal(gift_application)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(gift_applicationByte, &output.GiftApplications)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GiftDetailListUser(input *gift_applicationModel.Fields) interface{} {
	quantity, gift_application, err := r.GiftApplicationService.GiftDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &gift_applicationModel.AllGiftDetail{}
	output.Limit = input.Limit
	output.Page = input.Page
	output.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)

	gift_applicationByte, _ := json.Marshal(gift_application)
	err = json.Unmarshal(gift_applicationByte, &output.GiftApplication)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	// for _, value := range output.GiftApplication {
	// 	One_input := &gift_applicationModel.Field{}
	// 	One_input.GID = value.GID

	// 	detail, err := r.GiftApplicationService.GetByGIDDetail(One_input)
	// 	if err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			return code.GetCodeMessage(code.DoesNotExist, err)
	// 		}

	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// 	detailByte, _ := json.Marshal(detail)
	// 	err = json.Unmarshal(detailByte, &value.Detail)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return code.GetCodeMessage(code.InternalServerError, err)
	// 	}

	// }

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByGIDGiftDetailListUser(input *gift_applicationModel.Field) interface{} {
	gift_application, err := r.GiftApplicationService.GetByGIDGiftDetailListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &gift_applicationModel.GiftDetail{}
	gift_applicationByte, _ := json.Marshal(gift_application)
	err = json.Unmarshal(gift_applicationByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []gift_applicationModel.Review{}

	for _, value := range input {
		if value.Name == "單位主管審核" {
			gift_application, err := r.GiftApplicationService.GetByCaseID(value.CaseID)
			if err == nil {
				output := gift_applicationModel.Review{}
				gift_applicationByte, _ := json.Marshal(gift_application)
				err = json.Unmarshal(gift_applicationByte, &output)
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

func (r *resolver) GetByCaseIDtoViceTop(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []gift_applicationModel.Review{}

	for _, value := range input {
		if value.Name == "副總審核" {
			gift_application, err := r.GiftApplicationService.GetByCaseID(value.CaseID)
			if err == nil {
				output := gift_applicationModel.Review{}
				gift_applicationByte, _ := json.Marshal(gift_application)
				err = json.Unmarshal(gift_applicationByte, &output)
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

func (r *resolver) GetByCaseIDtoTop(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []gift_applicationModel.Review{}

	for _, value := range input {
		if value.Name == "總經理審核" {
			gift_application, err := r.GiftApplicationService.GetByCaseID(value.CaseID)
			if err == nil {
				output := gift_applicationModel.Review{}
				gift_applicationByte, _ := json.Marshal(gift_application)
				err = json.Unmarshal(gift_applicationByte, &output)
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

func (r *resolver) GetByCaseIDtoAttm(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []gift_applicationModel.Review{}

	for _, value := range input {
		if value.Name == "經辦結案確認" {
			gift_application, err := r.GiftApplicationService.GetByCaseID(value.CaseID)
			if err == nil {
				output := gift_applicationModel.Review{}
				gift_applicationByte, _ := json.Marshal(gift_application)
				err = json.Unmarshal(gift_applicationByte, &output)
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

func (r *resolver) GetByID(input *gift_applicationModel.Field) interface{} {
	gift_application, err := r.GiftApplicationService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &gift_applicationModel.Single{}
	gift_applicationByte, _ := json.Marshal(gift_application)
	err = json.Unmarshal(gift_applicationByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *gift_applicationModel.Updated) interface{} {
	_, err := r.GiftApplicationService.GetByID(&gift_applicationModel.Field{GID: input.GID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GiftApplicationService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *gift_applicationModel.Updated) interface{} {
	gift_application, err := r.GiftApplicationService.GetByID(&gift_applicationModel.Field{GID: input.GID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GiftApplicationService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, gift_application.GID)
}

func (r *resolver) Updated_Bonita(input *gift_applicationModel.Updated_Bonita) interface{} {
	_, err := r.GiftApplicationService.GetByID(&gift_applicationModel.Field{GID: input.GID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GiftApplicationService.Updated_Bonita(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, input.BonitaCaseID)
}
