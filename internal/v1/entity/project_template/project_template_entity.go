package project_template

import (
	model "eirc.app/internal/v1/structure/project_template"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.PtID != nil {
	// 	db.Where("pt_id = ?", input.PtID)
	// }

	if input.PtName != nil {
		db.Where("pt_name like %?%", *input.PtName)
	}

	if input.PtRemark != nil {
		db.Where("pt_remark = ?", input.PtRemark)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("pt_id = ?", input.PtID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("pt_id = ?", input.PtID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("pt_id = ?", input.PtID).Save(&input).Error

	return err
}
