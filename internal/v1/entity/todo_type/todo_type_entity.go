package todo_type

import (
	model "eirc.app/internal/v1/structure/todo_type"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.TtID != nil {
	// 	db.Where("tt_id = ?", input.TtID)
	// }

	// if input.Name != nil {
	// 	db.Where("name = ?", input.Name)
	// }

	// if input.UserID != nil {
	// 	db.Where("user_id = ?", input.UserID)
	// }

	// if input.CreateTime != nil {
	// 	db.Where("create_time = ?", input.CreateTime)
	// }

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error
	return amount, output, err
}


func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("tt_id = ?", input.TtID)
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByUserID(input *model.Field) (output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("user_id = ?", input.UserID).
	Order("create_time")
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.Find(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("tt_id = ?", input.TtID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("tt_id = ?", input.TtID).Save(&input).Error

	return err
}
