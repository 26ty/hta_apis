package project

import (
	"encoding/json"
	"errors"
	"strconv"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	projectModel "eirc.app/internal/v1/structure/project"
	Model "eirc.app/internal/v1/structure"
	// accountModel "eirc.app/internal/v1/structure/accounts"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *projectModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱
	// _, err := r.ProjectService.WithTrx(trx).GetByID(&projectModel.Field{PID: input.PID})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	project, err := r.ProjectService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, project.PID)
}

func (r *resolver) List(input *projectModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &projectModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, project, err := r.ProjectService.List(input)
	output.Total = quantity
	projectByte, err := json.Marshal(project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(projectByte, &output.Project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByProjectBonitaUserList(input *projectModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &projectModel.Bonita_ID_Lists{}
	project, err := r.ProjectService.GetByProjectBonitaUserList(input)
	projectByte, err := json.Marshal(project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(projectByte, &output.Project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) ProjectListUser(input *projectModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &projectModel.Project_Accounts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, project, err := r.ProjectService.ProjectListUser(input)
	output.Total = quantity
	projectByte, err := json.Marshal(project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(projectByte, &output.Project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) ProduceSalesListUser(input *projectModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &projectModel.Project_Accounts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, project, err := r.ProjectService.ProduceSalesListUser(input)
	output.Total = quantity
	projectByte, err := json.Marshal(project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(projectByte, &output.Project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) ProjectTemplateListUser(input *projectModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &projectModel.Project_Accounts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, project, err := r.ProjectService.ProjectTemplateListUser(input)
	output.Total = quantity
	projectByte, err := json.Marshal(project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(projectByte, &output.Project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) ProjectAuthorizationListUser(input *projectModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &projectModel.Project_Accounts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, project, err := r.ProjectService.ProjectAuthorizationListUser(input)
	output.Total = quantity
	projectByte, err := json.Marshal(project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(projectByte, &output.Project)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByProjectListUser(input *projectModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	project, err := r.ProjectService.GetByProjectListUser(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &projectModel.Project_Account{}
	projectByte, _ := json.Marshal(project)
	err = json.Unmarshal(projectByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *projectModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	project, err := r.ProjectService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &projectModel.Single{}
	projectByte, _ := json.Marshal(project)
	err = json.Unmarshal(projectByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByCaseIDtoPM(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []projectModel.Project_Account{}

	for _, value := range input {
		if value.Name == "PM待回報專案"{
			project, err := r.ProjectService.GetByCaseID(value.CaseID)
			if err == nil {
				output := projectModel.Project_Account{}
				projectByte, _ := json.Marshal(project)
				err = json.Unmarshal(projectByte, &output)
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
	// project, err := r.ProjectService.GetByID(input)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// output := &projectModel.Single{}
	// projectByte, _ := json.Marshal(project)
	// err = json.Unmarshal(projectByte, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) GetBonitaCaseListPMCompleted(input []Model.GetCaseListOutput) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []projectModel.Project_Account{}

	for _, value := range input {
		if value.Name == "專案完工送審"{
			project, err := r.ProjectService.GetByCaseID(value.CaseID)
			if err == nil {
				output := projectModel.Project_Account{}
				projectByte, _ := json.Marshal(project)
				err = json.Unmarshal(projectByte, &output)
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

func (r *resolver) GetB2BonitaCaseListTM(input []Model.GetCaseListOutput,userId string) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	arr := []projectModel.Tm_Return{} //回傳值
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
		if value.Name == "專案任務工作回報" && ans==0{
			project, err := r.ProjectService.GetByCaseIDBonitaUserID(value.CaseID,account_id)
			if err == nil {
				output := []projectModel.Tm_Return{}
				projectByte, _ := json.Marshal(project)
				err = json.Unmarshal(projectByte, &output)
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

func (r *resolver) GetB2BonitaCaseListCountersign(input []Model.GetCaseListOutput) interface{} {

	arr := []projectModel.Tm_Return{}

	for _, value := range input {
		if value.Name == "會簽"{
			project, err := r.ProjectService.GetByCaseIDTaskUserParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := projectModel.Tm_Return{}
				projectByte, _ := json.Marshal(project)
				err = json.Unmarshal(projectByte, &output)
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
	// //input.IsDeleted = util.PointerBool(false)
	// arr := []projectModel.Tm_Return{} //回傳值
	// pass := []string{} //要跳過的CASEID
	// ans := 0 //計算一圈input跟pass的重複次數

	// for _, value := range input {
	// 	for _,p := range pass{
	// 		if value.CaseID == p{
	// 			ans++
	// 		}
	// 	}
	// 	if value.Name == "會簽" && ans==0{
	// 		project, err := r.ProjectService.GetByCaseIDTaskUserStatus(value.CaseID,value.ParentCaseID)
	// 		if err == nil {
	// 			output := []projectModel.Tm_Return{}
	// 			projectByte, _ := json.Marshal(project)
	// 			err = json.Unmarshal(projectByte, &output)
	// 			if err != nil {
	// 				log.Error(err)
	// 				return code.GetCodeMessage(code.InternalServerError, err)
	// 			}
	// 			bonita := []Model.GetCaseListOutput{}
	// 			for _,value2 := range input{
	// 				if value2.CaseID == value.CaseID && value2.Name==value.Name{
	// 					bonita = append(bonita,value2)
	// 				}
	// 			}
	// 			i:=0
	// 			for _,value1 := range output{
	// 				caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
	// 				value1.BonitaCaseID = float32(caseid)
	// 				value1.BonitaTaskID = bonita[i].ID
	// 				value1.BonitaTaskName = bonita[i].Name
	// 				arr = append(arr, value1)
	// 				i++
	// 			}
	// 			pass = append(pass,value.CaseID)
	// 		}
	// 	}
	// 	ans=0	
	// }
	
	
	// return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) GetB2BonitaCaseListConfirm(input []Model.GetCaseListOutput) interface{} {

	arr := []projectModel.Tm_Return{}

	for _, value := range input {
		if value.Name == "確認會簽內容"{
			project, err := r.ProjectService.GetByCaseIDTaskUserParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := projectModel.Tm_Return{}
				projectByte, _ := json.Marshal(project)
				err = json.Unmarshal(projectByte, &output)
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
	// //input.IsDeleted = util.PointerBool(false)
	// arr := []projectModel.Tm_Return{} //回傳值
	// pass := []string{} //要跳過的CASEID
	// ans := 0 //計算一圈input跟pass的重複次數

	// for _, value := range input {
	// 	for _,p := range pass{
	// 		if value.CaseID == p{
	// 			ans++
	// 		}
	// 	}
	// 	if value.Name == "確認會簽內容" && ans==0{
	// 		project, err := r.ProjectService.GetByCaseIDTaskUserStatus(value.CaseID,value.ParentCaseID)
	// 		if err == nil {
	// 			output := []projectModel.Tm_Return{}
	// 			projectByte, _ := json.Marshal(project)
	// 			err = json.Unmarshal(projectByte, &output)
	// 			if err != nil {
	// 				log.Error(err)
	// 				return code.GetCodeMessage(code.InternalServerError, err)
	// 			}
	// 			bonita := []Model.GetCaseListOutput{}
	// 			for _,value2 := range input{
	// 				if value2.CaseID == value.CaseID && value2.Name==value.Name{
	// 					bonita = append(bonita,value2)
	// 				}
	// 			}
	// 			i:=0
	// 			for _,value1 := range output{
	// 				caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
	// 				value1.BonitaCaseID = float32(caseid)
	// 				value1.BonitaTaskID = bonita[i].ID
	// 				value1.BonitaTaskName = bonita[i].Name
	// 				arr = append(arr, value1)
	// 				i++
	// 			}
	// 			pass = append(pass,value.CaseID)
	// 		}
	// 	}
	// 	ans=0	
	// }
	
	
	// return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) GetB2BonitaCaseListDepartment(input []Model.GetCaseListOutput,userId string) interface{} {
	arr := []projectModel.Tm_Return{}

	for _, value := range input {
		if value.Name == "專案任務完工送審"{
			project, err := r.ProjectService.GetByCaseIDTaskUserParentcaseID(value.CaseID,value.ParentCaseID)
			if err == nil {
				output := projectModel.Tm_Return{}
				projectByte, _ := json.Marshal(project)
				err = json.Unmarshal(projectByte, &output)
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
	
	// //input.IsDeleted = util.PointerBool(false)
	// arr := []projectModel.Tm_Return{} //回傳值
	// pass := []string{} //要跳過的CASEID
	// ans := 0 //計算一圈input跟pass的重複次數

	// for _, value := range input {
	// 	for _,p := range pass{
	// 		if value.CaseID == p{
	// 			ans++
	// 		}
	// 	}
	// 	if value.Name == "專案任務完工送審" && ans==0{
	// 		project, err := r.ProjectService.GetByCaseIDTaskUserStatus(value.CaseID,value.ParentCaseID)
	// 		if err == nil {
	// 			output := []projectModel.Tm_Return{}
	// 			projectByte, _ := json.Marshal(project)
	// 			err = json.Unmarshal(projectByte, &output)
	// 			if err != nil {
	// 				log.Error(err)
	// 				return code.GetCodeMessage(code.InternalServerError, err)
	// 			}
	// 			bonita := []Model.GetCaseListOutput{}
	// 			for _,value2 := range input{
	// 				if value2.CaseID == value.CaseID && value2.Name==value.Name{
	// 					bonita = append(bonita,value2)
	// 				}
	// 			}
	// 			i:=0
	// 			for _,value1 := range output{
	// 				caseid, _ := strconv.ParseFloat(value.CaseID, 32) 
	// 				value1.BonitaCaseID = float32(caseid)
	// 				value1.BonitaTaskID = bonita[i].ID
	// 				value1.BonitaTaskName = bonita[i].Name
	// 				arr = append(arr, value1)
	// 				i++
	// 			}
	// 			pass = append(pass,value.CaseID)
	// 		}
	// 	}
	// 	ans=0	
	// }
	
	
	// return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) Deleted(input *projectModel.Updated) interface{} {
	_, err := r.ProjectService.GetByID(&projectModel.Field{PID: input.PID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProjectService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *projectModel.Updated) interface{} {
	project, err := r.ProjectService.GetByID(&projectModel.Field{PID: input.PID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProjectService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, project.PID)
}

func (r *resolver) Updated_Bonita(input *projectModel.Updated_Bonita) interface{} {
	_, err := r.ProjectService.GetByID(&projectModel.Field{PID: input.PID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProjectService.Updated_Bonita(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, input.BonitaCaseID)
}
