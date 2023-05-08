package attendee

import (
	model "eirc.app/internal/v1/structure/attendee"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	// if input.AID != nil {
	// 	db.Where("a_id = ?", input.AID)
	// }

	if input.MeetID != nil {
		db.Where("meet_id = ?", input.MeetID)
	}

	if input.UserID != nil {
		db.Where("user_id = ?", *input.UserID)
	}

	if input.Chairman != nil {
		db.Where("chairman = ?", input.Chairman)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("a_id = ?", input.AID)
	// if input.AID != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("a_id = ?", input.AID).Save(&input).Error

	return err
}
