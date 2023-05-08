package gift_application_details

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//部品零件贈送明細編號
	GdID string `gorm:"<-:create;primaryKey;column:gd_id;type:UUID;default:uuid_generate_v4();" json:"gd_id,omitempty"`
	//部品零件贈送申請單編號
	GiftID string `gorm:"<-:create;column:gift_id;type:UUID;not null;" json:"gift_id,omitempty"`
	//品號
	ProductNumber string `gorm:"column:product_number;type:TEXT;not null;" json:"product_number,omitempty"`
	//品名
	ProductName string `gorm:"column:product_name;type:TEXT;not null;" json:"product_name,omitempty"`
	//規格
	Specifications string `gorm:"column:specifications;type:TEXT;not null;" json:"specifications,omitempty"`
	//數量
	Quantity int `gorm:"column:quantity;type:INT4;not null;" json:"quantity,omitempty"`
	//單位成本
	UnitCost int `gorm:"column:unit_cost;type:INT4;not null;" json:"unit_cost,omitempty"`
	//成本(NTD)
	TotalCost int `gorm:"column:total_cost;type:INT4;not null;" json:"total_cost,omitempty"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//部品零件贈送明細編號
	GdID string `json:"gd_id,omitempty"`
	//部品零件贈送申請單編號
	GiftID string `json:"gift_id,omitempty"`
	//品號
	ProductNumber string `json:"product_number,omitempty"`
	//品名
	ProductName string `json:"product_name,omitempty"`
	//規格
	Specifications string `json:"specifications,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位成本
	UnitCost int `json:"unit_cost,omitempty"`
	//成本(NTD)
	TotalCost int `json:"total_cost,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//部品零件贈送明細編號
	GdID string `json:"gd_id,omitempty"`
	//部品零件贈送申請單編號
	GiftID string `json:"gift_id,omitempty"`
	//品號
	ProductNumber string `json:"product_number,omitempty"`
	//品名
	ProductName string `json:"product_name,omitempty"`
	//規格
	Specifications string `json:"specifications,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位成本
	UnitCost int `json:"unit_cost,omitempty"`
	//成本(NTD)
	TotalCost int `json:"total_cost,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//部品零件贈送申請單編號
	GiftID string `json:"gift_id,omitempty" binding:"required,uuid4" validate:"required"`
	//品號
	ProductNumber string `json:"product_number,omitempty" binding:"required" validate:"required"`
	//品名
	ProductName string `json:"product_name,omitempty" binding:"required" validate:"required"`
	//規格
	Specifications string `json:"specifications,omitempty" binding:"required" validate:"required"`
	//數量
	Quantity int `json:"quantity,omitempty" binding:"required" validate:"required"`
	//單位成本
	UnitCost int `json:"unit_cost,omitempty" binding:"required" validate:"required"`
	//成本(NTD)
	TotalCost int `json:"total_cost,omitempty" binding:"required" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//部品零件贈送明細編號
	GdID string `json:"gd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//品號
	ProductNumber string `json:"product_number,omitempty" from:"product_number"`
	//品名
	ProductName string `json:"product_name,omitempty" from:"product_name"`
	//規格
	Specifications string `json:"specifications,omitempty" from:"specifications"`
	//數量
	Quantity int `json:"quantity,omitempty" from:"quantity"`
	//單位成本
	UnitCost int `json:"unit_cost,omitempty" from:"unit_cost"`
	//成本(NTD)
	TotalCost int `json:"total_cost,omitempty" from:"total_cost"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	GiftApplicationDetails []*struct {
		//部品零件贈送明細編號
		GdID string `json:"gd_id,omitempty"`
		//部品零件贈送申請單編號
		GiftID string `json:"gift_id,omitempty"`
		//品號
		ProductNumber string `json:"product_number,omitempty"`
		//品名
		ProductName string `json:"product_name,omitempty"`
		//規格
		Specifications string `json:"specifications,omitempty"`
		//數量
		Quantity int `json:"quantity,omitempty"`
		//單位成本
		UnitCost int `json:"unit_cost,omitempty"`
		//成本(NTD)
		TotalCost int `json:"total_cost,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"gift_application_details"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//部品零件贈送明細編號
	GdID string `json:"gd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//品號
	ProductNumber string `json:"product_number,omitempty"`
	//品名
	ProductName string `json:"product_name,omitempty"`
	//規格
	Specifications string `json:"specifications,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位成本
	UnitCost int `json:"unit_cost,omitempty"`
	//成本(NTD)
	TotalCost int `json:"total_cost,omitempty"`
}

func (a *Table) TableName() string {
	return "gift_application_detail"
}
