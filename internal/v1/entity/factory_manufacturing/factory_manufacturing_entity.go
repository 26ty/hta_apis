package factory_manufacturing

import (
	model "eirc.app/internal/v1/structure/factory_manufacturings"
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

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) { //第一括號參數 第二括號回傳
	db := e.db.Model(&model.Table{}).Where("fm_id = ?", input.FmID)
	//first不會給空值
	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("fm_id = ?", input.FmID).Delete(&input).Error //model = table的意思

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("fm_id = ?", input.FmID).Save(&input).Error

	return err
}
