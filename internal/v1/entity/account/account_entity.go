package account

import (
	model "eirc.app/internal/v1/structure/accounts"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Users) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.Dep != nil {
	// 	db.Where("dep = ?", input.Dep)
	// }

	if input.CompanyID != nil {
		db.Where("company_id = ?", input.CompanyID)
	}

	if input.Account != nil {
		db.Where("account = ?", input.Account)
	}

	if input.Name != nil {
		db.Where("name like %?%", *input.Name)
	}


	err = db.Where("accounts.is_deleted = ?", false).Count(&amount).
		// Offset(int((input.Page - 1) * input.Limit)).
		// 	Limit(int(input.Limit)).
		Order("created_at desc").Find(&output).Error

	//db.Count(&count)
	return amount, output, err
}

func (e *entity) AccountNameList(input *model.Users) (amount int64, output []*model.Account_Name, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("accounts.account_id,accounts.bonita_user_id,accounts.bonita_manager_id,accounts.name,department.name as dep_name").
		Joins("left join personnel_affiliation on personnel_affiliation.user_id = accounts.account_id").
		Joins("left join department on department.d_id = personnel_affiliation.department_id").
		Where("accounts.is_deleted = ?", false).
		Count(&amount).
		Order("created_at desc").
		Find(&output).Error

	//db.Count(&count)
	return amount, output, err
}

func (e *entity) AccountNameDepartmentList(input *model.Users) (amount int64, output []*model.Account_Name, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("accounts.account_id,accounts.bonita_user_id,accounts.bonita_manager_id,accounts.name,department.name as dep_name,jobtitle.name as jobtitle_name,p_department.name as p_dep_name").
		Joins("left join personnel_affiliation on personnel_affiliation.user_id = accounts.account_id").
		Joins("left join department on department.d_id = personnel_affiliation.department_id").
		Joins("left join department as p_department on p_department.bonita_group_id = department.bonita_parent_group_id").
		Joins("left join jobtitle on jobtitle.j_id = personnel_affiliation.jobtitle_id").
		Where("accounts.is_deleted = ?", false).
		Count(&amount).
		Find(&output).Error

	//db.Count(&count)
	return amount, output, err
}


func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("account_id = ?", input.AccountID)
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByAccount(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("account = ?", input.Account)
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByBonitaUserID(userID string) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("bonita_user_id = ?", userID)
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Where("account_id = ?", input.AccountID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("account_id = ?", input.AccountID).Save(&input).Error

	return err
}

//為了導入EMAIL用的(暫時寫死)
func (e *entity) UpdatedCsv(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("account = ?", input.Account).Save(&input).Error

	return err
}
