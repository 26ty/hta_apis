package quotation_details

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//報價單明細編號
	QdID string `gorm:"<-:create;primaryKey;column:qd_id;type:UUID;default:uuid_generate_v4();" json:"qd_id,omitempty"`
	//報價單編號
	QuotationID string `gorm:"<-:create;column:quotation_id;type:UUID;not null;" json:"quotation_id,omitempty"`
	//項次
	No string `gorm:"column:no;type:TEXT;not null;" json:"no,omitempty"`
	//料號/品號規格
	MaterialType string `gorm:"column:material_type;type:TEXT;not null;" json:"material_type,omitempty"`
	//數量
	Quantity int `gorm:"column:quantity;type:INT4;not null;" json:"quantity,omitempty"`
	//單位
	Unit string `gorm:"column:unit;type:TEXT;not null;" json:"unit,omitempty"`
	//單價
	Price int `gorm:"column:price;type:INT4;not null;" json:"price,omitempty"`
	//合計
	TotalPrice int `gorm:"column:total_price;type:INT4;not null;" json:"total_price,omitempty"`
	//交貨日期
	DateForDelivery *time.Time `gorm:"column:date_for_delivery;type:DATE;not null;" json:"date_for_delivery,omitempty"`
	//備註
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//報價單明細編號
	QdID string `json:"qd_id,omitempty"`
	//報價單編號
	QuotationID string `json:"quotation_id,omitempty"`
	//項次
	No string `json:"no,omitempty"`
	//料號/品號規格
	MaterialType string `json:"material_type,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//單價
	Price int `json:"price,omitempty"`
	//合計
	TotalPrice int `json:"total_price,omitempty"`
	//交貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//報價單明細編號
	QdID string `json:"qd_id,omitempty"`
	//報價單編號
	QuotationID string `json:"quotation_id,omitempty"`
	//項次
	No string `json:"no,omitempty"`
	//料號/品號規格
	MaterialType string `json:"material_type,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//單價
	Price int `json:"price,omitempty"`
	//合計
	TotalPrice int `json:"total_price,omitempty"`
	//交貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//報價單編號
	QuotationID string `json:"quotation_id,omitempty" binding:"required,uuid4" validate:"required"`
	//項次
	No string `json:"no,omitempty" binding:"required" validate:"required"`
	//料號/品號規格
	MaterialType string `json:"material_type,omitempty" binding:"required" validate:"required"`
	//數量
	Quantity int `json:"quantity,omitempty" binding:"required" validate:"required"`
	//單位
	Unit string `json:"unit,omitempty" binding:"required" validate:"required"`
	//單價
	Price int `json:"price,omitempty" binding:"required" validate:"required"`
	//合計
	TotalPrice int `json:"total_price,omitempty" binding:"required" validate:"required"`
	//交貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//報價單明細編號
	QdID string `json:"qd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//項次
	No string `json:"no,omitempty" from:"no"`
	//料號/品號規格
	MaterialType string `json:"material_type,omitempty" from:"material_type"`
	//數量
	Quantity int `json:"quantity,omitempty" from:"quantity"`
	//單位
	Unit string `json:"unit,omitempty" from:"unit"`
	//單價
	Price int `json:"price,omitempty" from:"price"`
	//合計
	TotalPrice int `json:"total_price,omitempty" from:"total_price"`
	//交貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty" from:"date_for_delivery"`
	//備註
	Remark string `json:"remark,omitempty" from:"remark"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	QuotationDetails []*struct {
		//報價單明細編號
		QdID string `json:"qd_id,omitempty"`
		//報價單編號
		QuotationID string `json:"quotation_id,omitempty"`
		//項次
		No string `json:"no,omitempty"`
		//料號/品號規格
		MaterialType string `json:"material_type,omitempty"`
		//數量
		Quantity int `json:"quantity,omitempty"`
		//單位
		Unit string `json:"unit,omitempty"`
		//單價
		Price int `json:"price,omitempty"`
		//合計
		TotalPrice int `json:"total_price,omitempty"`
		//交貨日期
		DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"quotation_details"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//報價單明細編號
	QdID string `json:"qd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//項次
	No string `json:"no,omitempty"`
	//料號/品號規格
	MaterialType string `json:"material_type,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//單價
	Price int `json:"price,omitempty"`
	//合計
	TotalPrice int `json:"total_price,omitempty"`
	//交貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
}

func (a *Table) TableName() string {
	return "quotation_detail"
}
