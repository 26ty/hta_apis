package countersign_user

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	CuID string `gorm:"primary_key;column:cu_id;uuid_generate_v4()type:UUID;" json:"cu_id,omitempty"`
	// 專案編號
	CountersignId string `gorm:"column:countersign_id;type:UUID;" json:"countersign_id,omitempty"`
	// 
	UserId string `gorm:"column:user_id;type:UUID;" json:"user_id,omitempty"`
	// 
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	//
	DateForEstimatedCompletion time.Time `gorm:"column:date_for_estimated_completion;type:TIMESTAMP;" json:"date_for_estimated_completion,omitempty"`
	//
	DateForCompletion time.Time `gorm:"column:date_for_completion;type:TIMESTAMP;" json:"date_for_completion,omitempty"`
	// 
	Remark string `gorm:"column:remark;type:text;" json:"remark,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	//[ 7] bonita_parentcase_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaParentCaseID string `gorm:"column:bonita_parentcase_id;type:text;" json:"bonita_parentcase_id,omitempty"`
	//
	DateForEstimatedCompletionEmployee time.Time `gorm:"column:date_for_estimated_completion_employee;type:TIMESTAMP;" json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `gorm:"column:date_for_completion_employee;type:TIMESTAMP;" json:"date_for_completion_employee,omitempty"`

}

type CountersignUser_Account struct {
	// 編號
	CuID string `json:"cu_id,omitempty"`
	// 編號
	CsID string `json:"cs_id,omitempty"`
	// 專案編號
	CountersignName string `json:"countersign_name,omitempty"`
	// 
	UserId string `json:"user_id,omitempty"`
	// 部門表編號
	DID string `json:"d_id,omitempty"`
	// 
	DName string `json:"d_name,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	//父部門UUID
	ParentDID string `json:"parent_d_id,omitempty"`
	//父部門名稱
	ParentName string `json:"parent_name,omitempty"`
	// bonita父部門ID
	BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
	// bonita_user_id
	BonitaUserID string `json:"bonita_user_id,omitempty" `
	// bonita部門ID
	BonitaGroupID string `json:"bonita_group_id,omitempty"`
	//
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//
	DateForCompletion time.Time `json:"date_for_completion,omitempty"`
	// 
	Remark string `json:"remark,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`

}

type Base struct {
	// 編號
	CuID string `json:"cu_id,omitempty"`
	// 專案編號
	CountersignId string `json:"countersign_id,omitempty"`
	// 
	UserId string `json:"user_id,omitempty"`
	//
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//
	DateForCompletion time.Time `json:"date_for_completion,omitempty"`
	// 
	Remark string `json:"remark,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
}

type Single struct {
	// 編號
	CuID string `json:"cu_id,omitempty"`
	// 專案編號
	CountersignId string `json:"countersign_id,omitempty"`
	// 
	UserId string `json:"user_id,omitempty"`
	// 
	Creater string `json:"creater,omitempty"`
	//
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//
	DateForCompletion time.Time `json:"date_for_completion,omitempty"`
	// 
	Remark string `json:"remark,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
}

type Created struct {
	// 專案編號
	CountersignId string `json:"countersign_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 
	UserId string `json:"user_id,omitempty" binding:"required" validate:"required"`
	//
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//
	DateForCompletion time.Time `json:"date_for_completion,omitempty"`
	// 
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
	// 
	Remark string `json:"remark,omitempty"`
	//
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
}

type Field struct {
	// 編號
	CuID string `json:"cu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	CountersignId string `json:"countersign_id,omitempty" form:"countersign_id" binding:"omitempty,uuid4"`
	// 
	UserId string `json:"user_id,omitempty" form:"user_id" binding:"omitempty,uuid4"`
	//
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty" form:"date_for_estimated_completion" binding:"omitempty"`
	//
	DateForCompletion time.Time `json:"date_for_completion,omitempty" form:"date_for_completion" binding:"omitempty"`
	// 
	Remark string `json:"remark,omitempty" form:"remark" binding:"omitempty"`
	//
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type Documents struct {
	Field
	// 專案編號
	DocumentsID string `json:"documents_id,omitempty" form:"documents_id" binding:"omitempty,uuid4"`
}

type CountersignUser_Accounts struct {
	CountersignUser []*struct {
		// 編號
		CuID string `json:"cu_id,omitempty"`
		// 編號
		CsID string `json:"cs_id,omitempty"`
		// 
		UserId string `json:"user_id,omitempty"`
		// 專案編號
		CountersignName string `json:"countersign_name,omitempty"`
		// 部門表編號
		DID string `json:"d_id,omitempty"`
		// 
		DName string `json:"d_name,omitempty"`
		// 
		Name string `json:"name,omitempty"`
		//父部門UUID
		ParentDID string `json:"parent_d_id,omitempty"`
		//父部門名稱
		ParentName string `json:"parent_name,omitempty"`
		// bonita父部門ID
		BonitaParentGroupID string `json:"bonita_parent_group_id,omitempty"`
		// bonita_user_id
		BonitaUserID string `json:"bonita_user_id,omitempty" `
		// bonita部門ID
		BonitaGroupID string `json:"bonita_group_id,omitempty"`
		//
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//
		DateForCompletion time.Time `json:"date_for_completion,omitempty"`
		// 
		Remark string `json:"remark,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		//
		DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
		//
		DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
	} `json:"countersign_user"`
	model.OutTotal
}

type List struct {
	CountersignUser []*struct {
		// 編號
		CuID string `json:"cu_id,omitempty"`
		// 專案編號
		CountersignId string `json:"countersign_id,omitempty"`
		// 
		UserId string `json:"user_id,omitempty"`
		// 
		Creater string `json:"creater,omitempty"`
		//
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//
		DateForCompletion time.Time `json:"date_for_completion,omitempty"`
		// 
		Remark string `json:"remark,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		//
		DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
		//
		DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
	} `json:"countersign_user"`
	model.OutPage
}

type Updated struct {
	// 編號
	CuID string `json:"cu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 專案編號
	CountersignId *string `json:"countersign_id,omitempty" form:"countersign_id" binding:"omitempty,uuid4"`
	// 
	UserId *string `json:"user_id,omitempty" form:"user_id" binding:"omitempty,uuid4"`
	//
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty" form:"date_for_estimated_completion" binding:"omitempty"`
	//
	DateForCompletion time.Time `json:"date_for_completion,omitempty" form:"date_for_completion" binding:"omitempty"`
	// 
	Remark string `json:"remark,omitempty" form:"remark" binding:"omitempty"`
	//
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`
}

type Updated_Bonita struct {
	// 編號
	CuID string `json:"cu_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	BonitaParentCaseID string `json:"bonita_parentcase_id,omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "countersign_user"
}
