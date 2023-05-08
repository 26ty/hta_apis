package transaction_record

import (
	model "eirc.app/internal/v1/structure/transaction_record"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.Name != nil {
	// 	db.Where("name like %?%", *input.Name)
	// }

	// if input.Path != nil {
	// 	db.Where("path = ?", input.Path)
	// }

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByDocumentIDUserList(input *model.Fields) (amount int64, output []*model.Record_user_list, err error) {
	db := e.db.Model(&model.Table{})

	if input.DocumentID != "" {
		db.Where("document_id = ?", input.DocumentID)
	}

	err = db.Count(&amount).Joins("Accounts").Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("tr_id = ?", input.TrID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("tr_id = ?", input.TrID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("tr_id = ?", input.TrID).Save(&input).Error

	return err
}
