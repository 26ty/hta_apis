package meeting

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	//會議編號
	MID string `gorm:"primary_key;column:m_id;uuid_generate_v4()type:UUID;" json:"m_id,omitempty"`
	// 單據編號
	DocumentsID string `gorm:"column:documents_id;type:UUID;" json:"documents_id,omitempty"`
	// 會議名稱
	MName string `gorm:"column:m_name;type:TEXT;" json:"m_name,omitempty"`
	// 會議地點
	Room string `gorm:"column:room;type:TEXT;" json:"room,omitempty"`
	// 會議開始時間
	TimeForStart float32 `gorm:"column:time_for_start;type:double precision;" json:"time_for_start,omitempty"`
	// 會議結束時間
	TimeForEnd float32 `gorm:"column:time_for_end;type:double precision;" json:"time_for_end,omitempty"`
	// 會議日期
	DateForStart time.Time `gorm:"column:date_for_start;type:DATE;" json:"date_for_start,omitempty"`
	// 創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
	// 單據類型
	//OriginID string `gorm:"column:origin_id;type:UUID;" json:"origin_id,omitempty"`
}

type Base struct {
	// 會議編號
	MID string `json:"m_id,omitempty"`
	// 單據編號
	DocumentsID string `json:"documents_id,omitempty"`
	// 會議名稱
	MName string `json:"m_name,omitempty"`
	// 會議地點
	Room string `json:"room,omitempty"`
	// 會議開始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 會議結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 會議日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 單據類型
	OriginID string `json:"origin_id,omitempty"`
}

type MeetingListUser struct {
	// 會議編號
	MID string `json:"m_id,omitempty"`
	// 單據編號
	DocumentsID string `json:"documents_id,omitempty"`
	// 會議名稱
	MName string `json:"m_name,omitempty"`
	// 會議地點
	Room string `json:"room,omitempty"`
	// 會議開始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 會議結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 會議日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 負責人Attendee編號
	AID string `json:"a_id,omitempty"`
	// 負責人accounts編號
	AccountID string `json:"account_id,omitempty"`
	// 負責人姓名
	Name string `json:"name,omitempty"`
	// 負責人郵件
	Email string `json:"email,omitempty"`
	// 是否主席
	Chairman string `json:"chairman,omitempty"`
	//收過信件
	ReceiveEmail bool `json:"receive_email"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 單據類型
	OriginID string `json:"origin_id,omitempty"`
}

type AllMeetingListUserParticipant struct {
	Meeting []*struct {
		MeetingListUser
		Participant []*MeetingParticipant
	} `json:"meeting"`
	model.OutPage
}

type MeetingListUserParticipant struct {
	MeetingListUser
	Participant []*MeetingParticipant
}

//會議參與人員
type MeetingParticipant struct {
	// 負責人Attendee編號
	AID string `json:"a_id,omitempty"`
	//參與人員accounts編號
	ParticipantID string `json:"participant_id,omitempty"`
	//參與人員名稱
	ParticipantName string `json:"participant_name,omitempty"`
	//收過信件
	ReceiveEmail bool `json:"receive_email"`
	// 負責人郵件
	Email string `json:"email,omitempty"`
}

type Single struct {
	// 會議編號
	MID string `json:"m_id,omitempty"`
	// 單據編號
	DocumentsID string `json:"documents_id,omitempty"`
	// 會議名稱
	MName string `json:"m_name,omitempty"`
	// 會議地點
	Room string `json:"room,omitempty"`
	// 會議開始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 會議結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 會議日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
	// 創建時間
	CreateTime time.Time `json:"create_time"`
	// 單據類型
	OriginID string `json:"origin_id,omitempty"`
}

type Created struct {
	// 單據編號
	DocumentsID string `json:"documents_id" binding:"required" validate:"required"`
	// 會議名稱
	MName string `json:"m_name" binding:"required" validate:"required"`
	// 會議地點
	Room string `json:"room" binding:"required" validate:"required"`
	// 會議開始時間
	TimeForStart float32 `json:"time_for_start" binding:"required" validate:"required"`
	// 會議結束時間
	TimeForEnd float32 `json:"time_for_end" binding:"required" validate:"required"`
	// 會議日期
	DateForStart time.Time `json:"date_for_start" binding:"required" validate:"required"`
}

type Field struct {
	// 會議編號
	MID string `json:"m_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 單據編號
	DocumentsID string `json:"documents_id,omitempty" form:"documents_id" binding:"omitempty,uuid4"`
	// 會議名稱
	MName *string `json:"m_name,omitempty" form:"m_name"`
	// 會議地點
	Room *string `json:"room,omitempty" form:"room"`
	// 會議開始時間
	TimeForStart *float32 `json:"time_for_start,omitempty" form:"time_for_start"`
	// 會議結束時間
	TimeForEnd *float32 `json:"time_for_end,omitempty" form:"time_for_end"`
	// 會議日期
	DateForStart *string `json:"date_for_start,omitempty" form:"date_for_start"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.InPage
	model.User
}

type MeetingListUsers struct {
	Meeting []*struct {
		// 會議編號
		MID string `json:"m_id,omitempty"`
		// 單據編號
		DocumentsID string `json:"documents_id,omitempty"`
		// 會議名稱
		MName string `json:"m_name,omitempty"`
		// 會議地點
		Room string `json:"room,omitempty"`
		// 會議開始時間
		TimeForStart float32 `json:"time_for_start,omitempty"`
		// 會議結束時間
		TimeForEnd float32 `json:"time_for_end,omitempty"`
		// 會議日期
		DateForStart time.Time `json:"date_for_start,omitempty"`
		// 負責人編號
		AccountID string `json:"account_id,omitempty"`
		// 負責人姓名
		Name string `json:"name,omitempty"`
		// 是否主席
		Chairman string `json:"chairman,omitempty"`
		//收過信件
		ReceiveEmail bool `json:"receive_email"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 單據類型
		OriginID string `json:"origin_id,omitempty"`
	} `json:"meeting"`
	model.OutPage
}

type List struct {
	Meeting []*struct {
		// 會議編號
		MID string `json:"m_id,omitempty"`
		// 單據編號
		DocumentsID string `json:"documents_id,omitempty"`
		// 會議名稱
		MName string `json:"m_name,omitempty"`
		// 會議地點
		Room string `json:"room,omitempty"`
		// 會議開始時間
		TimeForStart float32 `json:"time_for_start,omitempty"`
		// 會議結束時間
		TimeForEnd float32 `json:"time_for_end,omitempty"`
		// 會議日期
		DateForStart time.Time `json:"date_for_start,omitempty"`
		// 創建時間
		CreateTime time.Time `json:"create_time"`
		// 單據類型
		OriginID string `json:"origin_id,omitempty"`
	} `json:"meeting"`
	model.OutPage
}

type Updated struct {
	// 會議編號
	MID string `json:"m_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 單據編號
	DocumentsID *string `json:"documents_id,omitempty" binding:"omitempty,uuid4"`
	// 會議名稱
	MName string `json:"m_name,omitempty"`
	// 會議地點
	Room string `json:"room,omitempty"`
	// 會議開始時間
	TimeForStart float32 `json:"time_for_start,omitempty"`
	// 會議結束時間
	TimeForEnd float32 `json:"time_for_end,omitempty"`
	// 會議日期
	DateForStart time.Time `json:"date_for_start,omitempty"`
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
	return "meeting"
}
