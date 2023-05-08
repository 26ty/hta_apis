package task_user

import (
	"time"
	"eirc.app/internal/v1/structure/accounts"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	TuID string `gorm:"primary_key;column:tu_id;uuid_generate_v4()type:UUID;" json:"tu_id,omitempty"`
	// 任務編號
	TaskID string `gorm:"column:task_id;type:UUID;" json:"task_id,omitempty"`
	// 負責人
	UserID string `gorm:"column:user_id;type:UUID;" json:"user_id,omitempty"`
	// 主要負責人
	Principal bool `gorm:"column:principal;type:BOOLEAN;" json:"principal,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	// 送審日
	DateForDelivery time.Time `gorm:"column:date_for_delivery;type:TIMESTAMP;" json:"date_for_delivery,omitempty"`
	// 狀態
	StatusTypeID string `gorm:"column:status_type_id;type:UUID;" json:"status_type_id,omitempty"`
	//[ 7] bonita_parentcase_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaParentCaseID string `gorm:"column:bonita_parentcase_id;type:text;" json:"bonita_parentcase_id,omitempty"`

	Accounts accounts.Table `gorm:"foreignkey:user_id;references:account_id"`

}

type Task_user_Labor_Hour struct {
	// 負責人編號
	AccountID string `json:"account_id,omitempty"`
	// 負責人姓名
	Name string `json:"name,omitempty"`
	// 負責人姓名
	Hour int `json:"hour,omitempty"`
}

type Task_user_Account struct {
	// 編號
	TuID string `json:"tu_id,omitempty"`
	// 任務編號
	TaskID string `json:"task_id,omitempty"`
	// 負責人
	UserID string `json:"user_id,omitempty"`
	// 主要負責人
	Principal bool `json:"principal,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 負責人姓名
	Accounts__name string `json:"name,omitempty"`
}

type Base struct {
	// 編號
	TuID string `json:"tu_id,omitempty"`
	// 任務編號
	TaskID string `json:"task_id,omitempty"`
	// 負責人
	UserID string `json:"user_id,omitempty"`
	// 主要負責人
	Principal bool `json:"principal,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 送審日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Single struct {
	// 編號
	TuID string `json:"tu_id,omitempty"`
	// 任務編號
	TaskID string `json:"task_id,omitempty"`
	// 負責人
	UserID string `json:"user_id,omitempty"`
	// 主要負責人
	Principal bool `json:"principal,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 送審日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Created struct {
	// 任務編號
	TaskID string `json:"task_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 負責人
	UserID string `json:"user_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 主要負責人
	Principal bool `json:"principal,omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty"`
}

type Field struct {
	// 編號
	TuID string `json:"tu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 任務編號
	TaskID *string `json:"task_id,omitempty" form:"task_id" binding:"omitempty,uuid4"`
	// 負責人
	UserID *string `json:"user_id,omitempty" form:"user_id" binding:"omitempty,uuid4"`
	// 主要負責人
	Principal *bool `json:"principal,omitempty" form:"principal" binding:"omitempty"`
	// 編號
	DocumentsID string `json:"documents_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
}

type Fields struct {
	Field
	model.InPage
}

type Created_List struct {
	Task_user []*Created `json:"task_user"`
}

type Updated_List struct {
	Task_user []*Updated `json:"task_user"`
}

type List struct {
	Task_user []*struct {
		// 編號
		TuID string `json:"tu_id,omitempty"`
		// 任務編號
		TaskID string `json:"task_id,omitempty"`
		// 負責人
		UserID string `json:"user_id,omitempty"`
		// 主要負責人
		Principal bool `json:"principal,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 負責人姓名
		Name string `json:"name"`
	} `json:"task_user"`
	model.OutPage
}

type Task_user_Labor_Hours struct {
	Task_user []*struct {
		// 負責人編號
		AccountID string `json:"account_id,omitempty"`
		// 負責人姓名
		Name string `json:"name,omitempty"`
		// 負責人姓名
		Hour int `json:"hour,omitempty"`
	} `json:"task_user"`
	model.OutTotal
}

type Updated struct {
	// 編號
	TuID string `json:"tu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 任務編號
	TaskID *string `json:"task_id,omitempty" binding:"omitempty,uuid4"`
	// 負責人
	UserID *string `json:"user_id,omitempty" binding:"omitempty,uuid4"`
	// 主要負責人
	Principal *bool `json:"principal,omitempty" binding:"omitempty"`
}

type Updated_Bonita struct {
	// 編號
	TuID string `json:"tu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	BonitaParentCaseID string `json:"bonita_parentcase_id,omitempty"`
}

type Updated_Review struct {
	// 編號
	TuID string `json:"tu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 送審日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" binding:"omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" binding:"omitempty"`
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
	return "task_user"
}
