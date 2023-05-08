package department

import (
	"time"

	model "eirc.app/internal/v1/structure"
	personnel_affiliation_model "eirc.app/internal/v1/structure/personnel_affiliation"

)

type Table struct {
	// 部門表編號
	DID string `gorm:"primaryKey;column:d_id;uuid_generate_v4()type:UUID;" json:"d_id,omitempty"`
	// 部門主管
	Manager string `gorm:"column:manager;type:UUID;" json:"manager,omitempty"`
	// 中文名稱
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 英文名稱
	EngName string `gorm:"column:eng_name;type:TEXT;" json:"eng_name,omitempty"`
	// 路徑
	Introduction string `gorm:"column:introduction;type:TEXT;" json:"introduction,omitempty"`
	// fax
	Fax string `gorm:"column:fax;type:TEXT;" json:"fax,omitempty"`
	// tel
	Tel string `gorm:"column:tel;type:TEXT;" json:"tel,omitempty"`
	// bonita部門ID
	BonitaGroupID string `gorm:"column:bonita_group_id;type:TEXT;" json:"bonita_group_id,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `gorm:"column:bonita_parent_group_id;type:TEXT;" json:"bonita_parent_group_id,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
}

type Base struct {
	// 部門表編號
	DID string `json:"d_id,omitempty"`
	// 部門主管
	Manager string `json:"manager,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 英文名稱
	EngName string `json:"eng_name,omitempty"`
	// 路徑
	Introduction string `json:"introduction,omitempty"`
	// fax
	Fax string `json:"fax,omitempty"`
	// tel
	Tel string `json:"tel,omitempty"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}


type Deparment_Users struct{
	Department []*struct {
		// 部門表編號
		DID string `json:"d_id,omitempty"`
		// 部門主管
		Manager string `json:"manager,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 英文名稱
		EngName string `json:"eng_name,omitempty"`
		// 路徑
		Introduction string `json:"introduction,omitempty"`
		// fax
		Fax string `json:"fax,omitempty"`
		// tel
		Tel string `json:"tel,omitempty"`
		// bonita部門ID
		BonitaGroupID string `json:"bonita_group_id,omitempty"`
		// bonita父部門ID
		BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
		//父部門UUID
		ParentDID string `gorm:"column:parent_d_id;type:UUID;" json:"parent_d_id,omitempty"`
		//父部門名稱
		ParentName string `gorm:"column:parent_name;type:TEXT;" json:"parent_name,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
		Users []*personnel_affiliation_model.Deparment_User
	} `json:"department"`
}

type Deparment_Account struct {
	// 部門表編號
	DID string `json:"d_id,omitempty"`
	// 部門主管
	Manager string `json:"manager,omitempty"`
	// 部門主管名稱
	ManagerName string `json:"manager_name,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 英文名稱
	EngName string `json:"eng_name,omitempty"`
	// 路徑
	Introduction string `json:"introduction,omitempty"`
	// fax
	Fax string `json:"fax,omitempty"`
	// tel
	Tel string `json:"tel,omitempty"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}

type Single struct {
	// 部門表編號
	DID string `json:"d_id,omitempty"`
	// 部門主管
	Manager string `json:"manager,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 英文名稱
	EngName string `json:"eng_name,omitempty"`
	// 路徑
	Introduction string `json:"introduction,omitempty"`
	// fax
	Fax string `json:"fax,omitempty"`
	// tel
	Tel string `json:"tel,omitempty"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
}

type Created struct {
	model.GonitaUser
	// 部門表編號
	DID string `json:"d_id"`
	// 部門主管
	Manager string `json:"manager"`
	// 中文名稱
	Name string `json:"name" binding:"required" validate:"required"`
	// 英文名稱
	EngName string `json:"eng_name"`
	// 路徑
	Introduction string `json:"introduction"`
	// fax
	Fax string `json:"fax"`
	// tel
	Tel string `json:"tel"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`

	//bonita顯示的名稱
	DisplayName string `json:"displayName,omitempty"`
	//bonita父部門
	ParentGroupId string `json:"parent_group_id,omitempty"`
}

type Field struct {
	// 部門表編號
	DID string `json:"d_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 部門主管
	Manager *string `json:"manager,omitempty" form:"manager" binding:"omitempty,uuid4"`
	// 中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 英文名稱
	EngName *string `json:"eng_name,omitempty" form:"eng_name"`
	// 路徑
	Introduction *string `json:"introduction,omitempty" form:"introduction"`
	// fax
	Fax *string `json:"fax,omitempty" form:"fax"`
	// tel
	Tel *string `json:"tel,omitempty" form:"tel"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
	//[ 7] create_time
	CreateTime *time.Time `json:"create_time,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type Deparment_Accounts struct {
	Department []*struct {
		// 部門表編號
		DID string `json:"d_id,omitempty"`
		// 部門主管
		Manager string `json:"manager,omitempty"`
		// 部門主管名稱
		ManagerName string `json:"manager_name,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 英文名稱
		EngName string `json:"eng_name,omitempty"`
		// 路徑
		Introduction string `json:"introduction,omitempty"`
		// fax
		Fax string `json:"fax,omitempty"`
		// tel
		Tel string `json:"tel,omitempty"`
		// bonita部門ID
		BonitaGroupID string `json:"bonita_group_id,omitempty"`
		// bonita父部門ID
		BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
	} `json:"department"`
	model.OutPage
}

type List struct {
	Department []*struct {
		// 部門表編號
		DID string `json:"d_id,omitempty"`
		// 部門主管
		Manager string `json:"manager,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 英文名稱
		EngName string `json:"eng_name,omitempty"`
		// 路徑
		Introduction string `json:"introduction,omitempty"`
		// fax
		Fax string `json:"fax,omitempty"`
		// tel
		Tel string `json:"tel,omitempty"`
		// bonita部門ID
		BonitaGroupID string `json:"bonita_group_id,omitempty"`
		// bonita父部門ID
		BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
	} `json:"department"`
	model.OutPage
}

type Updated struct {
	model.GonitaUser
	// 部門表編號
	DID string `json:"d_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 部門主管
	Manager string `json:"manager,omitempty" binding:"omitempty,uuid4"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 英文名稱
	EngName string `json:"eng_name,omitempty"`
	// 路徑
	Introduction string `json:"introduction,omitempty"`
	// fax
	Fax string `json:"fax,omitempty"`
	// tel
	Tel string `json:"tel,omitempty"`
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty" `

	//bonita顯示的名稱
	DisplayName string `json:"displayName,omitempty"`
	//bonita父部門
	ParentGroupId string `json:"parent_group_id,omitempty"`
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
	return "department"
}
