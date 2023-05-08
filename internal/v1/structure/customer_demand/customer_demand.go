package customer_demand

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//[ 0] cd_id
	CdId string `gorm:"primary_key;column:cd_id;uuid_generate_v4()type:UUID;" json:"cd_id,omitempty"`
	//[ 1] code
	Code string `gorm:"column:code;type:TEXT;" json:"code,omitempty"`
	//[ 1] project_id
	ProjectId string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	//[ 2] customer_id
	CustomerId string `gorm:"column:customer_id;type:UUID;" json:"customer_id,omitempty"`
	//[ 3] salesman_id
	SalesmanId string `gorm:"column:salesman_id;type:UUID;" json:"salesman_id,omitempty"`
	//[ 4] projectman_id
	ProjectmanId string `gorm:"column:projectman_id;type:UUID;" json:"projectman_id,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `gorm:"column:contact_person_id;type:UUID;" json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `gorm:"column:demand_content;type:TEXT;" json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `gorm:"column:suitable_content;type:TEXT;" json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `gorm:"column:other_content;type:TEXT;" json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `gorm:"column:budget;type:text;" json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `gorm:"column:date_for_recive;type:DATE;" json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `gorm:"column:date_for_estimated_start;type:DATE;" json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `gorm:"column:date_for_estimated_end;type:DATE;" json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `gorm:"column:date_for_actual_done;type:DATE;" json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `gorm:"column:machine_status_id;type:UUID;" json:"machine_status_id,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `gorm:"column:extend_rem;type:TEXT;" json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `gorm:"column:date_for_devlop;type:TIMESTAMPTZ;" json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `gorm:"column:est_quantity;type:INT4;" json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport bool `gorm:"column:eva_report;type:INT4;" json:"eva_report,omitempty"`
	//[ 7] status
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
	//[ 7] accept
	Accept bool `gorm:"column:accept;type:BOOLEAN;" json:"accept,omitempty"`
	//[ 4] creater
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
	//[ 7] bonita_case_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `gorm:"column:date_for_estimated_manager;type:DATE;" json:"date_for_estimated_manager,omitempty"`
	// 送審日
	DateForDelivery time.Time `gorm:"column:date_for_delivery;type:TIMESTAMP;" json:"date_for_delivery,omitempty"`
	// 是否填料
	Fill bool `gorm:"column:fill;type:BOOLEAN;" json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `gorm:"column:result_status;type:text;" json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `gorm:"column:result_content;type:text;" json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `gorm:"column:date_for_result;type:date;" json:"date_for_result,omitempty"`
	// 機台名稱
	ExtendTypeName string `gorm:"column:extend_type_name;type:text;" json:"extend_type_name,omitempty"`
}

type Create_Table struct {
	//[ 0] cd_id
	CdId string `gorm:"primary_key;column:cd_id;uuid_generate_v4()type:UUID;" json:"cd_id,omitempty"`
	//[ 1] project_id
	ProjectId string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	//[ 2] customer_id
	CustomerId string `gorm:"column:customer_id;type:UUID;" json:"customer_id,omitempty"`
	//[ 3] salesman_id
	SalesmanId string `gorm:"column:salesman_id;type:UUID;" json:"salesman_id,omitempty"`
	//[ 4] projectman_id
	ProjectmanId string `gorm:"column:projectman_id;type:UUID;" json:"projectman_id,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `gorm:"column:contact_person_id;type:UUID;" json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `gorm:"column:demand_content;type:TEXT;" json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `gorm:"column:suitable_content;type:TEXT;" json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `gorm:"column:other_content;type:TEXT;" json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `gorm:"column:budget;type:text;" json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `gorm:"column:date_for_recive;type:DATE;" json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `gorm:"column:date_for_estimated_start;type:DATE;" json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `gorm:"column:date_for_estimated_end;type:DATE;" json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `gorm:"column:date_for_actual_done;type:DATE;" json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `gorm:"column:machine_status_id;type:UUID;" json:"machine_status_id,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `gorm:"column:extend_rem;type:TEXT;" json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `gorm:"column:date_for_devlop;type:TIMESTAMPTZ;" json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `gorm:"column:est_quantity;type:INT4;" json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport bool `gorm:"column:eva_report;type:INT4;" json:"eva_report,omitempty"`
	//[ 7] status
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
	//[ 7] accept
	Accept bool `gorm:"column:accept;type:BOOLEAN;" json:"accept,omitempty"`
	//[ 4] creater
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time,omitempty"`
	//[ 7] bonita_case_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `gorm:"column:date_for_estimated_manager;type:DATE;" json:"date_for_estimated_manager,omitempty"`
	// 送審日
	DateForDelivery time.Time `gorm:"column:date_for_delivery;type:TIMESTAMP;" json:"date_for_delivery,omitempty"`
	// 是否填料
	Fill bool `gorm:"column:fill;type:BOOLEAN;" json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `gorm:"column:result_status;type:text;" json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `gorm:"column:result_content;type:text;" json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `gorm:"column:date_for_result;type:date;" json:"date_for_result,omitempty"`
	// 機台名稱
	ExtendTypeName string `gorm:"column:extend_type_name;type:text;" json:"extend_type_name,omitempty"`
}

type H_CR struct {
	//customer_demand表
	//編號
	CdID string `json:"cd_id,omitempty"`
	//單號
	Code           string `json:"code,omitempty"`
	Demand_content string `json:"demand_content,omitempty"`

	//countersign表
	//編號
	CsID string `json:"cs_id,omitempty"`

	//countersign_user表
	//編號
	CuID string `json:"cu_id,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//預定完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//預定完成日(員工)
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//完工送審
	DateForCompletion time.Time `json:"date_for_completion,omitempty"`
	//完工送審(員工)
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`

	//會簽人員
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

type CR struct {
	// customer_demend表
	//案件代號(客需單代號)
	CdID string `json:"cd_id,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	CdCode    string `json:"cd_code,omitempty"`
	//業務負責人(發文者)
	SalesmanID   string `json:"salesman_id,omitempty"`
	SalesmanName string `json:"salesman_name,omitempty"`
	//問題描述(客戶需求說明+適用製成說明)
	DemandContent   string `json:"demand_content,omitempty"`
	SuitableContent string `json:"suitable_content,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`

	//task表
	//task_id
	TID string `json:"t_id,omitempty"`
	//收件日期(任務創建日期)
	CreateTime time.Time `json:"create_time,omitempty"`
	//task_user_id
	TuID string `json:"tu_id,omitempty"`
	//任務負責人
	TaskUserID   string `json:"task_user_id,omitempty"`
	TaskUserName string `json:"task_user_name,omitempty"`
	//狀態
	Status string `json:"status,omitempty"`
	//處理情形
	Remark string `json:"remark,omitempty"`

	//labor_hour表
	//工時明細
	HID string `json:"h_id,omitempty"`
}

type Customer_Review_Task struct {
	// 編號
	TuID string `json:"tu_id,omitempty"`
	// 送出日
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	// 任務負責人
	Name string `json:"name,omitempty"`

	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//編號
	TCode string `json:"t_code,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//前一任務
	LastTask string `json:"last_task,omitempty"`
	//前一任務編號
	LastTaskCode string `json:"last_task_code,omitempty"`
	//前一任務名稱
	LastTaskName string `json:"last_task_name,omitempty"`
	//下一任務
	NextTask string `json:"next_task,omitempty"`
	//下一任務編號
	NextTaskCode string `json:"next_task_code,omitempty"`
	//下一任務名稱
	NextTaskName string `json:"next_task_name,omitempty"`
	//預計開始日
	TaskDateForEstimatedStart *time.Time `json:"task_date_for_estimated_start,omitempty"`
	//實際完成日
	DateForActualCompletion *time.Time `json:"date_for_actual_completion,omitempty"`
	//預計完成日
	DateForEstimatedCompletion *time.Time `json:"date_for_estimated_completion,omitempty"`
	//前一預計開始日
	LastDateForEstimatedStart *time.Time `json:"last_date_for_estimated_start,omitempty"`
	//前一實際完成日
	LastDateForActualCompletion *time.Time `json:"last_date_for_actual_completion,omitempty"`
	//前一預計完成日
	LastDateForEstimatedCompletion *time.Time `json:"last_date_for_estimated_completion,omitempty"`

	//編號
	CdID string `json:"cd_id,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//單號
	Code string `json:"code,omitempty"`
	//客戶名稱
	CustomerID   string `json:"customer_id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	//客戶需求說明
	DemandContent string `json:"demand_content,omitempty"`
	//適用製成說明
	SuitableContent string `json:"suitable_content,omitempty"`
	//其他說明
	OtherContent string `json:"other_content,omitempty"`
	//客戶預估時程(開始)
	DateForEstimatedStart *time.Time `json:"date_for_estimated_start,omitempty"`
	//客戶預估時程(結束)
	DateForEstimatedEnd *time.Time `json:"date_for_estimated_end,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `json:"date_for_recive,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string `json:"bonita_task_name,omitempty"`
}

type Customer_Review struct {
	//customer_demand表
	//編號
	CdID string `json:"cd_id,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//單號
	Code string `json:"code,omitempty"`
	//日期
	DateForRecive string `json:"date_for_recive,omitempty"`
	//客戶名稱
	CustomerID   string `json:"customer_id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	//客戶聯絡人
	ContactPersonID   string `json:"contact_person_id,omitempty"`
	ContactPersonName string `json:"contact_person_name,omitempty"`
	//電話
	ContactPersonPhone string `json:"contact_person_phone,omitempty"`
	//電子郵件
	ContactPersonEmail string `json:"contact_person_email,omitempty"`
	//客戶需求說明
	DemandContent string `json:"demand_content,omitempty"`
	//適用製成說明
	SuitableContent string `json:"suitable_content,omitempty"`
	//其他說明
	OtherContent string `json:"other_content,omitempty"`
	//客戶預算
	Budget string `json:"budget,omitempty"`
	//客戶預估時程(開始)
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//客戶預估時程(結束)
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	//機器狀態
	MachineStatusID string `json:"machine_status_id,omitempty"`
	//舊機衍生備註
	ExtendRem string `json:"extend_rem,omitempty"`
	//技術可行性評估報告，預計提交時間
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
	//未來三年預估數量
	EstQuantity string `json:"est_quantity,omitempty"`
	//不須提出技術可行性評估報告
	EvaReport bool `json:"eva_report,omitempty"`
	//發文單位
	SalesmanDep []*Dep `json:"salesman_dep,omitempty"`
	//發文者
	SalesmanID   string `json:"salesman_id,omitempty"`
	SalesmanName string `json:"salesman_name,omitempty"`
	//發文日期
	CreateTime time.Time `json:"create_time,omitempty"`
	//專案經理(PM)
	ProjectmanID   string `json:"projectman_id,omitempty"`
	ProjectmanName string `json:"projectman_name,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`

	//countersign_user表
	//會簽人員
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	//編號
	CuID string `json:"cu_id,omitempty"`
	//主旨
	Remark string `json:"remark,omitempty"`
	//預定完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//預定完成日(員工)
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//完工送審(員工)
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string `json:"bonita_task_name,omitempty"`
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

type Customer_Review2 struct {
	//customer_demand表
	//編號
	CdID string `json:"cd_id,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//單號
	Code string `json:"code,omitempty"`
	//發文單位
	SalesmanDep []*Dep `json:"salesman_dep,omitempty"`
	//發文者
	SalesmanID   string `json:"salesman_id,omitempty"`
	SalesmanName string `json:"salesman_name,omitempty"`
	//客戶名稱
	CustomerID   string `json:"customer_id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	//客戶需求說明
	DemandContent string `json:"demand_content,omitempty"`
	//適用製成說明
	SuitableContent string `json:"suitable_content,omitempty"`
	//其他說明
	OtherContent string `json:"other_content,omitempty"`
	//客戶預估時程(開始)
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//客戶預估時程(結束)
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
	//發文日期
	CreateTime time.Time `json:"create_time,omitempty"`

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string `json:"bonita_task_name,omitempty"`
}

type Customer_Demand_Account struct {
	//[ 0] cd_id
	CdId string `json:"cd_id,omitempty"`
	//[ 1] code
	Code string `json:"code,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//[ 1] project_code
	PCode string `json:"p_code,omitempty"`
	//[ 1] project_code
	PName string `json:"p_name,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `json:"machine_status_id,omitempty"`
	// 機台名稱
	ExtendTypeName string `json:"extend_type_name,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport bool `json:"eva_report,omitempty"`
	//[ 7] status
	Status string `json:"status,omitempty"`
	//[ 7] accept
	Accept bool `json:"accept,omitempty"`
	//[ 4] creater
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
	//[ 7] customer_id
	CustomerName string `json:"customer_name,omitempty"`
	//[ 8] salesman_id
	SalesmanName string `json:"salesman_name,omitempty"`
	//[ 9] projectman_id
	ProjectmanName string `json:"projectman_name,omitempty"`
	//[10] customer_id
	CustomerId string `json:"customer_id,omitempty"`
	//[11] salesman_id
	SalesmanId string `json:"salesman_id,omitempty"`
	//[12] projectman_id
	ProjectmanId string `json:"projectman_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
}

type Base struct {
	//[ 0] cd_id
	CdId string `json:"cd_id,omitempty"`
	//[ 1] code
	Code string `json:"code,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//[ 2] customer_id
	CustomerId string `json:"customer_id,omitempty"`
	//[ 3] salesman_id
	SalesmanId string `json:"salesman_id,omitempty"`
	//[ 4] projectman_id
	ProjectmanId string `json:"projectman_id,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `json:"machine_status_id,omitempty"`
	// 機台名稱
	ExtendTypeName string `json:"extend_type_name,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport bool `json:"eva_report,omitempty"`
	//[ 7] status
	Status string `json:"status,omitempty"`
	//[ 7] accept
	Accept bool `json:"accept,omitempty"`
	//[ 4] creater
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `json:"date_for_estimated_manager,omitempty"`
	// 送審日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
}

type Single struct {
	//[ 0] cd_id
	CdId string `json:"cd_id,omitempty"`
	//[ 1] code
	Code string `json:"code,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//[ 2] customer_id
	CustomerId string `json:"customer_id,omitempty"`
	//[ 3] salesman_id
	SalesmanId string `json:"salesman_id,omitempty"`
	//[ 4] projectman_id
	ProjectmanId string `json:"projectman_id,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `json:"machine_status_id,omitempty"`
	// 機台名稱
	ExtendTypeName string `json:"extend_type_name,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport bool `json:"eva_report,omitempty"`
	//[ 7] status
	Status string `json:"status,omitempty"`
	//[ 7] accept
	Accept bool `json:"accept,omitempty"`
	//[ 4] creater
	Creater string `json:"creater,omitempty"`
	//[ 7] create_time
	CreateTime time.Time `json:"create_time,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `json:"date_for_estimated_manager,omitempty"`
	// 送審日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
}

type Created struct {
	//[ 1] code
	//Code string `json:"code,omitempty"`
	//[ 2] customer_id
	CustomerId string `json:"customer_id,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//[ 3] salesman_id
	SalesmanId string `json:"salesman_id,omitempty" binding:"required" validate:"required"`
	//[ 4] projectman_id
	ProjectmanId string `json:"projectman_id,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `json:"machine_status_id,omitempty"`
	// 機台名稱
	ExtendTypeName string `json:"extend_type_name,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport bool `json:"eva_report,omitempty"`
	//[ 7] status
	Status string `json:"status,omitempty"`
	//[ 7] accept
	Accept bool `json:"accept,omitempty"`
	//[ 4] creater
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `json:"date_for_estimated_manager,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
}

type Field struct {
	//[ 0] cd_id
	CdId string `json:"cd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true" form:"cd_id"`
	//[ 1] code
	Code *string `json:"code,omitempty" form:"code"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//[ 2] customer_id
	CustomerId *string `json:"customer_id,omitempty" form:"customer_id"`
	//[ 3] salesman_id
	SalesmanId *string `json:"salesman_id,omitempty" form:"salesman_id"`
	//[ 4] projectman_id
	ProjectmanId *string `json:"projectman_id,omitempty" form:"projectman_id"`
	//[ 4] contact_person_id
	ContactPersonId *string `json:"contact_person_id,omitempty" form:"contact_person_id"`
	//[ 5] demand_content
	DemandContent *string `json:"demand_content,omitempty" form:"demand_content"`
	//[ 7] suitable_content
	SuitableContent *string `json:"suitable_content,omitempty" form:"suitable_content"`
	//[ 7] other_content
	OtherContent *string `json:"other_content,omitempty" form:"other_content"`
	//[ 7] budget
	Budget *string `json:"budget,omitempty" form:"budget"`
	//[ 7] date_for_recive
	DateForRecive *time.Time `json:"date_for_recive,omitempty" form:"date_for_recive"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart *time.Time `json:"date_for_estimated_start,omitempty" form:"date_for_estimated_start"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd *time.Time `json:"date_for_estimated_end,omitempty" form:"date_for_estimated_end"`
	//[ 7] date_for_actual_done
	DateForActualDone *time.Time `json:"date_for_actual_done,omitempty" form:"date_for_actual_done"`
	//[ 4] machine_status_id
	MachineStatusId *string `json:"machine_status_id,omitempty" form:"machine_status_id"`
	// 機台名稱
	ExtendTypeName string `json:"extend_type_name,omitempty" form:"extend_type_name"`
	//[ 7] extend_rem
	ExtendRem *string `json:"extend_rem,omitempty" form:"extend_rem"`
	//[ 7] date_for_devlop
	DateForDevlop *time.Time `json:"date_for_devlop,omitempty" form:"date_for_devlop"`
	//[ 7] est_quantity
	EstQuantity *int `json:"est_quantity,omitempty" form:"est_quantity"`
	//[ 7] eva_report
	EvaReport *bool `json:"eva_report,omitempty" form:"eva_report"`
	//[ 7] status
	Status *string `json:"status,omitempty" form:"status"`
	//[ 7] accept
	Accept *bool `json:"accept,omitempty" form:"accept"`
	//[ 4] creater
	Creater *string `json:"creater,omitempty" form:"creater"`
	//[ 7] create_time
	CreateTime *time.Time `json:"create_time,omitempty" form:"create_time"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `json:"date_for_estimated_manager,omitempty"`
	// 是否填料
	Fill bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type CRs struct {
	CustomerDemand []*struct {
		// customer_demend表
		//案件代號(客需單代號)
		CdID   string `json:"cd_id,omitempty"`
		CdCode string `json:"cd_code,omitempty"`
		//業務負責人(發文者)
		SalesmanID   string `json:"salesman_id,omitempty"`
		SalesmanName string `json:"salesman_name,omitempty"`
		//問題描述(客戶需求說明+適用製成說明)
		DemandContent   string `json:"demand_content,omitempty"`
		SuitableContent string `json:"suitable_content,omitempty"`
		// 是否填料
		Fill bool `json:"fill,omitempty"`
		// 結果狀態
		ResultStatus string `json:"result_status,omitempty"`
		// 結果說明
		ResultContent string `json:"result_content,omitempty"`
		// 再議日期
		DateForResult time.Time `json:"date_for_result,omitempty"`

		//task表
		//task_id
		TID string `json:"t_id,omitempty"`
		//收件日期(任務創建日期)
		CreateTime time.Time `json:"create_time,omitempty"`
		//task_user_id
		TuID string `json:"tu_id,omitempty"`
		//任務負責人
		TaskUserID   string `json:"task_user_id,omitempty"`
		TaskUserName string `json:"task_user_name,omitempty"`
		//狀態
		Status string `json:"status,omitempty"`
		//處理情形
		Remark string `json:"remark,omitempty"`

		//labor_hour表
		//工時明細
		HID string `json:"h_id,omitempty"`
	} `json:"customer_demand"`
	model.OutPage
}

type Customer_Demand_Accounts struct {
	CustomerDemand []*struct {
		//[ 0] cd_id
		CdId string `json:"cd_id,omitempty"`
		//[ 1] code
		Code string `json:"code,omitempty"`
		//[ 1] project_code
		ProjectId string `json:"project_id,omitempty"`
		//[ 1] project_code
		PCode string `json:"p_code,omitempty"`
		//[ 1] project_code
		PName string `json:"p_name,omitempty"`
		//[ 4] contact_person_id
		ContactPersonId string `json:"contact_person_id,omitempty"`
		//[ 5] demand_content
		DemandContent string `json:"demand_content,omitempty"`
		//[ 7] suitable_content
		SuitableContent string `json:"suitable_content,omitempty"`
		//[ 7] other_content
		OtherContent string `json:"other_content,omitempty"`
		//[ 7] budget
		Budget string `json:"budget,omitempty"`
		//[ 7] date_for_recive
		DateForRecive time.Time `json:"date_for_recive,omitempty"`
		//[ 7] date_for_estimated_start
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//[ 7] date_for_estimated_end
		DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
		//[ 7] date_for_actual_done
		DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
		//[ 4] machine_status_id
		MachineStatusId string `json:"machine_status_id,omitempty"`
		// 機台名稱
		ExtendTypeName string `json:"extend_type_name,omitempty"`
		//[ 7] extend_rem
		ExtendRem string `json:"extend_rem,omitempty"`
		//[ 7] date_for_devlop
		DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
		//[ 7] est_quantity
		EstQuantity int `json:"est_quantity,omitempty"`
		//[ 7] eva_report
		EvaReport bool `json:"eva_report,omitempty"`
		//[ 7] status
		Status string `json:"status,omitempty"`
		//[ 7] accept
		Accept bool `json:"accept,omitempty"`
		//[ 4] creater
		Creater string `json:"creater,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
		//[ 7] customer_id
		CustomerName string `json:"customer_name,omitempty"`
		//[ 8] salesman_id
		SalesmanName string `json:"salesman_name,omitempty"`
		//[ 9] projectman_id
		ProjectmanName string `json:"projectman_name,omitempty"`
		//[10] customer_id
		CustomerId string `json:"customer_id,omitempty"`
		//[11] salesman_id
		SalesmanId string `json:"salesman_id,omitempty"`
		//[12] projectman_id
		ProjectmanId string `json:"projectman_id,omitempty"`
		//[ 7] bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		// 是否填料
		Fill bool `json:"fill,omitempty"`
		// 結果狀態
		ResultStatus string `json:"result_status,omitempty"`
		// 結果說明
		ResultContent string `json:"result_content,omitempty"`
		// 再議日期
		DateForResult time.Time `json:"date_for_result,omitempty"`
	} `json:"customer_demand"`
	model.OutTotal
}

type List struct {
	CustomerDemand []*struct {
		//[ 0] cd_id
		CdId string `json:"cd_id,omitempty"`
		//[ 1] code
		Code string `json:"code,omitempty"`
		//[ 1] project_code
		ProjectId string `json:"project_id,omitempty"`
		//[ 2] customer_id
		CustomerId string `json:"customer_id,omitempty"`
		//[ 3] salesman_id
		SalesmanId string `json:"salesman_id,omitempty"`
		//[ 4] projectman_id
		ProjectmanId string `json:"projectman_id,omitempty"`
		//[ 4] contact_person_id
		ContactPersonId string `json:"contact_person_id,omitempty"`
		//[ 5] demand_content
		DemandContent string `json:"demand_content,omitempty"`
		//[ 7] suitable_content
		SuitableContent string `json:"suitable_content,omitempty"`
		//[ 7] other_content
		OtherContent string `json:"other_content,omitempty"`
		//[ 7] budget
		Budget string `json:"budget,omitempty"`
		//[ 7] date_for_recive
		DateForRecive time.Time `json:"date_for_recive,omitempty"`
		//[ 7] date_for_estimated_start
		DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
		//[ 7] date_for_estimated_end
		DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
		//[ 7] date_for_actual_done
		DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
		//[ 4] machine_status_id
		MachineStatusId string `json:"machine_status_id,omitempty"`
		// 機台名稱
		ExtendTypeName string `json:"extend_type_name,omitempty"`
		//[ 7] extend_rem
		ExtendRem string `json:"extend_rem,omitempty"`
		//[ 7] date_for_devlop
		DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
		//[ 7] est_quantity
		EstQuantity int `json:"est_quantity,omitempty"`
		//[ 7] eva_report
		EvaReport bool `json:"eva_report,omitempty"`
		//[ 7] status
		Status string `json:"status,omitempty"`
		//[ 7] accept
		Accept bool `json:"accept,omitempty"`
		//[ 4] creater
		Creater string `json:"creater,omitempty"`
		//[ 7] create_time
		CreateTime time.Time `json:"create_time,omitempty"`
		//[ 7] bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		//[ 7] date_for_estimated_manager
		DateForEstimatedManager time.Time `json:"date_for_estimated_manager,omitempty"`
		// 送審日
		DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
		// 是否填料
		Fill bool `json:"fill,omitempty"`
		// 結果狀態
		ResultStatus string `json:"result_status,omitempty"`
		// 結果說明
		ResultContent string `json:"result_content,omitempty"`
		// 再議日期
		DateForResult time.Time `json:"date_for_result,omitempty"`
	} `json:"customer_demand"`
	model.OutPage
}

type Updated struct {
	//[ 0] cd_id
	CdId string `json:"cd_id,omitempty"`
	//[ 1] code
	Code *string `json:"code,omitempty"`
	//[ 1] project_code
	ProjectId string `json:"project_id,omitempty"`
	//[ 2] customer_id
	CustomerId *string `json:"customer_id,omitempty"`
	//[ 3] salesman_id
	SalesmanId string `json:"salesman_id,omitempty"`
	//[ 4] projectman_id
	ProjectmanId string `json:"projectman_id,omitempty"`
	//[ 4] contact_person_id
	ContactPersonId string `json:"contact_person_id,omitempty"`
	//[ 5] demand_content
	DemandContent string `json:"demand_content,omitempty"`
	//[ 7] suitable_content
	SuitableContent string `json:"suitable_content,omitempty"`
	//[ 7] other_content
	OtherContent string `json:"other_content,omitempty"`
	//[ 7] budget
	Budget string `json:"budget,omitempty"`
	//[ 7] date_for_recive
	DateForRecive time.Time `json:"date_for_recive,omitempty"`
	//[ 7] date_for_estimated_start
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//[ 7] date_for_estimated_end
	DateForEstimatedEnd time.Time `json:"date_for_estimated_end,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`
	//[ 4] machine_status_id
	MachineStatusId string `json:"machine_status_id,omitempty"`
	// 機台名稱
	ExtendTypeName string `json:"extend_type_name,omitempty"`
	//[ 7] extend_rem
	ExtendRem string `json:"extend_rem,omitempty"`
	//[ 7] date_for_devlop
	DateForDevlop time.Time `json:"date_for_devlop,omitempty"`
	//[ 7] est_quantity
	EstQuantity int `json:"est_quantity,omitempty"`
	//[ 7] eva_report
	EvaReport *bool `json:"eva_report,omitempty"`
	//[ 7] status
	Status string `json:"status,omitempty"`
	//[ 7] accept
	Accept *bool `json:"accept,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] date_for_estimated_manager
	DateForEstimatedManager time.Time `json:"date_for_estimated_manager,omitempty"`
	// 是否填料
	Fill *bool `json:"fill,omitempty"`
	// 結果狀態
	ResultStatus string `json:"result_status,omitempty"`
	// 結果說明
	ResultContent string `json:"result_content,omitempty"`
	// 再議日期
	DateForResult time.Time `json:"date_for_result,omitempty"`
}

type Updated_Bonita struct {
	//[ 0] cd_id
	CdId string `json:"cd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 送審日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
	// // 創建者
	// BonitaTaskID float32 `json:"bonita_task_id,omitempty"`
	// // 狀態
	// Status string `json:"status,omitempty"`
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
	return "customer_demand"
}

func (a *Create_Table) TableName() string {
	return "customer_demand"
}
