package manufacture_user

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	MuID string `gorm:"primary_key;column:mu_id;uuid_generate_v4()type:UUID;" json:"mu_id,omitempty"`
	// 
	ManufactureID string `gorm:"column:manufacture_id;type:UUID;" json:"manufacture_id,omitempty"`
	// 
	UserID string `gorm:"column:user_id;type:UUID;" json:"user_id,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	MuID string `json:"mu_id,omitempty"`
	// 
	ManufactureID string `json:"manufacture_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type ManufactureAccounts struct {
	ManufactureAccount []*ManufactureAccount
}

type ManufactureAccount struct {
	// 編號
	MuID string `json:"mu_id,omitempty"`
	// 
	ManufactureID string `json:"manufacture_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// 
	Email string `json:"email,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	Deps []*Dep
}

type Dep struct {
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	DepartmentName string `json:"department_name,omitempty"`
	//
	JobtitleID string `json:"jobtitle_id,omitempty"`
	//
	JobtitleName string `json:"jobtitle_name,omitempty"`
}

type Single struct {
	// 編號
	MuID string `json:"mu_id,omitempty"`
	// 
	ManufactureID string `json:"manufacture_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 專案編號
	ManufactureID string `json:"manufacture_id,omitempty" binding:"required" validate:"required"`
	// 
	UserID string `json:"user_id,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	MuID string `json:"mu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ManufactureID string `json:"manufacture_id,omitempty" form:"manufacture_id" binding:"omitempty"`
	// 專案編號
	UserID *string `json:"user_id,omitempty" form:"user_id" binding:"omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	ManufactureUser []*struct {
		// 編號
		MuID string `json:"mu_id,omitempty"`
		// 
		ManufactureID string `json:"manufacture_id,omitempty"`
		// 
		UserID string `json:"user_id,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"manufacture_user"`
	model.OutPage
}

type Updated struct {
	// 編號
	MuID string `json:"mu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	ManufactureID *string `json:"manufacture_id,omitempty" form:"manufacture_id" binding:"omitempty"`
	// 專案編號
	UserID *string `json:"user_id,omitempty" form:"user_id" binding:"omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "manufacture_user"
}
