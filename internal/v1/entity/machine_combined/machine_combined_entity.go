package machine_combined

import (
	model "eirc.app/internal/v1/structure/machine_combined"
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

func (e *entity) MachineCombinedListLast(input *model.Fields) (amount int64, output []*model.Machine_Combined_Last, err error) {
	db := e.db.Model(&model.Table{}).
		Select("machine_combined.mc_id,machine_combined.project_id,machine_combined.mc_code,machine_combined.mc_number,machine_combined.mc_finished,machine_combined.last_mc,machine_combined.create_time,mcTable.mc_code as last_mc_code").
		Joins("left join machine_combined as mcTable on mcTable.mc_id = machine_combined.last_mc")

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByPIDMachineCombinedListLast(input *model.Fields) (amount int64, output []*model.Machine_Combined_Last, err error) {
	db := e.db.Model(&model.Table{}).
		Select("machine_combined.mc_id,machine_combined.project_id,machine_combined.mc_code,machine_combined.mc_number,machine_combined.mc_finished,machine_combined.last_mc,machine_combined.create_time,mcTable.mc_code as last_mc_code").
		Joins("left join machine_combined as mcTable on mcTable.mc_id = machine_combined.last_mc").
		Where("machine_combined.project_id = ?", input.ProjectID)

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("mc_id = ?", input.McID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("mc_id = ?",input.McID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("mc_id = ?",input.McID).Save(&input).Error

	return err
}
