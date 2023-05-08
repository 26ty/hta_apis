package antivirus_software

import (
	model "eirc.app/internal/v1/structure/antivirus_software"
	"strconv"
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
	db := e.db.Model(&model.Table{}).Where("as_id = ?", input.AsID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByCaseID(input string) (output *model.Review, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})

	err = db.Select("*").Where("antivirus_software.bonita_case_id = ?",value).First(&output).Error

	return output, err
}

func (e *entity) GetByPIDList(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("project_id = ?", input.ProjectID)

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("as_id = ?",input.AsID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("as_id = ?",input.AsID).Save(&input).Error

	return err
}
