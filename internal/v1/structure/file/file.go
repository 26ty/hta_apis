package file

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 檔案編號
	FID string `gorm:"primaryKey;column:f_id;uuid_generate_v4()type:UUID;" json:"f_id,omitempty"`
	// 任務人員
	DocumentsID string `gorm:"column:documents_id;type:UUID;" json:"documents_id,omitempty"`
	// 中文名稱
	FileName string `gorm:"column:file_name;type:TEXT;" json:"file_name,omitempty"`
	// 路徑
	FilePath string `gorm:"column:file_path;type:TEXT;" json:"file_path,omitempty"`
	//副檔名
	FileExtension string `gorm:"column:file_extension;type:TEXT;" json:"file_extension,omitempty"`
	//下載位置
	DownloadUrl string `gorm:"column:download_url;type:TEXT;" json:"download_url,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	// 創建人員
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
	// 檔案大小
	Size string `gorm:"column:size;type:TEXT;" json:"size,omitempty"`
}

type FilebydocumentId struct {
	//編號
	FID string `json:"f_id,omitempty"`
	//編號
	DocumentsID string `json:"documents_id,omitempty"`
	//附件名稱
	FileName string `json:"file_name,omitempty"`
	//附件路徑
	FilePath string `json:"file_path,omitempty"`
	//附件路徑
	FileExtension string `json:"file_extension,omitempty"`
	//下載位置
	DownloadUrl string `json:"download_url,omitempty"`
	//創建者id
	Creater string `json:"creater,omitempty"`
	//創建者名稱
	CreaterName string `json:"creater_name,omitempty"`
	//時間
	Create_time string `json:"create_time,omitempty"`
	// 檔案大小
	Size string `json:"size,omitempty"`
}

type Base struct {
	// 檔案編號
	FID string `json:"f_id,omitempty"`
	// 任務人員
	DocumentsID string `json:"documents_id,omitempty"`
	// 中文名稱
	FileName string `json:"file_name,omitempty"`
	// 路徑
	FilePath string `json:"file_path,omitempty"`
	//副檔名
	FileExtension string `json:"file_extension,omitempty"`
	//下載位置
	DownloadUrl string `json:"download_url,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 創建人員
	Creater string `json:"creater,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 檔案大小
	Size string `json:"size,omitempty"`
}

type Single struct {
	// 檔案編號
	FID string `json:"f_id,omitempty"`
	// 任務人員
	DocumentsID string `json:"documents_id,omitempty"`
	// 中文名稱
	FileName string `json:"file_name,omitempty"`
	// 路徑
	FilePath string `json:"file_path,omitempty"`
	//副檔名
	FileExtension string `json:"file_extension,omitempty"`
	//下載位置
	DownloadUrl string `json:"download_url,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 創建人員
	Creater string `json:"creater,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 檔案大小
	Size string `json:"size,omitempty"`
}

type Created struct {
	// 任務人員
	DocumentsID string `json:"documents_id,omitempty" binding:"required" validate:"required"`
	// 中文名稱
	FileName string `json:"file_name,omitempty" binding:"required" validate:"required"`
	// 路徑
	FilePath string `json:"file_path,omitempty"`
	// 副檔名
	FileExtension string `json:"file_extension,omitempty" binding:"required" validate:"required"`
	Base64    string `json:"base64,omitempty" binding:"required,base64" validate:"required"`
	// 創建人員
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
	//下載位置
	DownloadUrl string `json:"download_url,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 檔案大小
	Size string `json:"size,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 檔案編號
	FID string `json:"f_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 任務人員
	DocumentsID string `json:"documents_id,omitempty" form:"documents_id"  binding:"omitempty,uuid4"`
	// 中文名稱
	FileName string `json:"file_name,omitempty" form:"file_name"`
	// 路徑
	FilePath string `json:"file_path,omitempty" form:"file_path"`
	// 副檔名
	FileExtension string `json:"file_extension,omitempty" form:"file_extension"`
	//下載位置
	DownloadUrl string `json:"download_url,omitempty" form:"download_url"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 檔案大小
	Size string `json:"size,omitempty" form:"size"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type List struct {
	File []*struct {
		// 檔案編號
		FID string `json:"f_id,omitempty"`
		// 任務人員
		DocumentsID string `json:"documents_id,omitempty"`
		// 中文名稱
		FileName string `json:"file_name,omitempty"`
		// 路徑
		FilePath string `json:"file_path,omitempty"`
		//副檔名
		FileExtension string `json:"file_extension,omitempty"`
		//下載位置
		DownloadUrl string `json:"download_url,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 創建人員
		Creater string `json:"creater,omitempty"`
		// 檔案大小
		Size string `json:"size,omitempty"`
	} `json:"file"`
	model.OutPage
}

type FilebydocumentIds struct {
	File []*struct {
		//編號
		FID string `json:"f_id,omitempty"`
		//編號
		DocumentsID string `json:"documents_id,omitempty"`
		//附件名稱
		FileName string `json:"file_name,omitempty"`
		//附件路徑
		FilePath string `json:"file_path,omitempty"`
		//附件路徑
		FileExtension string `json:"file_extension,omitempty"`
		//下載位置
		DownloadUrl string `json:"download_url,omitempty"`
		//創建者id
		Creater string `json:"creater,omitempty"`
		//創建者名稱
		CreaterName string `json:"creater_name,omitempty"`
		//時間
		Create_time string `json:"create_time,omitempty"`
		// 檔案大小
		Size string `json:"size,omitempty"`
	} `json:"file"`
	model.OutTotal
}

type Updated struct {
	// 檔案編號
	FID string `json:"f_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 任務人員
	DocumentsID string `json:"documents_id,omitempty" binding:"omitempty,uuid4"`
	// 中文名稱
	FileName string `json:"file_name,omitempty"`
	// 路徑
	FilePath string `json:"file_path,omitempty"`
	//副檔名
	FileExtension string `json:"file_extension,omitempty"`
	//下載位置
	DownloadUrl string `json:"download_url,omitempty"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// 檔案大小
	Size string `json:"size,omitempty"`
}

type Login struct {
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password" binding:"required" validate:"required"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "file"
}
