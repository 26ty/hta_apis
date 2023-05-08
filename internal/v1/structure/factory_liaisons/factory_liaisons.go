package factory_liaisons

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//廠別聯絡人編號
	FlID string `gorm:"<-:create;primaryKey;column:fl_id;type:UUID;default:uuid_generate_v4();" json:"fl_id,omitempty"`
	//廠別編號
	FactoryID string `gorm:"column:factory_id;type:UUID;not null;" json:"factory_id,omitempty"`
	//總公司
	Co string `gorm:"column:co;type:TEXT;not null;" json:"co,omitempty"`
	//電話
	Tel string `gorm:"column:tel;type:TEXT;not null;" json:"tel,omitempty"`
	//聯絡人
	Liaison string `gorm:"column:liaison;type:TEXT;not null;" json:"liaison,omitempty"`
	//電子郵件
	Mail string `gorm:"column:mail;type:TEXT;not null;" json:"mail,omitempty"`
	//職稱
	Jobtitle string `gorm:"column:jobtitle;type:TEXT;not null;" json:"jobtitle,omitempty"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
}

type Base struct {
	//廠別聯絡人編號
	FlID string `json:"fl_id,omitempty"`
	//廠別編號
	FactoryID string `json:"factory_id,omitempty"`
	//總公司
	Co string `json:"co,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//職稱
	Jobtitle string `json:"jobtitle,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

type Single struct {
	//廠別聯絡人編號
	FlID string `json:"fl_id,omitempty"`
	//廠別編號
	FactoryID string `json:"factory_id,omitempty"`
	//總公司
	Co string `json:"co,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//職稱
	Jobtitle string `json:"jobtitle,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
}

// 放create時需輸入的欄位
type Created struct {
	//廠別編號
	FactoryID string `json:"factory_id,omitempty" binding:"required,uuid4" validate:"required"`
	//總公司
	Co string `json:"co,omitempty" binding:"required" validate:"required"`
	//電話
	Tel string `json:"tel,omitempty" binding:"required" validate:"required"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" binding:"required" validate:"required"`
	//職稱
	Jobtitle string `json:"jobtitle,omitempty" binding:"required" validate:"required"`
	//電子郵件
	Mail string `json:"mail,omitempty" binding:"required" validate:"required"`
}

type Created_List struct {
	Liaison []*Created `json:"liaison"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//廠別聯絡人編號
	FlID string `json:"fl_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//總公司
	Co string `json:"co,omitempty" from:"co"`
	//電話
	Tel string `json:"tel,omitempty" from:"tel"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" from:"liaison"`
	//電子郵件
	Mail string `json:"mail,omitempty" from:"mail"`
	//職稱
	Jobtitle string `json:"jobtitle,omitempty" from:"jobtitle"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	FactoryLiaison []*struct {
		//廠別聯絡人編號
		FlID string `json:"fl_id,omitempty"`
		//廠別編號
		FactoryID string `json:"factory_id,omitempty"`
		//總公司
		Co string `json:"co,omitempty"`
		//電話
		Tel string `json:"tel,omitempty"`
		//聯絡人
		Liaison string `json:"liaison,omitempty"`
		//電子郵件
		Mail string `json:"mail,omitempty"`
		//職稱
		Jobtitle string `json:"jobtitle,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"factory_liaisons"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//廠別聯絡人編號
	FlID string `json:"fl_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//總公司
	Co string `json:"co,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//職稱
	Jobtitle string `json:"jobtitle,omitempty"`
}

func (a *Table) TableName() string {
	return "factory_liaison"
}
