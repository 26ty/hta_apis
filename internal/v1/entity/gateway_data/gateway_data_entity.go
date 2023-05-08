package gateway_data

import (
	model "eirc.app/internal/v1/structure/gateway_data"
	gg_data_demand "eirc.app/internal/v1/structure/gg_data_demand"
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
	db := e.db.Model(&model.Table{}).Where("gd_id = ?", input.GdID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByClassificationTitle(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("classification = ? AND title = ?", input.Classification, input.Title)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByDataDemand(input string) (output *[]gg_data_demand.Review, err error) {

	db := e.db.Raw(input)

	err = db.First(&output).Error
	if err == nil{
		db = e.db.Raw(input)
		err = db.Find(&output).Error
	}


	return output, err
}


func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("gd_id = ?", input.GdID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("gd_id = ?", input.GdID).Save(&input).Error

	return err
}
