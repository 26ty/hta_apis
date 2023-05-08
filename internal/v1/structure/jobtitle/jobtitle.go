package jobtitle

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	JID string `gorm:"primaryKey;column:j_id;uuid_generate_v4()type:UUID;" json:"j_id,omitempty"`
	// 中文名稱
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// bonita role ID
	BonitaRoleID string `gorm:"column:bonita_role_id;type:text;" json:"bonita_role_id,omitempty"`
	// 
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
}

type Base struct {
	// 編號
	JID string `json:"j_id,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// bonita角色ID
	BonitaRoleID string `json:"bonita_role_id,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}


type Single struct {
	// 編號
	JID string `json:"j_id,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// bonita角色ID
	BonitaRoleID string `json:"bonita_role_id,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}

type Created struct {
	model.GonitaUser
	// 編號
	JID string `json:"j_id,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// bonita角色ID
	BonitaRoleID string `json:"bonita_role_id,omitempty" binding:"required" validate:"required"`
	// 創建者
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}

type Field struct {
	// 編號
	JID string `json:"j_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	// bonita角色ID
	BonitaRoleID string `json:"bonita_role_id,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type List struct {
	Jobtitle []*struct {
		// 編號
		JID string `json:"j_id,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// bonita角色ID
		BonitaRoleID string `json:"bonita_role_id,omitempty"`
		// 創建者
		Creater string `json:"creater,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
	} `json:"jobtitle"`
	model.OutPage
}

type Updated struct {
	model.GonitaUser
	// 編號
	JID string `json:"j_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// bonita角色ID
	BonitaRoleID string `json:"bonita_role_id,omitempty"`
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
	return "jobtitle"
}
