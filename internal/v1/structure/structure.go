package structure

type InPage struct {
	//頁數(請從1開始帶入)
	Page int64 `json:"page" binding:"required,gt=0" validate:"required" form:"page"`
	//筆數(請從1開始帶入,最高上限20)
	Limit int64 `json:"limit" binding:"required,gt=0" validate:"required" form:"limit"`
}

type User struct {
	// 負責人
	UserID string `json:"user_id,omitempty" binding:"omitempty,uuid4" form:"user_id"`
}

type User_TID struct {
	// 負責人
	UserID string `json:"user_id,omitempty" binding:"omitempty,uuid4" form:"user_id"`
	// 負責人
	TID string `json:"t_id,omitempty" binding:"omitempty,uuid4" form:"t_id"`
	//頁數(請從1開始帶入)
	Page int64 `json:"page" binding:"required,gt=0" validate:"required" form:"page"`
	//筆數(請從1開始帶入,最高上限20)
	Limit int64 `json:"limit" binding:"required,gt=0" validate:"required" form:"limit"`
}

type OutPage struct {
	//頁數結構
	InPage
	//總頁數
	Pages int64 `json:"pages"`
	//總筆數
	Total int64 `json:"total"`
}

type OutTotal struct {
	//總筆數
	Total int64 `json:"total"`
}

type GonitaUser struct {
	//bonita登入帳號
	Account string `json:"account"`
}

type GonitaMembership struct {
	//bonita登入帳號
	Account string `json:"account"`
	//bonita隸屬部門用的userid
	UserID string `json:"user_id"`
	//bonita隸屬部門用的group_id
	GroupID string `json:"group_id"`
	//bonita隸屬部門用的role_id
	RoleID string `json:"role_id"`
}

type GetDepartmentID struct {
	//Bonita部門Id
	ID string `json:"id"`
}

type GonitaCaseList struct {
	//筆數(請從1開始帶入,最高上限20)
	Rows string `json:"rows"`
	//
	UserID string `json:"user_id"`
}

type GetCaseListInput struct {
	//bonita登入帳號
	Account string `json:"account"`
	//
	UserID string `json:"user_id"`
	//gateway_data編號
	GdID string `json:"gd_id"`
}

type GetDetailListInput struct {
	//bonita登入帳號
	Account string `json:"account"`
	//
	UserID string `json:"user_id"`
	//
	CaseID string `json:"case_id"`
}

type GetCasekDetailOutput struct {
	DisplayName string `json:"displayName,omitempty"`
	//任務ID
	ID string `json:"id,omitempty"`
}

type CaseIDModelInput struct {
	GonitaUser
	//B2起單
	Pm string   `json:"pm,omitempty"`
	Tm []string `json:"tm,omitempty"`

	//C2起單
	Assistant string `json:"assistant,omitempty"`
	Recipient string `json:"recipient,omitempty"`

	//D啟單
	Money string `json:"money,omitempty"`
	Attm  string `json:"attm,omitempty"`
}

type GetCaseIDOutput struct {
	CaseID float32 `json:"caseId,omitempty"`
}

type ReviewInput struct {
	//B2_最高主管審核
	GmApprovalStatus *bool `json:"gmApprovalStatus,omitempty"`
	//B2_部門主管審核
	DmApprovalStatus *bool `json:"dmApprovalStatus,omitempty"`
	//B2_突發任務
	Tm []int `json:"tm,omitempty"`

	//A1&C2&工時異動&D_審核
	Status *bool `json:"status,omitempty"`
	//A1_新增任務
	Designee int `json:"designee,omitempty"`
	//A1_PM人選確認並負責RD部門勾選
	Department []int `json:"department,omitempty"`
	Pm         int   `json:"pm,omitempty"`
	//A1_指派各部門人員(可能1人或多人)
	Dstaff []int `json:"dstaff,omitempty"`

	//C2重起單
	Assistant string `json:"assistant,omitempty"`
	Recipient string `json:"recipient,omitempty"`

	//工時異動單號
	HmID string `json:"hm_id,omitempty"`

	//防毒軟體單號
	AsID string `json:"as_id,omitempty"`
}

type TransferTaskInput struct {
	TransferTask []*struct {
		//任務轉移
		BonitaTaskID string `json:"bonita_task_id,omitempty"`
		BonitaUserID string `json:"bonita_user_id,omitempty"`
	} `json:"transfer_task"`
}

type GetProcessIDList struct {
	//ProccessID
	ID string `json:"id,omitempty"`
	//名稱
	Name string `json:"name,omitempty"`
}

type GonitaListCount struct {
	//類別(A1 B2 C2....)
	Category string `json:"category,omitempty"`
	//數量
	Count int `json:"count,omitempty"`
	//任務名稱
	Name string `json:"name,omitempty"` //string
}

type GetCaseListOutput struct {
	//單據ID
	CaseID string `json:"caseId,omitempty"`
	//任務ID
	ID string `json:"id,omitempty"`
	//sub單據Id
	ParentCaseID string `json:"parentCaseId,omitempty"`
	//任務名稱
	Name string `json:"name,omitempty"` //string
	// DisplayDescription string `json:"displayDescription,omitempty"` //string
	// ExecutedBy float32 `json:"executedBy,omitempty"` //long
	// RootContainerId float32 `json:"rootContainerId,omitempty"` //long
	// Assigned_date time.Time `json:"assigned_date,omitempty"` //dateTime
	// DisplayName string `json:"displayName,omitempty"` //string
	// ExecutedBySubstitute float32 `json:"executedBySubstitute,omitempty"` //long
	// DueDate time.Time `json:"dueDate,omitempty"` //dateTime
	// Description string `json:"description,omitempty"` //string
	// Type string `json:"type"` //string
	// Priority string `json:"priority,omitempty"` //string
	// ActorId float32 `json:"actorId,omitempty"` //long
	// ProcessId float32 `json:"processId,omitempty"` //long
	// ReachedStateDate time.Time `json:"reached_state_date,omitempty"` //dateTime
	// RootCaseId float32 `json:"rootCaseId,omitempty"` //long
	// State string `json:"state,omitempty"` //string
	// ParentCaseId float32 `json:"parentCaseId,omitempty"`//long
	// LastUpdateDate time.Time `json:"last_update_date,omitempty"` //dateTime
	// AssignedID float32 `json:"assigned_id,omitempty"` //long
}

type GetCaseDetailListOutput struct {
	//customer_demand表
	//編號
	CdID string `json:"cd_id,omitempty"`
	//單號
	CdCode string `json:"cd_code,omitempty"`
	//客戶需求說明
	DemandContent string `json:"demand_content,omitempty"`

	//project表
	//編號
	PID string `json:"p_id,omitempty"`
	//單號
	PCode string `json:"p_code,omitempty"`
	//專案名稱
	PName string `json:"p_name,omitempty"`

	//manufactor_order表
	// 編號
	MID string `json:"m_id,omitempty"`
	//製令單號
	MCode string `json:"m_code,omitempty"`
	//製令名稱
	ProjectDetail string `json:"project_detail,omitempty"`

	//task表
	//編號
	TID string `json:"t_id,omitempty"`
	//本任務名稱
	TName string `json:"t_name,omitempty"`
	//任務負責人
	PrincipalName string `json:"principal_name,omitempty"`
	//任務負責人
	TuID string `json:"tu_id,omitempty"`

	//labor_hour表
	// 工時編號
	HmID string `json:"hm_id,omitempty"`
	// 原工時編號
	HourID string `json:"hour_id,omitempty"`
	// 工時歸屬
	HCategory string `json:"h_category,omitempty"`
	//提工時的人
	HCreaterName string `json:"h_creater_name,omitempty"`
	// 工時
	HLaborhour string `json:"h_laborhour,omitempty"`

	//類別(A1 B2 C2....)
	Category string `json:"category,omitempty"`
	//[ 7] bonita_case_id
	BonitaCaseID string `json:"bonita_case_id,omitempty"`
	//[ 7] bonita_case_id
	BonitaTaskID       string `json:"bonita_task_id,omitempty"`
	BonitaParentCaseID string `json:"bonita_parentcase_id,omitempty"`
	//任務名稱
	BonitaTaskName string `json:"bonita_task_name,omitempty"`
}

type SendEmailInput struct {
	//Port
	Port string `json:"port,omitempty" binding:"required" validate:"required"`
	//Host
	Host string `json:"host,omitempty" binding:"required" validate:"required"`
	//寄件人
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	//寄件人信箱
	Username string `json:"username,omitempty" binding:"required" validate:"required"`
	//寄件人密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	//收件人信箱
	To string `json:"to,omitempty" binding:"required" validate:"required"`
	//主旨
	Subject string `json:"subject,omitempty" binding:"required" validate:"required"`
	//內容
	Body string `json:"body,omitempty" binding:"required" validate:"required"`
}

// bonita 取得ProccessID
type GetProccessIDList struct {
	//ProccessID
	ID string `json:"id,omitempty"`
	//名稱
	Name string `json:"name,omitempty"`
}
