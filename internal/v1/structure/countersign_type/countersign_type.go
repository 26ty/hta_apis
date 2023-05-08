package countersign_type

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	CtID string `gorm:"primary_key;column:ct_id;uuid_generate_v4()type:UUID;" json:"ct_id,omitempty"`
	// 專案編號
	CountersignName string `gorm:"column:countersign_name;type:UUID;" json:"countersign_name,omitempty"`
	// 
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	CtID string `json:"ct_id,omitempty"`
	// 專案編號
	CountersignName string `json:"countersign_name,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	CtID string `json:"ct_id,omitempty"`
	// 專案編號
	CountersignName string `json:"countersign_name,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 專案編號
	CountersignName string `json:"countersign_name,omitempty" binding:"required" validate:"required"`
	// 
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	CtID string `json:"ct_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	CountersignName *string `json:"countersign_name,omitempty" form:"countersign_name" binding:"omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	CountersignType []*struct {
		// 編號
		CtID string `json:"ct_id,omitempty"`
		// 專案編號
		CountersignName string `json:"countersign_name,omitempty"`
		// 
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"countersign_type"`
	model.OutPage
}

type Updated struct {
	// 編號
	CtID string `json:"ct_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	CountersignName *string `json:"countersign_name,omitempty" form:"countersign_name" binding:"omitempty"`

}

type Token struct {
}

func (a *Table) TableName() string {
	return "countersign_type"
}
