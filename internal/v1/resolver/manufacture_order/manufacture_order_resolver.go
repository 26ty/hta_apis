package manufacture_order

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"
	"strconv"
	manufacture_orderModel "eirc.app/internal/v1/structure/manufacture_order"
	manufacture_userModel "eirc.app/internal/v1/structure/manufacture_user"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *manufacture_orderModel.Created) interface{} {
	defer trx.Rollback()

	manufacture_order, err := r.ManufactureOrderService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, manufacture_order.MID)
}

func (r *resolver) List(input *manufacture_orderModel.Fields) interface{} {
	output := &manufacture_orderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, manufacture_order, err := r.ManufactureOrderService.List(input)
	output.Total = quantity
	manufacture_orderByte, err := json.Marshal(manufacture_order)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(manufacture_orderByte, &output.ManufactureOrder)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

// func (r *resolver) ManufactureOrderCdListUser(input *manufacture_orderModel.Fields) interface{} {
// 	output := &manufacture_orderModel.ManufactureOrder_Cd_Accounts{}
// 	output.Limit = input.Limit
// 	output.Page = input.Page
// 	quantity, manufacture_order, err := r.ManufactureOrderService.ManufactureOrderCdListUser(input)
// 	output.Total = quantity
// 	manufacture_orderByte, err := json.Marshal(manufacture_order)
// 	if err != nil {
// 		log.Error(err)
// 		return code.GetCodeMessage(code.InternalServerError, err.Error())
// 	}

// 	output.Pages = util.Pagination(quantity, output.Limit)
// 	err = json.Unmarshal(manufacture_orderByte, &output.ManufactureOrder)
// 	if err != nil {
// 		log.Error(err)
// 		return code.GetCodeMessage(code.InternalServerError, err.Error())
// 	}

// 	return code.GetCodeMessage(code.Successful, output)
// }

func (r *resolver) ManufactureOrderProjectListUser(input *manufacture_orderModel.Fields) interface{} {
	output := &manufacture_orderModel.ManufactureOrder_Project_Accounts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, manufacture_order, err := r.ManufactureOrderService.ManufactureOrderProjectListUser(input)
	output.Total = quantity
	manufacture_orderByte, err := json.Marshal(manufacture_order)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(manufacture_orderByte, &output.ManufactureOrder)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *manufacture_orderModel.Field) interface{} {
	manufacture_order, err := r.ManufactureOrderService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &manufacture_orderModel.Single{}
	manufacture_orderByte, _ := json.Marshal(manufacture_order)
	err = json.Unmarshal(manufacture_orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByIDOne(input *manufacture_orderModel.Field) interface{} {
	output := &manufacture_orderModel.Ones{}
	manufacture_order, err := r.ManufactureOrderService.GetByIDOne(input)
	manufacture_orderByte, err := json.Marshal(manufacture_order)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(manufacture_orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	input_mu := &manufacture_userModel.Field{}
	input_mu.ManufactureID = manufacture_order.MID
	manufacture_user, err := r.ManufactureUserService.GetByManufactureID(input_mu)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	manufacture_userByte, err := json.Marshal(manufacture_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(manufacture_userByte, &output.ManufactureUser)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	// output := &manufacture_orderModel.Ones{}
	// manufacture_order, err := r.ManufactureOrderService.GetByIDOne(input)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// manufacture_orderByte, err := json.Marshal(manufacture_order)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }
	// err = json.Unmarshal(manufacture_orderByte, &output.ManufactureOrder)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	// input_mu := &manufacture_userModel.Field{}
	// input_mu.ManufactureID = manufacture_order.MID
	// manufacture_user, err := r.ManufactureUserService.GetByManufactureID(input_mu)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// manufacture_userByte, err := json.Marshal(manufacture_user)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }
	// err = json.Unmarshal(manufacture_userByte, &output.ManufactureUser)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByPIDList(input *manufacture_orderModel.Fields) interface{} {
	output := &manufacture_orderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, manufacture_order, err := r.ManufactureOrderService.GetByPIDList(input)
	output.Total = quantity
	manufacture_orderByte, err := json.Marshal(manufacture_order)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(manufacture_orderByte, &output.ManufactureOrder)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []manufacture_orderModel.Review{}

	for _, value := range input {
		if value.Name == "單位主管審核"{
			manufacture_order, err := r.ManufactureOrderService.GetByCaseID(value.CaseID)
			if err == nil {
				output := manufacture_orderModel.Review{}
				manufacture_orderByte, _ := json.Marshal(manufacture_order)
				err = json.Unmarshal(manufacture_orderByte, &output)
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

func (r *resolver) GetByCaseIDtoManufacture(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []manufacture_orderModel.Review{}

	for _, value := range input {
		if value.Name == "生管(製造)審核"{
			manufacture_order, err := r.ManufactureOrderService.GetByCaseID(value.CaseID)
			if err == nil {
				output := manufacture_orderModel.Review{}
				manufacture_orderByte, _ := json.Marshal(manufacture_order)
				err = json.Unmarshal(manufacture_orderByte, &output)
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
	arr := []manufacture_orderModel.Review{}

	for _, value := range input {
		if value.Name == "總經理審核"{
			manufacture_order, err := r.ManufactureOrderService.GetByCaseID(value.CaseID)
			if err == nil {
				output := manufacture_orderModel.Review{}
				manufacture_orderByte, _ := json.Marshal(manufacture_order)
				err = json.Unmarshal(manufacture_orderByte, &output)
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

func (r *resolver) GetByCaseIDtoConfirm(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []manufacture_orderModel.Review{}

	for _, value := range input {
		if value.Name == "確認單號開啟"{
			manufacture_order, err := r.ManufactureOrderService.GetByCaseID(value.CaseID)
			if err == nil {
				output := manufacture_orderModel.Review{}
				manufacture_orderByte, _ := json.Marshal(manufacture_order)
				err = json.Unmarshal(manufacture_orderByte, &output)
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

func (r *resolver) GetByCaseIDtoSave(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []manufacture_orderModel.Review{}

	for _, value := range input {
		if value.Name == "儲存製令單號"{
			manufacture_order, err := r.ManufactureOrderService.GetByCaseID(value.CaseID)
			if err == nil {
				output := manufacture_orderModel.Review{}
				manufacture_orderByte, _ := json.Marshal(manufacture_order)
				err = json.Unmarshal(manufacture_orderByte, &output)
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

func (r *resolver) Deleted(input *manufacture_orderModel.Updated) interface{} {
	_, err := r.ManufactureOrderService.GetByID(&manufacture_orderModel.Field{MID: input.MID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureOrderService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *manufacture_orderModel.Updated) interface{} {
	manufacture_order, err := r.ManufactureOrderService.GetByID(&manufacture_orderModel.Field{MID: input.MID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureOrderService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, manufacture_order.MID)
}

func (r *resolver) Updated_Bonita(input *manufacture_orderModel.Updated_Bonita) interface{} {
	_, err := r.ManufactureOrderService.GetByID(&manufacture_orderModel.Field{MID: input.MID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManufactureOrderService.Updated_Bonita(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, input.BonitaCaseID)
}
