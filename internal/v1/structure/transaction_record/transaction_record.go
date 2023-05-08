package transaction_record

import (
	"eirc.app/internal/v1/structure/accounts"
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 單據編號
	TrID string `gorm:"primaryKey;column:tr_id;uuid_generate_v4()type:UUID;" json:"tr_id,omitempty"`
	// 單據編號
	DocumentID string `gorm:"primaryKey;column:document_id;type:UUID;" json:"document_id,omitempty"`
	// 動作
	Actor string `gorm:"column:actor;type:TEXT;" json:"actor,omitempty"`
	// 內容
	Content string `gorm:"column:content;type:TEXT;" json:"content,omitempty"`
	// 流水號
	Rank int `gorm:"column:rank;type:INT4;" json:"rank,omitempty"`
	// 備註
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	// 建立者
	Creater  string         `gorm:"column:creater;uuid_generate_v4()type:UUID;" json:"creater,omitempty"`
	Accounts accounts.Table `gorm:"foreignkey:creater;references:account_id"`

	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 單據編號
	TrID string `json:"tr_id,omitempty"`
	// 單據編號
	DocumentID string `json:"document_id,omitempty"`
	// 動作
	Actor string `json:"actor,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 流水號
	Rank int `json:"rank,omitempty"`
	// 備註
	Remark string `json:"remark,omitempty"`
	// 建立者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 單據編號
	TrID string `json:"tr_id,omitempty"`
	// 單據編號
	DocumentID string `json:"document_id,omitempty"`
	// 動作
	Actor string `json:"actor,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 流水號
	Rank int `json:"rank,omitempty"`
	// 備註
	Remark string `json:"remark,omitempty"`
	// 建立者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 單據編號
	DocumentID string `json:"document_id,omitempty"`
	// 動作
	Actor string `json:"actor,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 備註
	Remark string `json:"remark,omitempty"`
	// 建立者
	Creater string `json:"creater,omitempty"`
}

type Field struct {
	// 單據編號
	TrID string `json:"tr_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 單據編號
	DocumentID string `json:"document_id,omitempty" form:"document_id" binding:"omitempty,uuid4"`
	// 動作
	Actor *string `json:"actor,omitempty" form:"actor"`
	// 內容
	Content *string `json:"content,omitempty" form:"content"`
	// 備註
	Remark string `json:"remark,omitempty" form:"remark"`
	// 流水號
	Rank *int `json:"rank,omitempty" form:"rank"`
	// 建立者
	Creater *string `json:"creater,omitempty" form:"creater"`
}

type Users struct {
	Field
	model.User
}

type Fields struct {
	Field
	model.InPage
}

type Record_user_list struct {
	// 單據編號
	TrID string `json:"tr_id,omitempty"`
	// 單據編號
	DocumentID string `json:"document_id,omitempty"`
	// 動作
	Actor string `json:"actor,omitempty"`
	// 內容
	Content string `json:"content,omitempty"`
	// 流水號
	Rank int `json:"rank,omitempty"`
	// 備註
	Remark string `json:"remark,omitempty"`
	// 建立者
	Creater string `json:"creater,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 建立者
	Accounts__name string `json:"creater_name,omitempty"`
}

type Record_user_lists struct {
	TransactionRecord []*struct {
		// 單據編號
		TrID string `json:"tr_id,omitempty"`
		// 單據編號
		DocumentID string `json:"document_id,omitempty"`
		// 動作
		Actor string `json:"actor,omitempty"`
		// 內容
		Content string `json:"content,omitempty"`
		// 流水號
		Rank int `json:"rank,omitempty"`
		// 備註
		Remark string `json:"remark,omitempty"`
		// 建立者
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 建立者
		CreaterName string `json:"creater_name,omitempty"`
	} `json:"transaction_record"`
	model.OutPage
}

type List struct {
	TransactionRecord []*struct {
		// 單據編號
		TrID string `json:"tr_id,omitempty"`
		// 單據編號
		DocumentID string `json:"document_id,omitempty"`
		// 動作
		Actor string `json:"actor,omitempty"`
		// 內容
		Content string `json:"content,omitempty"`
		// 流水號
		Rank int `json:"rank,omitempty"`
		// 備註
		Remark string `json:"remark,omitempty"`
		// 建立者
		Creater string `json:"creater,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"transaction_record"`
	model.OutPage
}

type Updated struct {
	// 單據編號
	TrID string `json:"tr_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 單據編號
	DocumentID *string `json:"document_id,omitempty" binding:"omitempty,uuid4"`
	// 動作
	Actor *string `json:"actor,omitempty"`
	// 內容
	Content *string `json:"content,omitempty"`
	// 備註
	Remark *string `json:"remark,omitempty"`
}

type Login struct {
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password" binding:"required" validate:"required"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "transaction_record"
}
