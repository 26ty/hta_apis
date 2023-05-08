package gift_application

import (
	"strconv"

	model "eirc.app/internal/v1/structure/gift_applications"
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

func (e *entity) GiftDetailListUser(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Joins("Account").Joins("Account_a").Joins("Account_attn").Preload("Detail").
		Order("gift_application.created_time desc").
		Offset(int((input.Page - 1) * input.Limit)).Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) { //第一括號參數 第二括號回傳
	db := e.db.Model(&model.Table{}).Where("g_id = ?", input.GID)
	//first不會給空值
	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByGIDGiftDetailListUser(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Joins("Account").Joins("Account_a").Joins("Account_attn").Preload("Detail").
		Where("g_id = ?", input.GID).First(&output).Error

	return output, err
}

func (e *entity) GetByCaseID(input string) (output *model.Review, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})
	//first不會給空值
	err = db.Joins("Account").Joins("Account_a").Joins("Account_attn").Where("bonita_case_id = ?", value).First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("g_id = ?", input.GID).Delete(&input).Error //model = table的意思

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("g_id = ?", input.GID).Save(&input).Error

	return err
}
