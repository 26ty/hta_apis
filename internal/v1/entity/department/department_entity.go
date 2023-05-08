package department

import (
	model "eirc.app/internal/v1/structure/department"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.DID != nil {
	// 	db.Where("d_id = ?", input.DID)
	// }

	if input.Manager != nil {
		db.Where("manager = ?", *input.Manager)
	}

	if input.Name != nil {
		db.Where("name like %?%", *input.Name)
	}

	if input.EngName != nil {
		db.Where("eng_name like %?%", input.EngName)
	}

	if input.Fax != nil {
		db.Where("fax like %?%", input.Fax)
	}

	if input.Tel != nil {
		db.Where("tel like %?%", input.Tel)
	}

	if input.Introduction != nil {
		db.Where("introduction like %?%", input.Introduction)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) A1Department(input *model.Field) (output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}) 

	err = db.
		Select("department.name,department.d_id,department.bonita_group_id,department.bonita_parent_group_id,department.create_time,department.manager,department.eng_name,department.introduction,department.fax,department.tel,parent_department.name as parent_name,parent_department.d_id as parent_d_id").
		Joins("left join department as parent_department on parent_department.bonita_group_id = department.bonita_parent_group_id").
		Where("department.d_id = '8711fe0c-62aa-427e-94b7-e0379ae0e908' OR department.d_id = '3fb4e9b0-544d-407f-a24e-15ead22706f2' OR department.d_id = 'dd8bd64c-050f-4bfa-993c-a5aeb9f91535' OR department.d_id = '9af6353c-6e53-4508-b097-aec80154d387' OR department.d_id = 'fc4ac1b8-fee5-41cc-ac20-0c55ee3792bb'").
		Order("department.create_time desc").Find(&output).Error

	return output, err
}

func (e *entity) AllDepartment(input *model.Field) (output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}) 

	err = db.
		Select("department.name,department.d_id,department.bonita_group_id,department.bonita_parent_group_id,department.create_time,department.manager,department.eng_name,department.introduction,department.fax,department.tel,parent_department.name as parent_name,parent_department.d_id as parent_d_id").
		Joins("left join department as parent_department on parent_department.bonita_group_id = department.bonita_parent_group_id").
		Order("department.create_time desc").Find(&output).Error

	return output, err
}

func (e *entity) DepartmentAccountList(input *model.Users) (amount int64, output []*model.Deparment_Account, err error) {
	db := e.db.Model(&model.Table{})

	err = db.
		Select("department.d_id,department.name,department.eng_name,department.introduction,department.tel,department.fax,department.manager,department.create_time,department.bonita_group_id,department.bonita_parent_group_id,accounts.name as manager_name").
		Joins("left join accounts on accounts.account_id = department.manager").
		Order("department.create_time desc").Count(&amount).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("d_id = ?", input.DID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("d_id = ?", input.DID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("d_id = ?", input.DID).Save(&input).Error

	return err
}
