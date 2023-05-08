package jig_demands

import (
	"time"

	model "eirc.app/internal/v1/structure"
	account "eirc.app/internal/v1/structure/accounts"
	jig_demand_detail "eirc.app/internal/v1/structure/jig_demand_details"
)

type Table struct {
	//治具需求單編號
	JID string `gorm:"<-:create;primaryKey;column:j_id;type:UUID;not null;default:uuid_generate_v4();" json:"j_id,omitempty"`

	Detail []jig_demand_detail.Table `gorm:"foreignkey:jig_id;references:j_id" json:"detail"`
	//單號
	JCode string `gorm:"->;column:j_code;type:TEXT;not null;default:add_jig_demand_code()" json:"j_code,omitempty"`
	//分類
	Kind string `gorm:"column:kind;type:TEXT;not null;" json:"kind,omitempty"`
	//類別
	Type string `gorm:"column:type;type:TEXT;not null;" json:"type,omitempty"`
	//待覆日期
	DateForRespond *time.Time `gorm:"column:date_for_respond;type:DATE;not null;" json:"date_for_respond,omitempty"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `gorm:"column:date_for_respond_of_limit;type:DATE;" json:"date_for_respond_of_limit,omitempty"`
	//客戶需求日
	DateForDemand *time.Time `gorm:"column:date_for_demand;type:DATE;not null;" json:"date_for_demand,omitempty"`
	//標題
	Title string `gorm:"column:title;type:TEXT;not null;" json:"title,omitempty"`
	//客戶名稱
	CustomerName string `gorm:"column:customer_name;type:TEXT;" json:"customer_name,omitempty"`
	//電話
	Tel string `gorm:"column:tel;type:TEXT;" json:"tel,omitempty"`
	//聯絡人
	Liaison string `gorm:"column:liaison;type:TEXT;" json:"liaison,omitempty"`
	//電子郵件
	Mail string `gorm:"column:mail;type:TEXT;" json:"mail,omitempty"`
	//機型
	MachineID string `gorm:"column:machine_id;type:UUID;" json:"machine_id,omitempty"`
	//專案代號
	ProjectCode string `gorm:"column:project_code;type:TEXT;" json:"project_code,omitempty"`
	//是否急件
	Urgent *bool `gorm:"column:urgent;type:bool;default:false;not null;" json:"urgent,omitempty"`
	//治具數量
	JigQuantity int `gorm:"column:jig_quantity;type:INT4;not null;" json:"jig_quantity,omitempty"`
	//摘要說明
	Summary string `gorm:"column:summary;type:TEXT;" json:"summary,omitempty"`
	//客戶通知日
	DateForNotify *time.Time `gorm:"column:date_for_notify;type:DATE;" json:"date_for_notify,omitempty"`
	//客圖/資訊
	DateForInformation *time.Time `gorm:"column:date_for_information;type:DATE;" json:"date_for_information,omitempty"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `gorm:"column:estimated_date_for_delivery;type:DATE;" json:"estimated_date_for_delivery,omitempty"`
	//PODate
	PODate *time.Time `gorm:"column:po_date;type:DATE;" json:"po_date,omitempty"`
	//內部訂單
	InnerOrder string `gorm:"column:inner_order;type:TEXT;" json:"inner_order,omitempty"`
	//贈送單號
	GiftCode string `gorm:"column:gift_code;type:TEXT;" json:"gift_code,omitempty"`
	//項目
	Item string `gorm:"column:item;type:TEXT;" json:"item,omitempty"`
	//客戶單號
	PoNo string `gorm:"column:po_no;type:TEXT;" json:"po_no,omitempty"`
	//規格
	Standard string `gorm:"column:standard;type:TEXT;" json:"standard,omitempty"`
	//申請狀態
	Status string `gorm:"column:status;type:TEXT;not null;" json:"status,omitempty"`
	//創建者
	Creater string `gorm:"<-:create;column:creater;type:UUID;not null;" json:"creater,omitempty"`

	Account account.Table `gorm:"foreignkey:creater;references:account_id"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
}

type Base struct {
	//治具需求單編號
	JID string `json:"j_id,omitempty"`

	Detail []jig_demand_detail.Base `json:"detail"`
	//單號
	JCode string `json:"j_code,omitempty"`
	//分類
	Kind string `json:"kind,omitempty"`
	//類別
	Type string `json:"type,omitempty"`
	//待覆日期
	DateForRespond *time.Time `json:"date_for_respond,omitempty"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty"`
	//客戶需求日
	DateForDemand *time.Time ` json:"date_for_demand,omitempty"`
	//標題
	Title string `json:"title,omitempty"`
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//治具數量
	JigQuantity int `json:"jig_quantity,omitempty"`
	//摘要說明
	Summary string `json:"summary,omitempty"`
	//客戶通知日
	DateForNotify *time.Time `json:"date_for_notify,omitempty"`
	//客圖/資訊
	DateForInformation *time.Time `json:"date_for_information,omitempty"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty"`
	//PODate
	PODate *time.Time `json:"po_date,omitempty"`
	//內部訂單
	InnerOrder string `json:"inner_order,omitempty"`
	//贈送單號
	GiftCode string `json:"gift_code,omitempty"`
	//項目
	Item string `json:"item,omitempty"`
	//客戶單號
	PoNo string `json:"po_no,omitempty"`
	//規格
	Standard string `json:"standard,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Single struct {
	//治具需求單編號
	JID string `json:"j_id,omitempty"`
	//單號
	JCode string `json:"j_code,omitempty"`
	//分類
	Kind string `json:"kind,omitempty"`
	//類別
	Type string `json:"type,omitempty"`
	//待覆日期
	DateForRespond *time.Time `json:"date_for_respond,omitempty"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty"`
	//客戶需求日
	DateForDemand *time.Time ` json:"date_for_demand,omitempty"`
	//標題
	Title string `json:"title,omitempty"`
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//治具數量
	JigQuantity int `json:"jig_quantity,omitempty"`
	//摘要說明
	Summary string `json:"summary,omitempty"`
	//客戶通知日
	DateForNotify *time.Time `json:"date_for_notify,omitempty"`
	//客圖/資訊
	DateForInformation *time.Time `json:"date_for_information,omitempty"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty"`
	//PODate
	PODate *time.Time `json:"po_date,omitempty"`
	//內部訂單
	InnerOrder string `json:"inner_order,omitempty"`
	//贈送單號
	GiftCode string `json:"gift_code,omitempty"`
	//項目
	Item string `json:"item,omitempty"`
	//客戶單號
	PoNo string `json:"po_no,omitempty"`
	//規格
	Standard string `json:"standard,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

// 放create時需輸入的欄位
type Created struct {
	//分類
	Kind string `json:"kind,omitempty" binding:"required" validate:"required"`
	//類別
	Type string `json:"type,omitempty" validate:"required"`
	//待覆日期
	DateForRespond *time.Time `json:"date_for_respond,omitempty" binding:"required" validate:"required"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty" validate:"required"`
	//客戶需求日
	DateForDemand *time.Time ` json:"date_for_demand,omitempty" binding:"required" validate:"required"`
	//標題
	Title string `json:"title,omitempty" binding:"required" validate:"required"`
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty" validate:"required"`
	//電話
	Tel string `json:"tel,omitempty" validate:"required"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" validate:"required"`
	//電子郵件
	Mail string `json:"mail,omitempty" validate:"required"`
	//機型
	MachineID string `json:"machine_id,omitempty" validate:"required"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty" validate:"required"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty" validate:"required"`
	//治具數量
	JigQuantity int `json:"jig_quantity,omitempty" validate:"required"`
	//摘要說明
	Summary string `json:"summary,omitempty" validate:"required"`
	//客戶通知日
	DateForNotify *time.Time `json:"date_for_notify,omitempty" validate:"required"`
	//客圖/資訊
	DateForInformation *time.Time `json:"date_for_information,omitempty" validate:"required"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty" validate:"required"`
	//PODate
	PODate *time.Time `json:"po_date,omitempty" validate:"required"`
	//內部訂單
	InnerOrder string `json:"inner_order,omitempty" validate:"required"`
	//贈送單號
	GiftCode string `json:"gift_code,omitempty" validate:"required"`
	//項目
	Item string `json:"item,omitempty" validate:"required"`
	//客戶單號
	PoNo string `json:"po_no,omitempty" validate:"required"`
	//規格
	Standard string `json:"standard,omitempty" validate:"required"`
	//申請狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	//創建者
	Creater string `json:"creater,omitempty" binding:"required,uuid4" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//治具需求單編號
	JID string `json:"j_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//分類
	Kind string `json:"kind,omitempty" from:"kind"`
	//類別
	Type string `json:"type,omitempty" from:"type"`
	//待覆日期
	DateForRespond *time.Time `json:"date_for_respond,omitempty" from:"date_for_respond"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty" from:"date_for_respond_of_limit"`
	//客戶需求日
	DateForDemand *time.Time ` json:"date_for_demand,omitempty" from:"date_for_demand"`
	//標題
	Title string `json:"title,omitempty" from:"title"`
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty" from:"customer_name"`
	//電話
	Tel string `json:"tel,omitempty" from:"tel"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" from:"liaison"`
	//電子郵件
	Mail string `json:"mail,omitempty" from:"mail"`
	//機型
	MachineID string `json:"machine_id,omitempty" from:"machine_id" binding:"omitempty,uuid4"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty" from:"project_code"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty" from:"urgent"`
	//治具數量
	JigQuantity int `json:"jig_quantity,omitempty" from:"jig_quantity"`
	//摘要說明
	Summary string `json:"summary,omitempty" from:"summary"`
	//客戶通知日
	DateForNotify *time.Time `json:"date_for_notify,omitempty" from:"date_for_notify"`
	//客圖/資訊
	DateForInformation *time.Time `json:"date_for_information,omitempty" from:"date_for_information"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty" from:"estimated_date_for_delivery"`
	//PODate
	PODate *time.Time `json:"po_date,omitempty" from:"po_date"`
	//內部訂單
	InnerOrder string `json:"inner_order,omitempty" from:"inner_order"`
	//贈送單號
	GiftCode string `json:"gift_code,omitempty" from:"gift_code"`
	//項目
	Item string `json:"item,omitempty" from:"item"`
	//客戶單號
	PoNo string `json:"po_no,omitempty" from:"po_no"`
	//規格
	Standard string `json:"standard,omitempty" from:"standard"`
	//申請狀態
	Status string `json:"status,omitempty" from:"status"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" from:"bonita_case_id"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type JigDemand_Account struct {
	//治具需求單編號
	JID string `json:"j_id,omitempty"`
	//單號
	JCode string `json:"j_code,omitempty"`
	//分類
	Kind string `json:"kind,omitempty"`
	//類別
	Type string `json:"type,omitempty"`
	//待覆日期
	DateForRespond *time.Time `json:"date_for_respond,omitempty"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty"`
	//客戶需求日
	DateForDemand *time.Time ` json:"date_for_demand,omitempty"`
	//標題
	Title string `json:"title,omitempty"`
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//治具數量
	JigQuantity int `json:"jig_quantity,omitempty"`
	//摘要說明
	Summary string `json:"summary,omitempty"`
	//客戶通知日
	DateForNotify *time.Time `json:"date_for_notify,omitempty"`
	//客圖/資訊
	DateForInformation *time.Time `json:"date_for_information,omitempty"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty"`
	//PODate
	PODate *time.Time `json:"po_date,omitempty"`
	//內部訂單
	InnerOrder string `json:"inner_order,omitempty"`
	//贈送單號
	GiftCode string `json:"gift_code,omitempty"`
	//項目
	Item string `json:"item,omitempty"`
	//客戶單號
	PoNo string `json:"po_no,omitempty"`
	//規格
	Standard string `json:"standard,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建者
	Account__name string `json:"creater_name,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type JD struct {
	//jig_demand表
	//治具需求單編號
	JID string `json:"j_id,omitempty"`
	//單號
	JCode string `json:"j_code,omitempty"`
	//業務負責人(salesman_id發文者creater)
	SalesmanID   string `json:"salesman_id,omitempty"`
	SalesmanName string `json:"salesman_name,omitempty"`
	//問題描述(標題+摘要說明)
	Title   string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`

	//task表
	//task_id
	TID string `json:"t_id,omitempty"`
	//收件日期(任務創建日期)
	CreatedTime time.Time `json:"created_time,omitempty"`

	//task_user表
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

type JDs struct {
	JigDemand []*JD `json:"jig_demands"`
	model.OutPage
}

type JigDetail struct {
	JigDemand_Account
	Detail []jig_demand_detail.Base `json:"detail"`
}

type AllJigDetail struct {
	JigDemand []*JigDetail `json:"jig_demands"`
	model.OutPage
}

type SearchJigDemand struct {
	SearchJigDemand []*struct {
		//治具需求單編號
		JID string `json:"j_id,omitempty"`
		//單號
		JCode string `json:"j_code,omitempty"`
		//待覆日期
		DateForRespond *time.Time `json:"date_for_respond,omitempty"`
		//客戶需求日
		DateForDemand *time.Time ` json:"date_for_demand,omitempty"`
		//標題
		Title string `json:"title,omitempty"`
		//客戶名稱
		CustomerName string `json:"customer_name,omitempty"`
		//電話
		Tel string `json:"tel,omitempty"`
		//專案代號
		ProjectCode string `json:"project_code,omitempty"`
		//客戶單號
		PoNo string `json:"po_no,omitempty"`
		//申請狀態
		Status string `json:"status,omitempty"`
		//創建者
		Account__name string `json:"creater,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"jig_demands"`
	model.OutPage
}

type List struct {
	JigDemands []*struct {
		//治具需求單編號
		JID string `json:"j_id,omitempty"`
		//單號
		JCode string `json:"j_code,omitempty"`
		//分類
		Kind string `json:"kind,omitempty"`
		//類別
		Type string `json:"type,omitempty"`
		//待覆日期
		DateForRespond *time.Time `json:"date_for_respond,omitempty"`
		//最晚待覆日期
		DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty"`
		//客戶需求日
		DateForDemand *time.Time ` json:"date_for_demand,omitempty"`
		//標題
		Title string `json:"title,omitempty"`
		//客戶名稱
		CustomerName string `json:"customer_name,omitempty"`
		//電話
		Tel string `json:"tel,omitempty"`
		//聯絡人
		Liaison string `json:"liaison,omitempty"`
		//電子郵件
		Mail string `json:"mail,omitempty"`
		//機型
		MachineID string `json:"machine_id,omitempty"`
		//專案代號
		ProjectCode string `json:"project_code,omitempty"`
		//是否急件
		Urgent *bool `json:"urgent,omitempty"`
		//治具數量
		JigQuantity int `json:"jig_quantity,omitempty"`
		//摘要說明
		Summary string `json:"summary,omitempty"`
		//客戶通知日
		DateForNotify *time.Time `json:"date_for_notify,omitempty"`
		//客圖/資訊
		DateForInformation *time.Time `json:"date_for_information,omitempty"`
		//預計出貨日
		EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty"`
		//PODate
		PODate *time.Time `json:"po_date,omitempty"`
		//內部訂單
		InnerOrder string `json:"inner_order,omitempty"`
		//贈送單號
		GiftCode string `json:"gift_code,omitempty"`
		//項目
		Item string `json:"item,omitempty"`
		//客戶單號
		PoNo string `json:"po_no,omitempty"`
		//規格
		Standard string `json:"standard,omitempty"`
		//申請狀態
		Status string `json:"status,omitempty"`
		//創建者
		Creater string `json:"creater,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
		//bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	} `json:"jig_demands"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//治具需求單編號
	JID string `json:"j_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//分類
	Kind string `json:"kind,omitempty"`
	//類別
	Type string `json:"type,omitempty"`
	//待覆日期
	DateForRespond *time.Time `json:"date_for_respond,omitempty"`
	//最晚待覆日期
	DateForRespondOfLimit *time.Time `json:"date_for_respond_of_limit,omitempty"`
	//客戶需求日
	DateForDemand *time.Time ` json:"date_for_demand,omitempty"`
	//標題
	Title string `json:"title,omitempty"`
	//客戶名稱
	CustomerName string `json:"customer_name,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty" binding:"omitempty,uuid4"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//治具數量
	JigQuantity int `json:"jig_quantity,omitempty"`
	//摘要說明
	Summary string `json:"summary,omitempty"`
	//客戶通知日
	DateForNotify *time.Time `json:"date_for_notify,omitempty"`
	//客圖/資訊
	DateForInformation *time.Time `json:"date_for_information,omitempty"`
	//預計出貨日
	EstimatedDateForDelivery *time.Time `json:"estimated_date_for_delivery,omitempty"`
	//PODate
	PODate *time.Time `json:"po_date,omitempty"`
	//內部訂單
	InnerOrder string `json:"inner_order,omitempty"`
	//贈送單號
	GiftCode string `json:"gift_code,omitempty"`
	//項目
	Item string `json:"item,omitempty"`
	//客戶單號
	PoNo string `json:"po_no,omitempty"`
	//規格
	Standard string `json:"standard,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Updated_Bonita struct {
	//治具需求單編號
	JID string `json:"j_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//CaseID
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

func (a *Table) TableName() string {
	return "jig_demand"
}
