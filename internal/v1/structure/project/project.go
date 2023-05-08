package project

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [uuid_generate_v4()]
	PID string `gorm:"primary_key;column:p_id;uuid_generate_v4()type:UUID;" json:"p_id,omitempty"`
	// 單號                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Code string `gorm:"column:code;type:TEXT;" json:"code,omitempty"`
	// 類別                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Type string `gorm:"column:type;type:TEXT;" json:"type,omitempty"`
	// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `gorm:"column:origin_id;type:UUID;" json:"origin_id,omitempty"`
	// 名字                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PName string `gorm:"column:p_name;type:TEXT;" json:"p_name,omitempty"`
	// 客戶單編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	CustomerID string `gorm:"column:customer_id;type:UUID;" json:"customer_id,omitempty"`
	// 業務編號                                    UUID                 null: false  primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	SalesmanID string `gorm:"column:salesman_id;type:UUID;" json:"salesman_id,omitempty"`
	// 客服編號                                  UUID                 null: false  primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	ServicemanID string `gorm:"column:serviceman_id;type:UUID;" json:"serviceman_id,omitempty"`
	// 專案負責人編號                                  UUID                 null: false  primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	ProjectmanID string `gorm:"column:projectman_id;type:UUID;" json:"projectman_id,omitempty"`
	// 狀態                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
	// 內部單號                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	InnerID string `gorm:"column:inner_id;type:TEXT;" json:"inner_id,omitempty"`
	// 專案起始日                                 DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForStart time.Time `gorm:"column:date_for_start;type:DATE;" json:"date_for_start,omitempty"`
	// 專案結束日                                   DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForEnd time.Time `gorm:"column:date_for_end;type:DATE;" json:"date_for_end,omitempty"`
	// 付款日                                  DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForPay time.Time `gorm:"column:date_for_pay;type:DATE;" json:"date_for_pay,omitempty"`
	// 出機日                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForDelivery time.Time `gorm:"column:date_for_delivery;type:DATE;" json:"date_for_delivery,omitempty"`
	// 驗收日                                DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	DateForCheck time.Time `gorm:"column:date_for_check;type:DATE;" json:"date_for_check,omitempty"`
	// 創建者                                        UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間                                    TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: [now()]
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
	//[ 7] bonita_case_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	// customer_code
	CustomerCode string `gorm:"column:customer_code;type:TEXT;" json:"customer_code,omitempty"`
	// machine_finished_number
	MachineFinishedNumber string `gorm:"column:machine_finished_number;type:TEXT;" json:"machine_finished_number,omitempty"`
	// jig_quantity
	JigQuantity int `gorm:"column:jig_quantity;type:INT4;" json:"jig_quantity,omitempty"`
	// machine_english
	MachineEnglish string `gorm:"column:machine_english;type:TEXT;" json:"machine_english,omitempty"`
	// machine_quantity
	MachineQuantity int `gorm:"column:machine_quantity;type:INT4;" json:"machine_quantity,omitempty"`
	// external_order
	ExternalOrder string `gorm:"column:external_order;type:TEXT;" json:"external_order,omitempty"`
	// internal_order
	InternalOrder string `gorm:"column:internal_order;type:TEXT;" json:"internal_order,omitempty"`
	// summary_description
	SummaryDescription string `gorm:"column:summary_description;type:TEXT;" json:"summary_description,omitempty"`
	// is_template
	IsTemplate bool `gorm:"column:is_template;type:Boolean;" json:"is_template,omitempty"`
	//銘板
	Nameplate string `gorm:"column:nameplate;type:TEXT;" json:"nameplate,omitempty"`
	//銷貨單號
	OrderNumber string `gorm:"column:order_number;type:TEXT;" json:"order_number,omitempty"`
}


type GetCaseListInput struct {
	model.GonitaUser
	model.GonitaCaseList
}


type Tm_Return struct {
	//編號
	TID string `json:"t_id,omitempty"`
	//單據來源編號
	DocumentsID string `json:"documents_id,omitempty"`
	//名字
	TName string `json:"t_name,omitempty"`
	//編號
	Code string `json:"code,omitempty"`
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
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty"`
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
	//內部訂單
	InnerID string `json:"inner_id,omitempty"`
	

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string  `json:"bonita_task_name,omitempty"`

	// 編號
	TuID string `json:"tu_id,omitempty"`
	// 送出日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
	// 任務負責人
	Name string `json:"name,omitempty"`
}


type Project_Account struct {
	TId string `json:"t_id,omitempty"`
	
	// 編號
	PID string `json:"p_id,omitempty"`
	// 單號
	Code string `json:"code,omitempty"`
	// customer_code
	CustomerCode string `json:"customer_code,omitempty"`
	// 類別
	Type string `json:"type,omitempty"`
	// 名字
	PName string `json:"p_name,omitempty"`
	// 客戶單編號
	CustomerID string `json:"customer_id,omitempty"`
	// 業務編號
	SalesmanID string `json:"salesman_id,omitempty"`
	// 客服編號
	ServicemanID string `json:"serviceman_id,omitempty"`
	// 專案負責人名稱
	ProjectmanID string `json:"projectman_id,omitempty"`
	// 客戶單名稱
	CustomerName string `json:"customer_name,omitempty"`
	// 業務名稱
	SalesmanName string `json:"salesman_name,omitempty"`
	// 客服名稱
	ServicemanName string `json:"serviceman_name,omitempty"`
	// 專案負責人名稱
	ProjectmanName string `json:"projectman_name,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//銷貨單號
	OrderNumber string `json:"order_number,omitempty"`
	//內部訂單
	InnerID string `json:"inner_id,omitempty"`
	// 專案起始日
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 專案結束日
	DateForEnd time.Time `json:"date_for_end,omitempty"`
	// 付款日
	DateForPay time.Time `json:"date_for_pay,omitempty"`
	// 出機日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
	// 驗收日
	DateForCheck time.Time `json:"date_for_check,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// machine_quantity
	MachineQuantity int `json:"machine_quantity,omitempty"`
	// is_template
	IsTemplate bool `json:"is_template,omitempty"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string  `json:"bonita_task_name,omitempty"`
}

type Bonita_ID_List struct {
	//編號
	PID string `json:"p_id,omitempty"`
	//名字
	ProjectId string `json:"projectman_id,omitempty"`
	//名字
	ProjectmanName string `json:"projectman_name,omitempty"`
	//名字
	ProjectmanBonitaId string `json:"projectman_bonita_id,omitempty"`
}

type Bonita_ID_Lists struct {
	Project []*struct {
		//編號
		PID string `json:"p_id,omitempty"`
		//名字
		ProjectId string `json:"projectman_id,omitempty"`
		//名字
		ProjectmanName string `json:"projectman_name,omitempty"`
		//名字
		ProjectmanBonitaId string `json:"projectman_bonita_id,omitempty"`
	} `json:"project"`
	model.OutPage
}

type Base struct {
	// 編號
	PID string `json:"p_id,omitempty"`
	// 單號
	Code string `json:"code,omitempty"`
	// 類別
	Type string `json:"type,omitempty"`
	// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `json:"origin_id,omitempty"`
	// 名字
	PName string `json:"p_name,omitempty"`
	// 客戶單編號
	CustomerID string `json:"customer_id,omitempty"`
	// 業務編號
	SalesmanID string `json:"salesman_id,omitempty"`
	// 客服編號
	ServicemanID string `json:"serviceman_id,omitempty"`
	// 專案負責人編號
	ProjectmanID string `json:"projectman_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//銷貨單號
	OrderNumber string `json:"order_number,omitempty"`
	//內部訂單
	InnerID string `json:"inner_id,omitempty"`
	// 專案起始日
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 專案結束日
	DateForEnd time.Time `json:"date_for_end,omitempty"`
	// 付款日
	DateForPay time.Time `json:"date_for_pay,omitempty"`
	// 出機日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
	// 驗收日
	DateForCheck time.Time `json:"date_for_check,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// customer_code
	CustomerCode string `json:"customer_code,omitempty"`
	// machine_finished_number
	MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
	// jig_quantity
	JigQuantity int `json:"jig_quantity,omitempty"`
	// machine_english
	MachineEnglish string `json:"machine_english,omitempty"`
	// machine_quantity
	MachineQuantity int `json:"machine_quantity,omitempty"`
	// external_order
	ExternalOrder string `json:"external_order,omitempty"`
	// internal_order
	InternalOrder string `json:"internal_order,omitempty"`
	// summary_description
	SummaryDescription string `json:"summary_description,omitempty"`
	// is_template
	IsTemplate bool `json:"is_template,omitempty"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`
}

type Single struct {
	// 編號
	PID string `json:"p_id,omitempty"`
	// 單號
	Code string `json:"code,omitempty"`
	// 類別
	Type string `json:"type,omitempty"`
	// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `json:"origin_id,omitempty"`
	// 名字
	PName string `json:"p_name,omitempty"`
	// 客戶單編號
	CustomerID string `json:"customer_id,omitempty"`
	// 業務編號
	SalesmanID string `json:"salesman_id,omitempty"`
	// 客服編號
	ServicemanID string `json:"serviceman_id,omitempty"`
	// 專案負責人編號
	ProjectmanID string `json:"projectman_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//銷貨單號
	OrderNumber string `json:"order_number,omitempty"`
	//內部訂單
	InnerID string `json:"inner_id,omitempty"`
	// 專案起始日
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 專案結束日
	DateForEnd time.Time `json:"date_for_end,omitempty"`
	// 付款日
	DateForPay time.Time `json:"date_for_pay,omitempty"`
	// 出機日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
	// 驗收日
	DateForCheck time.Time `json:"date_for_check,omitempty"`
	// 創建者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// customer_code
	CustomerCode string `json:"customer_code,omitempty"`
	// machine_finished_number
	MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
	// jig_quantity
	JigQuantity int `json:"jig_quantity,omitempty"`
	// machine_english
	MachineEnglish string `json:"machine_english,omitempty"`
	// machine_quantity
	MachineQuantity int `json:"machine_quantity,omitempty"`
	// external_order
	ExternalOrder string `json:"external_order,omitempty"`
	// internal_order
	InternalOrder string `json:"internal_order,omitempty"`
	// summary_description
	SummaryDescription string `json:"summary_description,omitempty"`
	// is_template
	IsTemplate bool `json:"is_template,omitempty"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`
}

type Created struct {
	// 單號
	Code string `json:"code,omitempty" binding:"required" validate:"required"`
	// 類別
	Type string `json:"type,omitempty" binding:"required" validate:"required"`
	// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `json:"origin_id,omitempty"`
	// 名字
	PName string `json:"p_name,omitempty" binding:"required" validate:"required"`
	// 客戶單編號
	CustomerID string `json:"customer_id,omitempty" binding:"required" validate:"required"`
	// 業務編號
	SalesmanID string `json:"salesman_id,omitempty" binding:"required" validate:"required"`
	// 客服編號
	ServicemanID string `json:"serviceman_id,omitempty" binding:"required" validate:"required"`
	// 專案負責人編號
	ProjectmanID string `json:"projectman_id,omitempty" binding:"required" validate:"required"`
	// 狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	//銷貨單號
	OrderNumber string `json:"order_number,omitempty"`
	//內部訂單
	InnerID string `json:"inner_id,omitempty"`
	// 專案起始日
	DateForStart time.Time `json:"date_for_start,omitempty" `
	// 專案結束日
	DateForEnd time.Time `json:"date_for_end,omitempty" `
	// 付款日
	DateForPay time.Time `json:"date_for_pay,omitempty" `
	// 出機日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty" `
	// 驗收日
	DateForCheck time.Time `json:"date_for_check,omitempty" `
	// 創建者
	Creater string `json:"creater,omitempty" binding:"required" validate:"required"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`

	// customer_code 待新增（客需單號）必填
	CustomerCode string `json:"customer_code,omitempty"`
	// machine_finished_number 待新增（機台總成品號）必填
	MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
	// jig_quantity 待新增（治具數量）必填
	JigQuantity int `json:"jig_quantity,omitempty"`
	// machine_english 待新增（機台英文名稱）
	MachineEnglish string `json:"machine_english,omitempty"`
	// machine_quantity 待新增（機台數量）
	MachineQuantity int `json:"machine_quantity,omitempty"`
	// external_order 待新增（外部訂單）
	ExternalOrder string `json:"external_order,omitempty"`
	// internal_order 待新增（內部訂單）
	InternalOrder string `json:"internal_order,omitempty"`
	// summary_description 待新增（摘要說明）
	SummaryDescription string `json:"summary_description,omitempty"`
}

type Field struct {
	// 編號
	PID string `json:"p_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 單號
	Code *string `json:"code,omitempty" form:"code"`
	// 類別
	Type *string `json:"type,omitempty" form:"type"`
	// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `json:"origin_id,omitempty"`
	// 名字
	PName *string `json:"p_name,omitempty" form:"p_name"`
	// 客戶單編號
	CustomerID *string `json:"customer_id,omitempty" form:"customer_id" `
	// 業務編號
	SalesmanID *string `json:"salesman_id,omitempty" form:"salesman_id"`
	// 客服編號
	ServicemanID *string `json:"serviceman_id,omitempty" form:"serviceman_id"`
	// 專案負責人編號
	ProjectmanID *string `json:"projectman_id,omitempty" form:"projectman_id"`
	// 狀態
	Status *string `json:"status,omitempty" form:"status"`
	//銷貨單號
	OrderNumber *string `json:"order_number,omitempty" form:"order_number"`
	//內部訂單
	InnerID *string `json:"inner_id,omitempty" form:"inner_id"`
	// 專案起始日
	DateForStart *time.Time `json:"date_for_start,omitempty" form:"date_for_start"`
	// 專案結束日
	DateForEnd *time.Time `json:"date_for_end,omitempty" form:"date_for_end"`
	// 付款日
	DateForPay *time.Time `json:"date_for_pay,omitempty" form:"date_for_pay"`
	// 出機日
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty" form:"date_for_delivery"`
	// 驗收日
	DateForCheck *time.Time `json:"date_for_check,omitempty" form:"date_for_check"`
	// bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" form:"bonita_case_id"`
	// customer_code 待新增（客需單號）必填
	CustomerCode string `json:"customer_code,omitempty" form:"customer_code"`
	// machine_finished_number 待新增（機台總成品號）必填
	MachineFinishedNumber string `json:"machine_finished_number,omitempty" form:"machine_finished_number"`
	// jig_quantity 待新增（治具數量）必填
	JigQuantity int `json:"jig_quantity,omitempty" form:"jig_quantity"`
	// machine_english 待新增（機台英文名稱）
	MachineEnglish string `json:"machine_english,omitempty" form:"machine_english"`
	// machine_quantity 待新增（機台數量）
	MachineQuantity int `json:"machine_quantity,omitempty" form:"machine_quantity"`
	// external_order 待新增（外部訂單）
	ExternalOrder string `json:"external_order,omitempty" form:"external_order"`
	// internal_order 待新增（內部訂單）
	InternalOrder string `json:"internal_order,omitempty" form:"internal_order"`
	// summary_description 待新增（摘要說明）
	SummaryDescription string `json:"summary_description,omitempty" form:"summary_description"`
	// is_template
	IsTemplate bool `json:"is_template,omitempty" form:"is_template"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`
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
	Project []*struct {
		// 編號
		PID string `json:"p_id,omitempty"`
		// 單號
		Code string `json:"code,omitempty"`
		// 類別
		Type string `json:"type,omitempty"`
		// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
		OriginID string `json:"origin_id,omitempty"`	
		// 名字
		PName string `json:"p_name,omitempty"`
		// 客戶單編號
		CustomerID string `json:"customer_id,omitempty" form:"customer_id"`
		// 業務編號
		SalesmanID string `json:"salesman_id,omitempty" form:"salesman_id"`
		// 客服編號
		ServicemanID string `json:"serviceman_id,omitempty"`
		// 專案負責人編號
		ProjectmanID string `json:"projectman_id,omitempty"`
		// 狀態
		Status string `json:"status,omitempty"`
		//銷貨單號
		OrderNumber string `json:"order_number,omitempty" `
		//內部訂單
		InnerID string `json:"inner_id,omitempty"`
		// 專案起始日
		DateForStart time.Time `json:"date_for_start,omitempty"`
		// 專案結束日
		DateForEnd time.Time `json:"date_for_end,omitempty"`
		// 付款日
		DateForPay time.Time `json:"date_for_pay,omitempty"`
		// 出機日
		DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
		// 驗收日
		DateForCheck time.Time `json:"date_for_check,omitempty"`
		// 創建者
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 創建者
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		// customer_code 待新增（客需單號）必填
		CustomerCode string `json:"customer_code,omitempty"`
		// machine_finished_number 待新增（機台總成品號）必填
		MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
		// jig_quantity 待新增（治具數量）必填
		JigQuantity int `json:"jig_quantity,omitempty"`
		// machine_english 待新增（機台英文名稱）
		MachineEnglish string `json:"machine_english,omitempty"`
		// machine_quantity 待新增（機台數量）
		MachineQuantity int `json:"machine_quantity,omitempty"`
		// external_order 待新增（外部訂單）
		ExternalOrder string `json:"external_order,omitempty"`
		// internal_order 待新增（內部訂單）
		InternalOrder string `json:"internal_order,omitempty"`
		// summary_description 待新增（摘要說明）
		SummaryDescription string `json:"summary_description,omitempty"`
		// is_template
		IsTemplate bool `json:"is_template,omitempty"`
		//銘板
		Nameplate string `json:"nameplate,omitempty"`
	} `json:"project"`
	model.OutPage
}

type Project_Accounts struct {
	Project []*struct {
		TId string `json:"t_id,omitempty"`

		// 編號
		PID string `json:"p_id,omitempty"`
		// 單號
		Code string `json:"code,omitempty"`
		// customer_code
		CustomerCode string `json:"customer_code,omitempty"`
		// 類別
		Type string `json:"type,omitempty"`
		// 名字
		PName string `json:"p_name,omitempty"`
		// 客戶單編號
		CustomerID string `json:"customer_id,omitempty"`
		// 業務編號
		SalesmanID string `json:"salesman_id,omitempty"`
		// 客服編號
		ServicemanID string `json:"serviceman_id,omitempty"`
		// 專案負責人名稱
		ProjectmanID string `json:"projectman_id,omitempty"`
		// 客戶單名稱
		CustomerName string `json:"customer_name,omitempty"`
		// 業務名稱
		SalesmanName string `json:"salesman_name,omitempty"`
		// 客服名稱
		ServicemanName string `json:"serviceman_name,omitempty"`
		// 專案負責人名稱
		ProjectmanName string `json:"projectman_name,omitempty"`
		// 狀態
		Status string `json:"status,omitempty"`
		//銷貨單號
		OrderNumber string `json:"order_number,omitempty" `
		//內部訂單
		InnerID string `json:"inner_id,omitempty"`
		// 專案起始日
		DateForStart time.Time `json:"date_for_start,omitempty"`
		// 專案結束日
		DateForEnd time.Time `json:"date_for_end,omitempty"`
		// 付款日
		DateForPay time.Time `json:"date_for_pay,omitempty"`
		// 出機日
		DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
		// 驗收日
		DateForCheck time.Time `json:"date_for_check,omitempty"`
		// 創建者
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 創建者
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		// machine_quantity
		MachineQuantity int `json:"machine_quantity,omitempty"`
		// is_template
		IsTemplate bool `json:"is_template,omitempty"`
		//銘板
		Nameplate string `json:"nameplate,omitempty"`
	} `json:"project"`
	model.OutPage
}

type Updated struct {
	// 編號
	PID string `json:"p_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 單號
	Code string `json:"code,omitempty"`
	// 類別
	Type string `json:"type,omitempty"`
	// 來源編號                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	OriginID string `json:"origin_id,omitempty"`
	// 名字
	PName string `json:"p_name,omitempty"`
	// 客戶單編號
	CustomerID string `json:"customer_id,omitempty" `
	// 業務編號
	SalesmanID string `json:"salesman_id,omitempty"`
	// 客服編號
	ServicemanID string `json:"serviceman_id,omitempty"`
	// 專案負責人編號
	ProjectmanID string `json:"projectman_id,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//銷貨單號
	OrderNumber string `json:"order_number" `
	//內部訂單
	InnerID string `json:"inner_id"`
	// 專案起始日
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 專案結束日
	DateForEnd time.Time `json:"date_for_end,omitempty"`
	// 付款日
	DateForPay time.Time `json:"date_for_pay,omitempty"`
	// 出機日
	DateForDelivery time.Time `json:"date_for_delivery,omitempty"`
	// 驗收日
	DateForCheck time.Time `json:"date_for_check,omitempty"`
	// 創建者
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// 待新增（客需單號）必填
	CustomerCode string `json:"customer_code,omitempty"`
	// 待新增（機台總成品號）必填
	MachineFinishedNumber string `json:"machine_finished_number,omitempty"`
	// 待新增（治具數量）必填
	JigQuantity int `json:"jig_quantity,omitempty"`
	// 待新增（機台英文名稱）
	MachineEnglish string `json:"machine_english,omitempty"`
	// 待新增（機台數量）
	MachineQuantity int `json:"machine_quantity,omitempty"`
	// 待新增（外部訂單）
	ExternalOrder string `json:"external_order,omitempty"`
	// 待新增（內部訂單）
	InternalOrder string `json:"internal_order,omitempty"`
	// 待新增（摘要說明）
	SummaryDescription string `json:"summary_description,omitempty"`
	// is_template
	IsTemplate *bool `json:"is_template,omitempty"`
	//銘板
	Nameplate string `json:"nameplate,omitempty"`
}
type Updated_Bonita struct {
	// 編號
	PID string `json:"p_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// // 創建者
	// BonitaTaskID float32 `json:"bonita_task_id,omitempty"`
	// // 狀態
	// Status string `json:"status,omitempty"`
}

func (a *Table) TableName() string {
	return "project"
}
