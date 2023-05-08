package labor_hour_modify

import (
	model "eirc.app/internal/v1/structure/labor_hour_modify"
	"strconv"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	if input.Category != nil {
		db.Where("category = ?", input.Category)
	}

	if input.Title != nil {
		db.Where("title = ?", *input.Title)
	}

	if input.Content != nil {
		db.Where("content like %?%", input.Content)
	}

	if input.Nature != nil {
		db.Where("nature = ?", *input.Nature)
	}

	if input.DateForStart != nil {
		db.Where("date_for_start = ?", input.DateForStart)
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
	db := e.db.Model(&model.Table{}).Where("hm_id = ?", input.HmID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByUserIdList(input *model.Field) (amount int64, output []*model.LaborHourModify, err error) {
	db := e.db.Model(&model.Table{}).
		Select("labor_hour_modify.create_time,task.t_id,task.code as t_code,task.t_name,project.p_id,project.p_name,project.projectman_id,pm_accounts.name as projectman_name,labor_hour_modify.bonita_case_id,status_type.status,labor_hour_modify.hm_id,labor_hour_modify.hour_id,labor_hour_modify.nature,labor_hour_modify.date_for_start,labor_hour_modify.time_for_start,labor_hour_modify.time_for_end,labor_hour_modify.laborhour,labor_hour_modify.category,labor_hour_modify.title,labor_hour_modify.content,c_accounts.name as creater_name,task.documents_id as task_documents_id,task.origin_id as task_origin_id,project.code as project_code,customer_demand_task.code as customer_demand_task_code,project.machine_finished_number,task_user.user_id as task_user_id,tu_accounts.name as task_user_name,countersign_user.user_id as countersign_user_id,cu_accounts.name as countersign_user_name,countersign.documents_id as countersign_documents_id,customer_demand_countersign.code as customer_demand_countersign_code").
		Joins("left join accounts as c_accounts on c_accounts.account_id = labor_hour_modify.creater").
		Joins("left join task_user on task_user.tu_id = labor_hour_modify.category").
		Joins("left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id").
		Joins("left join task on task.t_id = task_user.task_id").
		Joins("left join project on project.p_id = task.documents_id").
		Joins("left join accounts as pm_accounts on pm_accounts.account_id = project.projectman_id").
		Joins("left join customer_demand as customer_demand_task on customer_demand_task.cd_id = task.documents_id").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour_modify.category").
		Joins("left join accounts as cu_accounts on cu_accounts.account_id = countersign_user.user_id").
		Joins("left join countersign on countersign.cs_id = countersign_user.countersign_id").
		Joins("left join customer_demand as customer_demand_countersign on customer_demand_countersign.cd_id = countersign.documents_id").
		Joins("left join status_type on status_type.st_id = labor_hour_modify.status_type_id")

	err = db.Where("task_user.user_id = ? OR countersign_user.user_id = ?", input.UserID, input.UserID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByUserIdLaborHourModifyList(input *model.Field) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).
		//Select("labor_hour_modify.hm_id,labor_hour_modify.title,labor_hour_modify.content,labor_hour_modify.nature,labor_hour_modify.category,labor_hour_modify.time_for_start,labor_hour_modify.time_for_end,labor_hour_modify.date_for_start,labor_hour_modify.create_time").
		Joins("left join task_user on task_user.tu_id = labor_hour_modify.category").
		Joins("left join accounts on accounts.account_id = task_user.user_id").
		Joins("left join task on task.t_id = task_user.task_id")

	err = db.Where("user_id = ?", input.UserID).Where("t_id = ?", input.TID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByCuIdLaborHourModifyList(input *model.Field) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).
		//Select("labor_hour_modify.hm_id,labor_hour_modify.title,labor_hour_modify.content,labor_hour_modify.nature,labor_hour_modify.category,labor_hour_modify.time_for_start,labor_hour_modify.time_for_end,labor_hour_modify.date_for_start,labor_hour_modify.create_time").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour_modify.category").
		Joins("left join accounts on accounts.account_id = countersign_user.user_id")

	err = db.Where("cu_id = ?", input.CuID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByCaseID(input string) (output *model.ReviewByDepartment, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})

	err = db.Select("customer_demand.cd_id,customer_demand.code as cd_code,customer_demand.demand_content as cd_name,countersign.cs_id,countersign_user.remark as cs_name,department.name as d_name,project.p_id,project.code as p_code,project.p_name,project.projectman_id,p_accounts.name as projectman_name,task.t_id,task.code as t_code,task.t_name,labor_hour.title,labor_hour.content,labor_hour.nature,labor_hour.date_for_start,labor_hour.time_for_start,labor_hour.time_for_end,labor_hour.laborhour,labor_hour_modify.hm_id,labor_hour_modify.hour_id,labor_hour_modify.category,labor_hour_modify.title as m_title,labor_hour_modify.content as m_content,labor_hour_modify.nature as m_nature,labor_hour_modify.date_for_start as m_date_for_start,labor_hour_modify.time_for_start as m_time_for_start,labor_hour_modify.time_for_end as m_time_for_end,labor_hour_modify.laborhour as m_laborhour,labor_hour_modify.creater,accounts.name as creater_name,labor_hour_modify.create_time").
		Joins("left join accounts on accounts.account_id = labor_hour_modify.creater").
		Joins("left join labor_hour on labor_hour.h_id = labor_hour_modify.hour_id").
		Joins("left join task_user on task_user.tu_id = labor_hour_modify.category").
		Joins("left join task on task.t_id = task_user.task_id").
		Joins("left join project on project.p_id =task.documents_id").
		Joins("left join accounts as p_accounts on p_accounts.account_id = project.projectman_id").
		Joins("left join countersign_user on countersign_user.cu_id = labor_hour_modify.category").
		Joins("left join countersign on countersign.cs_id = countersign_user.countersign_id").
		Joins("left join department on department.d_id = countersign.department_id").
		Joins("left join customer_demand on countersign.documents_id = customer_demand.cd_id").
		Where("labor_hour_modify.bonita_case_id = ?",value).First(&output).Error


	return output, err
}


func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("hm_id = ?", input.HmID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("hm_id = ?", input.HmID).Save(&input).Error

	return err
}
