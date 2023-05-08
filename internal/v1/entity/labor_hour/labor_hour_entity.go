package labor_hour

import (
	model "eirc.app/internal/v1/structure/labor_hour"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	if input.Title != nil {
		db.Where("title = ?", *input.Title)
	}

	if input.Content != nil {
		db.Where("content like %?%", input.Content)
	}

	if input.Nature != nil {
		db.Where("nature = ?", *input.Nature)
	}

	if input.TimeForStart != nil {
		db.Where("time_for_start = ?", input.TimeForStart)
	}

	if input.TimeForEnd != nil {
		db.Where("time_for_end = ?", input.TimeForEnd)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("h_id = ?", input.HID)

	err = db.First(&output).Error

	return output, err
}
func (e *entity) GetByUserIdCategoryList(input *model.Field) (amount int64, output []*model.GetUserCategoryLabor, err error) {
	db := e.db.Model(&model.Table{}).
		Select("labor_hour.h_id,labor_hour.content,labor_hour.title,labor_hour.nature,labor_hour.time_for_start,labor_hour.time_for_end,labor_hour.date_for_start,labor_hour.category,labor_hour.laborhour,labor_hour.creater as user_id,c_accounts.name as user_name,task_user.tu_id,countersign_user.cu_id,countersign.cs_id,department.name as d_name,task.t_id,task.t_name,task.code as t_code,project.p_id,project.code as p_code,project.p_name,project.projectman_id as p_projectman_id,customer_demand_task.cd_id as cd_id_task,customer_demand_task.code as cd_code_task,customer_demand_task.projectman_id as cd_projectman_id_task,customer_demand_countersign.cd_id as cd_id_countersign,customer_demand_countersign.code as cd_code_countersign,customer_demand_countersign.projectman_id as cd_projectman_id_countersign").
		Joins("left join accounts as c_accounts on c_accounts.account_id = labor_hour.creater").
		Joins("left join task_user on task_user.tu_id = labor_hour.category").
		Joins("left join task on task.t_id = task_user.task_id").
		Joins("left join project on project.p_id = task.documents_id").
		Joins("left join customer_demand as customer_demand_task on customer_demand_task.cd_id = task.documents_id").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour.category").
		Joins("left join countersign on countersign.cs_id = countersign_user.countersign_id").
		Joins("left join department on department.d_id = countersign.cs_id").
		Joins("left join customer_demand as customer_demand_countersign on customer_demand_countersign.cd_id = countersign.documents_id")

	err = db.Where("labor_hour.creater = ? AND labor_hour.category = ? AND labor_hour.date_for_start = ?", input.UserID,input.Category,input.DateForStart).
			Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByUserIdCategory(input *model.Field) (amount int64, output []*model.GetUserCategoryLabor, err error) {
	db := e.db.Model(&model.Table{}).
		Select("labor_hour.h_id,labor_hour.content,labor_hour.title,labor_hour.nature,labor_hour.time_for_start,labor_hour.time_for_end,labor_hour.date_for_start,labor_hour.category,labor_hour.laborhour,labor_hour.creater as user_id,c_accounts.name as user_name,task_user.tu_id,countersign_user.cu_id,countersign.cs_id,department.name as d_name,task.t_id,task.t_name,task.code as t_code,project.p_id,project.code as p_code,project.p_name,project.projectman_id as p_projectman_id,customer_demand_task.cd_id as cd_id_task,customer_demand_task.code as cd_code_task,customer_demand_task.projectman_id as cd_projectman_id_task,customer_demand_countersign.cd_id as cd_id_countersign,customer_demand_countersign.code as cd_code_countersign,customer_demand_countersign.projectman_id as cd_projectman_id_countersign").
		Joins("left join accounts as c_accounts on c_accounts.account_id = labor_hour.creater").
		Joins("left join task_user on task_user.tu_id = labor_hour.category").
		Joins("left join task on task.t_id = task_user.task_id").
		Joins("left join project on project.p_id = task.documents_id").
		Joins("left join customer_demand as customer_demand_task on customer_demand_task.cd_id = task.documents_id").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour.category").
		Joins("left join countersign on countersign.cs_id = countersign_user.countersign_id").
		Joins("left join department on department.d_id = countersign.cs_id").
		Joins("left join customer_demand as customer_demand_countersign on customer_demand_countersign.cd_id = countersign.documents_id")

	err = db.Where("labor_hour.creater = ? AND labor_hour.category = ?", input.UserID,input.Category).
			Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByUserIdMonthList(input *model.Field_Month) (amount int64, output []*model.GetUserAllLabor, err error) {
	db := e.db.Model(&model.Table{}).
		Select("meeting.origin_id as m_origin_id,labor_hour.category,labor_hour.creater as user_id,c_accounts.name as user_name,task_user.tu_id,countersign_user.cu_id,countersign.cs_id,department.name as d_name,task.t_id,task.origin_id as t_origin_id,task.t_name,task.code as t_code,project.p_id,project.code as p_code,project.p_name,project.projectman_id as p_projectman_id,customer_demand_task.cd_id as cd_id_task,customer_demand_task.code as cd_code_task,customer_demand_task.projectman_id as cd_projectman_id_task,customer_demand_countersign.cd_id as cd_id_countersign,customer_demand_countersign.code as cd_code_countersign,customer_demand_countersign.projectman_id as cd_projectman_id_countersign").
		Joins("left join accounts as c_accounts on c_accounts.account_id = labor_hour.creater").
		Joins("left join task_user on task_user.tu_id = labor_hour.category").
		Joins("left join meeting on meeting.m_id = labor_hour.category").
		Joins("left join task on task.t_id = task_user.task_id").
		Joins("left join project on project.p_id = task.documents_id").
		Joins("left join customer_demand as customer_demand_task on customer_demand_task.cd_id = task.documents_id").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour.category").
		Joins("left join countersign on countersign.cs_id = countersign_user.countersign_id").
		Joins("left join department on department.d_id = countersign.cs_id").
		Joins("left join customer_demand as customer_demand_countersign on customer_demand_countersign.cd_id = countersign.documents_id")

	LastDateOfMonth := input.FirstDate.AddDate(0, 1, -1)
	err = db.Where("labor_hour.creater = ? AND labor_hour.date_for_start >= ? AND labor_hour.date_for_start <= ? ", input.UserID,input.FirstDate,LastDateOfMonth).
			Group("meeting.origin_id,labor_hour.category,labor_hour.creater,c_accounts.name,task_user.tu_id,countersign_user.cu_id,countersign.cs_id,department.name,task.t_id,task.t_name,task.code,project.p_id,project.code,project.p_name,project.projectman_id,customer_demand_task.cd_id,customer_demand_task.code,customer_demand_task.projectman_id,customer_demand_countersign.cd_id,customer_demand_countersign.code,customer_demand_countersign.projectman_id").
			Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByUserIdOneLaborhour(input *model.Field_Month) (output []*model.GetUserOneLabor, err error) {
	db := e.db.Model(&model.Table{})

	LastDateOfMonth := input.FirstDate.AddDate(0, 1, -1)

	err = db.Select("labor_hour.date_for_start,SUM(labor_hour.laborhour) as laborhour").
			Where("labor_hour.creater = ? AND labor_hour.category = ? AND labor_hour.date_for_start >= ? AND labor_hour.date_for_start <= ? ", input.UserID, input.Category,input.FirstDate,LastDateOfMonth).
			Group("labor_hour.category,labor_hour.date_for_start").Find(&output).Error

	return output, err
}

func (e *entity) GetByUserIdOneSumLaborhour(input *model.Field_Month) (output *model.GetUserOneSumLabor, err error) {
	db := e.db.Model(&model.Table{})

	LastDateOfMonth := input.FirstDate.AddDate(0, 1, -1)

	err = db.Select("SUM(labor_hour.laborhour) as sum_of_laborhour").
			Where("labor_hour.creater = ? AND labor_hour.category = ? AND labor_hour.date_for_start >= ? AND labor_hour.date_for_start <= ? ", input.UserID, input.Category,input.FirstDate,LastDateOfMonth).
			Group("labor_hour.category").Find(&output).Error

	return output, err
}

func (e *entity) GetByUserIdMonthSumList(input *model.Field_Month) (output []*model.GetUserAllSumLabor, err error) {
	db := e.db.Model(&model.Table{})

	LastDateOfMonth := input.FirstDate.AddDate(0, 1, -1)

	err = db.Select("date_for_start,SUM(labor_hour.laborhour) as sum_laborhour").
			Where("labor_hour.creater = ? AND labor_hour.date_for_start >= ? AND labor_hour.date_for_start <= ? ", input.UserID,input.FirstDate,LastDateOfMonth).
			Group("date_for_start").Find(&output).Error

	return output, err
}

func (e *entity) GetByUserIdList(input *model.Field) (amount int64, output []*model.LaborHour, err error) {
	db := e.db.Model(&model.Table{}).
		Select("labor_hour.h_id,labor_hour.nature,labor_hour.date_for_start,labor_hour.time_for_start,labor_hour.time_for_end,labor_hour.laborhour,labor_hour.category,labor_hour.title,labor_hour.content,c_accounts.name as creater_name,task.documents_id as task_documents_id,task.origin_id as task_origin_id,project.code as project_code,customer_demand_task.code as customer_demand_task_code,project.machine_finished_number,task_user.user_id as task_user_id,tu_accounts.name as task_user_name,countersign_user.user_id as countersign_user_id,cu_accounts.name as countersign_user_name,countersign.documents_id as countersign_documents_id,customer_demand_countersign.code as customer_demand_countersign_code").
		Joins("left join accounts as c_accounts on c_accounts.account_id = labor_hour.creater").
		Joins("left join task_user on task_user.tu_id = labor_hour.category").
		Joins("left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id").
		Joins("left join task on task.t_id = task_user.task_id").
		Joins("left join project on project.p_id = task.documents_id").
		Joins("left join customer_demand as customer_demand_task on customer_demand_task.cd_id = task.documents_id").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour.category").
		Joins("left join accounts as cu_accounts on cu_accounts.account_id = countersign_user.user_id").
		Joins("left join countersign on countersign.cs_id = countersign_user.countersign_id").
		Joins("left join customer_demand as customer_demand_countersign on customer_demand_countersign.cd_id = countersign.documents_id")

	err = db.Where("labor_hour.creater = ?", input.UserID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByUserIdLaborHourList(input *model.Field) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).
		//Select("labor_hour.h_id,labor_hour.title,labor_hour.content,labor_hour.nature,labor_hour.category,labor_hour.time_for_start,labor_hour.time_for_end,labor_hour.date_for_start,labor_hour.create_time").
		Joins("left join task_user on task_user.tu_id = labor_hour.category").
		Joins("left join accounts on accounts.account_id = task_user.user_id").
		Joins("left join task on task.t_id = task_user.task_id")

	err = db.Where("user_id = ?", input.UserID).Where("t_id = ?", input.TID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByCuIdLaborHourList(input *model.Field) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).
		//Select("labor_hour.h_id,labor_hour.title,labor_hour.content,labor_hour.nature,labor_hour.category,labor_hour.time_for_start,labor_hour.time_for_end,labor_hour.date_for_start,labor_hour.create_time").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour.category").
		Joins("left join accounts on accounts.account_id = countersign_user.user_id")

	err = db.Where("cu_id = ?", input.CuID).Count(&amount).Find(&output).Error

	return amount, output, err
}


func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("h_id = ?", input.HID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("h_id = ?", input.HID).Save(&input).Error

	return err
}
