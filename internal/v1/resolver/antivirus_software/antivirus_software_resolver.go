package antivirus_software

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	antivirus_softwareModel "eirc.app/internal/v1/structure/antivirus_software"
	Model "eirc.app/internal/v1/structure"
	"strconv"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *antivirus_softwareModel.Created) interface{} {
	defer trx.Rollback()

	antivirus_software, err := r.AntivirusSoftwareService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, antivirus_software.AsID)
}

func (r *resolver) List(input *antivirus_softwareModel.Fields) interface{} {
	output := &antivirus_softwareModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, antivirus_software, err := r.AntivirusSoftwareService.List(input)
	output.Total = quantity
	antivirus_softwareByte, err := json.Marshal(antivirus_software)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(antivirus_softwareByte, &output.AntivirusSoftware)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *antivirus_softwareModel.Field) interface{} {
	antivirus_software, err := r.AntivirusSoftwareService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &antivirus_softwareModel.Single{}
	antivirus_softwareByte, _ := json.Marshal(antivirus_software)
	err = json.Unmarshal(antivirus_softwareByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCaseIDtoTop(input []Model.GetCaseListOutput) interface{} {
	arr := []antivirus_softwareModel.Review{}

	for _, value := range input {
		if value.Name == "主管審核"{
			customer_demand, err := r.AntivirusSoftwareService.GetByCaseID(value.CaseID)
			if err == nil {
				output := antivirus_softwareModel.Review{}
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

func (r *resolver) GetByPIDList(input *antivirus_softwareModel.Fields) interface{} {
	output := &antivirus_softwareModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, antivirus_software, err := r.AntivirusSoftwareService.GetByPIDList(input)
	output.Total = quantity
	antivirus_softwareByte, err := json.Marshal(antivirus_software)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(antivirus_softwareByte, &output.AntivirusSoftware)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *antivirus_softwareModel.Updated) interface{} {
	_, err := r.AntivirusSoftwareService.GetByID(&antivirus_softwareModel.Field{AsID: input.AsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.AntivirusSoftwareService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *antivirus_softwareModel.Updated) interface{} {
	antivirus_software, err := r.AntivirusSoftwareService.GetByID(&antivirus_softwareModel.Field{AsID: input.AsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.AntivirusSoftwareService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, antivirus_software.AsID)
}
