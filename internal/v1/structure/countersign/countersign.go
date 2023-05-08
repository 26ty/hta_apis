package countersign

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	CsID string `gorm:"primary_key;column:cs_id;uuid_generate_v4()type:UUID;" json:"cs_id,omitempty"`
	// 專案編號
	DocumentsID string `gorm:"column:documents_id;type:UUID;" json:"documents_id,omitempty"`
	// 
	DepartmentID string `gorm:"column:department_id;type:UUID;" json:"department_id,omitempty"`
	// 
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	CsID string `json:"cs_id,omitempty"`
	// 專案編號
	DocumentsID string `json:"documents_id,omitempty"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	CsID string `json:"cs_id,omitempty"`
	// 專案編號
	DocumentsID string `json:"documents_id,omitempty"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 專案編號
	DocumentsID string `json:"documents_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 
	DepartmentID string `json:"department_id,omitempty" binding:"required" validate:"required"`
	// 
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	CsID string `json:"cs_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	DocumentsID *string `json:"documents_id,omitempty" form:"documents_id" binding:"omitempty,uuid4"`
	// 
	DepartmentID *string `json:"department_id,omitempty" form:"department_id" binding:"omitempty,uuid4"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Countersign []*struct {
		// 編號
		CsID string `json:"cs_id,omitempty"`
		// 專案編號
		DocumentsID string `json:"documents_id,omitempty"`
		// 
		DepartmentID string `json:"department_id,omitempty"`
		// 
		Creater string `json:"creater,omitempty"`
		// bonita部門ID
		BonitaGroupID string `json:"bonita_group_id,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"countersign"`
	model.OutPage
}

type Updated struct {
	// 編號
	CsID string `json:"cs_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	DocumentsID *string `json:"documents_id,omitempty" form:"documents_id" binding:"omitempty,uuid4"`
	// 
	DepartmentID *string `json:"department_id,omitempty" form:"department_id" binding:"omitempty,uuid4"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "countersign"
}
