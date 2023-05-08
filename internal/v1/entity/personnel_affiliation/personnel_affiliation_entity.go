package personnel_affiliation

import (
	model "eirc.app/internal/v1/structure/personnel_affiliation"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.PaID != nil {
	// 	db.Where("pa_id = ?", input.PaID)
	// }


	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("pa_id = ?", input.PaID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByDepartmentID(input *model.Field) (output []*model.Deparment_User, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("accounts.account_id,accounts.name,accounts.email,accounts.bonita_user_id,department.d_id,department.name as d_name,parent_department.name as parent_name,parent_department.d_id as parent_d_id").
		Joins("left join accounts on accounts.account_id = personnel_affiliation.user_id").
		Joins("left join department on department.d_id = personnel_affiliation.department_id").
		Joins("left join department as parent_department on parent_department.bonita_group_id = department.bonita_parent_group_id").
		Where("personnel_affiliation.department_id = ? ",input.DepartmentID).
		Find(&output).Error

	return output, err
}

func (e *entity) GetByParentDepartmentID(bonita_group_id string) (output []*model.Deparment_User, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("accounts.account_id,accounts.name,accounts.email,accounts.bonita_user_id,department.d_id,department.name as d_name,parent_department.name as parent_name,parent_department.d_id as parent_d_id").
		Joins("left join accounts on accounts.account_id = personnel_affiliation.user_id").
		Joins("left join department on department.d_id = personnel_affiliation.department_id").
		Joins("left join department as parent_department on parent_department.bonita_group_id = department.bonita_parent_group_id ").
		Where("department.bonita_group_id = ? OR department.bonita_parent_group_id = ? ",bonita_group_id,bonita_group_id).
		Find(&output).Error

	return output, err
}

func (e *entity) GetByUserID(input *model.Field) (output []*model.Affiliation_Account, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("personnel_affiliation.pa_id,personnel_affiliation.user_id,accounts.name,accounts.bonita_user_id,personnel_affiliation.department_id,department.name as department_name,department.bonita_group_id,personnel_affiliation.jobtitle_id,jobtitle.name as jobtitle_name,jobtitle.bonita_role_id,personnel_affiliation.creater,personnel_affiliation.create_time").
		Joins("left join accounts on accounts.account_id = personnel_affiliation.user_id").
		Joins("left join department on department.d_id = personnel_affiliation.department_id").
		Joins("left join jobtitle on jobtitle.j_id = personnel_affiliation.jobtitle_id").
		Where("personnel_affiliation.user_id = ? ",input.UserID).
		Order("personnel_affiliation.create_time desc").Find(&output).Error

	return output, err
}


func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("pa_id = ?", input.PaID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("pa_id = ?", input.PaID).Save(&input).Error

	return err
}
