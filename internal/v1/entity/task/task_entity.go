package task

import (
	model "eirc.app/internal/v1/structure/task"
)

func (e *entity) Created(input *model.Create_Table) (err error) {
	err = e.db.Model(&model.Create_Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.PID != nil {
	// 	db.Where("p_id = ?", input.PID)
	// }

	// if input.DocumentsID != nil {
	// 	db.Where("documents_id = ?", input.DocumentsID)
	// }

	if input.TName != nil {
		db.Where("t_name like %?%", input.TName)
	}

	if input.Remark != nil {
		db.Where("remark = ?", input.Remark)
	}

	if input.Landmark != nil {
		db.Where("landmark = ?", input.Landmark)
	}

	if input.File != nil {
		db.Where("file = ?", input.File)
	}

	if input.Rank != nil {
		db.Where("rank = ?", input.Rank)
	}

	if input.LastTask != nil {
		db.Where("last_task = ?", input.LastTask)
	}

	if input.DateForEstimatedStart != nil {
		db.Where("date_for_estimated_start = ?", input.DateForEstimatedStart)
	}

	if input.DateForActualCompletion != nil {
		db.Where("date_for_actual_completion = ?", input.DateForActualCompletion)
	}

	if input.DateForEstimatedCompletion != nil {
		db.Where("date_for_estimated_completion = ?", input.DateForEstimatedCompletion)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByIDListTaskHour(input *model.Users) (amount int64, output []*model.Task_Account_Labor_Hour, err error) {
	db := e.db.Model(&model.Table{}).
		Select("accounts.account_id,accounts.name,t_name,(labor_hour.time_for_end-labor_hour.time_for_start) as hour").
		Joins("left join task_user on task_user.task_id = task.t_id").
		Joins("left join labor_hour on labor_hour.category = task_user.tu_id").
		Joins("left join accounts on accounts.account_id = task_user.user_id")

	err = db.Where("task_user.user_id = ?", input.UserID).
		Where("task.documents_id = ?", input.DocumentsID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByTaskListUser(input *model.Users) (amount int64, output []*model.Task_Account, err error) {
	db := e.db.Model(&model.Table{})
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.
		Select("status_type.status,accounts.bonita_user_id,task.t_id,task.code,task.quantity,task.documents_id,task.t_name,task.remark,task.landmark,task.file,task.last_task,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.origin_id,task_user.tu_id,task_user.principal,accounts.account_id,accounts.name,task.create_time").
		Joins("left join task_user on task.t_id=task_user.task_id").
		Joins("left join status_type on task_user.status_type_id=status_type.st_id").
		Joins("left join accounts on accounts.account_id=task_user.user_id").
		Where("t_id = ?", input.TID).
		Order("task.create_time desc").Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetTaskListHourByUserID(input *model.Field) (amount int64, output []*model.Task_Hour_User, err error) {
	db := e.db.Model(&model.Table{}).
		Select("sum(labor_hour.time_for_end-labor_hour.time_for_start) as hour,accounts.name,accounts.account_id").
		Joins("left join task_user on task_user.task_id = task.t_id").
		Joins("left join labor_hour on labor_hour.category = task_user.tu_id").
		Joins("left join accounts on accounts.account_id = task_user.user_id").
		Joins("left join project on project.p_id = task.documents_id").
		Where("task.documents_id = ?", input.DocumentsID).
		Group("user_id,name,account_id")

	// db := e.db.Raw("select SUM(labor_hour.time_for_end-labor_hour.time_for_start) as hour,accounts.name from task inner join task_user on task_user.task_id = task.t_id inner join labor_hour on labor_hour.category = task_user.tu_id inner join accounts on accounts.account_id = task_user.user_id inner join project on project.p_id = task.documents_id").
	// 	Where("documents_id = ?", input.DocumentsID).
	// 	Group("task_user.user_id,accounts.name")

	err = db.Count(&amount).Find(&output).Error

	//db.Count(&count)
	return amount, output, err
}

//資源右邊
func (e *entity) GetByTaskListHourDocumentsAndUserID(input *model.Field) (amount int64, output []*model.Task_Hour_User, err error) {
	db := e.db.Model(&model.Table{}).
		Select("sum(labor_hour.time_for_end-labor_hour.time_for_start) as hour,accounts.name,accounts.account_id,task.t_name").
		Joins("left join task_user on task_user.task_id = task.t_id").
		Joins("left join labor_hour on labor_hour.category = task_user.tu_id").
		Joins("left join accounts on accounts.account_id = task_user.user_id").
		Joins("left join project on project.p_id = task.documents_id").
		Where("task.documents_id = ?", input.DocumentsID).
		Where("accounts.account_id = ?", input.AccountID).
		Group("t_name,name,category,account_id")

	err = db.Count(&amount).Find(&output).Error

	//db.Count(&count)
	return amount, output, err
}

func (e *entity) TaskListUser(input *model.Users) (amount int64, output []*model.Task_Account, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("status_type.status,accounts.bonita_user_id,task.t_id,task.code,task.rank,task.hierarchy,task.quantity,task.documents_id,task.t_name,task.remark,task.landmark,task.file,task.last_task,task.date_for_estimated_start,task.date_for_actual_completion,task.date_for_estimated_completion,task.origin_id,task_user.tu_id,accounts.account_id,accounts.name,task.create_time").
		Joins("left join task_user on task.t_id=task_user.task_id").
		Joins("left join status_type on task_user.status_type_id=status_type.st_id").
		Joins("left join accounts on accounts.account_id=task_user.user_id").
		Joins("left join project on project.p_id=task.documents_id").
		Where("task_user.principal=true").
		Where("task.documents_id = ?", input.DocumentsID).
		Order("task.create_time desc").Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByOriginIDAndUserID(input *model.Users) (amount int64, output []*model.Task_OriginId, err error) {
	db := e.db.Model(&model.Table{}).
		Select("task.landmark,task.todo_status,todo_type.name as t_type,project.p_id,project.code as p_code,project.p_name,task.code,task.quantity,project.p_name,project.code,task.t_id,task.documents_id,task.t_name,task.remark,task.default_labor_hour,task.origin_id,task.date_for_estimated_start,task.date_for_estimated_completion,task.date_for_actual_completion,task_user.tu_id,task_user.principal,accounts.account_id,accounts.name").
		Joins("left join task_user on task.t_id = task_user.task_id").
		Joins("left join accounts on task_user.user_id = accounts.account_id").
		Joins("left join project on task.documents_id = project.p_id").
		Joins("left join todo_type on task.todo_type_id = todo_type.tt_id").
		Where("task.origin_id = ?", input.OriginID).
		Where("task_user.user_id = ?", input.UserID)

	err = db.Count(&amount).
		Order("task.create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByTIDTaskListUser(input *model.Fields) (amount int64, output []*model.Task_User_Account, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("task.t_id,task.code,task_user.tu_id,task_user.principal,accounts.account_id,accounts.name,task.create_time").
		Joins("left join task_user on task.t_id=task_user.task_id").
		Joins("left join accounts on accounts.account_id=task_user.user_id").
		Joins("left join project on project.p_id=task.documents_id").
		Where("t_id = ?", input.TID).
		Order("task.create_time desc").Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

//專案範本
func (e *entity) GetByDocumentIDTaskListLast(input *model.Fields) (amount int64, output []*model.Task_Template_Last, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("task.t_id,task.code,task.t_name,task.documents_id,task.landmark,task.date_for_estimated_start,task.date_for_estimated_completion,task.last_task,tTable.t_name as last_t_name,task.default_date,task.default_labor_hour,task.create_time").
		Joins("left join task as tTable on tTable.t_id = task.last_task").
		Joins("left join project_template on project_template.pt_id = task.documents_id").
		Where("project_template.pt_id = ?", input.PtID).
		Order("task.create_time").Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByDocumentIDTaskList(input *model.Field) (amount int64, output []*model.Task_Template, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("task.t_id,task.code,task.t_name,task.date_for_estimated_start,task.date_for_estimated_completion,task.landmark,task.default_date,task.default_labor_hour,task.last_task,lastTaskTable.code as last_task_code,lastTaskTable.t_name as last_task_name").
		Joins("left join project on project.p_id=task.documents_id").
		Joins("left join task as lastTaskTable on lastTaskTable.t_id = task.last_task").
		Where("task.documents_id = ?", input.DocumentsID).
		Order("task.create_time").Count(&amount).Find(&output).Error

	return amount, output, err
}


func (e *entity) GetByIDTaskBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error) {

	db := e.db.Model(&model.Table{})

	err = db.Select("task.t_id,task_user.user_id,task_user.principal,task.documents_id,accounts.name as task_user_name,accounts.bonita_user_id as task_bonita_id").
		Joins("left join task_user on task_user.task_id = task.t_id").
		Joins("left join accounts on accounts.account_id = task_user.user_id").
		Joins("left join project on project.p_id = task.documents_id").
		Where("task.documents_id = ?", input.DocumentsID).
		Order("task.create_time desc").Find(&output).Error

	//db.Count(&count)
	return output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("t_id = ?", input.TID)
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("t_id = ?", input.TID).Save(&input).Error

	return err
}
