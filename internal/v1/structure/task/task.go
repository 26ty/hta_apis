package task

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//[ 0] t_id                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	TID string `gorm:"primary_key;column:t_id;uuid_generate_v4()type:UUID;" json:"t_id,omitempty"`
	//[ 1] documents_id                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	DocumentsID string `gorm:"column:documents_id;type:UUID;" json:"documents_id,omitempty"`
	//[ 2] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	TName string `gorm:"column:t_name;type:TEXT;" json:"t_name,omitempty"`
	//[ 3] remark                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//[ 4] landmark                                       BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Landmark bool `gorm:"column:landmark;type:BOOLEAN;" json:"landmark,omitempty"`
	//[ 5] file                                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	File bool `gorm:"column:file;type:BOOLEAN;" json:"file,omitempty"`
	//[ 6] create_time                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [now()]
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
	//[ 7] rank                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Rank int `gorm:"column:rank;type:INT4;" json:"rank,omitempty"`
	//[ 8] last_task                                      UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	LastTask string `gorm:"column:last_task;type:UUID;" json:"last_task,omitempty"`
	//[ 9] date_for_estimated_start                       DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForEstimatedStart time.Time `gorm:"column:date_for_estimated_start;type:DATE;" json:"date_for_estimated_start,omitempty"`
	//[10] date_for_actual_completion                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForActualCompletion time.Time `gorm:"column:date_for_actual_completion;type:DATE;" json:"date_for_actual_completion,omitempty"`
	//[11] date_for_ estimated _completion                DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForEstimatedCompletion time.Time `gorm:"column:date_for_estimated_completion;type:DATE;" json:"date_for_estimated_completion,omitempty"`
	//[12] origin_id                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `gorm:"column:origin_id;type:UUID;" json:"origin_id,omitempty"`
	//[13] code                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	Code string `gorm:"column:code;type:UUID;" json:"code,omitempty"`
	//[ 7] quantity
	Quantity int `gorm:"column:quantity;type:INT4;" json:"quantity,omitempty"`
	//[ 7] default_date
	DefaultDate int `gorm:"column:default_date;type:INT4;" json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `gorm:"column:default_labor_hour;type:INT4;" json:"default_labor_hour,omitempty"`
	//[ 7] hierarchy
	Hierarchy int `gorm:"column:hierarchy;type:INT4;" json:"hierarchy,omitempty"`
	//[12] todo_type_id                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	TodoTypeID string `gorm:"column:todo_type_id;type:UUID;" json:"todo_type_id,omitempty"`
	//[ 5] todo_status                                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	TodoStatus bool `gorm:"column:todo_status;type:BOOLEAN;" json:"todo_status,omitempty"`

}

type Create_Table struct {
	//[ 0] t_id                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	TID string `gorm:"primary_key;column:t_id;uuid_generate_v4()type:UUID;" json:"t_id,omitempty"`
	//[ 1] documents_id                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	DocumentsID string `gorm:"column:documents_id;type:UUID;" json:"documents_id,omitempty"`
	//[ 2] name                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	TName string `gorm:"column:t_name;type:TEXT;" json:"t_name,omitempty"`
	//[ 3] remark                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//[ 4] landmark                                       BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Landmark bool `gorm:"column:landmark;type:BOOLEAN;" json:"landmark,omitempty"`
	//[ 5] file                                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	File bool `gorm:"column:file;type:BOOLEAN;" json:"file,omitempty"`
	//[ 6] create_time                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [now()]
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
	// //[ 7] rank(資料庫預存程序自動新增)                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	// Rank int `gorm:"column:rank;type:INT4;" json:"rank,omitempty"`
	//[ 8] last_task                                      UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	LastTask string `gorm:"column:last_task;type:UUID;" json:"last_task,omitempty"`
	//[ 9] date_for_estimated_start                       DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForEstimatedStart time.Time `gorm:"column:date_for_estimated_start;type:DATE;" json:"date_for_estimated_start,omitempty"`
	//[10] date_for_actual_completion                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForActualCompletion time.Time `gorm:"column:date_for_actual_completion;type:DATE;" json:"date_for_actual_completion,omitempty"`
	//[11] date_for_ estimated _completion                DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForEstimatedCompletion time.Time `gorm:"column:date_for_estimated_completion;type:DATE;" json:"date_for_estimated_completion,omitempty"`
	//[12] origin_id                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `gorm:"column:origin_id;type:UUID;" json:"origin_id,omitempty"`
	// //[13] code(資料庫預存程序自動新增)                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	// Code string `gorm:"column:code;type:UUID;" json:"code,omitempty"`
	//[ 7] quantity
	Quantity int `gorm:"column:quantity;type:INT4;" json:"quantity,omitempty"`
	//[ 7] default_date
	DefaultDate int `gorm:"column:default_date;type:INT4;" json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `gorm:"column:default_labor_hour;type:INT4;" json:"default_labor_hour,omitempty"`
	//[ 7] hierarchy
	Hierarchy int `gorm:"column:hierarchy;type:INT4;" json:"hierarchy,omitempty"`
	//[12] todo_type_id                                   UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	TodoTypeID string `gorm:"column:todo_type_id;type:UUID;" json:"todo_type_id,omitempty"`
	//[ 5] todo_status                                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	TodoStatus bool `gorm:"column:todo_status;type:BOOLEAN;" json:"todo_status,omitempty"`

}

type Task_Account_Labor_Hour struct {
	//名字
	TName string `json:"t_name,omitempty"`
	// 人員編號
	AccountID string `json:"account_id,omitempty"`
	// 負責人姓名
	Name string `json:"name,omitempty"`
	// 工時
	Hour int `json:"hour,omitempty"`
}

type Task_Account struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	// code
	Code string `json:"code,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//里程碑
	Landmark bool `json:"landmark,omitempty"`
	//附件
	File bool `json:"file,omitempty"`
	//流水號
	Rank int `json:"rank,omitempty"`
	//[ 7] quantity
	Quantity int `json:"quantity,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	// 任務負責人編號
	TuID string `json:"tu_id,omitempty"`
	//是否主要負責人
	Principal bool `json:"principal,omitempty"`
	// 人員編號
	AccountID string `json:"account_id,omitempty"`
	// 負責人姓名
	Name string `json:"name,omitempty"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//實際完成日
	DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	//[ 6] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaUserID string `json:"bonita_user_id,omitempty"`
	//todo_status
	TodoStatus bool `json:"todo_status,omitempty"`
	//status
	Status string `json:"status,omitempty"`
}

type Task_User_Account struct {
	//編號
	TID string `json:"t_id,omitempty"`
	// code
	Code string `json:"code,omitempty"`
	// 任務負責人編號
	TuID string `json:"tu_id,omitempty"`
	//是否主要負責人
	Principal bool `json:"principal,omitempty"`
	// 人員編號
	AccountID string `json:"account_id,omitempty"`
	// 負責人姓名
	Name string `json:"name,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
}

type Task_Template struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	// code
	Code string `json:"code,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//里程碑
	Landmark bool `json:"landmark,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//前一任務Code
	LastTaskCode string `json:"last_task_code,omitempty"`
	//前一任務名字
	LastTaskName string `json:"last_task_name,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time,omitempty"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
}

type Task_Template_Last struct {
	// 範本編號
	PtID string `json:"pt_id,omitempty"`
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	// code
	Code string `json:"code,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//里程碑
	Landmark bool `json:"landmark,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//前一任務名稱
	LastTName string `json:"last_t_name,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time,omitempty"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
}

type Bonita_ID_List struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//編號
	PID string `json:"p_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	UserId string `json:"user_id,omitempty"`
	//名字
	TaskUserName string `json:"task_user_name,omitempty"`
	//名字
	TaskBonitaId string `json:"task_bonita_id,omitempty"`
	//名字
	ProjectId string `json:"projectman_id,omitempty"`
	//名字
	ProjectmanName string `json:"projectman_name,omitempty"`
	//名字
	ProjectmanBonitaId string `json:"projectman_bonita_id,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	// 主要負責人
	Principal bool `json:"principal,omitempty"`
}

type Base struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//里程碑
	Landmark bool `json:"landmark,omitempty"`
	//附件
	File bool `json:"file,omitempty"`
	//流水號
	Rank int `json:"rank,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	//編號
	Code string `json:"code,omitempty"`
	//[ 7] quantity
	Quantity int `json:"quantity,omitempty"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//實際完成日
	DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time,omitempty"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	//todo_type_id
	TodoTypeID string `json:"todo_type_id,omitempty"`
	//todo_status
	TodoStatus bool `json:"todo_status,omitempty"`
}

type Task_Hour_User struct {
	//編號
	AccountID string `json:"account_id,omitempty"`
	//名字
	Name string `json:"name,omitempty"`
	//時數
	Hour int `json:"hour,omitempty"`
	//任務名稱
	TName string `json:"t_name,omitempty"`
}

type Single struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//里程碑
	Landmark bool `json:"landmark,omitempty"`
	//附件
	File bool `json:"file,omitempty"`
	//流水號
	Rank int `json:"rank,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	//編號
	Code string `json:"code,omitempty"`
	//[ 7] quantity
	Quantity int `json:"quantity,omitempty"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//實際完成日
	DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	//todo_type_id
	TodoTypeID string `json:"todo_type_id,omitempty"`
	//todo_status
	TodoStatus bool `json:"todo_status,omitempty"`
}

type Created struct {
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty" binding:"required,uuid4" validate:"required"`
	//名字
	TName string `json:"t_name,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty"`
	//里程碑
	Landmark bool `json:"landmark,omitempty" `
	//附件
	File bool `json:"file,omitempty"`
	//[ 7] quantity
	Quantity int `json:"quantity,omitempty"`
	//流水號(資料庫自動新增)
	//Rank int `json:"rank,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty" binding:"required" validate:"required"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start" binding:"required" validate:"required"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty" binding:"required" validate:"required"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	//todo_type_id
	TodoTypeID string `json:"todo_type_id,omitempty" binding:"required" validate:"required"`
	//todo_status
	TodoStatus bool `json:"todo_status,omitempty"`
}

type Field struct {
	// 範本編號
	PtID string `json:"pt_id,omitempty"`
	//編號
	TID string `json:"t_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty" form:"documents_id" binding:"omitempty,uuid4"`
	//人員編號
	AccountID string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4"`
	//名字
	TName *string `json:"t_name,omitempty" form:"t_name"`
	//備註
	Remark *string `json:"remark,omitempty" form:"remark"`
	//里程碑
	Landmark *bool `json:"landmark,omitempty" form:"landmark"`
	//附件
	File *bool `json:"file,omitempty" form:"file"`
	//流水號
	Rank *int `json:"rank,omitempty" form:"rank"`
	//[ 7] quantity
	Quantity int `json:"quantity,omitempty" form:"quantity"`
	//前一任務
	LastTask *string `json:"last_task,omitempty" form:"last_task"`
	//來源編號
	OriginID string `json:"origin_id,omitempty" form:"origin_id"`
	//預計開始日
	DateForEstimatedStart *time.Time `json:"date_for_estimated_start,omitempty" form:"date_for_estimated_start"`
	//實際完成日
	DateForActualCompletion *time.Time `json:"date_for_actual_completion,omitempty" form:"date_for_actual_completion"`
	//預計完成日
	DateForEstimatedCompletion *time.Time `json:"date_for_estimated_completion,omitempty" form:"date_for_estimated_completion"`
	// code
	Code string `json:"code,omitempty" form:"code"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty" form:"default_date"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty" form:"default_labor_hour"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty" form:"hierarchy"`
	//todo_type_id
	TodoTypeID string `json:"todo_type_id,omitempty" form:"todo_type_id"`
	//todo_status
	TodoStatus bool `json:"todo_status,omitempty" form:"todo_status"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type Created_List struct {
	Task []*Created `json:"task"`
}

type Updated_List struct {
	Task []*Updated `json:"task"`
}

type List struct {
	Task []*struct {
		//編號
		TID string `json:"t_id,omitempty"`
		//單據來源編號
		DocumentsID string `json:"documents_id,omitempty"`
		//名字
		TName string `json:"t_name,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//里程碑
		Landmark bool `json:"landmark,omitempty"`
		//附件
		File bool `json:"file,omitempty"`
		//流水號
		Rank int `json:"rank,omitempty"`
		//前一任務
		LastTask string `json:"last_task,omitempty"`
		//來源編號
		OriginID string `json:"origin_id,omitempty"`
		//編號
		Code string `json:"code,omitempty"`
		//[ 7] quantity
		Quantity int `json:"quantity,omitempty"`
		//預計開始日
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//實際完成日
		DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
		//預計完成日
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//建立時間
		CreateTime time.Time `json:"create_time"`
		//[ 7] default_date
		DefaultDate int `json:"default_date,omitempty"`
		//[ 7] default_labor_hour
		DefaultLaborHour int `json:"default_labor_hour,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
		//todo_type_id
		TodoTypeID string `json:"todo_type_id,omitempty"`
		//todo_status
		TodoStatus bool `json:"todo_status,omitempty"`
	} `json:"task"`
	model.OutPage
}

type Task_Account_Labor_Hours struct {
	Task []*struct {
		//名字
		TName string `json:"t_name,omitempty"`
		// 人員編號
		AccountID string `json:"account_id,omitempty"`
		// 負責人姓名
		Name string `json:"name,omitempty"`
		// 工時
		Hour int `json:"hour,omitempty"`
	} `json:"task"`
	model.OutTotal
}

type Task_Templates struct {
	Task []*struct {
		//編號
		TID string `json:"t_id,omitempty"`
		// code
		Code string `json:"code,omitempty"`
		//名字
		TName string `json:"t_name,omitempty"`
		//里程碑
		Landmark bool `json:"landmark,omitempty"`
		//前一任務
		LastTask string `json:"last_task,omitempty"`
		//前一任務Code
		LastTaskCode string `json:"last_task_code,omitempty"`
		//前一任務名字
		LastTaskName string `json:"last_task_name,omitempty"`
		//預計開始日
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//預計完成日
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//建立時間
		CreateTime time.Time `json:"create_time,omitempty"`
		//[ 7] default_date
		DefaultDate int `json:"default_date,omitempty"`
		//[ 7] default_labor_hour
		DefaultLaborHour int `json:"default_labor_hour,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
	} `json:"task"`
	model.OutPage
}

type Task_Accounts struct {
	Task []*struct {
		//編號
		TID string `json:"t_id,omitempty"`
		// code
		Code string `json:"code,omitempty"`
		//單據來源編號
		DocumentsID string `json:"documents_id,omitempty"`
		//名字
		TName string `json:"t_name,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//里程碑
		Landmark bool `json:"landmark,omitempty"`
		//附件
		File bool `json:"file,omitempty"`
		//流水號
		Rank int `json:"rank,omitempty"`
		//[ 7] quantity
		Quantity int `json:"quantity,omitempty"`
		//前一任務
		LastTask string `json:"last_task,omitempty"`
		//來源編號
		OriginID string `json:"origin_id,omitempty"`
		// 任務負責人編號
		TuID string `json:"tu_id,omitempty"`
		//是否主要負責人
		Principal bool `json:"principal,omitempty"`
		// 人員編號
		AccountID string `json:"account_id,omitempty"`
		// 負責人姓名
		Name string `json:"name,omitempty"`
		//預計開始日
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//實際完成日
		DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
		//預計完成日
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//建立時間
		CreateTime time.Time `json:"create_time,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
		//[ 6] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
		BonitaUserID string `json:"bonita_user_id,omitempty"`
		//todo_status
		TodoStatus bool `json:"todo_status,omitempty"`
		//status
		Status string `json:"status,omitempty"`
	} `json:"task"`
	model.OutPage
}

type Bonita_ID_Lists struct {
	Task []*struct {
		//編號
		TID string `json:"t_id,omitempty"`
		//編號
		PID string `json:"p_id,omitempty"`
		//單據來源編號
		DocumentsID string `json:"documents_id,omitempty"`
		//名字
		UserId string `json:"user_id,omitempty"`
		//名字
		TaskUserName string `json:"task_user_name,omitempty"`
		//名字
		TaskBonitaId string `json:"task_bonita_id,omitempty"`
		//名字
		ProjectId string `json:"projectman_id,omitempty"`
		//名字
		ProjectmanName string `json:"projectman_name,omitempty"`
		//名字
		ProjectmanBonitaId string `json:"projectman_bonita_id,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
		// 主要負責人
		Principal bool `json:"principal,omitempty"`
	} `json:"task"`
	model.OutPage
}

type Task_Template_Lasts struct {
	Task []*struct {
		// 範本編號
		PtID string `json:"pt_id,omitempty"`
		//編號
		TID string `json:"t_id,omitempty"`
		//單據來源編號
		DocumentsID string `json:"documents_id,omitempty"`
		// code
		Code string `json:"code,omitempty"`
		//名字
		TName string `json:"t_name,omitempty"`
		//里程碑
		Landmark bool `json:"landmark,omitempty"`
		//前一任務
		LastTask string `json:"last_task,omitempty"`
		//來源編號
		OriginID string `json:"origin_id,omitempty"`
		//預計開始日
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//預計完成日
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//前一任務名稱
		LastTName string `json:"last_t_name,omitempty"`
		//建立時間
		CreateTime time.Time `json:"create_time,omitempty"`
		//[ 7] default_date
		DefaultDate int `json:"default_date,omitempty"`
		//[ 7] default_labor_hour
		DefaultLaborHour int `json:"default_labor_hour,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
	} `json:"task"`
	model.OutPage
}
type Task_OriginId struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//編號
	PID string `json:"p_id,omitempty"`
	//編號
	Code string `json:"code,omitempty"`
	//編號
	PCode string `json:"p_code,omitempty"`
	//編號
	PName string `json:"p_name,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//預計工時
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	// 任務負責人編號
	TuID string `json:"tu_id,omitempty"`
	//是否主要負責人
	Principal bool `json:"principal"`
	// 人員編號
	AccountID string `json:"account_id,omitempty"`
	// 負責人姓名
	Name string `json:"name,omitempty"`
	//預計開始date_for_estimated_start
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//預計完成date_for_estimated_completion
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//實際完成日
	DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
	//建立時間
	CreateTime time.Time `json:"create_time,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	// todo_type分類名稱
	TType string `json:"t_type,omitempty"`
	// 待辦事項是否完成
	TodoStatus bool `json:"todo_status"`
	//里程碑
	Landmark bool `json:"landmark,omitempty"`
}

type Task_OriginIds struct {
	Task []*struct {
		//編號
		TID string `json:"t_id,omitempty"`
		//編號
		PID string `json:"p_id,omitempty"`
		//編號
		Code string `json:"code,omitempty"`
		//編號
		PCode string `json:"p_code,omitempty"`
		//編號
		PName string `json:"p_name,omitempty"`
		//單據來源編號
		DocumentsID string `json:"documents_id,omitempty"`
		//名字
		TName string `json:"t_name,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//預計工時
		DefaultLaborHour int `json:"default_labor_hour,omitempty"`
		//來源編號
		OriginID string `json:"origin_id,omitempty"`
		// 任務負責人編號
		TuID string `json:"tu_id,omitempty"`
		//是否主要負責人
		Principal bool `json:"principal"`
		// 人員編號
		AccountID string `json:"account_id,omitempty"`
		// 負責人姓名
		Name string `json:"name,omitempty"`
		//預計開始date_for_estimated_start
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//預計完成date_for_estimated_completion
		DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
		//實際完成日
		DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
		//建立時間
		CreateTime time.Time `json:"create_time,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
		// todo_type分類名稱
		TType string `json:"t_type,omitempty"`
		// 待辦事項是否完成
		TodoStatus bool `json:"todo_status"`
		//里程碑
		Landmark bool `json:"landmark"`
	} `json:"task"`
	model.OutTotal
}

type Task_Hour_Users struct {
	Task []*struct {
		//編號
		AccountID string `json:"account_id,omitempty"`
		//名字
		Name string `json:"name,omitempty"`
		//時數
		Hour int `json:"hour,omitempty"`
		//任務名稱
		TName string `json:"t_name,omitempty"`
	} `json:"task"`
	model.OutTotal
}

type Task_User_Accounts struct {
	Task []*struct {
		//編號
		TID string `json:"t_id,omitempty"`
		//單據來源編號
		DocumentsID string `json:"documents_id,omitempty"`
		//名字
		UserId string `json:"user_id,omitempty"`
		//名字
		TaskUserName string `json:"task_user_name,omitempty"`
		//名字
		TaskBonitaId string `json:"task_bonita_id,omitempty"`
		//名字
		ProjectmanName string `json:"projectman_name,omitempty"`
		//名字
		ProjectmanBonitaId string `json:"projectman_bonita_id,omitempty"`
		//階層
		Hierarchy int `json:"hierarchy,omitempty"`
	} `json:"task"`
	model.OutPage
}

type Updated struct {
	//編號
	TID string `json:"t_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty" binding:"omitempty,uuid4"`
	//名字
	TName string `json:"t_name,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//里程碑
	Landmark *bool `json:"landmark,omitempty"`
	//附件
	File *bool `json:"file,omitempty"`
	//流水號
	Rank int `json:"rank,omitempty"`
	//[ 7] quantity
	Quantity int `json:"quantity,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//來源編號
	OriginID string `json:"origin_id,omitempty"`
	//預計開始日
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//實際完成日
	DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//[ 7] default_date
	DefaultDate int `json:"default_date,omitempty"`
	//[ 7] default_labor_hour
	DefaultLaborHour int `json:"default_labor_hour,omitempty"`
	//階層
	Hierarchy int `json:"hierarchy,omitempty"`
	//todo_type_id
	TodoTypeID string `json:"todo_type_id,omitempty"`
	//todo_status
	TodoStatus bool `json:"todo_status,omitempty"`
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
	return "task"
}

func (a *Create_Table) TableName() string {
	return "task"
}
