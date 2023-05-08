package machine_combined

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	McID string `gorm:"primary_key;column:mc_id;uuid_generate_v4()type:UUID;" json:"mc_id,omitempty"`
	// 專案編號
	ProjectID string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	// 
	McCode string `gorm:"column:mc_code;type:text;" json:"mc_code,omitempty"`
	// 
	McNumber int `gorm:"column:mc_number;type:integer;" json:"mc_number,omitempty"`
	// 
	McFinished string `gorm:"column:mc_finished;type:text;" json:"mc_finished,omitempty"`
	// 
	LastMc string `gorm:"column:last_mc;type:UUID;" json:"last_mc,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Machine_Combined_Last struct {
	// 編號
	McID string `json:"mc_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 
	McCode string `json:"mc_code,omitempty"`
	// 
	McNumber int `json:"mc_number,omitempty"`
	// 
	McFinished string `json:"mc_finished,omitempty"`
	// 
	LastMc string `json:"last_mc,omitempty"`
	// 
	LastMcCode string `json:"last_mc_code,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Base struct {
	// 編號
	McID string `json:"mc_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 
	McCode string `json:"mc_code,omitempty"`
	// 
	McNumber int `json:"mc_number,omitempty"`
	// 
	McFinished string `json:"mc_finished,omitempty"`
	// 
	LastMc string `json:"last_mc,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	McID string `json:"mc_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 
	McCode string `json:"mc_code,omitempty"`
	// 
	McNumber int `json:"mc_number,omitempty"`
	// 
	McFinished string `json:"mc_finished,omitempty"`
	// 
	LastMc string `json:"last_mc,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 專案編號
	ProjectID string `json:"project_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 
	McCode string `json:"mc_code,omitempty" binding:"required" validate:"required"`
	// 
	McNumber int `json:"mc_number,omitempty" binding:"required" validate:"required"`
	// 
	McFinished string `json:"mc_finished,omitempty" binding:"required" validate:"required"`
	// 
	LastMc string `json:"last_mc,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	McID string `json:"mc_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	// 
	McCode *string `json:"mc_code,omitempty" form:"mc_code" binding:"omitempty"`
	// 
	McNumber *int `json:"mc_number,omitempty" form:"mc_number" binding:"omitempty"`
	// 
	McFinished *string `json:"mc_finished,omitempty" form:"mc_finished" binding:"omitempty"`
	// 
	LastMc *string `json:"last_mc,omitempty" form:"last_mc" binding:"omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type Machine_Combined_Lasts struct {
	MachineCombined []*struct {
		// 編號
		McID string `json:"mc_id,omitempty"`
		// 專案編號
		ProjectID string `json:"project_id,omitempty"`
		// 
		McCode string `json:"mc_code,omitempty"`
		// 
		McNumber int `json:"mc_number,omitempty"`
		// 
		McFinished string `json:"mc_finished,omitempty"`
		// 
		LastMc string `json:"last_mc,omitempty"`
		// 
		LastMcCode string `json:"last_mc_code,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"machine_combined"`
	model.OutPage
}

type List struct {
	MachineCombined []*struct {
		// 編號
		McID string `json:"mc_id,omitempty"`
		// 專案編號
		ProjectID string `json:"project_id,omitempty"`
		// 
		McCode string `json:"mc_code,omitempty"`
		// 
		McNumber int `json:"mc_number,omitempty"`
		// 
		McFinished string `json:"mc_finished,omitempty"`
		// 
		LastMc string `json:"last_mc,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"machine_combined"`
	model.OutPage
}

type Updated struct {
	// 編號
	McID string `json:"mc_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ProjectID *string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	// 
	McCode *string `json:"mc_code,omitempty" form:"mc_code" binding:"omitempty"`
	// 
	McNumber *int `json:"mc_number,omitempty" form:"mc_number" binding:"omitempty"`
	// 
	McFinished *string `json:"mc_finished,omitempty" form:"mc_finished" binding:"omitempty"`
	// 
	LastMc *string `json:"last_mc,omitempty" form:"last_mc" binding:"omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "machine_combined"
}
