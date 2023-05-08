package factory_manufacturings

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//廠別生產列表編號
	FmID string `gorm:"<-:create;primaryKey;column:fm_id;type:UUID;default:uuid_generate_v4();" json:"fm_id,omitempty"`
	//廠別編號
	FactoryID string `gorm:"column:factory_id;type:UUID;not null;" json:"factory_id,omitempty"`
	//機型
	MachineID string `gorm:"column:machine_id;type:UUID;not null;" json:"machine_id,omitempty"`
	//專案
	ProjectID string `gorm:"column:project_id;type:UUID;not null;" json:"project_id,omitempty"`
	//製令
	OrderID string `gorm:"column:order_id;type:UUID;not null;" json:"order_id,omitempty"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//廠別生產列表編號
	FmID string `json:"fm_id,omitempty"`
	//廠別編號
	FactoryID string `json:"factory_id,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案
	ProjectID string `json:"project_id,omitempty"`
	//製令
	OrderID string `json:"order_id,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//廠別生產列表編號
	FmID string `json:"fm_id,omitempty"`
	//廠別編號
	FactoryID string `json:"factory_id,omitempty"`
	//機型
	MachineID string `json:"machine_id,omitempty"`
	//專案
	ProjectID string `json:"project_id,omitempty"`
	//製令
	OrderID string `json:"order_id,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//廠別編號
	FactoryID string `json:"factory_id,omitempty" binding:"required,uuid4" validate:"required"`
	//機型
	MachineID string `json:"machine_id,omitempty" binding:"required,uuid4" validate:"required"`
	//專案
	ProjectID string `json:"project_id,omitempty" binding:"required,uuid4" validate:"required"`
	//製令
	OrderID string `json:"order_id,omitempty" binding:"required,uuid4" validate:"required"`
}

type Created_List struct {
	Manufacturing []*Created `json:"manufacturing"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//廠別生產列表編號
	FmID string `json:"fm_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//機型
	MachineID string `json:"machine_id,omitempty" from:"machine_id" binding:"omitempty,uuid4"`
	//專案
	ProjectID string `json:"project_id,omitempty" from:"project_id" binding:"omitempty,uuid4"`
	//製令
	OrderID string `json:"order_id,omitempty" from:"order_id" binding:"omitempty,uuid4"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	FactoryManufacturings []*struct {
		//廠別生產列表編號
		FmID string `json:"fm_id,omitempty"`
		//廠別編號
		FactoryID string `json:"factory_id,omitempty"`
		//機型
		MachineID string `json:"machine_id,omitempty"`
		//專案
		ProjectID string `json:"project_id,omitempty"`
		//製令
		OrderID string `json:"order_id,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"factory_manufacturings"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//廠別生產列表編號
	FmID string `json:"fm_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//機型
	MachineID string `json:"machine_id,omitempty" binding:"omitempty,uuid4"`
	//專案
	ProjectID string `json:"project_id,omitempty" binding:"omitempty,uuid4"`
	//製令
	OrderID string `json:"order_id,omitempty" binding:"omitempty,uuid4"`
}

func (a *Table) TableName() string {
	return "factory_manufacturing"
}
