package customer_demand

import (
	"strconv"

	model "eirc.app/internal/v1/structure/customer_demand"
)

func (e *entity) Created(input *model.Create_Table) (err error) {
	err = e.db.Model(&model.Create_Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) CustomerDemandListUser(input *model.Field) (amount int64, output []*model.Customer_Demand_Account, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("customer_demand.cd_id,customer_demand.code,customer_demand.project_id,project.code as p_code,project.p_name,customer_demand.contact_person_id,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.budget,customer_demand.date_for_recive,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,customer_demand.date_for_actual_done,customer_demand.machine_status_id,customer_demand.extend_type_name,customer_demand.extend_rem,customer_demand.date_for_devlop,customer_demand.est_quantity,customer_demand.eva_report,customer_demand.status,customer_demand.accept,customer_demand.creater,customer_demand.create_time,customer_demand.customer_id,c_accounts.name as customer_name,customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.projectman_id,p_accounts.name as projectman_name,customer_demand.bonita_case_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result").
		Joins("left join project on project.p_id = customer_demand.project_id").
		Joins("left join accounts as c_accounts on c_accounts.account_id = customer_demand.customer_id").
		Joins("left join accounts as p_accounts on p_accounts.account_id = customer_demand.projectman_id").
		Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
		Order("customer_demand.create_time desc").Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByCuIDCustomerDemandListUser(input *model.Field) (output *model.Customer_Demand_Account, err error) {
	db := e.db.Raw("select customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,customer_demand.bonita_case_id,customerTable.cd_id,customer_demand.contact_person_id,customer_demand.project_id,project.code as p_code,project.p_name,customer_demand.suitable_content,customer_demand.other_content,customer_demand.budget,customer_demand.date_for_estimated_start,customer_demand.machine_status_id,customer_demand.extend_type_name,customer_demand.extend_rem,customer_demand.est_quantity,customer_demand.eva_report,customer_demand.accept,customer_demand.creater,customer_demand.create_time,customer_demand.code,customer_demand.demand_content,customer_demand.date_for_recive,customer_demand.date_for_devlop,customer_demand.date_for_estimated_end,customer_demand.date_for_actual_done,customer_demand.status,customer_demand.salesman_id,customer_demand.customer_id,customer_demand.projectman_id,customerTable.projectman_name,customerTable.salesman_name,customerTable.customer_name from (select projectmanTable.cd_id,projectmanTable.salesman_name,projectmanTable.projectman_name,accounts.name as customer_name from (select salesmanTable.cd_id,salesmanTable.customer_id,salesmanTable.salesman_name,accounts.name as projectman_name from (select customer_demand.cd_id,customer_demand.customer_id,customer_demand.projectman_id,accounts.name as salesman_name from customer_demand left join accounts on accounts.account_id = customer_demand.salesman_id)as salesmanTable left join accounts on accounts.account_id = salesmanTable.projectman_id) as projectmanTable left join accounts on accounts.account_id = projectmanTable.customer_id) as customerTable left join customer_demand on customer_demand.cd_id = customerTable.cd_id left join project on customer_demand.project_id = project.p_id where customerTable.cd_id = ?", input.CdId)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("cd_id = ?", input.CdId)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByUserIDListCR(input *model.Users) (amount int64, output []*model.CR, err error) {
	db := e.db.Model(&model.Table{})

	// err = db.Select("customer_demand.project_id,customer_demand.cd_id,customer_demand.code as cd_code,customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.demand_content,task.t_id,task.create_time,task_user.tu_id,task_user.user_id as task_user_id,tu_accounts.name as task_user_name,task_user.status_type_id,status_type.status,task.remark,labor_hour.h_id ").
	// 	Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
	// 	Joins("left join task on task.documents_id = customer_demand.cd_id").
	// 	Joins("left join task_user on task_user.task_id = task.t_id").
	// 	Joins("left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id").
	// 	Joins("left join status_type on status_type.st_id = task_user.status_type_id").
	// 	Joins("left join labor_hour on labor_hour.category = task_user.tu_id").
	// 	Where("task_user.user_id = ?", input.UserID).Count(&amount).Find(&output).Error

	err = db.Select("customer_demand.project_id,customer_demand.cd_id,customer_demand.code as cd_code,customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.demand_content,task.t_id,task.create_time,task_user.tu_id,task_user.user_id as task_user_id,tu_accounts.name as task_user_name,task_user.status_type_id,status_type.status,task.remark,customer_demand.suitable_content").
		Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
		Joins("left join task on task.documents_id = customer_demand.cd_id").
		Joins("left join task_user on task_user.task_id = task.t_id").
		Joins("left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id").
		Joins("left join status_type on status_type.st_id = task_user.status_type_id").
		Where("task_user.user_id = ?", input.UserID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByUserIDListHCR(input *model.Users) (amount int64, output []*model.H_CR, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("customer_demand.cd_id,customer_demand.code,customer_demand.demand_content,countersign.cs_id,countersign_user.cu_id,countersign_user.remark,countersign_user.date_for_estimated_completion,countersign_user.date_for_estimated_completion_employee,countersign_user.date_for_completion,countersign_user.date_for_completion_employee,countersign_user.user_id,cu_accounts.name as user_name").
		Joins("left join countersign on countersign.documents_id =customer_demand.cd_id").
		Joins("left join countersign_user on countersign_user.countersign_id =countersign.cs_id").
		Joins("left join accounts as cu_accounts on cu_accounts.account_id =countersign_user.user_id").
		Where("countersign_user.user_id = ?", input.UserID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByCaseID(input string) (output *model.Customer_Review, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})

	err = db.Select("customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,customer_demand.cd_id,customer_demand.code,customer_demand.date_for_recive,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.budget,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,customer_demand.machine_status_id,customer_demand.extend_rem,customer_demand.date_for_devlop,customer_demand.est_quantity,customer_demand.eva_report,customer_demand.create_time,customer_demand.projectman_id,p_accounts.name as projectman_name,customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.customer_id,c_accounts.name as customer_name,customer_demand.contact_person_id,cp_accounts.name as contact_person_name,cp_accounts.email as contact_person_email,cp_accounts.phone as contact_person_phone,customer_demand.bonita_case_id ").
		Joins("left join accounts as p_accounts on p_accounts.account_id = customer_demand.projectman_id ").
		Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
		Joins("left join accounts as c_accounts on c_accounts.account_id = customer_demand.customer_id").
		Joins("left join accounts as cp_accounts on cp_accounts.account_id = customer_demand.contact_person_id").
		Where("customer_demand.bonita_case_id = ?", value).First(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDCountersignUserID(input string, userID string) (output *model.Customer_Review, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})

	err = db.Select("countersign_user.date_for_estimated_completion_employee,countersign_user.date_for_completion_employee,customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,customer_demand.cd_id,customer_demand.code,customer_demand.date_for_recive,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.budget,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,customer_demand.machine_status_id,customer_demand.extend_rem,customer_demand.date_for_devlop,customer_demand.est_quantity,customer_demand.eva_report,customer_demand.create_time,customer_demand.projectman_id,p_accounts.name as projectman_name,customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.customer_id,c_accounts.name as customer_name,customer_demand.contact_person_id,cp_accounts.name as contact_person_name,cp_accounts.email as contact_person_email,cp_accounts.phone as contact_person_phone,customer_demand.bonita_case_id,countersign_user.cu_id,countersign_user.remark,countersign_user.date_for_estimated_completion,countersign_user.user_id,cu_accounts.name as user_name ").
		Joins("left join accounts as p_accounts on p_accounts.account_id = customer_demand.projectman_id ").
		Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
		Joins("left join accounts as c_accounts on c_accounts.account_id = customer_demand.customer_id").
		Joins("left join accounts as cp_accounts on cp_accounts.account_id = customer_demand.contact_person_id").
		Joins("left join countersign on countersign.documents_id = customer_demand.cd_id").
		Joins("left join countersign_user on countersign_user.countersign_id = countersign.cs_id").
		Joins("left join accounts as cu_accounts on cu_accounts.account_id = countersign_user.user_id").
		Where("customer_demand.bonita_case_id = ?", value).
		Where("countersign_user.user_id = ?", userID).First(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDCountersignParentcaseID(input string, bonita_parentcase_id string) (output *model.Customer_Review, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})

	err = db.Select("countersign_user.date_for_estimated_completion_employee,countersign_user.date_for_completion_employee,customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,customer_demand.cd_id,customer_demand.code,customer_demand.date_for_recive,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.budget,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,customer_demand.machine_status_id,customer_demand.extend_rem,customer_demand.date_for_devlop,customer_demand.est_quantity,customer_demand.eva_report,customer_demand.create_time,customer_demand.projectman_id,p_accounts.name as projectman_name,customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.customer_id,c_accounts.name as customer_name,customer_demand.contact_person_id,cp_accounts.name as contact_person_name,cp_accounts.email as contact_person_email,cp_accounts.phone as contact_person_phone,customer_demand.bonita_case_id,countersign_user.cu_id,countersign_user.remark,countersign_user.date_for_estimated_completion,countersign_user.user_id,cu_accounts.name as user_name ").
		Joins("left join accounts as p_accounts on p_accounts.account_id = customer_demand.projectman_id ").
		Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
		Joins("left join accounts as c_accounts on c_accounts.account_id = customer_demand.customer_id").
		Joins("left join accounts as cp_accounts on cp_accounts.account_id = customer_demand.contact_person_id").
		Joins("left join countersign on countersign.documents_id = customer_demand.cd_id").
		Joins("left join countersign_user on countersign_user.countersign_id = countersign.cs_id").
		Joins("left join accounts as cu_accounts on cu_accounts.account_id = countersign_user.user_id").
		Where("customer_demand.bonita_case_id = ?", value).
		Where("countersign_user.bonita_parentcase_id = ?", bonita_parentcase_id).First(&output).Error

	return output, err
}

func (e *entity) GetByCaseID2(input string) (output *model.Customer_Review2, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	// db := e.db.Raw("select customer_demand.cd_id,customer_demand.code,customer_demand.customer_id,accounts.name,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end from customer_demand inner join accounts on customer_demand.customer_id= accounts.account_id where customer_demand.bonita_case_id = ?", value)

	err = db.Select("customer_demand.salesman_id,s_accounts.name as salesman_name,customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,customer_demand.cd_id,customer_demand.code,customer_demand.customer_id,accounts.name,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end ").
		Joins("left join accounts on customer_demand.customer_id= accounts.account_id ").
		Joins("left join accounts as s_accounts on s_accounts.account_id = customer_demand.salesman_id").
		Where("customer_demand.bonita_case_id = ?", value).First(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDTaskUserStatus(caseId string, bonita_parentcase_id string) (output *model.Customer_Review_Task, err error) {
	value, _ := strconv.Atoi(caseId)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select customer_demand.date_for_actual_done,customer_demand.date_for_devlop,customer_demand.date_for_recive,customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,task_user.tu_id,tu_accounts.name,task_user.date_for_delivery,customer_demand.cd_id,customer_demand.code,customer_demand.customer_id,accounts.name as customer_name,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,task.t_id,task.documents_id,task.t_name,task.code as t_code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.next_task,task.next_task_code,task.next_task_name,task.date_for_estimated_start as task_date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion from customer_demand left join accounts on accounts.account_id = customer_demand.customer_id left join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,nexttask.t_id as next_task,nexttask.t_name as next_task_name,nexttask.code as next_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_task_name,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion from task as nexttask right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.t_id = nexttask.last_task) as task on task.documents_id = customer_demand.cd_id left join task_user on task.t_id = task_user.task_id left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id where customer_demand.bonita_case_id = ? AND task_user.bonita_parentcase_id= ? ", value, bonita_parentcase_id)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDTaskUser(caseId string, userID string, status_type_id string) (output []*model.Customer_Review_Task, err error) {
	value, _ := strconv.Atoi(caseId)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select customer_demand.date_for_actual_done,customer_demand.date_for_devlop,customer_demand.date_for_recive,customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,task_user.tu_id,tu_accounts.name,task_user.date_for_delivery,customer_demand.cd_id,customer_demand.code,customer_demand.customer_id,accounts.name as customer_name,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,task.t_id,task.documents_id,task.t_name,task.code as t_code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.next_task,task.next_task_code,task.next_task_name,task.date_for_estimated_start as task_date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion from customer_demand left join accounts on accounts.account_id = customer_demand.customer_id left join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,nexttask.t_id as next_task,nexttask.t_name as next_task_name,nexttask.code as next_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_task_name,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion from task as nexttask right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.t_id = nexttask.last_task) as task on task.documents_id = customer_demand.cd_id left join task_user on task.t_id = task_user.task_id left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id where customer_demand.bonita_case_id = ?  and task_user.user_id = ? AND task_user.status_type_id= ?", value, userID, status_type_id)

	err = db.First(&output).Error

	if err == nil {
		db = e.db.Raw("select customer_demand.project_id,customer_demand.fill,customer_demand.result_status,customer_demand.result_content,customer_demand.date_for_result,task_user.tu_id,tu_accounts.name,task_user.date_for_delivery,customer_demand.cd_id,customer_demand.code,customer_demand.customer_id,accounts.name as customer_name,customer_demand.demand_content,customer_demand.suitable_content,customer_demand.other_content,customer_demand.date_for_estimated_start,customer_demand.date_for_estimated_end,task.t_id,task.documents_id,task.t_name,task.code as t_code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.next_task,task.next_task_code,task.next_task_name,task.date_for_estimated_start as task_date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion from customer_demand left join accounts on accounts.account_id = customer_demand.customer_id left join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,nexttask.t_id as next_task,nexttask.t_name as next_task_name,nexttask.code as next_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_task_name,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion from task as nexttask right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.t_id = nexttask.last_task) as task on task.documents_id = customer_demand.cd_id left join task_user on task.t_id = task_user.task_id left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id where customer_demand.bonita_case_id = ?  and task_user.user_id = ? AND task_user.status_type_id= ?", value, userID, status_type_id)
		err = db.Find(&output).Error
	}

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("cd_id = ?", input.CdId).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("cd_id = ?", input.CdId).Save(&input).Error

	return err
}
