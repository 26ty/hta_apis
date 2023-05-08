package jig_demand

import (
	model "eirc.app/internal/v1/structure/jig_demands"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error
	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_time desc").Find(&output).Error //desc 由大到小

	return amount, output, err
}

func (e *entity) JigDetailListUser(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Joins("Account").Preload("Detail").
		Order("jig_demand.created_time desc").
		Offset(int((input.Page - 1) * input.Limit)).Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByJIDJigDetailListUser(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Joins("Account").Preload("Detail").
		Where("j_id = ?", input.JID).First(&output).Error

	return output, err
}

func (e *entity) GetByUserIDListJD(input *model.Users) (amount int64, output []*model.JD, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("jig_demand.j_id,jig_demand.j_code,jig_demand.creater as salesman_id,s_accounts.name as salesman_name,jig_demand.title,task.t_id,task.create_time,task_user.tu_id,task_user.user_id as task_user_id,tu_accounts.name as task_user_name,task_user.status_type_id,status_type.status,task.remark,labor_hour.h_id").
		Joins("left join accounts as s_accounts on s_accounts.account_id = jig_demand.creater").
		Joins("left join task on task.documents_id = jig_demand.j_id").
		Joins("left join task_user on task_user.task_id = task.t_id").
		Joins("left join accounts as tu_accounts on tu_accounts.account_id = task_user.user_id").
		Joins("left join status_type on status_type.st_id = task_user.status_type_id").
		Joins("left join labor_hour on labor_hour.category = task_user.tu_id").
		Where("task_user.user_id = ?", input.UserID).Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) { //第一括號參數 第二括號回傳
	db := e.db.Model(&model.Table{}).Where("j_id = ?", input.JID)
	//first不會給空值
	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("j_id = ?", input.JID).Delete(&input).Error //model = table的意思

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("j_id = ?", input.JID).Save(&input).Error

	return err
}
