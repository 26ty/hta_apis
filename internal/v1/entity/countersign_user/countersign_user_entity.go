package countersign_user

import (
	model "eirc.app/internal/v1/structure/countersign_user"
	countersign_model "eirc.app/internal/v1/structure/countersign"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByIDCountersignUserListUser(input *model.Documents) (amount int64, output []*model.CountersignUser_Account, err error) {
	db := e.db.Model(&countersign_model.Table{}).
			Select("countersign.cs_id,department.d_id,department.name as d_name,department.bonita_group_id,parent_department.d_id as parent_d_id,parent_department.name as parent_name,department.bonita_parent_group_id,countersign_user.cu_id,countersign_user.user_id,accounts.name,accounts.bonita_user_id,countersign_user.date_for_estimated_completion,countersign_user.date_for_completion,countersign_user.remark,countersign_user.create_time,countersign_user.date_for_estimated_completion_employee,countersign_user.date_for_completion_employee").
			Joins("left join department on countersign.department_id = department.d_id").
			Joins("left join department as parent_department on parent_department.bonita_group_id = department.bonita_parent_group_id ").
			Joins("left join countersign_user on countersign_user.countersign_id = countersign.cs_id").
			Joins("left join accounts on countersign_user.user_id = accounts.account_id")

	if input.DocumentsID != "" {
		db.Where("countersign.documents_id = ?", input.DocumentsID)
	}
	
	if input.CountersignId != "" {
		db.Where("countersign_user.countersign_id = ?", input.CountersignId)
	}

	if input.CuID != "" {
		db.Where("countersign_user.cu_id = ?", input.CuID)
	}

	err = db.Count(&amount).Find(&output).Error

	return amount, output, err
}


func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("cu_id = ?", input.CuID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("cu_id = ?",input.CuID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("cu_id = ?",input.CuID).Save(&input).Error

	return err
}
