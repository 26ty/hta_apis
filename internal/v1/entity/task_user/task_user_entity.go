package task_user

import (
	model "eirc.app/internal/v1/structure/task_user"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Task_user_Account, err error) {
	db := e.db.Model(&model.Table{})
	// if input.AID != nil {
	// 	db.Where("a_id = ?", input.AID)
	// }

	if input.UserID != nil {
		db.Where("user_id = ?", input.UserID)
	}

	if input.TaskID != nil {
		db.Where("task_id = ?", *input.TaskID)
	}

	err = db.Count(&amount).Joins("Accounts").Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByDocumnetIDListHour(input *model.Field) (amount int64, output []*model.Task_user_Labor_Hour, err error) {
	db := e.db.Model(&model.Table{})
	
	err = db.Select("accounts.account_id,accounts.name,COALESCE(SUM(labor_hour.time_for_end-labor_hour.time_for_start),0) as hour").
	Joins("left join task on task.t_id = task_user.task_id").
	Joins("left join labor_hour on labor_hour.category = task_user.tu_id").
	Joins("left join accounts on accounts.account_id = task_user.user_id").
	Group("account_id,name").Where("documents_id = ?", input.DocumentsID).Count(&amount).Find(&output).Error
	
	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("tu_id = ?", input.TuID)
	// if input.AID != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetName(input *model.Field) (output *model.Task_user_Account, err error) {
	db := e.db.Model(&model.Table{})
	// if input.AID != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.Joins("Accounts").Where("tu_id = ?", input.TuID).First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("tu_id = ?", input.TuID).Save(&input).Error

	return err
}

func (e *entity) Updated_Bonita(input *model.Updated_Bonita) (err error) {
	err = e.db.Model(&model.Table{}).Where("tu_id = ?", input.TuID).
	Update("bonita_parentcase_id", input.BonitaParentCaseID).Error

	return err
}
