package login

import (
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/jwe"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	accountsModel "eirc.app/internal/v1/structure/accounts"
	jweModel "eirc.app/internal/v1/structure/jwe"
	loginsModel "eirc.app/internal/v1/structure/logins"
	Model "eirc.app/internal/v1/structure"

	"gorm.io/gorm"
)

func (r *resolver) Web(input *loginsModel.Login) interface{} {
	acknowledge, accounts, err := r.Account.AcknowledgeAccount(&accountsModel.Field{
		Account:   util.PointerString(input.Account),
		Password:  util.PointerString(input.Password),
		CompanyID: util.PointerString(input.CompanyID),
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if acknowledge == false {
		return code.GetCodeMessage(code.PermissionDenied, "Incorrect account password")
	}

	token, err := r.JWE.Created(&jweModel.JWE{
		AccountID: accounts[0].AccountID,
		CompanyID: input.CompanyID,
		Name:      accounts[0].Name,
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, token)

}

func (r *resolver) Refresh(input *jweModel.Refresh) interface{} {
	privateKey := "-----BEGIN PRIVATE KEY-----\n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\nix/Ccf7Enn38xwIXaZjumpva1dg7bdc=\n-----END PRIVATE KEY-----"
	j := &jwe.JWT{
		PrivateKey: privateKey,
		Token:      input.RefreshToken,
	}

	if len(j.Token) == 0 {
		return code.GetCodeMessage(code.JWTRejected, "refresh_token is error")
	}

	j, err := j.Verify()
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.JWTRejected, "refresh_token is error")
	}

	account, err := r.Account.GetByID(&accountsModel.Field{
		AccountID: j.Other["account_id"].(string),
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.JWTRejected, "refresh_token is error")
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	token, err := r.JWE.Created(&jweModel.JWE{
		AccountID: account.AccountID,
		CompanyID: account.CompanyID,
		Name:      account.Name,
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	token.RefreshToken = input.RefreshToken
	return code.GetCodeMessage(code.Successful, token)
}

func (r *resolver) GetBonitaCaseCount(input []Model.GetCaseListOutput) interface{} {
	arr := []Model.GonitaListCount{}//待審核單據列表
	num := 0 //計算是否已存在待審核單據列表
	result := 0 //紀錄要存在arr的哪個index
	category := "" //流程類別
	for _, input_value := range input {

		_, cd_err := r.CustomerDemandService.GetByCaseID(input_value.CaseID)
		_, p_err := r.ProjectService.GetByCaseID(input_value.CaseID)
		_, mo_err := r.ManufactureOrderService.GetByCaseID(input_value.CaseID)
		_, lh_err := r.LaborHourModifyService.GetByCaseID(input_value.CaseID)
		
		if cd_err == nil{
			category = "A1"
		}else if p_err == nil{
			category = "B2"
		}else if mo_err == nil {
			category = "C2"
		}else if lh_err == nil {
			category = "Labor"
		}else{
			category = ""
		}

		if category != "" {
			for index,arr_value := range arr{
				if input_value.Name == arr_value.Name && arr_value.Category == category{
					num++ //已存在待審核單據列表
					result = index //已存在待審核單據列表的位置(index)
					break
				}
			}
			if num != 0{
				arr[result].Count = arr[result].Count + 1
				num = 0 //重新計算下個是否已存在待審核單據列表
			}else{
				output := Model.GonitaListCount{} 
				output.Category = category
				output.Count = 1
				output.Name = input_value.Name
				arr = append(arr, output) //將此項目加入待審核單據列表
				num = 0 //重新計算下個是否已存在待審核單據列表
			}
			
		}
		
	}

	return code.GetCodeMessage(code.Successful, arr)

}

func (r *resolver) GetBonitaCaseDetail(input []Model.GetCaseListOutput) interface{} {
	arr := []Model.GetCaseDetailListOutput{}//待審核單據列表

	category := "" //流程類別
	for _, input_value := range input {

		customer_demand, cd_err := r.CustomerDemandService.GetByCaseID(input_value.CaseID)
		project, p_err := r.ProjectService.GetByCaseID(input_value.CaseID)
		manufacture, mo_err := r.ManufactureOrderService.GetByCaseID(input_value.CaseID)
		labor_hour, lh_err := r.LaborHourModifyService.GetByCaseID(input_value.CaseID)
		
		if cd_err == nil{
			category = "A1"
		}else if p_err == nil{
			category = "B2"
		}else if mo_err == nil {
			category = "C2"
		}else if lh_err == nil {
			category = "Labor"
		}else{
			category = ""
		}
		output := Model.GetCaseDetailListOutput{}
		output.Category = category
		output.BonitaTaskName = input_value.Name
		output.BonitaTaskID = input_value.ID
		output.BonitaCaseID = input_value.CaseID
		output.BonitaParentCaseID = input_value.ParentCaseID
		switch category {
		case "A1": //case內自帶break方法，執行完就會跳出switch
			output.CdID = customer_demand.CdID
			output.CdCode = customer_demand.Code
			output.DemandContent = customer_demand.DemandContent 

			customer_demand_task, cd_task_err := r.CustomerDemandService.GetByCaseIDTaskUserStatus(input_value.CaseID,input_value.ParentCaseID)
			if cd_task_err == nil{
				output.TID = customer_demand_task.TID
				output.TName = customer_demand_task.TName
				output.PrincipalName = customer_demand_task.Name
				output.TuID  = customer_demand_task.TuID  
			}
			arr = append(arr, output) //將此項目加入待審核單據列表
		case "B2":
			output.PID = project.PID
			output.PCode = project.Code
			output.PName  = project.PName 

			project_task, p_task_err := r.ProjectService.GetByCaseIDTaskUserParentcaseID(input_value.CaseID,input_value.ParentCaseID)
			if p_task_err == nil{
				output.TID = project_task.TID
				output.TName = project_task.TName
				output.PrincipalName = project_task.Name
				output.TuID  = project_task.TuID 
			}
			arr = append(arr, output) //將此項目加入待審核單據列表
		case "C2":
			output.MID = manufacture.MID
			output.MCode = manufacture.Code
			output.ProjectDetail  = manufacture.ProjectDetail 
			arr = append(arr, output) //將此項目加入待審核單據列表
		case "Labor":
			output.HmID = labor_hour.HmID
			output.HourID = labor_hour.HourID
			output.HCategory = labor_hour.Category
			output.HCreaterName = labor_hour.CreaterName
			output.HLaborhour = labor_hour.Laborhour
			arr = append(arr, output) //將此項目加入待審核單據列表
		}

	}

	return code.GetCodeMessage(code.Successful, arr)

}