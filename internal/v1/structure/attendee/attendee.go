package attendee

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 編號
	AID string `gorm:"primaryKey;column:a_id;uuid_generate_v4()type:UUID;" json:"a_id,omitempty"`
	// 會議編號
	MeetID string `gorm:"column:meet_id;type:UUID;" json:"meet_id,omitempty"`
	// 參與者編號
	UserID string `gorm:"column:user_id;type:UUID;" json:"user_id,omitempty"`
	// 主席
	Chairman bool `gorm:"column:chairman;type:BOOLEAN;" json:"chairman,omitempty"`
	//收過信件
	ReceiveEmail bool `gorm:"column:receive_email;type:BOOLEAN;" json:"receive_email,omitempty"`
	// 創建時間Z
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
}

type Base struct {
	// 編號
	AID string `json:"a_id,omitempty"`
	// 會議編號
	MeetID string `json:"meet_id,omitempty"`
	// 參與者編號
	UserID string `json:"user_id,omitempty"`
	// 主席
	Chairman bool `json:"chairman,omitempty"`
	//收過信件
	ReceiveEmail bool `json:"receive_email,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Single struct {
	// 編號
	AID string `json:"a_id,omitempty"`
	// 會議編號
	MeetID string `json:"meet_id,omitempty"`
	// 參與者編號
	UserID string `json:"user_id,omitempty"`
	// 主席
	Chairman bool `json:"chairman,omitempty"`
	//收過信件
	ReceiveEmail bool `json:"receive_email,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
}

type Created struct {
	// 編號
	MeetID string `json:"meet_id" binding:"required,uuid4" validate:"required"`
	// 參與者編號
	UserID string `json:"user_id,omitempty" binding:"required" validate:"required"`
	// 主席
	Chairman bool `json:"chairman"`
	//收過信件
	ReceiveEmail bool `json:"receive_email"`
}

type Field struct {
	// 編號
	AID string `json:"a_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 會議編號
	MeetID *string `json:"meet_id,omitempty" form:"meet_id" binding:"omitempty,uuid4"`
	// 參與者編號
	UserID *string `json:"user_id,omitempty" form:"user_id" binding:"omitempty,uuid4"`
	// 主席
	Chairman *bool `json:"chairman,omitempty" form:"chairman"`
	//收過信件
	ReceiveEmail bool `json:"receive_email,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Attendee []*struct {
		// 編號
		AID string `json:"a_id,omitempty"`
		// 會議編號
		MeetID string `json:"meet_id,omitempty"`
		// 參與者編號
		UserID string `json:"user_id,omitempty"`
		// 主席
		Chairman bool `json:"chairman,omitempty"`
		//收過信件
		ReceiveEmail bool `json:"receive_email,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
	} `json:"attendee"`
	model.OutPage
}

type Updated_List struct {
	Attendee []*Updated `json:"attendee"`
}

type Updated struct {
	// 編號
	AID string `json:"a_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 會議編號
	MeetID *string `json:"meet_id,omitempty" binding:"omitempty,uuid4"`
	// 參與者編號
	UserID *string `json:"user_id,omitempty"`
	// 主席
	Chairman *bool `json:"chairman"`
	//收過信件
	ReceiveEmail bool `json:"receive_email"`
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
	return "attendee"
}

