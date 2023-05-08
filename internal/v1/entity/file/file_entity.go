package file

import (
	model "eirc.app/internal/v1/structure/file"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	// if input.BucketName != "" {
	// 	db.Where("bucket_name like %?%", input.BucketName)
	// }

	// if len(input.S3ID) != 0 {
	// 	db.Where("s3_id = ?", input.S3ID)
	// }

	err = db.Where("file.is_deleted = false").Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{})
	db.Where("f_id = ?", input.FID).Where("file.is_deleted = false")

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByDocumentID(input *model.Field) (amount int64, output []*model.FilebydocumentId, err error) {
	db := e.db.Model(&model.Table{})
	db.Select("file.size,file.f_id,file.download_url,file.documents_id,file.file_name,file.file_path,file.file_extension,file.creater,accounts.name as creater_name,file.create_time").
		Joins("left join accounts on accounts.account_id = file.creater").
		Joins("left join project on project.p_id = file.documents_id").
		Joins("left join customer_demand on customer_demand.cd_id = file.documents_id").
		Joins("left join task on task.t_id = file.documents_id").
		Joins("left join countersign on countersign.cs_id = file.documents_id").
		Where("file.documents_id = ? OR task.documents_id = ? OR countersign.documents_id = ?", input.DocumentsID,input.DocumentsID,input.DocumentsID)

	err = db.Where("file.is_deleted = false").Count(&amount).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByDocumentIDUserID(input *model.Users) (amount int64, output []*model.FilebydocumentId, err error) {
	db := e.db.Model(&model.Table{})
	db.Select("file.size,file.f_id,file.download_url,file.documents_id,file.file_name,file.file_path,file.file_extension,file.creater,accounts.name as creater_name,file.create_time").
		Joins("left join accounts on accounts.account_id = file.creater").
		Joins("left join project on project.p_id = file.documents_id").
		Joins("left join customer_demand on customer_demand.cd_id = file.documents_id").
		Joins("left join task on task.t_id = file.documents_id").
		Joins("left join countersign on countersign.cs_id = file.documents_id").
		Where("file.documents_id = ? AND file.creater = ?", input.DocumentsID,input.UserID)

	err = db.Where("file.is_deleted = false").Count(&amount).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) Updated(input *model.Table) (err error) {
	db := e.db.Model(&model.Table{})
	db.Where("f_id = ?", input.FID)

	err = db.Save(&input).Error

	return err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).
		Where("f_id = ?", input.FID).
		Delete(&input).Error

	return err
}