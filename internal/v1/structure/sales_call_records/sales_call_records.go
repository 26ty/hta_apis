package sales_call_records

import (
	"time"

	model "eirc.app/internal/v1/structure"
	account "eirc.app/internal/v1/structure/accounts"
)

type Table struct {
	//業務顧客拜訪紀錄編號
	SID string `gorm:"<-:create;primaryKey;column:s_id;type:UUID;default:uuid_generate_v4();" json:"s_id,omitempty"`
	//客戶/廠別
	CustomerName string `gorm:"column:customer_name;type:TEXT;not null;" json:"customer,omitempty"`
	//業務人員
	SalesmanID string `gorm:"column:salesman_id;type:UUID;not null;" json:"salesman_id,omitempty"`

	Account_s account.Table `gorm:"foreignkey:salesman_id;references:account_id"`
	//機器狀態
	MachineStatusID string `gorm:"column:machine_status_id;type:TEXT;not null;" json:"machine_status_id,omitempty"`
	//拜訪時間
	StartTime *time.Time `gorm:"column:start_time;type:TIMESTAMP;not null;" json:"start,omitempty"`
	//機型
	MachineID string `gorm:"column:machine_id;type:UUID;" json:"machine_id,omitempty"`
	//專案
	ProjectID string `gorm:"column:project_id;type:UUID;" json:"project_id,omitempty"`
	//製令
	OrderID string `gorm:"column:order_id;type:UUID;" json:"order_id,omitempty"`
	//拜訪紀錄
	Record string `gorm:"column:record;type:TEXT;not null;" json:"record,omitempty"`
	//創建者
	Creater string `gorm:"<-:create;column:creater;type:UUID;not null;" json:"creater,omitempty"`

	Account account.Table `gorm:"foreignkey:creater;references:account_id"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//業務顧客拜訪紀錄編號
	SID string `json:"s_id,omitempty"`
	//客戶/廠別
	CustomerName string `json:"customer,omitempty"`
	//業務人員
	SalesmanID string `json:"salesman_id,omitempty"`
	//機器狀態
	MachineStatusID string ` json:"machine_status_id,omitempty"`
	//拜訪時間
	StartTime *time.Time `json:"start,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案
	ProjectID string `json:"project_id,omitempty"`
	//製令
	OrderID string `json:"order_id,omitempty"`
	//拜訪紀錄
	Record string `json:"record,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//業務顧客拜訪紀錄編號
	SID string `json:"s_id,omitempty"`
	//客戶/廠別
	CustomerName string `json:"customer,omitempty"`
	//業務人員
	SalesmanID string `json:"salesman_id,omitempty"`
	//機器狀態
	MachineStatusID string ` json:"machine_status_id,omitempty"`
	//拜訪時間
	StartTime *time.Time `json:"start,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案
	ProjectID string `json:"project_id,omitempty"`
	//製令
	OrderID string `json:"order_id,omitempty"`
	//拜訪紀錄
	Record string `json:"record,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//客戶/廠別
	CustomerName string `json:"customer,omitempty" binding:"required" valistart:"required"`
	//業務人員
	SalesmanID string `json:"salesman_id,omitempty" binding:"required,uuid4" valistart:"required"`
	//機器狀態
	MachineStatusID string ` json:"machine_status_id,omitempty" binding:"required" valistart:"required"`
	//拜訪時間
	StartTime *time.Time `json:"start,omitempty" binding:"required" valistart:"required"`
	//機型
	MachineID string `json:"machine_id,omitempty" valistart:"required"`
	//專案
	ProjectID string `json:"project_id,omitempty" valistart:"required"`
	//製令
	OrderID string `json:"order_id,omitempty" valistart:"required"`
	//拜訪紀錄
	Record string `json:"record,omitempty" binding:"required" valistart:"required"`
	//創建者
	Creater string `json:"creater,omitempty" binding:"required,uuid4" valistart:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//業務顧客拜訪紀錄編號
	SID string `json:"s_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//客戶/廠別
	CustomerName string `json:"customer,omitempty" from:"customer_name"`
	//機器狀態
	MachineStatusID string ` json:"machine_status_id,omitempty" from:"machine_status_id"`
	//拜訪時間
	StartTime *time.Time `json:"start,omitempty" from:"start_time"`
	//機型
	MachineID string `json:"machine_id,omitempty" from:"machine_id" binding:"omitempty,uuid4"`
	//專案
	ProjectID string `json:"project_id,omitempty" from:"project_id" binding:"omitempty,uuid4"`
	//製令
	OrderID string `json:"order_id,omitempty" from:"order_id" binding:"omitempty,uuid4"`
	//拜訪紀錄
	Record string `json:"record,omitempty" from:"record"`
}

type Fields struct {
	Field
	model.InPage
}

type SalesCallRecord_Account struct {
	//業務顧客拜訪紀錄編號
	SID string `json:"s_id,omitempty"`
	//客戶/廠別
	CustomerName string `json:"customer,omitempty"`
	//業務人員
	Account_s__name string `json:"salesman,omitempty"`
	//機器狀態
	MachineStatusID string ` json:"machine_status_id,omitempty"`
	//拜訪時間
	StartTime *time.Time `json:"start,omitempty"`
	//標題
	//Title string `json:"title,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案
	ProjectID string `json:"project_id,omitempty"`
	//製令
	OrderID string `json:"order_id,omitempty"`
	//拜訪紀錄
	Record string `json:"record,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建者
	Account__name string `json:"creater_name,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type List struct {
	SalesCallRecords []*SalesCallRecord_Account `json:"sales_call_records"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//客戶/廠別
	CustomerName string `json:"customer,omitempty"`
	//業務顧客拜訪紀錄編號
	SID string `json:"s_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//機器狀態
	MachineStatusID string ` json:"machine_status_id,omitempty"`
	//拜訪時間
	StartTime *time.Time `json:"start,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty" binding:"omitempty,uuid4"`
	//專案
	ProjectID string `json:"project_id,omitempty" binding:"omitempty,uuid4"`
	//製令
	OrderID string `json:"order_id,omitempty" binding:"omitempty,uuid4"`
	//拜訪紀錄
	Record string `json:"record,omitempty"`
}

func (a *Table) TableName() string {
	return "sales_call_record"
}
