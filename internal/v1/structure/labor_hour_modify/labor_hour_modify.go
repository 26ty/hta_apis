package labor_hour_modify

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//[ 0] title                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Title string `gorm:"column:title;type:TEXT;" json:"title,omitempty"`
	//[ 1] content                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Content string `gorm:"column:content;type:TEXT;" json:"content,omitempty"`
	//[ 2] nature                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Nature string `gorm:"column:nature;type:TEXT;" json:"nature,omitempty"`
	//[ 3] create_time                                    TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [now()]
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
	//[ 4] hm_id                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	HmID string `gorm:"primary_key;column:hm_id;uuid_generate_v4()type:UUID;" json:"hm_id,omitempty"`
	//[ 4] hour_id                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	HourID string `gorm:"column:hour_id;type:UUID;" json:"hour_id,omitempty"`
	//[ 5] category                                       UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	Category string `gorm:"column:category;type:UUID;" json:"category,omitempty"`
	//[ 6] time_for_start                                 double precision                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TimeForStart float32 `gorm:"column:time_for_start;type:double precision;" json:"time_for_start,omitempty"`
	//[ 7] date_for_start                                 DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForStart time.Time `gorm:"column:date_for_start;type:DATE;" json:"date_for_start,omitempty"`
	//[ 8] time_for_end                                   double precision                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TimeForEnd float32 `gorm:"column:time_for_end;type:double precision;" json:"time_for_end,omitempty"`
	//[ 8] labor_hour                                   double precision                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Laborhour float32 `gorm:"column:laborhour;type:double precision;" json:"laborhour,omitempty"`
	// 創建者
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	//[ 7] bonita_case_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	// 狀態
	StatusTypeID string `gorm:"column:status_type_id;type:UUID;" json:"status_type_id,omitempty"`

}

//直屬主管送審資料
type ReviewByDepartment struct {
	//labor_hour表
	Title string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Nature string `json:"nature,omitempty"`
	DateForStart string `json:"date_for_start,omitempty"`
	TimeForStart float32 `json:"time_for_start,omitempty"`
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	Laborhour string `json:"laborhour,omitempty"`
	
	//labor_hour_modify表
	HmID string `json:"hm_id,omitempty"`
	HourID string `json:"hour_id,omitempty"`
	Category string `json:"category,omitempty"`
	MTitle string `json:"m_title,omitempty"`
	MContent string `json:"m_content,omitempty"`
	MNature string `json:"m_nature,omitempty"`
	MDateForStart string `json:"m_date_for_start,omitempty"`
	MTimeForStart float32 `json:"m_time_for_start,omitempty"`
	MTimeForEnd float32 `json:"m_time_for_end,omitempty"`
	MLaborhour string `json:"m_laborhour,omitempty"`
	//創建者id
	Creater string `json:"creater,omitempty"`
	//創建者名稱
	CreaterName string `json:"creater_name,omitempty"`
	//創建時間
	Create_time string `json:"create_time,omitempty"`

	//工時單據來源(專案)
	PID string `json:"p_id,omitempty"`
	PCode string `json:"p_code,omitempty"`
	PName string `json:"p_name,omitempty"`
	ProjectmanId string `json:"projectman_id,omitempty"`
	ProjectmanName string `json:"projectman_name,omitempty"`
	
	//任務id
	TID string `json:"t_id,omitempty"`
	TCode string `json:"t_code,omitempty"`
	TName string `json:"t_name,omitempty"`

	//工時單據來源(客需單)
	CdID string `json:"cd_id,omitempty"`
	CdCode string `json:"cd_code,omitempty"`
	CdName string `json:"cd_name,omitempty"`
	
	//會簽id
	CsID string `json:"cs_id,omitempty"`
	CsName string `json:"cs_name,omitempty"`
	DName string `json:"d_name,omitempty"`

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string  `json:"bonita_task_name,omitempty"`
}

type LaborHourModify struct {
	// 工時編號
	HmID string `json:"hm_id,omitempty"`
	// 原工時編號
	HourID string `json:"hour_id,omitempty"`
	// 工時歸屬
	Category string `json:"category,omitempty"`
	// 主題
	Title string `json:"title,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 性質
	Nature string `json:"nature,omitempty"`
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 起始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 工時
	Laborhour string `json:"laborhour,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	//
	CreaterName string `json:"creater_name,omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態名稱
	Status string `json:"status,omitempty" `

	//單據來源編號
	TaskDocumentsID string `json:"task_documents_id,omitempty"`
	// code(project任務)
	ProjectCode string `json:"project_code,omitempty"`
	// code(customer_demand任務)
	CustomerDemandTaskCode string `json:"customer_demand_task_code,omitempty"`
	//來源編號
	TaskOriginID string `json:"task_origin_id,omitempty"`
	// (project)machine_finished_number
	MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
	//單據來源編號
	CountersignDocumentsID string `json:"countersign_documents_id,omitempty"`
	// code(customer_demand會簽)
	CustomerDemandCountersignCode string `json:"customer_demand_countersign_code,omitempty"`

	//
	TaskUserID string `json:"task_user_id,omitempty"`
	//
	TaskUserName string `json:"task_user_name,omitempty"`

	//
	CountersignUserID string `json:"countersign_user_id,omitempty"`
	//
	CountersignUserName string `json:"countersign_user_name,omitempty"`

	//project
	PId string `json:"p_id,omitempty"`
	PName string `json:"p_name,omitempty"`
	ProjectmanId string `json:"projectman_id,omitempty"`
	ProjectmanName string `json:"projectman_name,omitempty"`
	
	//task
	TId string `json:"t_id,omitempty"`
	TCode string `json:"t_code,omitempty"`
	TName string `json:"t_name,omitempty"`
}

type Base struct {
	// 工時編號
	HmID string `json:"hm_id,omitempty"`
	// 原工時編號
	HourID string `json:"hour_id,omitempty"`
	// 工時歸屬
	Category string `json:"category,omitempty"`
	// 主題
	Title string `json:"title,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 性質
	Nature string `json:"nature,omitempty"`
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 起始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 工時
	Laborhour float32 `json:"laborhour,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Single struct {
	// 工時編號
	HmID string `json:"hm_id,omitempty"`
	// 原工時編號
	HourID string `json:"hour_id,omitempty"`
	// 工時歸屬
	Category string `json:"category,omitempty"`
	// 主題
	Title string `json:"title,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 性質
	Nature string `json:"nature,omitempty"`
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 起始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 工時
	Laborhour float32 `json:"laborhour,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Created struct {
	// 原工時編號
	HourID string `json:"hour_id" binding:"required" validate:"required"`
	// 工時歸屬
	Category string `json:"category" binding:"required" validate:"required"`
	// 主題
	Title string `json:"title" binding:"required" validate:"required"`
	// 內容
	Content string `json:"content" binding:"required" validate:"required"`
	// 性質
	Nature string `json:"nature" binding:"required" validate:"required"`
	// 日期
	DateForStart time.Time `json:"date_for_start" binding:"required" validate:"required"`
	// 起始時間
	TimeForStart float32 `json:"time_for_start" binding:"required" validate:"required"`
	// 結束時間
	TimeForEnd float32 `json:"time_for_end" binding:"required" validate:"required"`
	// 工時
	Laborhour float32 `json:"laborhour,omitempty"`
	// 創建者
	Creater string  `json:"creater" binding:"required" validate:"required"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Field struct {
	// 工時編號
	HmID string `json:"hm_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true" form:"hm_id"`
	// 原工時編號
	HourID *string `json:"hour_id,omitempty" form:"hour_id" binding:"omitempty,uuid4"`
	// 工時歸屬
	Category *string `json:"category,omitempty" form:"category" binding:"omitempty,uuid4"`
	// 主題
	Title *string `json:"title,omitempty" form:"title"`
	// 內容
	Content *string `json:"content,omitempty" form:"content"`
	// 性質
	Nature *string `json:"nature,omitempty" form:"nature"`
	// 日期
	DateForStart *string `json:"date_for_start,omitempty" form:"date_for_start"`
	// 起始時間
	TimeForStart *float32 `json:"time_for_start,omitempty" form:"time_for_start"`
	// 結束時間
	TimeForEnd *float32 `json:"time_for_end,omitempty" form:"time_for_end"`
	//任務編號
	TID string `json:"t_id,omitempty" `
	//任務編號
	CuID string `json:"cu_id,omitempty" `
	//人員編號
	UserID string `json:"user_id,omitempty" binding:"omitempty,uuid4" form:"user_id"`
	// 工時
	Laborhour float32 `json:"laborhour,omitempty" form:"laborhour"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Fields struct {
	Field
	model.InPage
}

type LaborHourModifys struct {
	LaborHourModify []*struct {
		// 工時編號
		HmID string `json:"hm_id,omitempty"`
		// 原工時編號
		HourID string `json:"hour_id,omitempty"`
		// 工時歸屬
		Category string `json:"category,omitempty"`
		// 主題
		Title string `json:"title,omitempty"`
		// 內容
		Content string `json:"content,omitempty"`
		// 性質
		Nature string `json:"nature,omitempty"`
		// 日期
		DateForStart time.Time `json:"date_for_start,omitempty"`
		// 起始時間
		TimeForStart float32 `json:"time_for_start,omitempty"`
		// 結束時間
		TimeForEnd float32 `json:"time_for_end,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 工時
		Laborhour string `json:"laborhour,omitempty"`
		//
		Creater string `json:"creater,omitempty"`
		//
		CreaterName string `json:"creater_name,omitempty"`
		//[ 7] bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		// 狀態ID
		StatusTypeID string `json:"status_type_id,omitempty" `
		// 狀態名稱
		Status string `json:"status,omitempty" `

		//單據來源編號
		TaskDocumentsID string `json:"task_documents_id,omitempty"`
		// code(project任務)
		ProjectCode string `json:"project_code,omitempty"`
		// code(customer_demand任務)
		CustomerDemandTaskCode string `json:"customer_demand_task_code,omitempty"`
		//來源編號
		TaskOriginID string `json:"task_origin_id,omitempty"`
		// (project)machine_finished_number
		MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
		//單據來源編號
		CountersignDocumentsID string `json:"countersign_documents_id,omitempty"`
		// code(customer_demand會簽)
		CustomerDemandCountersignCode string `json:"customer_demand_countersign_code,omitempty"`

		//
		TaskUserID string `json:"task_user_id,omitempty"`
		//
		TaskUserName string `json:"task_user_name,omitempty"`

		//
		CountersignUserID string `json:"countersign_user_id,omitempty"`
		//
		CountersignUserName string `json:"countersign_user_name,omitempty"`

		//project
		PId string `json:"p_id,omitempty"`
		PName string `json:"p_name,omitempty"`
		ProjectmanId string `json:"projectman_id,omitempty"`
		ProjectmanName string `json:"projectman_name,omitempty"`
		
		//task
		TId string `json:"t_id,omitempty"`
		TCode string `json:"t_code,omitempty"`
		TName string `json:"t_name,omitempty"`
	} `json:"labor_hour_modify"`
	model.OutPage
}

type List struct {
	LaborHourModify []*struct {
		// 工時編號
		HmID string `json:"hm_id,omitempty"`
		// 原工時編號
		HourID string `json:"hour_id,omitempty"`
		// 工時歸屬
		Category string `json:"category,omitempty"`
		// 主題
		Title string `json:"title,omitempty"`
		// 內容
		Content string `json:"content,omitempty"`
		// 性質
		Nature string `json:"nature,omitempty"`
		// 日期
		DateForStart time.Time `json:"date_for_start,omitempty"`
		// 起始時間
		TimeForStart float32 `json:"time_for_start,omitempty"`
		// 結束時間
		TimeForEnd float32 `json:"time_for_end,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 工時
		Laborhour float32 `json:"laborhour,omitempty"`
		//
		Creater string `json:"creater,omitempty"`
		//[ 7] bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		// 狀態
		StatusTypeID string `json:"status_type_id,omitempty" `
	} `json:"labor_hour_modify"`
	model.OutPage
}

type Updated struct {
	// 工時編號
	HmID string `json:"hm_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 原工時編號
	HourID *string `json:"hour_id,omitempty" binding:"omitempty,uuid4"`
	// 工時歸屬
	Category *string `json:"category,omitempty" binding:"omitempty,uuid4"`
	// 主題
	Title *string `json:"title,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 性質
	Nature string `json:"nature,omitempty"`
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 起始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 工時
	Laborhour float32 `json:"laborhour,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" `
}

type Updated_Review struct {
	// 編號
	HmID string `json:"hm_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 狀態
	StatusTypeID string `json:"status_type_id,omitempty" binding:"required" validate:"required"`
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
	return "labor_hour_modify"
}
