package personnel_affiliation

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	PaID string `gorm:"primaryKey;column:pa_id;uuid_generate_v4()type:UUID;" json:"pa_id,omitempty"`
	// 
	UserID string `gorm:"column:user_id;type:UUID;" json:"user_id,omitempty"`
	// 
	DepartmentID string `gorm:"column:department_id;type:UUID;" json:"department_id,omitempty"`
	// 
	JobtitleID string `gorm:"column:jobtitle_id;type:UUID;" json:"jobtitle_id,omitempty"`
	// 
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
}

type Affiliation_Account struct {
	// 編號
	PaID string `json:"pa_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	Name string `json:"name,omitempty"`
	BonitaUserID string `json:"bonita_user_id"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	DepartmentName string `json:"department_name,omitempty"`
	BonitaGroupID string `json:"bonita_group_id"`
	// 
	JobtitleID string `json:"jobtitle_id,omitempty"`
	JobtitleName string `json:"jobtitle_name,omitempty"`
	BonitaRoleID string `json:"bonita_role_id,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`

}

//部門人員篩選
type Deparment_User struct{
	// 編號
	AccountID string `json:"account_id,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	//[ 6] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaUserID string `json:"bonita_user_id,omitempty"`
	//部門編號
	DID string `json:"d_id,omitempty"`
	//部門名稱
	DName string `json:"d_name,omitempty"`
	//父部門UUID
	ParentDID string `json:"parent_d_id,omitempty"`
	//父部門名稱
	ParentName string `json:"parent_name,omitempty"`
}

type Base struct {
	// 編號
	PaID string `json:"pa_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	JobtitleID string `json:"jobtitle_id,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}


type Single struct {
	// 編號
	PaID string `json:"pa_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	JobtitleID string `json:"jobtitle_id,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}

type Created struct {
	model.GonitaUser
	// 編號
	PaID string `json:"pa_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty" binding:"required" validate:"required"`
	// 
	DepartmentID string `json:"department_id,omitempty" binding:"required" validate:"required"`
	// 
	JobtitleID string `json:"jobtitle_id,omitempty" binding:"required" validate:"required"`
	// 
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`

	// bonita人員ID
	BonitaUserID string `json:"bonita_user_id" binding:"required" validate:"required"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id" binding:"required" validate:"required"`
	// bonita角色ID
	BonitaRoleID string `json:"bonita_role_id,omitempty" binding:"required" validate:"required"`

}

type Field struct {
	// 編號
	PaID string `json:"pa_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	JobtitleID string `json:"jobtitle_id,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}


type List struct {
	PersonnelAffiliation []*struct {
		// 編號
		PaID string `json:"pa_id,omitempty"`
		// 
		UserID string `json:"user_id,omitempty"`
		// 
		DepartmentID string `json:"department_id,omitempty"`
		// 
		JobtitleID string `json:"jobtitle_id,omitempty"`
		// 
		Creater string `json:"creater,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
	} `json:"personnel_affiliation"`
	model.OutPage
}

type Updated struct {
	model.GonitaUser
	// 編號
	PaID string `json:"pa_id,omitempty"`
	// 
	UserID string `json:"user_id,omitempty"`
	// 
	DepartmentID string `json:"department_id,omitempty"`
	// 
	JobtitleID string `json:"jobtitle_id,omitempty"`

	// bonita人員ID
	BonitaUserID string `json:"bonita_user_id"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id"`
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
	return "personnel_affiliation"
}
