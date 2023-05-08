package accounts

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is companies database table struct
type Table struct {
	// 編號
	AccountID string `gorm:"primaryKey;column:account_id;uuid_generate_v4()type:UUID;" json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `gorm:"column:company_id;type:UUID;" json:"company_id,omitempty"`
	// 帳號
	Account string `gorm:"column:account;type:VARCHAR;" json:"account,omitempty"`
	// 中文名稱
	Name string `gorm:"column:name;type:VARCHAR;" json:"name,omitempty"`
	// 密碼
	Password string `gorm:"column:pwd;type:VARCHAR;" json:"password,omitempty"`
	// 角色編號
	//RoleID string `gorm:"column:role_id;type:VARCHAR;" json:"role_id,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:update_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `gorm:"column:update_by;type:UUID;" json:"updated_by,omitempty"`
	//[ 3] dep                                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	//Dep string `gorm:"column:dep;type:UUID;" json:"dep,omitempty"`
	//[ 4] phone                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Phone string `gorm:"column:phone;type:TEXT;" json:"phone,omitempty"`
	//[ 5] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Email string `gorm:"column:email;type:TEXT;" json:"email,omitempty"`
	//[ 6] status                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Status bool `gorm:"column:status;type:bool;true" json:"status,omitempty"`
	//[ 7] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaUserID string `gorm:"column:bonita_user_id;type:TEXT;" json:"bonita_user_id,omitempty"`
	//[ 7] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaManagerID string `gorm:"column:bonita_manager_id;type:TEXT;" json:"bonita_manager_id,omitempty"`

}

type Create_Table struct {
	// 編號
	AccountID string `gorm:"primaryKey;column:account_id;uuid_generate_v4()type:UUID;" json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `gorm:"column:company_id;type:UUID;" json:"company_id,omitempty"`
	// 帳號
	Account string `gorm:"column:account;type:VARCHAR;" json:"account,omitempty"`
	// 中文名稱
	Name string `gorm:"column:name;type:VARCHAR;" json:"name,omitempty"`
	// 密碼
	Password string `gorm:"column:pwd;type:VARCHAR;" json:"password,omitempty"`
	// 角色編號
	//RoleID string `gorm:"column:role_id;type:VARCHAR;" json:"role_id,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//[ 3] dep                                            UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	//Dep string `gorm:"column:dep;type:UUID;" json:"dep,omitempty"`
	//[ 4] phone                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Phone string `gorm:"column:phone;type:TEXT;" json:"phone,omitempty"`
	//[ 5] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Email string `gorm:"column:email;type:TEXT;" json:"email,omitempty"`
	//[ 7] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaUserID string `gorm:"column:bonita_user_id;type:TEXT;" json:"bonita_user_id,omitempty"`
	//[ 7] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaManagerID string `gorm:"column:bonita_manager_id;type:TEXT;" json:"bonita_manager_id,omitempty"`
}

type Base struct {
	// 編號
	AccountID string `json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 角色編號
	//RoleID string `json:"role_id,omitempty"`
	// 所屬部門
	//Dep string `json:"dep,omitempty"`
	// 電話
	Phone string `json:"phone,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 帳號狀態
	Status bool `json:"status,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
	// BonitaUserID
	BonitaUserID string `json:"bonita_user_id,omitempty"`
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`
}

type Account_Name struct {
	// 編號
	AccountID string `json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 角色編號
	//RoleID string `json:"role_id,omitempty"`
	// 所屬部門
	//Dep string `json:"dep,omitempty"`
	// BonitaUserID
	BonitaUserID string `json:"bonita_user_id,omitempty"`
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`
	//所屬部門名稱
	DepName string `json:"dep_name,omitempty"`
	//所屬部門的父部門名稱
	PDepName string `json:"p_dep_name,omitempty"`
	//職稱
	JobtitleName string `json:"jobtitle_name,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
}

type Account_Names struct {
	Accounts []*struct {
		// 編號
		AccountID string `json:"account_id,omitempty"`
		// 公司ID
		CompanyID string `json:"company_id,omitempty"`
		// 帳號
		Account string `json:"account,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 角色編號
		//RoleID string `json:"role_id,omitempty"`
		// 所屬部門
		//Dep string `json:"dep,omitempty"`
		// BonitaUserID
		BonitaUserID string `json:"bonita_user_id,omitempty"`
		// BonitaUserID
		BonitaManagerID string `json:"bonita_manager_id,omitempty"`
		//所屬部門名稱
		DepName string `json:"dep_name,omitempty"`
		//所屬部門的父部門名稱
		PDepName string `json:"p_dep_name,omitempty"`
		//職稱
		JobtitleName string `json:"jobtitle_name,omitempty"`
		// 創建時間
		CreatedAt time.Time `json:"created_at"`
	} `json:"accounts"`
	model.OutPage
}

type Single struct {
	// 編號
	AccountID string `json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 角色編號
	//RoleID string `json:"role_id,omitempty"`
	// 所屬部門
	//Dep string `json:"dep,omitempty"`
	// 電話
	Phone string `json:"phone,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 帳號狀態
	Status bool `json:"status,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
	//[ 6] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaUserID string `json:"bonita_user_id,omitempty"`
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`
}

type Created struct {
	// 公司ID
	CompanyID string `json:"company_id"`
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 中文名稱
	Name string `json:"name" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password" binding:"required" validate:"required"`
	// 角色編號
	//RoleID string `json:"role_id"`
	// 所屬部門
	//Dep string `json:"dep" binding:"required" validate:"required"`
	// 電話
	Phone string `json:"phone"`
	// 電子郵件
	Email string `json:"email"`
	// 創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
	//[ 6] bonita_user_id                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	BonitaUserID string `json:"bonita_user_id,omitempty"`
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`

	//bonita顯示的名稱
	UserName string `json:"userName,omitempty"`
	//bonita父部門
	ManagerID string `json:"manager_id,omitempty"`
	//bonita是否啟用使用者
	Enabled string `json:"enabled,omitempty"`
	// 創建者帳號
	CreatedAccount string `json:"created_account,omitempty"`
}

type Field struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty" form:"company_id" `
	// 帳號
	Account *string `json:"account,omitempty" form:"account"`
	// 密碼
	Password *string `json:"password" swaggerignore:"true"`
	// 中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 角色編號
	//RoleID *string `json:"role_id,omitempty" form:"role_id"`
	// 所屬部門
	//Dep *string `json:"dep,omitempty" form:"dep"`
	// 電話
	Phone *string `json:"phone,omitempty" form:"phone"`
	// 電子郵件
	Email *string `json:"email,omitempty" form:"email"`
	// 是否刪除
	Status *bool `json:"status,omitempty" swaggerignore:"true"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
	// bonita_user_id
	BonitaUserID *string `json:"bonita_user_id,omitempty" form:"bonita_user_id"`
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type Users struct {
	Field
	model.User
}

type List struct {
	Accounts []*struct {
		// 編號
		AccountID string `json:"account_id,omitempty"`
		// 公司ID
		CompanyID string `json:"company_id,omitempty"`
		// 帳號
		Account string `json:"account,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 角色編號
		//RoleID string `json:"role_id,omitempty"`
		// 所屬部門
		//Dep string `json:"dep,omitempty"`
		// 電話
		Phone string `json:"phone,omitempty"`
		// 電子郵件
		Email string `json:"email,omitempty"`
		// 帳號狀態
		Status bool `json:"status,omitempty"`
		// bonita_user_id
		BonitaUserID *string `json:"bonita_user_id,omitempty"`
		// BonitaUserID
		BonitaManagerID string `json:"bonita_manager_id,omitempty"`
	} `json:"accounts"`
	model.OutPage
}

type Updated struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 組織ID
	CompanyID *string `json:"company_id,omitempty" `
	// 中文名稱
	Name *string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 角色編號
	//RoleID *string `json:"role_id,omitempty"`
	// 所屬部門
	//Dep string `json:"dep" `
	// 電話
	Phone string `json:"phone" `
	// 電子郵件
	Email string `json:"email" `
	// 帳號狀態
	Status bool `json:"status"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
	// bonita_user_id
	BonitaUserID *string `json:"bonita_user_id,omitempty" `
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`

	//bonita顯示的名稱
	UserName string `json:"userName,omitempty"`
	//bonita父部門
	ManagerID string `json:"manager_id,omitempty"`
	//bonita是否啟用使用者
	Enabled string `json:"enabled,omitempty"`
	// 創建者帳號
	CreatedAccount string `json:"created_account,omitempty"`
}

type UpdatedCsv struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 組織ID
	CompanyID *string `json:"company_id,omitempty" `
	// 中文名稱
	Name *string `json:"name,omitempty"`
	// 中文名稱
	Account *string `json:"account,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 角色編號
	//RoleID *string `json:"role_id,omitempty"`
	// 所屬部門
	//Dep string `json:"dep" `
	// 電話
	Phone string `json:"phone" `
	// 電子郵件
	Email string `json:"email" `
	// 帳號狀態
	Status bool `json:"status"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
	// bonita_user_id
	BonitaUserID *string `json:"bonita_user_id,omitempty" `
	// BonitaUserID
	BonitaManagerID string `json:"bonita_manager_id,omitempty"`

	//bonita顯示的名稱
	UserName string `json:"userName,omitempty"`
	//bonita父部門
	ManagerID string `json:"manager_id,omitempty"`
	//bonita是否啟用使用者
	Enabled string `json:"enabled,omitempty"`
	// 創建者帳號
	CreatedAccount string `json:"created_account,omitempty"`
}

type UpdatedCsvList struct {
	Account []*UpdatedCsv `json:"account"`
}

func (a *Table) TableName() string {
	return "accounts"
}

func (a *Create_Table) TableName() string {
	return "accounts"
}

