package quotation

import (
	model "eirc.app/internal/v1/structure/quotations"
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

func (e *entity) QuotationDetailListUser(input *model.Fields) (amount int64, output []*model.QuotationDetail, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Joins("Account").Preload("Detail").
		Order("quotation.created_time desc").
		Offset(int((input.Page - 1) * input.Limit)).Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByQIDQuotationDetailListUser(input *model.Field) (output *model.QuotationDetail, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Joins("Account").Preload("Detail").
		Where("q_id = ?", input.QID).First(&output).Error

	return output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) { //第一括號參數 第二括號回傳
	db := e.db.Model(&model.Table{}).Where("q_id = ?", input.QID)
	//first不會給空值
	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("q_id = ?", input.QID).Delete(&input).Error //model = table的意思

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("q_id = ?", input.QID).Save(&input).Error

	return err
}
