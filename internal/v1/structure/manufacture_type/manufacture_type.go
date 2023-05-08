package manufacture_type

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	MtID string `gorm:"primary_key;column:mt_id;uuid_generate_v4()type:UUID;" json:"mt_id,omitempty"`
	// 
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	MtID string `json:"mt_id,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	MtID string `json:"mt_id,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 專案編號
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	MtID string `json:"mt_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	Name *string `json:"name,omitempty" form:"name" binding:"omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	ManufactureType []*struct {
		// 編號
		MtID string `json:"mt_id,omitempty"`
		// 
		Name string `json:"name,omitempty"`
		// 
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"manufacture_type"`
	model.OutPage
}

type Updated struct {
	// 編號
	MtID string `json:"mt_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	Name *string `json:"name,omitempty" form:"name" binding:"omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "manufacture_type"
}
