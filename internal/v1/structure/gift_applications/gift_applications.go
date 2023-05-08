package gift_applications

import (
	"time"

	model "eirc.app/internal/v1/structure"
	account "eirc.app/internal/v1/structure/accounts"
	gift_application_detail "eirc.app/internal/v1/structure/gift_application_details"
)

type Table struct {
	//部品零件贈送申請單編號
	GID string `gorm:"<-:create;primaryKey;column:g_id;type:UUID;default:uuid_generate_v4();" json:"g_id,omitempty"`

	Detail []gift_application_detail.Table `gorm:"foreignkey:gift_id;references:g_id" json:"detail"`
	//單號
	GCode string `gorm:"->;column:g_code;type:TEXT;not null;default:add_gift_application_code()" json:"g_code,omitempty"`
	//ERP請購單號
	ErpID string `gorm:"column:erp_id;type:UUID;" json:"erp_id,omitempty"`
	//申請者
	ApplicantID string `gorm:"column:applicant_id;type:UUID;not null;" json:"applicant_id,omitempty"`

	Account_a account.Table `gorm:"foreignkey:applicant_id;references:account_id"`
	//是否急件
	Urgent *bool `gorm:"column:urgent;type:bool;default:false;not null;" json:"urgent,omitempty"`
	//經辦
	AttnID string `gorm:"column:attn_id;type:UUID;not null;" json:"attn_id,omitempty"`

	Account_attn account.Table `gorm:"foreignkey:attn_id;references:account_id"`
	//需求者
	Purchaser string `gorm:"column:purchaser;type:TEXT;" json:"purchaser,omitempty"`
	//贈送原因
	ReasonForGiving string `gorm:"column:reason_for_giving;type:TEXT;not null;" json:"reason_for_giving,omitempty"`
	//保固
	Warranty string `gorm:"column:warranty;type:TEXT;not null;" json:"warranty,omitempty"`
	//專案代號
	ProjectCode string `gorm:"column:project_code;type:TEXT;" json:"project_code,omitempty"`
	//客戶別
	CustomerCode string `gorm:"column:customer_code;type:TEXT;" json:"customer_code,omitempty"`
	//贈送說明
	Description string `gorm:"column:description;type:TEXT;" json:"description,omitempty"`
	//作廢事由
	ReasonForVoiding string `gorm:"column:reason_for_voiding;type:TEXT;" json:"reason_for_voiding,omitempty"`
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
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty"`
	//單號
	GCode string `json:"g_code,omitempty"`

	Detail []gift_application_detail.Base `json:"detail"`
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty"`
	//申請者
	ApplicantID string `json:"applicant_id,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//經辦
	AttnID string `json:"attn_id,omitempty"`
	//需求者
	Purchaser string `json:"purchaser,omitempty"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty"`
	//保固
	Warranty string `json:"warranty,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty"`
	//贈送說明
	Description string `json:"description,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Single struct {
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty"`
	//單號
	GCode string `json:"g_code,omitempty"`
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty"`
	//申請者
	ApplicantID string `json:"applicant_id,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//經辦
	AttnID string `json:"attn_id,omitempty"`
	//需求者
	Purchaser string `json:"purchaser,omitempty"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty"`
	//保固
	Warranty string `json:"warranty,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty"`
	//贈送說明
	Description string `json:"description,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

// 放create時需輸入的欄位
type Created struct {
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty" validate:"required"`
	//申請者
	ApplicantID string `json:"applicant_id,omitempty" binding:"required,uuid4" validate:"required"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty" binding:"required" validate:"required"`
	//經辦
	AttnID string `json:"attn_id,omitempty" validate:"required"`
	//需求者
	Purchaser string `json:"purchaser,omitempty" validate:"required"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty" validate:"required"`
	//保固
	Warranty string `json:"warranty,omitempty" validate:"required"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty" validate:"required"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty" validate:"required"`
	//贈送說明
	Description string `json:"description,omitempty" validate:"required"`
	//申請狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty" validate:"required"`
	//創建者
	Creater string `json:"creater,omitempty" binding:"required,uuid4" validate:"required"`
}

// get(gatbyid)、patch、delete共用
// 放id及可變動欄位
type Field struct {
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty" from:"erp_id" binding:"omitempty,uuid4"`
	//申請者
	ApplicantID string `json:"applicant_id,omitempty" from:"applicant_id" binding:"omitempty,uuid4"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty" from:"urgent"`
	//經辦
	AttnID string `json:"attn_id,omitempty" from:"attn_id" binding:"omitempty,uuid4"`
	//需求者
	Purchaser string `json:"purchaser,omitempty" from:"purchaser"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty" from:"reason_for_giving"`
	//保固
	Warranty string `json:"warranty,omitempty" from:"warranty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty" from:"project_code"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty" from:"customer_code"`
	//贈送說明
	Description string `json:"description,omitempty" from:"description"`
	//申請狀態
	Status string `json:"status,omitempty" from:"status"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty" from:"reason_for_voiding"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty" from:"bonita_case_id"`
}

type Fields struct {
	Field
	model.InPage
}

type GiftApplication_Account struct {
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty"`
	//單號
	GCode string `json:"g_code,omitempty"`
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty"`
	//申請者
	Account_a__name string `json:"applicant,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//經辦
	Account_attn__name string `json:"attn,omitempty"`
	//需求者
	Purchaser string `json:"purchaser,omitempty"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty"`
	//保固
	Warranty string `json:"warranty,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty"`
	//贈送說明
	Description string `json:"description,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建者
	Account__name string `json:"creater_name,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type GiftDetail struct {
	GiftApplication_Account
	Detail []gift_application_detail.Base `json:"detail,omitempty"`
	// gorm:"foreignkey:gift_id;references:g_id"
}

type AllGiftDetail struct {
	GiftApplication []*GiftDetail `json:"gift_applications"`
	model.OutPage
}

type List struct {
	GiftApplications []*struct {
		//部品零件贈送申請單編號
		GID string `json:"g_id,omitempty"`
		//單號
		GCode string `json:"g_code,omitempty"`
		//ERP請購單號
		ErpID string `json:"erp_id,omitempty"`
		//申請者
		ApplicantID string `json:"applicant_id,omitempty"`
		//是否急件
		Urgent *bool `json:"urgent,omitempty"`
		//經辦
		AttnID string `json:"attn_id,omitempty"`
		//需求者
		Purchaser string `json:"purchaser,omitempty"`
		//贈送原因
		ReasonForGiving string `json:"reason_for_giving,omitempty"`
		//保固
		Warranty string `json:"warranty,omitempty"`
		//專案代號
		ProjectCode string `json:"project_code,omitempty"`
		//客戶別
		CustomerCode string `json:"customer_code,omitempty"`
		//贈送說明
		Description string `json:"description,omitempty"`
		//申請狀態
		Status string `json:"status,omitempty"`
		//作廢事由
		ReasonForVoiding string `json:"reason_for_voiding,omitempty"`
		//創建者
		Creater string `json:"creater,omitempty"`
		//創建時間
		CreatedTime time.Time `json:"created_time"`
		//bonita_case_id
		BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	} `json:"gift_applications"`
	model.OutPage
}

// 放id及可變動欄位
type Updated struct {
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty" binding:"omitempty,uuid4"`
	//申請者
	ApplicantID string `json:"applicant_id,omitempty" binding:"omitempty,uuid4"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//經辦
	AttnID string `json:"attn_id,omitempty" binding:"omitempty,uuid4"`
	//需求者
	Purchaser string `json:"purchaser,omitempty"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty"`
	//保固
	Warranty string `json:"warranty,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty"`
	//贈送說明
	Description string `json:"description,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty"`
	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Updated_Bonita struct {
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//CaseID
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
}

type Review struct {
	//gift_application
	//部品零件贈送申請單編號
	GID string `json:"g_id,omitempty"`
	//單號
	GCode string `json:"g_code,omitempty"`
	//ERP請購單號
	ErpID string `json:"erp_id,omitempty"`
	//申請者
	Account_a__name string `json:"applicant,omitempty"`
	//是否急件
	Urgent *bool `json:"urgent,omitempty"`
	//經辦
	Account_attn__name string `json:"attn,omitempty"`
	//需求者
	Purchaser string `json:"purchaser,omitempty"`
	//贈送原因
	ReasonForGiving string `json:"reason_for_giving,omitempty"`
	//保固
	Warranty string `json:"warranty,omitempty"`
	//專案代號
	ProjectCode string `json:"project_code,omitempty"`
	//客戶別
	CustomerCode string `json:"customer_code,omitempty"`
	//贈送說明
	Description string `json:"description,omitempty"`
	//申請狀態
	Status string `json:"status,omitempty"`
	//作廢事由
	ReasonForVoiding string `json:"reason_for_voiding,omitempty"`
	//創建者
	Creater string `json:"creater,omitempty"`
	//創建者
	Account__name string `json:"creater_name,omitempty"`
	//創建時間
	CreatedTime time.Time `json:"created_time"`

	//bonita_case_id
	BonitaCaseID float32 `json:"bonita_case_id,omitempty"`
	//bonita_case_id
	BonitaTaskID string `json:"bonita_task_id,omitempty"`
	//任務名稱
	BonitaTaskName string `json:"bonita_task_name,omitempty"`
}

func (a *Table) TableName() string {
	return "gift_application"
}
