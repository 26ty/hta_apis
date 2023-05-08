package jig_demand_details

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//治具需求單明細編號
	JdID string `gorm:"<-:create;primaryKey;column:jd_id;type:UUID;default:uuid_generate_v4();" json:"jd_id,omitempty"`
	//治具需求單編號
	JigID string `gorm:"column:jig_id;type:UUID;not null;" json:"jig_id,omitempty"`
	//產品需求
	ProductDemand string `gorm:"column:product_demand;type:TEXT;not null;" json:"product_demand,omitempty"`
	//總成品號
	ProductNumber string `gorm:"column:product_number;type:TEXT;" json:"product_number,omitempty"`
	//是否急件
	Urgent *bool `gorm:"column:urgent;type:bool;default:false;not null;" json:"urgent,omitempty"`
	//數量
	Quantity int `gorm:"column:quantity;type:int4;not null;" json:"quantity,omitempty"`
	//單位
	Unit string `gorm:"column:unit;type:TEXT;not null;" json:"unit,omitempty"`
	//備註
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//設計預交日
	EstimatedDateForDesign *time.Time `gorm:"column:estimated_date_for_design;type:DATE;" json:"estimated_date_for_design,omitempty"`
	//設計取得客圖/資訊
	AcquiredDateForInformation *time.Time `gorm:"column:acquired_date_for_information;type:DATE;" json:"acquired_date_for_information,omitempty"`
	//設計完工日
	DateForDesign *time.Time `gorm:"column:date_for_design;type:DATE;" json:"date_for_design,omitempty"`
	//完工數量
	CompletedQuantity int `gorm:"column:completed_quantity;type:int4;" json:"completed_quantity,omitempty"`
	//告知日期
	DateForInform *time.Time `gorm:"column:date_for_inform;type:DATE;" json:"date_for_inform,omitempty"`
	//回復備註
	Response string `gorm:"column:response;type:TEXT;" json:"response,omitempty"`
	//內部訂單日
	DateForInnerOrder *time.Time `gorm:"column:date_for_inner_order;type:DATE;" json:"date_for_inner_order,omitempty"`
	//請購完成日
	DateForPurchase *time.Time `gorm:"column:date_for_purchase;type:DATE;" json:"date_for_purchase,omitempty"`
	//發包完成日
	DateForPackage *time.Time `gorm:"column:date_for_package;type:DATE;" json:"date_for_package,omitempty"`
	//預計入料日
	EstimatedDateForIncoming *time.Time `gorm:"column:estimated_date_for_incoming;type:DATE;" json:"estimated_date_for_incoming,omitempty"`
	//入料完成日
	DateForIncoming *time.Time `gorm:"column:date_for_incoming;type:DATE;" json:"date_for_incoming,omitempty"`
	//入庫交期日
	EstimatedDateForInventory *time.Time `gorm:"column:estimated_date_for_inventory;type:DATE;" json:"estimated_date_for_inventory,omitempty"`
	//入庫完工日
	DateForInventory *time.Time `gorm:"column:date_for_inventory;type:DATE;" json:"date_for_inventory,omitempty"`
	//出貨日期
	DateForDelivery *time.Time `gorm:"column:date_for_delivery;type:DATE;" json:"date_for_delivery,omitempty"`
	//治具訂單數量
	OrderQuantity int `gorm:"column:order_quantity;type:INT4;" json:"order_quantity,omitempty"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//治具需求單明細編號
	JdID string `json:"jd_id,omitempty"`
	//治具需求單編號
	JigID string `json:"jig_id,omitempty"`
	//產品需求
	ProductDemand string `json:"product_demand,omitempty"`
	//總成品號
	ProductNumber string `json:"product_number,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//設計預交日
	EstimatedDateForDesign *time.Time `json:"estimated_date_for_design,omitempty"`
	//設計取得客圖/資訊
	AcquiredDateForInformation *time.Time `json:"acquired_date_for_information,omitempty"`
	//設計完工日
	DateForDesign *time.Time `json:"date_for_design,omitempty"`
	//完工數量
	CompletedQuantity int `json:"completed_quantity,omitempty"`
	//告知日期
	DateForInform *time.Time `json:"date_for_inform,omitempty"`
	//回復備註
	Response string `json:"response,omitempty"`
	//內部訂單日
	DateForInnerOrder *time.Time `json:"date_for_inner_order,omitempty"`
	//請購完成
	DateForPurchase *time.Time `json:"date_for_purchase,omitempty"`
	//發包完成
	DateForPackage *time.Time `json:"date_for_package,omitempty"`
	//預計入料日
	EstimatedDateForIncoming *time.Time `json:"estimated_date_for_incoming,omitempty"`
	//入料完成
	DateForIncoming *time.Time `json:"date_for_incoming,omitempty"`
	//入庫交期
	EstimatedDateForInventory *time.Time `json:"estimated_date_for_inventory,omitempty"`
	//入庫完工
	DateForInventory *time.Time `json:"date_for_inventory,omitempty"`
	//出貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	//治具訂單數量
	OrderQuantity int `json:"order_quantity,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//治具需求單明細編號
	JdID string `json:"jd_id,omitempty"`
	//治具需求單編號
	JigID string `json:"jig_id,omitempty"`
	//產品需求
	ProductDemand string `json:"product_demand,omitempty"`
	//總成品號
	ProductNumber string `json:"product_number,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//設計預交日
	EstimatedDateForDesign *time.Time `json:"estimated_date_for_design,omitempty"`
	//設計取得客圖/資訊
	AcquiredDateForInformation *time.Time `json:"acquired_date_for_information,omitempty"`
	//設計完工日
	DateForDesign *time.Time `json:"date_for_design,omitempty"`
	//完工數量
	CompletedQuantity int `json:"completed_quantity,omitempty"`
	//告知日期
	DateForInform *time.Time `json:"date_for_inform,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//回復備註
	Response string `json:"response,omitempty"`
	//內部訂單日
	DateForInnerOrder *time.Time `json:"date_for_inner_order,omitempty"`
	//請購完成
	DateForPurchase *time.Time `json:"date_for_purchase,omitempty"`
	//發包完成
	DateForPackage *time.Time `json:"date_for_package,omitempty"`
	//預計入料日
	EstimatedDateForIncoming *time.Time `json:"estimated_date_for_incoming,omitempty"`
	//入料完成
	DateForIncoming *time.Time `json:"date_for_incoming,omitempty"`
	//入庫交期
	EstimatedDateForInventory *time.Time `json:"estimated_date_for_inventory,omitempty"`
	//入庫完工
	DateForInventory *time.Time `json:"date_for_inventory,omitempty"`
	//出貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	//治具訂單數量
	OrderQuantity int `json:"order_quantity,omitempty"`
}

// 放create時需輸入的欄位
type Created struct {
	//治具需求單編號
	JigID string `json:"jig_id,omitempty" binding:"required,uuid4" validate:"required"`
	//產品需求
	ProductDemand string `json:"product_demand,omitempty" binding:"required" validate:"required"`
	//總成品號
	ProductNumber string `json:"product_number,omitempty" validate:"required"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty" binding:"required" validate:"required"`
	//數量
	Quantity int `json:"quantity,omitempty" binding:"required" validate:"required"`
	//單位
	Unit string `json:"unit,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" validate:"required"`
	//回復備註
	Response string `json:"response,omitempty" validate:"required"`
	//設計預交日
	EstimatedDateForDesign *time.Time `json:"estimated_date_for_design,omitempty" validate:"required"`
	//設計取得客圖/資訊
	AcquiredDateForInformation *time.Time `json:"acquired_date_for_information,omitempty" validate:"required"`
	//設計完工日
	DateForDesign *time.Time `json:"date_for_design,omitempty" validate:"required"`
	//完工數量
	CompletedQuantity int `json:"completed_quantity,omitempty" validate:"required"`
	//告知日期
	DateForInform *time.Time `json:"date_for_inform,omitempty" validate:"required"`
	//內部訂單日
	DateForInnerOrder *time.Time `json:"date_for_inner_order,omitempty" validate:"required"`
	//請購完成
	DateForPurchase *time.Time `json:"date_for_purchase,omitempty" validate:"required"`
	//發包完成
	DateForPackage *time.Time `json:"date_for_package,omitempty" validate:"required"`
	//預計入料日
	EstimatedDateForIncoming *time.Time `json:"estimated_date_for_incoming,omitempty" validate:"required"`
	//入料完成
	DateForIncoming *time.Time `json:"date_for_incoming,omitempty" validate:"required"`
	//入庫交期
	EstimatedDateForInventory *time.Time `json:"estimated_date_for_inventory,omitempty" validate:"required"`
	//入庫完工
	DateForInventory *time.Time `json:"date_for_inventory,omitempty" validate:"required"`
	//出貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty" validate:"required"`
	//治具訂單數量
	OrderQuantity int `json:"order_quantity,omitempty" validate:"required"`
}

type Created_List struct {
	Detail []*Created `json:"detail"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//治具需求單明細編號
	JdID string `json:"jd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//治具需求單編號
	JigID string `json:"jig_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//產品需求
	ProductDemand string `json:"product_demand,omitempty" from:"product_demand"`
	//總成品號
	ProductNumber string `json:"product_number,omitempty" from:"product_number"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty" from:"urgent"`
	//數量
	Quantity int `json:"quantity,omitempty" from:"quantity"`
	//單位
	Unit string `json:"unit,omitempty" from:"unit"`
	//備註
	Remark string `json:"remark,omitempty" from:"remark"`
	//設計預交日
	EstimatedDateForDesign *time.Time `json:"estimated_date_for_design,omitempty" from:"estimated_date_for_design,omitempty"`
	//設計取得客圖/資訊
	AcquiredDateForInformation *time.Time `json:"acquired_date_for_information,omitempty" from:"acquired_date_for_information"`
	//設計完工日
	DateForDesign *time.Time `json:"date_for_design,omitempty" from:"date_for_design"`
	//完工數量
	CompletedQuantity int `json:"completed_quantity,omitempty" from:"completed_quantity"`
	//告知日期
	DateForInform *time.Time `json:"date_for_inform,omitempty" from:"date_for_inform"`
	//回復備註
	Response string `json:"response,omitempty" from:"response"`
	//內部訂單日
	DateForInnerOrder *time.Time `json:"date_for_inner_order,omitempty" from:"date_for_inner_order"`
	//請購完成
	DateForPurchase *time.Time `json:"date_for_purchase,omitempty" from:"date_for_purchase"`
	//發包完成
	DateForPackage *time.Time `json:"date_for_package,omitempty" from:"date_for_package"`
	//預計入料日
	EstimatedDateForIncoming *time.Time `json:"estimated_date_for_incoming,omitempty" from:"estimated_date_for_incoming"`
	//入料完成
	DateForIncoming *time.Time `json:"date_for_incoming,omitempty" from:"date_for_incoming"`
	//入庫交期
	EstimatedDateForInventory *time.Time `json:"estimated_date_for_inventory,omitempty" from:"estimated_date_for_inventory"`
	//入庫完工
	DateForInventory *time.Time `json:"date_for_inventory,omitempty" from:"date_for_inventory"`
	//出貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty" from:"date_for_delivery"`
	//治具訂單數量
	OrderQuantity int `json:"order_quantity,omitempty" from:"order_quantity"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	JigDemandDetails []*struct {
		//治具需求單明細編號
		JdID string `json:"jd_id,omitempty"`
		//治具需求單編號
		JigID string `json:"jig_id,omitempty"`
		//產品需求
		ProductDemand string `json:"product_demand,omitempty"`
		//總成品號
		ProductNumber string `json:"product_number,omitempty"`
		//是否急件
		Urgent *bool `json:"urgent,omitempty"`
		//數量
		Quantity int `json:"quantity,omitempty"`
		//單位
		Unit string `json:"unit,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//設計預交日
		EstimatedDateForDesign *time.Time `json:"estimated_date_for_design,omitempty"`
		//設計取得客圖/資訊
		AcquiredDateForInformation *time.Time `json:"acquired_date_for_information,omitempty"`
		//設計完工日
		DateForDesign *time.Time `json:"date_for_design,omitempty"`
		//完工數量
		CompletedQuantity int `json:"completed_quantity,omitempty"`
		//告知日期
		DateForInform *time.Time `json:"date_for_inform,omitempty"`
		//內部訂單日
		DateForInnerOrder *time.Time `json:"date_for_inner_order,omitempty"`
		//請購完成
		DateForPurchase *time.Time `json:"date_for_purchase,omitempty"`
		//發包完成
		DateForPackage *time.Time `json:"date_for_package,omitempty"`
		//預計入料日
		EstimatedDateForIncoming *time.Time `json:"estimated_date_for_incoming,omitempty"`
		//入料完成
		DateForIncoming *time.Time `json:"date_for_incoming,omitempty"`
		//入庫交期
		EstimatedDateForInventory *time.Time `json:"estimated_date_for_inventory,omitempty"`
		//入庫完工
		DateForInventory *time.Time `json:"date_for_inventory,omitempty"`
		//回復備註
		Response string `json:"response,omitempty"`
		//出貨日期
		DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
		//治具訂單數量
		OrderQuantity int `json:"order_quantity,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"jig_demand_details"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//治具需求單明細編號
	JdID string `json:"jd_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//治具需求單編號
	JigID string `json:"jig_id,omitempty" binding:"omitempty,uuid4"`
	//產品需求
	ProductDemand string `json:"product_demand,omitempty"`
	//總成品號
	ProductNumber string `json:"product_number,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//設計預交日
	EstimatedDateForDesign *time.Time `json:"estimated_date_for_design,omitempty"`
	//設計取得客圖/資訊
	AcquiredDateForInformation *time.Time `json:"acquired_date_for_information,omitempty"`
	//設計完工日
	DateForDesign *time.Time `json:"date_for_design,omitempty"`
	//完工數量
	CompletedQuantity int `json:"completed_quantity,omitempty"`
	//告知日期
	DateForInform *time.Time `json:"date_for_inform,omitempty"`
	//回復備註
	Response string `json:"response,omitempty"`
	//內部訂單日
	DateForInnerOrder *time.Time `json:"date_for_inner_order,omitempty"`
	//請購完成
	DateForPurchase *time.Time `json:"date_for_purchase,omitempty"`
	//發包完成
	DateForPackage *time.Time `json:"date_for_package,omitempty"`
	//預計入料日
	EstimatedDateForIncoming *time.Time `json:"estimated_date_for_incoming,omitempty"`
	//入料完成
	DateForIncoming *time.Time `json:"date_for_incoming,omitempty"`
	//入庫交期
	EstimatedDateForInventory *time.Time `json:"estimated_date_for_inventory,omitempty"`
	//入庫完工
	DateForInventory *time.Time `json:"date_for_inventory,omitempty"`
	//出貨日期
	DateForDelivery *time.Time `json:"date_for_delivery,omitempty"`
	//治具訂單數量
	OrderQuantity int `json:"order_quantity,omitempty"`
}

func (a *Table) TableName() string {
	return "jig_demand_detail"
}
