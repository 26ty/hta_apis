package customer_demand

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	"strconv"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"
	personnel_affiliationModel "eirc.app/internal/v1/structure/personnel_affiliation"
	customer_demandModel "eirc.app/internal/v1/structure/customer_demand"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *customer_demandModel.Created) interface{} {
	defer trx.Rollback()

	customer_demand, err := r.CustomerDemandService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, customer_demand.CdId)
}

func (r *resolver) List(input *customer_demandModel.Fields) interface{} {
	output := &customer_demandModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, customer_demand, err := r.CustomerDemandService.List(input)
	output.Total = quantity
	customer_demandByte, err := json.Marshal(customer_demand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(customer_demandByte, &output.CustomerDemand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) CustomerDemandListUser(input *customer_demandModel.Field) interface{} {
	output := &customer_demandModel.Customer_Demand_Accounts{}
	quantity, customer_demand, err := r.CustomerDemandService.CustomerDemandListUser(input)
	output.Total = quantity
	customer_demandByte, err := json.Marshal(customer_demand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(customer_demandByte, &output.CustomerDemand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCuIDCustomerDemandListUser(input *customer_demandModel.Field) interface{} {
	customer_demand, err := r.CustomerDemandService.GetByCuIDCustomerDemandListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &customer_demandModel.Customer_Demand_Account{}
	customer_demandByte, _ := json.Marshal(customer_demand)
	err = json.Unmarshal(customer_demandByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIDListCR(input *customer_demandModel.Users) interface{} {
	output := &customer_demandModel.CRs{}
	quantity, customer_demand, err := r.CustomerDemandService.GetByUserIDListCR(input)
	output.Total = quantity
	customer_demandByte, err := json.Marshal(customer_demand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(customer_demandByte, &output.CustomerDemand)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserIDListHCR(input *customer_demandModel.Users) interface{} {
	_,customer_demand, err := r.CustomerDemandService.GetByUserIDListHCR(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &[]customer_demandModel.H_CR{}
	customer_demandByte, _ := json.Marshal(customer_demand)
	err = json.Unmarshal(customer_demandByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *customer_demandModel.Field) interface{} {
	customer_demand, err := r.CustomerDemandService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &customer_demandModel.Single{}
	customer_demandByte, _ := json.Marshal(customer_demand)
	err = json.Unmarshal(customer_demandByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCaseIDtoDepartment(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "業務經理審核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoDirector(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "回覆簽核意見並指定專案經理"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "PM人選確認並負責RD部門勾選"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoDispatch(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "指派各部門人員(可能1人或多人)"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoEvaluation(input []Model.GetCaseListOutput,userId string) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	account, err_account := r.AccountService.GetByBonitaUserID(string(userId)) //用bonita_user_id換account_id
	if err_account != nil {
		if errors.Is(err_account, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err_account)
		}

		log.Error(err_account)
		return code.GetCodeMessage(code.InternalServerError, err_account)
	}
	account_id := account.AccountID

	for _, value := range input {
		if value.Name == "送交評估報告"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDCountersignUserID(value.CaseID,account_id)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
				output.BonitaCaseID = float32(caseid)
				output.BonitaTaskID = value.ID
				output.BonitaTaskName = value.Name
				arr = append(arr, output)

				err_update := r.CountersignUserService.Updated_Bonita(output.CuID,value.ParentCaseID) 
				if err_update != nil {
					if errors.Is(err_update, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err_update)
					}

					log.Error(err_update)
					return code.GetCodeMessage(code.InternalServerError, err_update)
				}
			}
		}	
	}

	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) GetByCaseIDtoCountersign(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "主管審核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDCountersignParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoPMEvaluation(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review2{}

	for _, value := range input {
		if value.Name == "提交評估報告"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDCountersignParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review2{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoBusiness(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "業務簽核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDCountersignParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoBusinessManager(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "業務經理簽核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDCountersignParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoBusinessDirector(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "副總簽核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDCountersignParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoTaskFinish(input []Model.GetCaseListOutput,userId string) interface{} {
	// //input.IsDeleted = util.PointerBool(false)
	// arr := []customer_demandModel.Customer_Review_Task{}

	// account, err_account := r.AccountService.GetByBonitaUserID(string(userId)) //用bonita_user_id換account_id
	// if err_account != nil {
	// 	if errors.Is(err_account, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err_account)
	// 	}

	// 	log.Error(err_account)
	// 	return code.GetCodeMessage(code.InternalServerError, err_account)
	// }
	// account_id := account.AccountID
	// for _, value := range input {
	// 	if value.Name == "任務工作送審"{
	// 		customer_demand, err := r.CustomerDemandService.GetByCaseIDTaskUser(value.CaseID,account_id,"98c543c6-944e-4860-b323-166ed5f3920e")
	// 		if err == nil {
	// 			output := customer_demandModel.Customer_Review_Task{}
	// 			customer_demandByte, _ := json.Marshal(customer_demand)
	// 			err = json.Unmarshal(customer_demandByte, &output)
	// 			if err != nil {
	// 				log.Error(err)
	// 				return code.GetCodeMessage(code.InternalServerError, err)
	// 			}
	// 			caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
	// 			output.BonitaCaseID = float32(caseid)
	// 			output.BonitaTaskID = value.ID
	// 			output.BonitaTaskName = value.Name
	// 			arr = append(arr, output)

	// 			err_update := r.TaskUserService.Updated_Bonita(output.TuID,value.ParentCaseID) 
	// 			if err_update != nil {
	// 				if errors.Is(err_update, gorm.ErrRecordNotFound) {
	// 					return code.GetCodeMessage(code.DoesNotExist, err_update)
	// 				}

	// 				log.Error(err_update)
	// 				return code.GetCodeMessage(code.InternalServerError, err_update)
	// 			}
	// 		}
	// 	}	
	// }

	// return code.GetCodeMessage(code.Successful, arr)
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review_Task{}//回傳值
	pass := []string{} //要跳過的CASEID
	ans := 0 //計算一圈input跟pass的重複次數

	account, err_account := r.AccountService.GetByBonitaUserID(string(userId)) //用bonita_user_id換account_id
	if err_account != nil {
		if errors.Is(err_account, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err_account)
		}

		log.Error(err_account)
		return code.GetCodeMessage(code.InternalServerError, err_account)
	}
	account_id := account.AccountID
	for _, value := range input {
		for _,p := range pass{
			if value.CaseID == p{
				ans++
			}
		}
		if value.Name == "任務工作送審" && ans==0{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDTaskUser(value.CaseID,account_id,"98c543c6-944e-4860-b323-166ed5f3920e")
			if err == nil {
				output := []customer_demandModel.Customer_Review_Task{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				bonita := []Model.GetCaseListOutput{}
				for _,value2 := range input{
					if value2.CaseID == value.CaseID && value2.Name==value.Name{
						bonita = append(bonita,value2)
					}
				}
				i:=0
				for _,value1 := range output{
					caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
					value1.BonitaCaseID = float32(caseid)
					value1.BonitaTaskID = bonita[i].ID
					value1.BonitaTaskName = bonita[i].Name
					value1.Name = account.Name
					arr = append(arr, value1)
					err_update := r.TaskUserService.Updated_Bonita(value1.TuID,bonita[i].ParentCaseID) 
					if err_update != nil {
						if errors.Is(err_update, gorm.ErrRecordNotFound) {
							return code.GetCodeMessage(code.DoesNotExist, err_update)
						}

						log.Error(err_update)
						return code.GetCodeMessage(code.InternalServerError, err_update)
					}
					i++
				}
				pass = append(pass,value.CaseID)
			}
		}
		ans=0	
	}
	
	
	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) GetByCaseIDtoTaskFinishManager(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review_Task{}

	for _, value := range input {
		if value.Name == "直屬主管審核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseIDTaskUserStatus(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review_Task{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
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

func (r *resolver) GetByCaseIDtoBusinessClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "結案"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoDepartmentClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "業務經理結案審核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoDirectorClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review2{}

	for _, value := range input {
		if value.Name == "副總結案審核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID2(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review2{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
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

func (r *resolver) GetByCaseIDtoTopClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "總經理結案審核"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoProductionClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "製造部主管通知"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoCountersignClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "回報製令完工"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) GetByCaseIDtoPMClose(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []customer_demandModel.Customer_Review{}

	for _, value := range input {
		if value.Name == "製令結案"{
			customer_demand, err := r.CustomerDemandService.GetByCaseID(value.CaseID)
			if err == nil {
				output := customer_demandModel.Customer_Review{}
				customer_demandByte, _ := json.Marshal(customer_demand)
				err = json.Unmarshal(customer_demandByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: output.SalesmanID})
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return code.GetCodeMessage(code.DoesNotExist, err)
					}

					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}

				personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
				err = json.Unmarshal(personnel_affiliationByte, &output.SalesmanDep)
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

func (r *resolver) Deleted(input *customer_demandModel.Updated) interface{} {
	_, err := r.CustomerDemandService.GetByID(&customer_demandModel.Field{CdId: input.CdId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CustomerDemandService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *customer_demandModel.Updated) interface{} {
	customer_demand, err := r.CustomerDemandService.GetByID(&customer_demandModel.Field{CdId: input.CdId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CustomerDemandService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, customer_demand.CdId)
}

func (r *resolver) Updated_Bonita(input *customer_demandModel.Updated_Bonita) interface{} {
	_, err := r.CustomerDemandService.GetByID(&customer_demandModel.Field{CdId: input.CdId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.CustomerDemandService.Updated_Bonita(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, input.BonitaCaseID)
}
