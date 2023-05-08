package plug_in

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	PiID string `gorm:"primary_key;column:pi_id;uuid_generate_v4()type:UUID;" json:"pi_id,omitempty"`
	// 專案編號
	ProjectID string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	// 內容
	Content string `gorm:"column:content;type:text;" json:"content,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	PiID string `json:"pi_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	PiID string `json:"pi_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 專案編號
	ProjectID string `json:"project_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 內容
	Content string `json:"content,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	PiID string `json:"pi_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	// 內容
	Content *string `json:"content,omitempty" form:"content" binding:"omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	PlugIn []*struct {
		// 編號
		PiID string `json:"pi_id,omitempty"`
		// 專案編號
		ProjectID string `json:"project_id,omitempty"`
		// 內容
		Content string `json:"content,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"plug_in"`
	model.OutPage
}

type Updated struct {
	// 編號
	PiID string `json:"pi_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ProjectID *string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	// 內容
	Content *string `json:"content,omitempty" form:"content" binding:"omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "plug_in"
}
