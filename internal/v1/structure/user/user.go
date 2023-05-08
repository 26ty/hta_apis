package user

import (
	"time"

	model "eirc.app/internal/v1/structure"
)
type Table struct {
	//[ 0] u_id                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	UID string `gorm:"primary_key;column:u_id;uuid_generate_v4()type:UUID;" json:"u_id,omitempty"`
	//[ 1] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	//[ 2] auth_id                                        UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	AuthID string `gorm:"column:auth_id;type:UUID;" json:"auth_id,omitempty"`
	//[ 3] dep                                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	Dep string `gorm:"column:dep;type:UUID;" json:"dep,omitempty"`
	//[ 4] phone                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Phone string `gorm:"column:phone;type:TEXT;" json:"phone,omitempty"`
	//[ 5] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Email string `gorm:"column:email;type:TEXT;" json:"email,omitempty"`
	//[ 6] status                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
	//[ 7] create_time                                    TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [now()]
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
}

type Base struct {
	// 編號
	UID string `json:"u_id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty"`
	// 特殊權限
	AuthID string `json:"auth_id,omitempty"`
	// 所屬部門
	Dep string `json:"dep,omitempty"`
	// 電話
	Phone string `json:"phone,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 帳號狀態
	Status string `json:"status,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	UID string `json:"u_id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty"`
	// 特殊權限
	AuthID string `json:"auth_id,omitempty"`
	// 所屬部門
	Dep string `json:"dep,omitempty"`
	// 電話
	Phone string `json:"phone,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 帳號狀態
	Status string `json:"status,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 特殊權限
	AuthID string `json:"auth_id" binding:"required" validate:"required"`
	// 姓名
	Name string `json:"name" binding:"required" validate:"required"`
	// 所屬部門
	Dep string `json:"dep" binding:"required" validate:"required"`
	// 電話
	Phone string `json:"phone" binding:"required" validate:"required"`
	// 電子郵件
	Email string `json:"email" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	UID string `json:"u_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 任務人員
	AuthID *string `json:"auth_id,omitempty" form:"auth_id" binding:"omitempty,uuid4"`
	// 姓名
	Name *string `json:"name,omitempty" form:"name"`
	// 所屬部門
	Dep *string `json:"dep,omitempty" form:"dep"`
	// 電話
	Phone *string `json:"phone,omitempty" form:"phone"`
	// 電子郵件
	Email *string `json:"email,omitempty" form:"email"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	User []*struct {
		// 編號
		UID string `json:"u_id,omitempty"`
		// 姓名
		Name string `json:"name,omitempty"`
		// 特殊權限
		AuthID string `json:"auth_id,omitempty"`
		// 所屬部門
		Dep string `json:"dep,omitempty"`
		// 電話
		Phone string `json:"phone,omitempty"`
		// 電子郵件
		Email string `json:"email,omitempty"`
		// 帳號狀態
		Status string `json:"status,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"user"`
	model.OutPage
}

type Updated struct {
	// 編號
	UID string `json:"u_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 特殊權限
	AuthID string `json:"auth_id" binding:"required" validate:"required"`
	// 姓名
	Name string `json:"name" binding:"required" validate:"required"`
	// 所屬部門
	Dep string `json:"dep" binding:"required" validate:"required"`
	// 電話
	Phone string `json:"phone" binding:"required" validate:"required"`
	// 電子郵件
	Email string `json:"email" binding:"required" validate:"required"`
	// 帳號狀態
	Status bool `json:"status" binding:"required" validate:"required"`
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
	return "user"
}
