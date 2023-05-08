package gateway_data

import (
	"time"
	model "eirc.app/internal/v1/structure"
)

// Table struct is companies database table struct
type Table struct {
	// 編號
	GdID string `gorm:"primaryKey;column:gd_id;uuid_generate_v4()type:UUID;" json:"gd_id,omitempty"`
	// 類別(A1、B2、C2)
	Classification string `gorm:"column:classification;type:UUID;" json:"classification,omitempty"`
	// 標題(單位主管審核)
	Title string `gorm:"column:title;type:TEXT;" json:"title,omitempty"`
	// 資料需求(預存程序名稱)
	DataDemand string `gorm:"column:data_demand;type:TEXT;" json:"data_demand,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	GdID string `json:"gd_id,omitempty"`
	// 類別(A1、B2、C2)
	Classification string `json:"classification,omitempty"`
	// 標題(單位主管審核)
	Title string `json:"title,omitempty"`
	// 資料需求(預存程序名稱)
	DataDemand string `json:"data_demand,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	GdID string `json:"gd_id,omitempty"`
	// 類別(A1、B2、C2)
	Classification string `json:"classification,omitempty"`
	// 標題(單位主管審核)
	Title string `json:"title,omitempty"`
	// 資料需求(預存程序名稱)
	DataDemand string `json:"data_demand,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 類別(A1、B2、C2)
	Classification string `json:"classification,omitempty" binding:"required" validate:"required"`
	// 標題(單位主管審核)
	Title string `json:"title,omitempty" binding:"required" validate:"required"`
	// 資料需求(預存程序名稱)
	DataDemand string `json:"data_demand,omitempty" binding:"required" validate:"required"`
}

type Field struct {
	// 編號
	GdID string `json:"gd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 類別(A1、B2、C2)
	Classification string `json:"classification,omitempty" form:"classification" `
	// 標題(單位主管審核)
	Title string `json:"title,omitempty" form:"title" `
	// 資料需求(預存程序名稱)
	DataDemand string `json:"data_demand,omitempty" form:"data_demand" `
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	GatewayData []*struct {
		// 編號
		GdID string `json:"gd_id,omitempty"`
		// 類別(A1、B2、C2)
		Classification string `json:"classification,omitempty"`
		// 標題(單位主管審核)
		Title string `json:"title,omitempty"`
		// 資料需求(預存程序名稱)
		DataDemand string `json:"data_demand,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"gateway_data"`
	model.OutPage
}

type Updated struct {
	// 編號
	GdID string `json:"gd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 類別(A1、B2、C2)
	Classification string `json:"classification,omitempty"`
	// 標題(單位主管審核)
	Title string `json:"title,omitempty"`
	// 資料需求(預存程序名稱)
	DataDemand string `json:"data_demand,omitempty"`
}

func (a *Table) TableName() string {
	return "gateway_data"
}

