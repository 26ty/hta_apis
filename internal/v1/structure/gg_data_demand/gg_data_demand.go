package gg_data_demand

import (
	"time"
)

type Review struct {
	//C2
	// manufacture_order表
	// 編號
	MID string `json:"m_id,omitempty"`
	//單號(客需單或專案任務)
	ProjectID string `json:"project_id,omitempty"`
	ProjectCode string `json:"project_code,omitempty"`
	ProjectDetail string `json:"project_detail,omitempty"`
	CustomerDemandCode string `json:"customer_demand_code,omitempty"`
	//製令、專案單號
	Code string `json:"code,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//主見品號
	OrderName string `json:"order_name,omitempty"`
	//需求數量
	Amount string `json:"amount,omitempty"`
	//客戶名稱(製令、專案、客需)
	CustomerID string `json:"customer_id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	//出貨地點
	ShipmentLocation string `json:"shipment_location,omitempty"`
	//製令開啟期限
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//製令關閉期限
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//預計出貨日期
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
	//內部訂單編號(製令、專案、客需)
	InnerID string `json:"inner_id,omitempty"`
	//其他相關單據
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//製令、專案備註
	Remark string `json:"remark,omitempty"`
	//發文者
	Creater string `json:"creater,omitempty"`
	CreaterName string `json:"creater_name,omitempty"`
	//發文單位
	
	//業務助理
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	SalesAssistantName string `json:"sales_assistant_name,omitempty"`
	//收文者
	RecipientID string `json:"recipient_id,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
	//業務負責人
	SalesmanID string `json:"salesman_id,omitempty"`
	SalesmanName string `json:"salesman_name,omitempty"`
	//副本

	//B2(Tm_Return)
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
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
	DateForEstimatedStart time.Time `json:"date_for_estimated_start,omitempty"`
	//實際完成日
	DateForActualCompletion time.Time `json:"date_for_actual_completion,omitempty"`
	//預計完成日
	DateForEstimatedCompletion time.Time `json:"date_for_estimated_completion,omitempty"`
	//前一預計開始日
	LastDateForEstimatedStart time.Time `json:"last_date_for_estimated_start,omitempty"`
	//前一實際完成日
	LastDateForActualCompletion time.Time `json:"last_date_for_actual_completion,omitempty"`
	//前一預計完成日
	LastDateForEstimatedCompletion time.Time `json:"last_date_for_estimated_completion,omitempty"`

	//專案編號
	PID string `json:"p_id,omitempty"`
	//專案代號
	PCode string `json:"p_code,omitempty"`
	//專案名稱
	PName string `json:"p_name,omitempty"`
	//專案負責人
	ProjectmanName string `json:"projectman_name,omitempty"`
	//客服負責人
	ServicemanName string `json:"serviceman_name,omitempty"`
	// 專案起始日
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 專案結束日
	DateForEnd time.Time `json:"date_for_end,omitempty"`
	// 機台總成品號
	MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
	// 機台英文名稱
	MachineEnglish string `json:"machine_english,omitempty"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`
	//銷貨單號
	OrderNumber string `json:"order_number,omitempty"`

	// 編號
	TuID string `json:"tu_id,omitempty"`
	// 送出日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
	// 任務負責人
	Name string `json:"name,omitempty"`

	//B2(Project_Account)
	// customer_code
	CustomerCode string `json:"customer_code,omitempty"`
	// 類別
	Type string `json:"type,omitempty"`
	// 客服編號
	ServicemanID string `json:"serviceman_id,omitempty"`
	// 專案負責人名稱
	ProjectmanID string `json:"projectman_id,omitempty"`
	// 付款日
	DateForPay time.Time `json:"date_for_pay,omitempty"`
	// 驗收日
	DateForCheck time.Time `json:"date_for_check,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// machine_quantity
	MachineQuantity int `json:"machine_quantity,omitempty"`
	// is_template
	IsTemplate bool `json:"is_template,omitempty"`

	//A1(Customer_Reivew)
	//customer_demand表
	//編號
	CdID string `json:"cd_id,omitempty"`
	//日期
	DateForRecive string `json:"date_for_recive,omitempty"`
	//客戶聯絡人
	ContactPersonID string `json:"contact_person_id,omitempty"`
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
	// //發文單位
	// SalesmanDep []*Dep `json:"salesman_dep,omitempty"`
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
	UserID string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	//編號
	CuID string `json:"cu_id,omitempty"`
	//預定完成日(員工)
	DateForEstimatedCompletionEmployee time.Time `json:"date_for_estimated_completion_employee,omitempty"`
	//完工送審(員工)
	DateForCompletionEmployee time.Time `json:"date_for_completion_employee,omitempty"`

	//A1(Customer_Review_Task)
	//編號
	TCode string `json:"t_code,omitempty"`
	//預計開始日
	TaskDateForEstimatedStart *time.Time `json:"task_date_for_estimated_start,omitempty"`
	//[ 7] date_for_actual_done
	DateForActualDone time.Time `json:"date_for_actual_done,omitempty"`

	//Bonita
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string  `json:"bonita_task_name,omitempty"`
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

type Reviews struct {
	Review
	//發文單位
	SalesmanDep []*Dep `json:"salesman_dep,omitempty"`
}
