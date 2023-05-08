package quotations

import (
	"time"

	model "eirc.app/internal/v1/structure"
	account "eirc.app/internal/v1/structure/accounts"
	quotation_detail "eirc.app/internal/v1/structure/quotation_details"
)

type Table struct {
	//報價單編號
	QID string `gorm:"<-:create;primaryKey;column:q_id;type:UUID;default:uuid_generate_v4();" json:"q_id,omitempty"`

	Detail []quotation_detail.Table `gorm:"foreignkey:quotation_id;references:q_id" json:"Detail,omitempty"`
	//單號
	QCode string `gorm:"->;column:q_code;type:TEXT;not null;default:add_quotation_code()" json:"q_code,omitempty"`
	//幣別
	CurrencyType string `gorm:"column:currency_type;type:TEXT;not null;" json:"currency_type,omitempty"`
	//收件人
	To string `gorm:"column:to;type:TEXT;not null;" json:"to,omitempty"`
	//收件公司
	Attn string `gorm:"column:attn;type:TEXT;not null;" json:"attn,omitempty"`
	//寄件人
	From string `gorm:"column:from;type:TEXT;not null;" json:"from,omitempty"`
	//電話1
	Tel1 string `gorm:"column:tel1;type:TEXT;" json:"tel1,omitempty"`
	//電話2
	Tel2 string `gorm:"column:tel2;type:TEXT;" json:"tel2,omitempty"`
	//傳真
	Fax string `gorm:"column:fax;type:TEXT;" json:"fax,omitempty"`
	//報價日期
	Date *time.Time `gorm:"column:date;type:DATE;not null;" json:"date,omitempty"`
	//備註
	Remark string `gorm:"column:remark;type:TEXT;" json:"remark,omitempty"`
	//申請狀態
	Status string `gorm:"column:status;type:TEXT;not null;" json:"status,omitempty"`
	//創建者
	Creater string `gorm:"<-:create;column:creater;type:UUID;not null;" json:"creater,omitempty"`

	Account account.Table `gorm:"foreignkey:creater;references:account_id"`
	//創建時間
	CreatedTime time.Time `gorm:"<-:create;column:created_time;type:TIMESTAMPTZ;not null;" json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `gorm:"column:bonita_case_id;type:double precision;" json:"bonita_case_id,omitempty"`
}

type Base struct {
	//報價單編號
	QID string `json:"q_id,omitempty"`
	//單號
	QCode string `json:"q_code,omitempty"`
	//收件人
	To string `json:"to,omitempty"`
	//幣別
	CurrencyType string `json:"currency_type,omitempty"`
	//收件公司
	Attn string `json:"attn,omitempty"`
	//寄件人
	From string `json:"from,omitempty"`
	//電話1
	Tel1 string `json:"tel1,omitempty"`
	//電話2
	Tel2 string `json:"tel2,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//報價日期
	Date *time.Time `json:"date,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Single struct {
	//報價單編號
	QID string `json:"q_id,omitempty"`
	//單號
	QCode string `json:"q_code,omitempty"`
	//收件人
	To string `json:"to,omitempty"`
	//幣別
	CurrencyType string `json:"currency_type,omitempty"`
	//收件公司
	Attn string `json:"attn,omitempty"`
	//寄件人
	From string `json:"from,omitempty"`
	//電話1
	Tel1 string `json:"tel1,omitempty"`
	//電話2
	Tel2 string `json:"tel2,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//報價日期
	Date *time.Time `json:"date,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

// 放create時需輸入的欄位
type Created struct {
	//收件人
	To string `json:"to,omitempty" binding:"required" validate:"required"`
	//幣別
	CurrencyType string `json:"currency_type,omitempty" binding:"required" validate:"required"`
	//收件公司
	Attn string `json:"attn,omitempty" binding:"required" validate:"required"`
	//寄件人
	From string `json:"from,omitempty" binding:"required" validate:"required"`
	//電話1
	Tel1 string `json:"tel1,omitempty" validate:"required"`
	//電話2
	Tel2 string `json:"tel2,omitempty" validate:"required"`
	//傳真
	Fax string `json:"fax,omitempty" validate:"required"`
	//報價日期
	Date *time.Time `json:"date,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" validate:"required"`
	//申請狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	//創建者
	Creater string `json:"creater,omitempty" binding:"required,uuid4" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//報價單編號
	QID string `json:"q_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//收件人
	To string `json:"to,omitempty" from:"to"`
	//幣別
	CurrencyType string `json:"currency_type,omitempty" from:"currency_type"`
	//收件公司
	Attn string `json:"attn,omitempty" from:"attn"`
	//寄件人
	From string `json:"from,omitempty" from:"from"`
	//電話1
	Tel1 string `json:"tel1,omitempty" from:"tel1"`
	//電話2
	Tel2 string `json:"tel2,omitempty" from:"tel2"`
	//傳真
	Fax string `json:"fax,omitempty" from:"fax"`
	//報價日期
	Date *time.Time `json:"date,omitempty" from:"date"`
	//備註
	Remark string `json:"remark,omitempty" from:"remark"`
	//申請狀態
	Status string `json:"status,omitempty" from:"status"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" from:"bonita_case_id"`
}

type Fields struct {
	Field
	model.InPage
}

type Quotation_Account struct {
	//報價單編號
	QID string `json:"q_id,omitempty"`
	//單號
	QCode string `json:"q_code,omitempty"`
	//收件人
	To string `json:"to,omitempty"`
	//幣別
	CurrencyType string `json:"currency_type,omitempty"`
	//收件公司
	Attn string `json:"attn,omitempty"`
	//寄件人
	From string `json:"from,omitempty"`
	//電話1
	Tel1 string `json:"tel1,omitempty"`
	//電話2
	Tel2 string `json:"tel2,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//報價日期
	Date *time.Time `json:"date,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建者
	Account__name string `json:"creater_name,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type QuotationDetail struct {
	Quotation_Account
	Detail []quotation_detail.Table `gorm:"foreignkey:quotation_id;references:q_id" json:"Detail,omitempty"`
}

type AllQuotationDetail struct {
	Quotation []*QuotationDetail `json:"quotations"`
	model.OutPage
}

type List struct {
	Quotations []*struct {
		//報價單編號
		QID string `json:"q_id,omitempty"`
		//單號
		QCode string `json:"q_code,omitempty"`
		//收件人
		To string `json:"to,omitempty"`
		//幣別
		CurrencyType string `json:"currency_type,omitempty"`
		//收件公司
		Attn string `json:"attn,omitempty"`
		//寄件人
		From string `json:"from,omitempty"`
		//電話1
		Tel1 string `json:"tel1,omitempty"`
		//電話2
		Tel2 string `json:"tel2,omitempty"`
		//傳真
		Fax string `json:"fax,omitempty"`
		//報價日期
		Date *time.Time `json:"date,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//申請狀態
		Status string `json:"status,omitempty"`
		//創建者
		Creater string `json:"creater,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
		//bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	} `json:"quotations"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//報價單編號
	QID string `json:"q_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//收件人
	To string `json:"to,omitempty"`
	//幣別
	CurrencyType string `json:"currency_type,omitempty"`
	//收件公司
	Attn string `json:"attn,omitempty"`
	//寄件人
	From string `json:"from,omitempty"`
	//電話1
	Tel1 string `json:"tel1,omitempty"`
	//電話2
	Tel2 string `json:"tel2,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//報價日期
	Date *time.Time `json:"date,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Updated_Bonita struct {
	//報價單編號
	QID string `json:"q_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//CaseID
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

func (a *Table) TableName() string {
	return "quotation"
}
