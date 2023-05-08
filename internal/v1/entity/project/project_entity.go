package project

import (
	model "eirc.app/internal/v1/structure/project"
	"strconv"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.PID != nil {
	// 	db.Where("p_id = ?", input.PID)
	// }

	if input.Code != nil {
		db.Where("code = ?", input.Code)
	}

	if input.Type != nil {
		db.Where("type like %?%", input.Type)
	}

	if input.PName != nil {
		db.Where("p_name = ?", input.PName)
	}

	if input.CustomerID != nil {
		db.Where("customer_id = ?", input.CustomerID)
	}

	if input.SalesmanID != nil {
		db.Where("salesman_id = ?", input.SalesmanID)
	}

	if input.ServicemanID != nil {
		db.Where("serviceman_id = ?", input.ServicemanID)
	}

	if input.ProjectmanID != nil {
		db.Where("projectman_id = ?", input.ProjectmanID)
	}

	if input.Status != nil {
		db.Where("status = ?", input.Status)
	}

	if input.InnerID != nil {
		db.Where("inner_id = ?", input.InnerID)
	}

	if input.DateForStart != nil {
		db.Where("date_for_start = ?", input.DateForStart)
	}

	if input.DateForEnd != nil {
		db.Where("date_for_end = ?", input.DateForEnd)
	}

	if input.DateForPay != nil {
		db.Where("date_for_pay = ?", input.DateForPay)
	}

	if input.DateForDelivery != nil {
		db.Where("date_for_start = ?", input.DateForDelivery)
	}

	if input.DateForCheck != nil {
		db.Where("date_for_check = ?", input.DateForCheck)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByProjectBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error) {

	db := e.db.Model(&model.Table{})

	err = db.Select("project.p_id,project.projectman_id,accounts.name as projectman_name,accounts.bonita_user_id as projectman_bonita_id").
		Joins("left join accounts on project.projectman_id = accounts.account_id ").
		Where("p_id = ?", input.PID).
		Order("project.create_time desc").Find(&output).Error

	//db.Count(&count)
	return output, err
}

func (e *entity) ProjectListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) {

	db := e.db.Model(&model.Table{})

	err = db.Select("project.p_id,project.code,project.customer_code,project.type,project.p_name,project.customer_id,project.salesman_id,project.serviceman_id,project.projectman_id,c_accounts.name as customer_name,s_accounts.name as salesman_name,sv_accounts.name as serviceman_name,p_accounts.name as projectman_name,project.status,project.inner_id,project.date_for_start,project.date_for_end,project.date_for_pay,project.date_for_delivery,project.date_for_check,project.creater,project.create_time").
		Joins("left join accounts as c_accounts on c_accounts.account_id = project.customer_id ").
		Joins("left join accounts as s_accounts on s_accounts.account_id = project.salesman_id ").
		Joins("left join accounts as sv_accounts on sv_accounts.account_id = project.serviceman_id ").
		Joins("left join accounts as p_accounts on p_accounts.account_id = project.projectman_id ").
		Where("project.origin_id='5451f88e-6d83-44c6-96c3-cd1d049249f7' or project.origin_id='1e6913f5-55be-413a-94a5-68f8cc67d5b2'").
		Order("project.create_time desc").Find(&output).Count(&amount).Error

	return amount, output, err
}

func (e *entity) ProduceSalesListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) {


	db := e.db.Model(&model.Table{})

	err = db.Select("project.is_template,project.machine_quantity,project.p_id,project.code,project.customer_code,project.type,project.p_name,project.customer_id,project.salesman_id,project.serviceman_id,project.projectman_id,p_accounts.name as projectman_name,c_accounts.name as customer_name,s_accounts.name as salesman_name,sv_accounts.name as serviceman_name,project.status,project.inner_id,project.date_for_start,project.date_for_end,project.date_for_pay,project.date_for_delivery,project.date_for_check,project.creater,project.create_time").
		Joins("left join accounts as c_accounts on c_accounts.account_id = project.customer_id ").
		Joins("left join accounts as s_accounts on s_accounts.account_id = project.salesman_id ").
		Joins("left join accounts as sv_accounts on sv_accounts.account_id = project.serviceman_id ").
		Joins("left join accounts as p_accounts on p_accounts.account_id = project.projectman_id ").
		Where("project.origin_id='5451f88e-6d83-44c6-96c3-cd1d049249f7'").
		Order("project.create_time desc").Find(&output).Count(&amount).Error
	return amount, output, err
}

func (e *entity) ProjectTemplateListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) {

	err = e.db.Raw("select project.p_id,project.code,project.customer_code,project.type,project.p_name,project.customer_id,project.salesman_id,project.serviceman_id,project.projectman_id,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.status,project.inner_id,project.date_for_start,project.date_for_end,project.date_for_pay,project.date_for_delivery,project.date_for_check,project.creater,project.create_time from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id where project.origin_id='ef242726-7b97-4943-9318-5eb27c1bb8b5'").
		Order("create_time desc").Find(&output).Error

	e.db.Model(&model.Table{}).Where("project.origin_id='ef242726-7b97-4943-9318-5eb27c1bb8b5'").Count(&amount)
	return amount, output, err
}

func (e *entity) ProjectAuthorizationListUser(input *model.Fields) (amount int64, output []*model.Project_Account, err error) {

	db := e.db.Model(&model.Table{})

	err = db.Select("project.p_id,project.code,project.projectman_id,project.date_for_start,project.origin_id,accounts.name as projectman_name,task.t_id").
		Joins("left join accounts on accounts.account_id = project.projectman_id").
		Joins("left join task on task.documents_id = project.p_id").
		Where("project.origin_id ='7f7daf49-ccb2-4ee4-9ad4-dec3d7b7bb4f'").
		Order("project.create_time desc").Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByProjectListUser(input *model.Field) (output *model.Project_Account, err error) {

	db := e.db.Raw("select project.p_id,project.code,project.customer_code,project.type,project.p_name,project.customer_id,project.salesman_id,project.serviceman_id,project.projectman_id,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.status,project.order_number,project.inner_id,project.date_for_start,project.date_for_end,project.date_for_pay,project.date_for_delivery,project.date_for_check,project.creater,project.create_time from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id where project.p_id = ?", input.PID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("p_id = ?", input.PID)
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByCaseID(input string) (output *model.Project_Account, err error) {
	value, _ := strconv.Atoi(input)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select project.p_id,project.code,project.customer_code,project.type,project.p_name,project.customer_id,project.salesman_id,project.serviceman_id,project.projectman_id,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.status,project.inner_id,project.date_for_start,project.date_for_end,project.date_for_pay,project.date_for_delivery,project.date_for_check,project.creater,project.create_time from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id where bonita_case_id = ?", value)

	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDBonitaUserID(caseId string,userId string) (output []*model.Tm_Return, err error) {
	value, _ := strconv.Atoi(caseId)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select project.order_number,project.inner_id,task_user.tu_id,task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion,project.p_id,project.code as p_code,project.p_name,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.date_for_start,project.date_for_end,project.machine_finished_number,project.machine_english from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.documents_id=project.p_id left join task_user on task.t_id = task_user.task_id where project.bonita_case_id = ? and task_user.user_id = ? and task_user.status_type_id='98c543c6-944e-4860-b323-166ed5f3920e'", value,userId)

	err = db.Find(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDTaskUserParentcaseID(caseId string,bonita_parentcase_id string) (output *model.Tm_Return, err error) {
	value, _ := strconv.Atoi(caseId)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select task_user.date_for_delivery,task_user.tu_id,tu_accounts.name,task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion,project.p_id,project.code as p_code,project.p_name,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.date_for_start,project.date_for_end,project.machine_finished_number,project.machine_english from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.documents_id=project.p_id left join task_user on task.t_id = task_user.task_id left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id where project.bonita_case_id= ? AND task_user.bonita_parentcase_id= ? ", value,bonita_parentcase_id)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDTaskUserStatus2(caseId string,status_type_id string,status_type_id2 string) (output []*model.Tm_Return, err error) {
	value, _ := strconv.Atoi(caseId)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select task_user.tu_id,tu_accounts.name,task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion,project.p_id,project.code as p_code,project.p_name,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.date_for_start,project.date_for_end,project.machine_finished_number,project.machine_english from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.documents_id=project.p_id left join task_user on task.t_id = task_user.task_id left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id where project.bonita_case_id= ? AND (task_user.status_type_id= ? OR task_user.status_type_id= ?)", value,status_type_id,status_type_id2)

	err = db.Find(&output).Error

	return output, err
}

func (e *entity) GetByCaseIDTaskUserDepartment(caseId string,status_type_id string,dep string) (output []*model.Tm_Return, err error) {
	value, _ := strconv.Atoi(caseId)
	// db := e.db.Model(&model.Table{}).Where("bonita_case_id = ?", value)

	db := e.db.Raw("select task_user.tu_id,tu_accounts.name,task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,task.last_task_code,task.last_task_name,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.last_date_for_estimated_start,task.last_date_for_actual_completion,task.last_date_for_estimated_completion,project.p_id,project.code as p_code,project.p_name,accounts.name as customer_name,serviceman.salesman_name,serviceman.serviceman_name,serviceman.projectman_name,project.date_for_start,project.date_for_end,project.machine_finished_number,project.machine_english from (select projectman.p_id,projectman.salesman_name,projectman.projectman_name,accounts.name as serviceman_name,projectman.customer_id from (select sales.p_id,sales.salesman_name,accounts.name as projectman_name,sales.serviceman_id,sales.customer_id from (select project.p_id,accounts.name as salesman_name,project.projectman_id,project.serviceman_id,project.salesman_id,project.customer_id from project left join accounts on accounts.account_id = project.salesman_id) as sales left join accounts on accounts.account_id = sales.projectman_id) as projectman left join accounts on accounts.account_id = projectman.serviceman_id) as serviceman left join accounts on accounts.account_id = serviceman.customer_id left join project on project.p_id = serviceman.p_id right join (select task.t_id,task.documents_id,task.t_name,task.code,task.remark,task.last_task,lasttask.t_name as last_task_name,lasttask.code as last_task_code,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,lasttask.date_for_estimated_start as last_date_for_estimated_start,lasttask.date_for_actual_completion as last_date_for_actual_completion,lasttask.date_for_estimated_completion as last_date_for_estimated_completion from task left join task as lasttask on lasttask.t_id = task.last_task)as task on task.documents_id=project.p_id left join task_user on task.t_id = task_user.task_id left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id where project.bonita_case_id= ? AND task_user.status_type_id= ? AND task_user_name.dep = ?", value,status_type_id,dep)

	err = db.Find(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("p_id = ?", input.PID).Save(&input).Error

	return err
}
