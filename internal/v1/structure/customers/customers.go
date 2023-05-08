package customers

import (
	"time"

	model "eirc.app/internal/v1/structure"
	account "eirc.app/internal/v1/structure/accounts"
)

// Table struct is database table struct
type Table struct {
	//客戶編號
	CID string `gorm:"<-:create;primaryKey;column:c_id;type:UUID;default:uuid_generate_v4();" json:"c_id,omitempty"`
	//簡稱
	ShortName string `gorm:"column:short_name;type:TEXT;not null;" json:"short_name,omitempty"`
	//英文名稱
	EngName string `gorm:"column:eng_name;type:TEXT;" json:"eng_name,omitempty"`
	//中文名稱
	Name string `gorm:"column:name;type:TEXT;not null;" json:"name,omitempty"`
	//郵遞區號
	ZipCode string `gorm:"column:zip_code;type:TEXT;not null;" json:"zip_code,omitempty"`
	//地址
	Address string `gorm:"column:address;type:TEXT;not null;" json:"address,omitempty"`
	//電話
	Tel string `gorm:"column:tel;type:TEXT;not null;" json:"tel,omitempty"`
	//傳真
	Fax string `gorm:"column:fax;type:TEXT;" json:"fax,omitempty"`
	//地圖
	Map string `gorm:"column:map;type:TEXT;" json:"map,omitempty"`
	//聯絡人
	Liaison string `gorm:"column:liaison;type:TEXT;" json:"liaison,omitempty"`
	//電子郵件
	Mail string `gorm:"column:mail;type:TEXT;" json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `gorm:"column:liaison_phone;type:TEXT;" json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `gorm:"column:tax_id_number;type:TEXT;not null;" json:"tax_id_number,omitempty"`
	//備註
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//創建者
	Creater string `gorm:"<-:create;column:creater;type:UUID;not null;" json:"creater,omitempty"`

	Account account.Table `gorm:"foreignkey:creater;references:account_id"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	//客戶編號
	CID string `json:"c_id,omitempty"`
	//簡稱
	ShortName string `json:"short_name,omitempty"`
	//英文名稱
	EngName string `json:"eng_name,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty"`
	//地址
	Address string `json:"address,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//地圖
	Map string `json:"map,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//客戶編號
	CID string `json:"c_id,omitempty"`
	//簡稱
	ShortName string `json:"short_name,omitempty"`
	//英文名稱
	EngName string `json:"eng_name,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty"`
	//地址
	Address string `json:"address,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//地圖
	Map string `json:"map,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//簡稱
	ShortName string `json:"short_name,omitempty" binding:"required" validate:"required"`
	//英文名稱
	EngName string `json:"eng_name,omitempty" validate:"required"`
	//中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty" binding:"required" validate:"required"`
	//地址
	Address string `json:"address,omitempty" binding:"required" validate:"required"`
	//電話
	Tel string `json:"tel,omitempty" binding:"required" validate:"required"`
	//傳真
	Fax string `json:"fax,omitempty" validate:"required"`
	//地圖
	Map string `json:"map,omitempty" validate:"required"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" validate:"required"`
	//電子郵件
	Mail string `json:"mail,omitempty" validate:"required"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty" validate:"required"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" validate:"required"`
	//創建者
	Creater string `json:"creater,omitempty" binding:"required,uuid4" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//客戶編號
	CID string `json:"c_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//簡稱
	ShortName string `json:"short_name,omitempty" from:"short_name"`
	//英文名稱
	EngName string `json:"eng_name,omitempty" from:"eng_name"`
	//中文名稱
	Name string `json:"name,omitempty" from:"name"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty" from:"zip_code"`
	//地址
	Address string `json:"address,omitempty" from:"address"`
	//電話
	Tel string `json:"tel,omitempty" from:"tel"`
	//傳真
	Fax string `json:"fax,omitempty" from:"fax"`
	//地圖
	Map string `json:"map,omitempty" from:"map"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" from:"liaison"`
	//電子郵件
	Mail string `json:"mail,omitempty" from:"mail"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty" from:"liaison_phone"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty" from:"tax_id_number"`
	//備註
	Remark string `json:"remark,omitempty" from:"remark"`
}

type Fields struct {
	Field
	model.InPage
}

type Customer_Account struct {
	//客戶編號
	CID string `json:"c_id,omitempty"`
	//簡稱
	ShortName string `json:"short_name,omitempty"`
	//英文名稱
	EngName string `json:"eng_name,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty"`
	//地址
	Address string `json:"address,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//地圖
	Map string `json:"map,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//創建者
	Account__name string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type List struct {
	Customers []*Customer_Account `json:"customers"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//客戶編號
	CID string `json:"c_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//簡稱
	ShortName string `json:"short_name,omitempty"`
	//英文名稱
	EngName string `json:"eng_name,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty"`
	//地址
	Address string `json:"address,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//地圖
	Map string `json:"map,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
}

func (a *Table) TableName() string {
	return "customer"
}
