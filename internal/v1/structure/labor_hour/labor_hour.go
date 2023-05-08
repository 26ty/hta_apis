package labor_hour

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
	//[ 4] h_id                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	HID string `gorm:"primary_key;column:h_id;uuid_generate_v4()type:UUID;" json:"h_id,omitempty"`
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

}

//工時表編輯(不用Groupby)
type GetUserCategoryLabor struct {
	//工時單據來源 project customer表
	//DocumentsID string `json:"documents_id,omitempty"` //p_id or cd_id //無法同時給所以分成PID跟CdID
	PID string `json:"p_id,omitempty"`//p_id
	CdIDTask string `json:"cd_id_task,omitempty"`//cd_id
	CdIDCountersign string `json:"cd_id_countersign,omitempty"`//cd_id
	//DocumentsCode string `json:"documents_code,omitempty"` //無法同時給所以分成PCode跟CdCode
	PCode string `json:"p_code,omitempty"`//p_code
	CdCodeTask string `json:"cd_code_task,omitempty"`//cd_code
	CdCodeCountersign string `json:"cd_code_countersign,omitempty"`//cd_code
	//DocumentsName string `json:"documents_name,omitempty"`//單據名稱 //無法同時給所以分成PName跟CdName
	PName string `json:"p_name,omitempty"`//p_name
	//CdName string `json:"cd_name,omitempty"`//cd_name //客需單沒有Name
	//DocumentsmanID string `json:"documentsman_id,omitempty"`//專案經理or客需單負責人 //無法同時給所以分成PProjectmanID跟CdProjectmanID
	PProjectmanID string `json:"p_projectman_id,omitempty"`//p_projectman_id
	CdProjectmanIDTask string `json:"cd_projectman_id_task,omitempty"`//cd_projectman_id
	CdProjectmanIDCountersign string `json:"cd_projectman_id_countersign,omitempty"`//cd_projectman_id
	
	//任務task表
	TID string `json:"t_id,omitempty"`//任務id
	TCode string `json:"t_code,omitempty"` // 任務代號
	TName string `json:"t_name,omitempty"` //任務名稱
	//會簽countersign表
	CsID string `json:"cs_id,omitempty"`//會簽id
	//部門department表
	DName string `json:"d_name,omitempty"` //會簽單位名稱
	
	//任務task_user表
	TuID string `json:"tu_id,omitempty"`//任務id
	//會簽countersign_user表
	CuID string `json:"cu_id,omitempty"`//會簽id
	
	
	//labor_hour表
	HID string `json:"h_id,omitempty"`//工時表id 
	Category string `json:"category,omitempty"`
	Laborhour string `json:"laborhour,omitempty"` //工時總數
	DateForStart time.Time `json:"date_for_start,omitempty"` // 日期
	Title string `json:"title,omitempty"`	// 主題 
	Content string `json:"content,omitempty"`	// 內容  
	Nature string `json:"nature,omitempty"`	// 性質  
	TimeForStart float32 `json:"time_for_start,omitempty"`	// 起始時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`	// 結束時間

	//account表
	UserID string `json:"user_id,omitempty"` //該使用者id
	//Account string `json:"account,omitempty"`//使用者帳號 //帳號不應亂給所以取消
	UserName string `json:"user_name,omitempty"`//使用者名稱
}

//工時表
type GetUserAllLabor struct {
	//工時單據來源 project customer表
	//DocumentsID string `json:"documents_id,omitempty"` //p_id or cd_id //無法同時給所以分成PID跟CdID
	PID string `json:"p_id,omitempty"`//p_id
	CdIDTask string `json:"cd_id_task,omitempty"`//cd_id
	CdIDCountersign string `json:"cd_id_countersign,omitempty"`//cd_id
	//DocumentsCode string `json:"documents_code,omitempty"` //無法同時給所以分成PCode跟CdCode
	PCode string `json:"p_code,omitempty"`//p_code
	CdCodeTask string `json:"cd_code_task,omitempty"`//cd_code
	CdCodeCountersign string `json:"cd_code_countersign,omitempty"`//cd_code
	//DocumentsName string `json:"documents_name,omitempty"`//單據名稱 //無法同時給所以分成PName跟CdName
	PName string `json:"p_name,omitempty"`//p_name
	//CdName string `json:"cd_name,omitempty"`//cd_name //客需單沒有Name
	//DocumentsmanID string `json:"documentsman_id,omitempty"`//專案經理or客需單負責人 //無法同時給所以分成PProjectmanID跟CdProjectmanID
	PProjectmanID string `json:"p_projectman_id,omitempty"`//p_projectman_id
	CdProjectmanIDTask string `json:"cd_projectman_id_task,omitempty"`//cd_projectman_id
	CdProjectmanIDCountersign string `json:"cd_projectman_id_countersign,omitempty"`//cd_projectman_id
	
	//任務task表
	TID string `json:"t_id,omitempty"`//任務id
	TCode string `json:"t_code,omitempty"` // 任務代號
	TName string `json:"t_name,omitempty"` //任務名稱
	TOriginID string `json:"t_origin_id,omitempty"` //任務來源單號

	//會簽countersign表
	CsID string `json:"cs_id,omitempty"`//會簽id
	//部門department表
	DName string `json:"d_name,omitempty"` //會簽單位名稱
	
	//任務task_user表
	TuID string `json:"tu_id,omitempty"`//任務id
	//會簽countersign_user表
	CuID string `json:"cu_id,omitempty"`//會簽id

	//Meeting類型
	MOriginID string `json:"m_origin_id,omitempty"`
	
	//labor_hour表
	//HID string `json:"h_id,omitempty"`//工時表id //與Laborhour互相衝突無法GroupBy所以取消
	Category string `json:"category,omitempty"`
	//Laborhour string `json:"laborhour,omitempty"` //單一筆任務的工時 如果有重複提報需累計該任務已提報的工時總數
	//DateLaborHour string `json:"datelaborhour,omitempty"`//同一天的工時累計 //有另外給了所以取消
	DateForStart time.Time `json:"date_for_start,omitempty"` // 日期
	//Title string `json:"title,omitempty"`	// 主題  //影響Groupby結果所以取消
	//Content string `json:"content,omitempty"`	// 內容  //影響Groupby結果所以取消
	//Nature string `json:"nature,omitempty"`	// 性質  //影響Groupby結果所以取消

	//account表
	UserID string `json:"user_id,omitempty"` //該使用者id
	//Account string `json:"account,omitempty"`//使用者帳號 //帳號不應亂給所以取消
	UserName string `json:"user_name,omitempty"`//使用者名稱
}

//工時表個項總工時
type GetUserOneSumLabor struct {
	SumOfLaborhour string `json:"sum_of_laborhour,omitempty"`
}

//工時表個項工時
type GetUserOneLabor struct {
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 工時
	Laborhour string `json:"laborhour,omitempty"`
}

//工時表工時加總每天
type GetUserAllSumLabor struct {
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 工時
	SumLaborhour string `json:"sum_laborhour,omitempty"`
}

type LaborHour struct {
	// 工時編號
	HID string `json:"h_id,omitempty"`
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
	//
	CreaterName string `json:"creater_name,omitempty"`

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
}

type Base struct {
	// 工時編號
	HID string `json:"h_id,omitempty"`
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
}

type Single struct {
	// 工時編號
	HID string `json:"h_id,omitempty"`
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
}

type Created struct {
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
}

type Field struct {
	// 工時編號
	HID string `json:"h_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true" form:"h_id"`
	// 工時歸屬
	Category string `json:"category,omitempty" form:"category" binding:"omitempty,uuid4"`
	// 主題
	Title *string `json:"title,omitempty" form:"title"`
	// 內容
	Content *string `json:"content,omitempty" form:"content"`
	// 性質
	Nature *string `json:"nature,omitempty" form:"nature"`
	// 日期
	DateForStart time.Time `json:"date_for_start,omitempty" form:"date_for_start"`
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
}

type Fields struct {
	Field
	model.InPage
}
type Field_Month struct {
	Field
	FirstDate time.Time `json:"first_date"`
}

type GetUserAllLabors struct {
	LaborHour []*struct {
		//工時單據來源 project customer表
		//DocumentsID string `json:"documents_id,omitempty"` //p_id or cd_id //無法同時給所以分成PID跟CdID
		PID string `json:"p_id,omitempty"`//p_id
		CdIDTask string `json:"cd_id_task,omitempty"`//cd_id
		CdIDCountersign string `json:"cd_id_countersign,omitempty"`//cd_id
		//DocumentsCode string `json:"documents_code,omitempty"` //無法同時給所以分成PCode跟CdCode
		PCode string `json:"p_code,omitempty"`//p_code
		CdCodeTask string `json:"cd_code_task,omitempty"`//cd_code
		CdCodeCountersign string `json:"cd_code_countersign,omitempty"`//cd_code
		//DocumentsName string `json:"documents_name,omitempty"`//單據名稱 //無法同時給所以分成PName跟CdName
		PName string `json:"p_name,omitempty"`//p_name
		//CdName string `json:"cd_name,omitempty"`//cd_name //客需單沒有Name
		//DocumentsmanID string `json:"documentsman_id,omitempty"`//專案經理or客需單負責人 //無法同時給所以分成PProjectmanID跟CdProjectmanID
		PProjectmanID string `json:"p_projectman_id,omitempty"`//p_projectman_id
		CdProjectmanIDTask string `json:"cd_projectman_id_task,omitempty"`//cd_projectman_id
		CdProjectmanIDCountersign string `json:"cd_projectman_id_countersign,omitempty"`//cd_projectman_id
		
		//任務task表
		TID string `json:"t_id,omitempty"`//任務id
		TCode string `json:"t_code,omitempty"` // 任務代號
		TName string `json:"t_name,omitempty"` //任務名稱
		TOriginID string `json:"t_origin_id,omitempty"` //任務來源單號

		//會簽countersign表
		CsID string `json:"cs_id,omitempty"`//會簽id
		//部門department表
		DName string `json:"d_name,omitempty"` //會簽單位名稱
		
		//任務task_user表
		TuID string `json:"tu_id,omitempty"`//任務id
		//會簽countersign_user表
		CuID string `json:"cu_id,omitempty"`//會簽id

		//Meeting類型
		MOriginID string `json:"m_origin_id,omitempty"`
		
		//labor_hour表
		//HID string `json:"h_id,omitempty"`//工時表id //與Laborhour互相衝突無法GroupBy所以取消
		Category string `json:"category,omitempty"`
		//Laborhour string `json:"laborhour,omitempty"` //單一筆任務的工時 如果有重複提報需累計該任務已提報的工時總數
		//DateLaborHour string `json:"datelaborhour,omitempty"`//同一天的工時累計
		DateForStart time.Time `json:"date_for_start,omitempty"` // 日期
		//Title string `json:"title,omitempty"`	// 主題
		//Content string `json:"content,omitempty"`	// 內容
		//Nature string `json:"nature,omitempty"`	// 性質

		//account表
		UserID string `json:"user_id,omitempty"` //該使用者id
		//Account string `json:"account,omitempty"`//使用者帳號 //帳號不應亂給所以取消
		UserName string `json:"user_name,omitempty"`//使用者名稱
		DateOfLaborhour []*GetUserOneLabor `json:"date_of_laborhour"`
		SumOfLaborhour string `json:"sum_of_laborhour,omitempty"`
	} `json:"labor_hour"`
	DateOfSum []*GetUserAllSumLabor `json:"date_of_sum"`
}

type LaborHours struct {
	LaborHour []*struct {
		// 工時編號
		HID string `json:"h_id,omitempty"`
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
		//
		CreaterName string `json:"creater_name,omitempty"`

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
	} `json:"labor_hour"`
	model.OutPage
}

type List struct {
	LaborHour []*struct {
		// 工時編號
		HID string `json:"h_id,omitempty"`
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
	} `json:"labor_hour"`
	model.OutPage
}

type Updated struct {
	// 工時編號
	HID string `json:"h_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
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
	return "labor_hour"
}
