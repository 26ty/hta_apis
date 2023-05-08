package manufacture_order

import (
	"time"

	model "eirc.app/internal/v1/structure"
	ManufactureUserModel "eirc.app/internal/v1/structure/manufacture_user"

)

type Table struct {
	// 編號
	MID string `gorm:"primary_key;column:m_id;uuid_generate_v4()type:UUID;" json:"m_id,omitempty"`
	//
	Code string `gorm:"column:code;type:TEXT;" json:"code,omitempty"`
	//
	ProjectID string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	//
	OrderName string `gorm:"column:order_name;type:TEXT;" json:"order_name,omitempty"`
	//
	Amount int `gorm:"column:amount;type:INT;" json:"amount,omitempty"`
	//
	ShipmentLocation string `gorm:"column:shipment_location;type:TEXT;" json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `gorm:"column:date_for_open;type:TIMESTAMPTZ;" json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `gorm:"column:date_for_close;type:TIMESTAMPTZ;" json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `gorm:"column:date_for_estimated_shipment;type:TIMESTAMPTZ;" json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
	//
	CopyFile string `gorm:"column:copy_file;type:UUID;" json:"copy_file,omitempty"`
	//
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//
	SalesAssistantID string `gorm:"column:sales_assistant_id;type:UUID;" json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `gorm:"column:recipient_id;type:UUID;" json:"recipient_id,omitempty"`
	//
	SalesmanID string `gorm:"column:salesman_id;type:UUID;" json:"salesman_id,omitempty"`
	//
	InnerID string `gorm:"column:inner_id;type:text;" json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `gorm:"column:other_document_code;type:text;" json:"other_document_code,omitempty"`
	//
	CustomerID string `gorm:"column:customer_id;type:uuid;" json:"customer_id,omitempty"`
	//
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	//[ 7] bonita_case_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	// project_detail
	ProjectDetail string `gorm:"column:project_detail;type:TEXT;" json:"project_detail,omitempty"`
}

type Create_Table struct {
	// 編號
	MID string `gorm:"primary_key;column:m_id;uuid_generate_v4()type:UUID;" json:"m_id,omitempty"`
	//
	ProjectID string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	//
	OrderName string `gorm:"column:order_name;type:TEXT;" json:"order_name,omitempty"`
	//
	Amount int `gorm:"column:amount;type:INT;" json:"amount,omitempty"`
	//
	ShipmentLocation string `gorm:"column:shipment_location;type:TEXT;" json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `gorm:"column:date_for_open;type:TIMESTAMPTZ;" json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `gorm:"column:date_for_close;type:TIMESTAMPTZ;" json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `gorm:"column:date_for_estimated_shipment;type:TIMESTAMPTZ;" json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `gorm:"column:status;type:TEXT;" json:"status,omitempty"`
	//
	CopyFile string `gorm:"column:copy_file;type:UUID;" json:"copy_file,omitempty"`
	//
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//
	SalesAssistantID string `gorm:"column:sales_assistant_id;type:UUID;" json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `gorm:"column:recipient_id;type:UUID;" json:"recipient_id,omitempty"`
	//
	SalesmanID string `gorm:"column:salesman_id;type:UUID;" json:"salesman_id,omitempty"`
	//
	InnerID string `gorm:"column:inner_id;type:text;" json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `gorm:"column:other_document_code;type:text;" json:"other_document_code,omitempty"`
	//
	CustomerID string `gorm:"column:customer_id;type:uuid;" json:"customer_id,omitempty"`
	//
	Creater string `gorm:"column:creater;type:UUID;" json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	//[ 7] bonita_case_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
	// project_detail
	ProjectDetail string `gorm:"column:project_detail;type:TEXT;" json:"project_detail,omitempty"`
}

type Review struct {
	// manufacture_order表
	// 編號
	MID string `json:"m_id,omitempty"`
	//單號(客需單或專案任務)
	ProjectID string `json:"project_id,omitempty"`
	ProjectCode string `json:"project_code,omitempty"`
	ProjectDetail string `json:"project_detail,omitempty"`
	CustomerDemandCode string `json:"customer_demand_code,omitempty"`
	//製令單號
	Code string `json:"code,omitempty"`
	// 狀態
	Status string `json:"status,omitempty"`
	//主見品號
	OrderName string `json:"order_name,omitempty"`
	//需求數量
	Amount string `json:"amount,omitempty"`
	//客戶名稱
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
	//內部訂單編號
	InnerID string `json:"inner_id,omitempty"`
	//其他相關單據
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//備註
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

	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string  `json:"bonita_task_name,omitempty"`
}

type One struct {
	// 編號
	MID string `json:"m_id,omitempty"`
	//
	Code string `json:"code,omitempty"`
	//
	ProjectID string `json:"project_id,omitempty"`
	ProjectCode string `json:"project_code,omitempty"`
	ProjectDetail string `json:"project_detail,omitempty"`
	CustomerDemandCode string `json:"customer_demand_code,omitempty"`
	//
	OrderName string `json:"order_name,omitempty"`
	//
	Amount int `json:"amount,omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty"`
	//
	Remark string `json:"remark,omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	// 中文名稱
	CreaterName string `json:"creater_name,omitempty"`
}

type Ones struct {
	// 編號
	MID string `json:"m_id,omitempty"`
	//
	Code string `json:"code,omitempty"`
	//
	ProjectID string `json:"project_id,omitempty"`
	ProjectCode string `json:"project_code,omitempty"`
	ProjectDetail string `json:"project_detail,omitempty"`
	CustomerDemandCode string `json:"customer_demand_code,omitempty"`
	//
	OrderName string `json:"order_name,omitempty"`
	//
	Amount int `json:"amount,omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty"`
	//
	Remark string `json:"remark,omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	// 中文名稱
	CreaterName string `json:"creater_name,omitempty"`
	//副本參與人
	ManufactureUser []*ManufactureUserModel.ManufactureAccount `json:"manufacture_user"`
}

type Base struct {
	// 編號
	MID string `json:"m_id,omitempty"`
	//
	Code string `json:"code,omitempty"`
	//
	ProjectID string `json:"project_id,omitempty"`
	//
	OrderName string `json:"order_name,omitempty"`
	//
	Amount int `json:"amount,omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty"`
	//
	Remark string `json:"remark,omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	// 中文名稱
	CreaterName string `json:"creater_name,omitempty"`
	// project_detail
	ProjectDetail string `json:"project_detail,omitempty"`
}

type Single struct {
	// 編號
	MID string `json:"m_id,omitempty"`
	//
	Code string `json:"code,omitempty"`
	//
	ProjectID string `json:"project_id,omitempty"`
	//
	OrderName string `json:"order_name,omitempty"`
	//
	Amount int `json:"amount,omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty"`
	//
	Remark string `json:"remark,omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	// 中文名稱
	CreaterName string `json:"creater_name,omitempty"`
	// project_detail
	ProjectDetail string `json:"project_detail,omitempty"`
}

type Created struct {
	//
	Code string `json:"code,omitempty" `
	//
	ProjectID string `json:"project_id,omitempty" binding:"required,uuid4" validate:"required"`
	//
	OrderName string `json:"order_name,omitempty" `
	//
	Amount int `json:"amount,omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty"`
	//
	Remark string `json:"remark,omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty" `
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	// project_detail
	ProjectDetail string `json:"project_detail,omitempty"`
}

// type ManufactureOrder_Cd_Account struct {
// 	// 編號
// 	MID string `json:"m_id,omitempty"`
// 	//
// 	Code string `json:"code,omitempty"`
// 	//
// 	ProjectID string `json:"project_id,omitempty"`
// 	//
// 	OrderName string `json:"order_name,omitempty"`
// 	//
// 	DateForOpen time.Time `json:"date_for_open,omitempty"`
// 	//
// 	DateForClose time.Time `json:"date_for_close,omitempty"`
// 	//
// 	Status string `json:"status,omitempty"`
// 	//
// 	RecipientID string `json:"recipient_id,omitempty"`
// 	//
// 	SalesmanID string `json:"salesman_id,omitempty"`
// 	//
// 	Creater string `json:"creater,omitempty"`
// 	//
// 	SalesmanName string `json:"salesman_name,omitempty"`
// 	//
// 	RecipientName string `json:"recipient_name,omitempty"`
// 	//
// 	CreaterName string `json:"creater_name,omitempty"`
// 	//
// 	CdCode string `json:"cd_code,omitempty"`
// 	//[ 7] bonita_case_id
// 	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
// }

type ManufactureOrder_Project_Account struct {
	// 編號
	MID string `json:"m_id,omitempty"`
	//
	Code string `json:"code,omitempty"`
	//
	ProjectID string `json:"project_id,omitempty"`
	//
	OrderName string `json:"order_name,omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//
	Creater string `json:"creater,omitempty"`
	//
	SalesmanName string `json:"salesman_name,omitempty"`
	//
	RecipientName string `json:"recipient_name,omitempty"`
	//
	CreaterName string `json:"creater_name,omitempty"`
	//
	PCode string `json:"p_code,omitempty"`
	// project_detail
	ProjectDetail string `json:"project_detail,omitempty"`
	//
	CdCode string `json:"cd_code,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Field struct {
	// 編號
	MID string `json:"m_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	Code string `json:"code,omitempty" form:"code" binding:"omitempty"`
	//
	ProjectID string `json:"project_id,omitempty" form:"project_id" binding:"omitempty,uuid4"`
	//
	OrderName string `json:"order_name,omitempty"  form:"order_name" binding:"omitempty"`
	//
	Amount int `json:"amount,omitempty"  form:"amount" binding:"omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"  form:"shipment_location" binding:"omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"  form:"date_for_open" binding:"omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"  form:"date_for_close" binding:"omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"  form:"date_for_estimated_shipment" binding:"omitempty"`
	//
	Status string `json:"status,omitempty"  form:"status" binding:"omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty"  form:"copy_file"`
	//
	Remark string `json:"remark,omitempty" form:"remark" binding:"omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	// bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" form:"bonita_case_id"`
	// project_detail
	ProjectDetail string `json:"project_detail,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

// type ManufactureOrder_Cd_Accounts struct {
// 	ManufactureOrder []*struct {
// 		// 編號
// 		MID string `json:"m_id,omitempty"`
// 		//
// 		Code string `json:"code,omitempty"`
// 		//
// 		ProjectID string `json:"project_id,omitempty"`
// 		//
// 		OrderName string `json:"order_name,omitempty"`
// 		//
// 		DateForOpen time.Time `json:"date_for_open,omitempty"`
// 		//
// 		DateForClose time.Time `json:"date_for_close,omitempty"`
// 		//
// 		Status string `json:"status,omitempty"`
// 		//
// 		RecipientID string `json:"recipient_id,omitempty"`
// 		//
// 		SalesmanID string `json:"salesman_id,omitempty"`
// 		//
// 		Creater string `json:"creater,omitempty"`
// 		//
// 		SalesmanName string `json:"salesman_name,omitempty"`
// 		//
// 		RecipientName string `json:"recipient_name,omitempty"`
// 		//
// 		CreaterName string `json:"creater_name,omitempty"`
// 		//
// 		CdCode string `json:"cd_code,omitempty"`
// 		// 
// 		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
// 	} `json:"manufacture_order"`
// 	model.OutPage
// }

type ManufactureOrder_Project_Accounts struct {
	ManufactureOrder []*struct {
		// 編號
		MID string `json:"m_id,omitempty"`
		//
		Code string `json:"code,omitempty"`
		//
		ProjectID string `json:"project_id,omitempty"`
		//
		OrderName string `json:"order_name,omitempty"`
		//
		DateForOpen time.Time `json:"date_for_open,omitempty"`
		//
		DateForClose time.Time `json:"date_for_close,omitempty"`
		//
		Status string `json:"status,omitempty"`
		//
		RecipientID string `json:"recipient_id,omitempty"`
		//
		SalesmanID string `json:"salesman_id,omitempty"`
		//
		Creater string `json:"creater,omitempty"`
		//
		SalesmanName string `json:"salesman_name,omitempty"`
		//
		RecipientName string `json:"recipient_name,omitempty"`
		//
		CreaterName string `json:"creater_name,omitempty"`
		//
		PCode string `json:"p_code,omitempty"`
		// project_detail
		ProjectDetail string `json:"project_detail,omitempty"`
		//
		CdCode string `json:"cd_code,omitempty"`
		// 
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	} `json:"manufacture_order"`
	model.OutPage
}

type List struct {
	ManufactureOrder []*struct {
		// 編號
		MID string `json:"m_id,omitempty"`
		//
		Code string `json:"code,omitempty"`
		//
		ProjectID string `json:"project_id,omitempty"`
		//
		OrderName string `json:"order_name,omitempty"`
		//
		Amount int `json:"amount,omitempty"`
		//
		ShipmentLocation string `json:"shipment_location,omitempty"`
		//
		DateForOpen time.Time `json:"date_for_open,omitempty"`
		//
		DateForClose time.Time `json:"date_for_close,omitempty"`
		//
		DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"`
		//
		Status string `json:"status,omitempty"`
		//
		CopyFile string `json:"copy_file,omitempty"`
		//
		Remark string `json:"remark,omitempty"`
		//
		SalesAssistantID string `json:"sales_assistant_id,omitempty"`
		//
		RecipientID string `json:"recipient_id,omitempty"`
		//
		SalesmanID string `json:"salesman_id,omitempty"`
		//
		InnerID string `json:"inner_id,omitempty"`
		//
		OtherDocumentCode string `json:"other_document_code,omitempty"`
		//
		CustomerID string `json:"customer_id,omitempty"`
		// 
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
		//
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 中文名稱
		CreaterName string `json:"creater_name,omitempty"`
		// project_detail
		ProjectDetail string `json:"project_detail,omitempty"`
	} `json:"manufacture_order"`
	model.OutPage
}

type Updated struct {
	// 編號
	MID string `json:"m_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	Code string `json:"code,omitempty" form:"code" binding:"omitempty"`
	//
	ProjectID string `json:"project_id,omitempty"`
	//
	OrderName string `json:"order_name,omitempty"  form:"order_name" binding:"omitempty"`
	//
	Amount int `json:"amount,omitempty"  form:"amount" binding:"omitempty"`
	//
	ShipmentLocation string `json:"shipment_location,omitempty"  form:"shipment_location" binding:"omitempty"`
	//
	DateForOpen time.Time `json:"date_for_open,omitempty"  form:"date_for_open" binding:"omitempty"`
	//
	DateForClose time.Time `json:"date_for_close,omitempty"  form:"date_for_close" binding:"omitempty"`
	//
	DateForEstimatedShipment time.Time `json:"date_for_estimated_shipment,omitempty"  form:"date_for_estimated_shipment" binding:"omitempty"`
	//
	Status string `json:"status,omitempty"  form:"status" binding:"omitempty"`
	//
	CopyFile string `json:"copy_file,omitempty" `
	//
	Remark string `json:"remark,omitempty" form:"remark" binding:"omitempty"`
	//
	SalesAssistantID string `json:"sales_assistant_id,omitempty"`
	//
	RecipientID string `json:"recipient_id,omitempty"`
	//
	SalesmanID string `json:"salesman_id,omitempty"`
	//
	InnerID string `json:"inner_id,omitempty"`
	//
	OtherDocumentCode string `json:"other_document_code,omitempty"`
	//
	CustomerID string `json:"customer_id,omitempty"`
	// 
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// project_detail
	ProjectDetail string `json:"project_detail,omitempty"`
}

type Updated_Bonita struct {
	// 編號
	MID string `json:"m_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	// // 創建者
	// BonitaTaskID float32 `json:"bonita_task_id,omitempty"`
	// // 狀態
	// Status string `json:"status,omitempty"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "manufacture_order"
}

func (a *Create_Table) TableName() string {
	return "manufacture_order"
}

