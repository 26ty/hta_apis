package antivirus_software

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	AsID string `gorm:"primary_key;column:as_id;uuid_generate_v4()type:UUID;" json:"as_id,omitempty"`
	// 專案編號
	ProjectID string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	// 
	AsName string `gorm:"column:as_name;type:text;" json:"as_name,omitempty"`
	// 
	SoftwareNumber string `gorm:"column:software_number;type:text;" json:"software_number,omitempty"`
	// 
	MachineNumber string `gorm:"column:machine_number;type:text;" json:"machine_number,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	// 狀態
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
}

type Base struct {
	// 編號
	AsID string `json:"as_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 
	AsName string `json:"as_name,omitempty"`
	// 
	SoftwareNumber string `json:"software_number,omitempty"`
	// 
	MachineNumber string `json:"machine_number,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
}

type Review struct {
	// 編號
	AsID string `json:"as_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 
	AsName string `json:"as_name,omitempty"`
	// 
	SoftwareNumber string `json:"software_number,omitempty"`
	// 
	MachineNumber string `json:"machine_number,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`

	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string  `json:"bonita_task_name,omitempty"`
}

type Single struct {
	// 編號
	AsID string `json:"as_id,omitempty"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty"`
	// 
	AsName string `json:"as_name,omitempty"`
	// 
	SoftwareNumber string `json:"software_number,omitempty"`
	// 
	MachineNumber string `json:"machine_number,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
}

type Created struct {
	// 專案編號
	ProjectID string `json:"project_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 
	AsName string `json:"as_name,omitempty" binding:"required" validate:"required"`
	// 
	SoftwareNumber string `json:"software_number,omitempty" binding:"required" validate:"required"`
	// 
	MachineNumber string `json:"machine_number,omitempty" binding:"required" validate:"required"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	AsID string `json:"as_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ProjectID string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	// 
	AsName *string `json:"as_name,omitempty" form:"as_name" binding:"omitempty"`
	// 
	SoftwareNumber *string `json:"software_number,omitempty" form:"software_number" binding:"omitempty"`
	// 
	MachineNumber *string `json:"machine_number,omitempty" form:"machine_number" binding:"omitempty"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" form:"bonita_case_id" binding:"omitempty"`
	// 狀態
	Status string `json:"status,omitempty" form:"status" binding:"omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	AntivirusSoftware []*struct {
		// 編號
		AsID string `json:"as_id,omitempty"`
		// 專案編號
		ProjectID string `json:"project_id,omitempty"`
		// 
		AsName string `json:"as_name,omitempty"`
		// 
		SoftwareNumber string `json:"software_number,omitempty"`
		// 
		MachineNumber string `json:"machine_number,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		//[ 7] bonita_case_id 
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		// 狀態
		Status string `json:"status,omitempty"`
	} `json:"antivirus_software"`
	model.OutPage
}

type Updated struct {
	// 編號
	AsID string `json:"as_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ProjectID *string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	// 
	AsName *string `json:"as_name,omitempty" form:"as_name" binding:"omitempty"`
	// 
	SoftwareNumber *string `json:"software_number,omitempty" form:"software_number" binding:"omitempty"`
	// 
	MachineNumber *string `json:"machine_number,omitempty" form:"machine_number" binding:"omitempty"`
	//[ 7] bonita_case_id 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" form:"bonita_case_id" binding:"omitempty"`
	// 狀態
	Status string `json:"status,omitempty" form:"status" binding:"omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "antivirus_software"
}
