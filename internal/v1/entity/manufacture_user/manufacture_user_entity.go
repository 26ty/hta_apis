package manufacture_user

import (
	model "eirc.app/internal/v1/structure/manufacture_user"
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

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("mu_id = ?", input.MuID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByManufactureID(input *model.Field) (output []*model.ManufactureAccount, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("manufacture_user.mu_id,manufacture_user.manufacture_id,manufacture_user.user_id,accounts.name,accounts.email,manufacture_user.create_time").
	Joins("left join accounts on accounts.account_id = manufacture_user.user_id").
	Where("manufacture_id = ?", input.ManufactureID).Find(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("mu_id = ?",input.MuID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("mu_id = ?",input.MuID).Save(&input).Error

	return err
}
