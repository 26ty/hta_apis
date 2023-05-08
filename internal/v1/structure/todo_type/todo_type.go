package todo_type

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is companies database table struct
type Table struct {
	//代號
	TtID string `gorm:"primary_key;column:tt_id;uuid_generate_v4()type:UUID;" json:"tt_id,omitempty"`
	//分類名稱
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	//使用者id
	UserID string `gorm:"column:user_id;type:UUID;" json:"user_id,omitempty"`
	//創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
}


type Base struct {
	//代號
	TtID string `json:"tt_id,omitempty"`
	//分類名稱
	Name string `json:"name,omitempty"`
	//使用者id
	UserID string `json:"user_id,omitempty"`
	//創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	//代號
	TtID string `json:"tt_id,omitempty"`
	//分類名稱
	Name string `json:"name,omitempty"`
	//使用者id
	UserID string `json:"user_id,omitempty"`
	//創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	//分類名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	//使用者id
	UserID string `json:"user_id,omitempty" binding:"required" validate:"required" `
}

type Field struct {
	//代號
	TtID string `json:"tt_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//分類名稱
	Name string `json:"name,omitempty" form:"name"`
	//使用者id
	UserID string `json:"user_id,omitempty"  form:"user_id"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Todo_type []*struct {
	//代號
	TtID string `json:"tt_id,omitempty"`
	//分類名稱
	Name string `json:"name,omitempty"`
	//使用者id
	UserID string `json:"user_id,omitempty"`
	//創建時間
	CreateTime time.Time `json:"create_time"`
	} `json:"todo_type"`
	model.OutPage
}

type Updated struct {
	//代號
	TtID string `json:"tt_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//分類名稱
	Name string `json:"name,omitempty"`
	//使用者id
	UserID string `json:"user_id,omitempty"`
}

func (a *Table) TableName() string {
	return "todo_type"
}

// func (a *Create_Table) TableName() string {
// 	return "todo_type"
// }

