package project_template

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 範本編號
	PtID string `gorm:"primaryKey;column:pt_id;uuid_generate_v4()type:UUID;" json:"pt_id,omitempty"`
	// 範本代號
	PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
	// 範本名稱
	PtName string `gorm:"column:pt_name;type:TEXT;" json:"pt_name,omitempty"`
	// 範本被註
	PtRemark string `gorm:"column:pt_remark;type:TEXT;" json:"pt_remark,omitempty"`
	// 創建者
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 範本編號
	PtID string `json:"pt_id,omitempty"`
	// 範本代號
	PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
	// 範本名稱
	PtName string `json:"pt_name,omitempty"`
	// 範本被註
	PtRemark string `json:"pt_remark,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 範本編號
	PtID string `json:"pt_id,omitempty"`
	// 範本代號
	PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
	// 範本名稱
	PtName string `json:"pt_name,omitempty"`
	// 範本被註
	PtRemark string `json:"pt_remark,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 範本代號
	PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
	// 範本名稱
	PtName string `json:"pt_name,omitempty" binding:"required" validate:"required"`
	// 範本被註
	PtRemark string `json:"pt_remark,omitempty" binding:"required" validate:"required"`
	// 創建者
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 範本編號
	PtID string `json:"pt_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 範本代號
	PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
	// 範本名稱
	PtName *string `json:"pt_name,omitempty" form:"pt_name"`
	// 範本被註
	PtRemark *string `json:"pt_remark,omitempty" form:"pt_remark"`
	// 創建者
	Creater *string `json:"creater,omitempty" form:"creater"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	ProjectTemplate []*struct {
		// 範本編號
		PtID string `json:"pt_id,omitempty"`
		// 範本代號
		PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
		// 範本名稱
		PtName string `json:"pt_name,omitempty"`
		// 範本被註
		PtRemark string `json:"pt_remark,omitempty"`
		// 創建者
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"project_template"`
	model.OutPage
}

type Updated struct {
	// 範本編號
	PtID string `json:"pt_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 範本代號
	PtCode string `gorm:"column:pt_code;type:TEXT;" json:"pt_code,omitempty"`
	// 範本名稱
	PtName *string `json:"pt_name,omitempty"`
	// 範本被註
	PtRemark *string `json:"pt_remark,omitempty"`
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
	return "project_template"
}
