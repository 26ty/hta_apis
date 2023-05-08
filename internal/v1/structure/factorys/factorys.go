package factorys

import (
	"time"

	model "eirc.app/internal/v1/structure"
	account "eirc.app/internal/v1/structure/accounts"

	customer "eirc.app/internal/v1/structure/customers"
	factory_liaison "eirc.app/internal/v1/structure/factory_liaisons"
	factory_manufacturing "eirc.app/internal/v1/structure/factory_manufacturings"
)

type Table struct {
	//客戶廠別編號
	FID string `gorm:"<-:create;primaryKey;column:f_id;type:UUID;default:uuid_generate_v4();" json:"f_id,omitempty"`

	Manufacturing []factory_manufacturing.Table `gorm:"foreignkey:factory_id;references:f_id" json:"manufacturing"`
	Liaison       []factory_liaison.Table       `gorm:"foreignkey:factory_id;references:f_id" json:"liaison"`
	//客戶
	CustomerID string `gorm:"column:customer_id;type:UUID;not null;" json:"customer_id,omitempty"`

	Customer customer.Table `gorm:"foreignkey:customer_id;references:c_id"`
	//地區
	Area string `gorm:"column:area;type:TEXT;not null;" json:"area,omitempty"`
	//地點
	Location string `gorm:"column:location;type:TEXT;not null;" json:"location,omitempty"`
	//廠別
	Factory string `gorm:"column:factory;type:TEXT;not null;" json:"factory,omitempty"`
	//備註
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//是否啟用
	Enable *bool `gorm:"column:enable;type:bool;default:false;not null;" json:"enable,omitempty"`
	//創建者
	Creater string `gorm:"<-:create;column:creater;type:UUID;not null;" json:"creater,omitempty"`

	Account account.Table `gorm:"foreignkey:account_id;references:creater"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//客戶廠別編號
	FID string `json:"f_id,omitempty"`

	Manufacturing []factory_manufacturing.Base `json:"manufacturing"`
	Liaison       []factory_liaison.Base       `json:"liaison"`
	//客戶
	CustomerID string `json:"customer_id,omitempty"`
	//地區
	Area string `json:"area,omitempty"`
	//地點
	Location string `json:"location,omitempty"`
	//廠別
	Factory string `json:"factory,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//是否啟用
	Enable *bool `json:"enable,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//客戶廠別編號
	FID string `json:"f_id,omitempty"`
	//客戶
	CustomerID string `json:"customer_id,omitempty"`
	//地區
	Area string `json:"area,omitempty"`
	//地點
	Location string `json:"location,omitempty"`
	//廠別
	Factory string `json:"factory,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//是否啟用
	Enable *bool `json:"enable,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//客戶
	CustomerID string `json:"customer_id,omitempty" binding:"required,uuid4" validate:"required"`
	//地區
	Area string `json:"area,omitempty" binding:"required" validate:"required"`
	//地點
	Location string `json:"location,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" validate:"required"`
	//是否啟用
	Enable *bool `json:"enable,omitempty" binding:"required" validate:"required"`
	//廠別
	Factory string `json:"factory,omitempty" binding:"required" validate:"required"`
	//創建者
	Creater string `json:"creater,omitempty" binding:"required,uuid4" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//客戶廠別編號
	FID string `json:"f_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//客戶
	CustomerID string `json:"customer_id,omitempty" from:"customer_id" binding:"omitempty,uuid4"`
	//地區
	Area string `json:"area,omitempty" from:"area"`
	//地點
	Location string `json:"location,omitempty" from:"location"`
	//廠別
	Factory string `json:"factory,omitempty" from:"factory"`
	//備註
	Remark string `json:"remark,omitempty" from:"remark"`
	//是否啟用
	Enable *bool `json:"enable,omitempty" from:"enable"`
}

type Fields struct {
	Field
	model.InPage
}

type Factory_Account struct {
	//客戶廠別編號
	FID string `json:"f_id,omitempty"`
	//客戶ID
	CustomerID string `json:"customer_id,omitempty"`
	//客戶
	Customer__short_name string `json:"customer,omitempty"`
	//地區
	Area string `json:"area,omitempty"`
	//地點
	Location string `json:"location,omitempty"`
	//廠別
	Factory string `json:"factory,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//是否啟用
	Enable *bool `json:"enable,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建者
	Account__name string `json:"creater_name,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type FLM struct {
	Factory_Account
	Manufacturing []factory_manufacturing.Base `json:"manufacturing,omitempty"`
	Liaison       []factory_liaison.Base       `json:"liaison,omitempty"`
}

type AllFLM struct {
	Factory []*FLM `json:"factorys"`
	model.OutPage
}

type SearchFactory struct {
	SearchFactory []*struct {
		Factory_Account
		Manufacturing []factory_manufacturing.Base `json:"manufacturing,omitempty"`
	} `json:"factorys"`
	model.OutPage
}

type List struct {
	Factorys []*struct {
		//客戶廠別編號
		FID string `json:"f_id,omitempty"`
		//客戶
		CustomerID string `json:"customer_id,omitempty"`
		//地區
		Area string `json:"area,omitempty"`
		//地點
		Location string `json:"location,omitempty"`
		//廠別
		Factory string `json:"factory,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//是否啟用
		Enable *bool `json:"enable,omitempty"`
		//創建者
		Creater string `json:"creater,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"factorys"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//客戶廠別編號
	FID string `json:"f_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//客戶
	CustomerID string `json:"customer_id,omitempty" binding:"omitempty,uuid4"`
	//地區
	Area string `json:"area,omitempty"`
	//地點
	Location string `json:"location,omitempty"`
	//廠別
	Factory string `json:"factory,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//是否啟用
	Enable *bool `json:"enable,omitempty"`
}

func (a *Table) TableName() string {
	return "factory"
}
