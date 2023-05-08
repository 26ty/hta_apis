package gateway_data

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	gateway_dataModel "eirc.app/internal/v1/structure/gateway_data"
	gg_data_demandModel "eirc.app/internal/v1/structure/gg_data_demand"
	personnel_affiliationModel "eirc.app/internal/v1/structure/personnel_affiliation"
	Model "eirc.app/internal/v1/structure"
	"encoding/json"
	"strconv"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *gateway_dataModel.Created) interface{} {
	defer trx.Rollback()

	gateway_data, err := r.GatewayDataService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, gateway_data.GdID)
}

func (r *resolver) List(input *gateway_dataModel.Fields) interface{} {
	output := &gateway_dataModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, gateway_data, err := r.GatewayDataService.List(input)
	output.Total = quantity
	gateway_dataByte, err := json.Marshal(gateway_data)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(gateway_dataByte, &output.GatewayData)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *gateway_dataModel.Field) interface{} {
	gateway_data, err := r.GatewayDataService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &gateway_dataModel.Single{}
	gateway_dataByte, _ := json.Marshal(gateway_data)
	err = json.Unmarshal(gateway_dataByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByClassificationTitle(input *gateway_dataModel.Field) interface{} {
	gateway_data, err := r.GatewayDataService.GetByClassificationTitle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &gateway_dataModel.Single{}
	gateway_dataByte, _ := json.Marshal(gateway_data)
	err = json.Unmarshal(gateway_dataByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	
	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDataDemand(input []Model.GetCaseListOutput,gdID string,userId string) interface{} {
	demand, err := r.GatewayDataService.GetByID(&gateway_dataModel.Field{GdID:gdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}//確認Gateway_data有無此ID
	account, err_account := r.AccountService.GetByBonitaUserID(string(userId)) //用bonita_user_id換account_id
	if err_account != nil {
		if errors.Is(err_account, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err_account)
		}

		log.Error(err_account)
		return code.GetCodeMessage(code.InternalServerError, err_account)
	}//用bonita_user_id換account_id
	account_id := account.AccountID
	arr := []gg_data_demandModel.Reviews{}//輸出

	for _, value := range input {
		if value.Name == demand.Title{
			bonita := "ARRAY["
			num := 1 //計算陣列數量(用來判斷是否需要,)
			for _,value2 := range input{
				if value2.CaseID == value.CaseID && value2.Name==value.Name{
					if num > 1{
						bonita = bonita +","
					}
					bonita = bonita +"'"+ value2.ParentCaseID +"'"
					num++
				}
			}
			bonita = bonita+"]"
			sql := "select * from "
			data := sql+demand.DataDemand +"("+value.CaseID+",'"+account_id+"','"+value.ParentCaseID+"',"+bonita+")" 
			gateway_data, err := r.GatewayDataService.GetByDataDemand(data)
			if err == nil {
				output := &[]gg_data_demandModel.Reviews{}
				gateway_dataByte, _ := json.Marshal(gateway_data)
				err = json.Unmarshal(gateway_dataByte, &output)
				if err != nil {
					log.Error(err)
					return code.GetCodeMessage(code.InternalServerError, err)
				}
				for _,value1 := range *output{
					if value1.SalesmanID != ""{
						personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(&personnel_affiliationModel.Field{UserID: value1.SalesmanID})
						if err != nil {
							if errors.Is(err, gorm.ErrRecordNotFound) {
								return code.GetCodeMessage(code.DoesNotExist, err)
							}

							log.Error(err)
							return code.GetCodeMessage(code.InternalServerError, err)
						}

						personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
						err = json.Unmarshal(personnel_affiliationByte, &value1.SalesmanDep)
						if err != nil {
							log.Error(err)
							return code.GetCodeMessage(code.InternalServerError, err)
						}
					}
					caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
					value1.BonitaCaseID = float32(caseid)
					value1.BonitaTaskID = value.ID
					value1.BonitaTaskName = value.Name
					arr = append(arr, value1)
				}
				
			}

		}	
	}

	return code.GetCodeMessage(code.Successful, arr)

	// sql := "select * from "
	// data := sql+demand.DataDemand +"("+strconv.Itoa(2283)+")" 

	// gateway_data, err := r.GatewayDataService.GetByDataDemand(data)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// output := &gateway_dataModel.C2_Review{}
	// gateway_dataByte, _ := json.Marshal(gateway_data)
	// err = json.Unmarshal(gateway_dataByte, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *gateway_dataModel.Updated) interface{} {
	_, err := r.GatewayDataService.GetByID(&gateway_dataModel.Field{GdID: input.GdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GatewayDataService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *gateway_dataModel.Updated) interface{} {
	gateway_data, err := r.GatewayDataService.GetByID(&gateway_dataModel.Field{GdID: input.GdID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.GatewayDataService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, gateway_data.GdID)
}
