package countersign

import (
	model "eirc.app/internal/v1/structure/countersign"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Single, err error) {
	db := e.db.Model(&model.Table{}).
	Select("countersign.cs_id,countersign.documents_id,countersign.department_id,countersign.creater,department.bonita_group_id,countersign.create_time").
	Joins("left join department on department.d_id = countersign.department_id")

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("cs_id = ?", input.CsID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("cs_id = ?",input.CsID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("cs_id = ?",input.CsID).Save(&input).Error

	return err
}
