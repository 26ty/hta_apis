package bpm

import (
	"encoding/json"
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"github.com/gin-gonic/gin"

	gonita "eirc.app/internal/pkg/gonita"
	model "eirc.app/internal/v1/structure"
	accounts "eirc.app/internal/v1/structure/accounts"
	department "eirc.app/internal/v1/structure/department"
)

// GetProcessID
// 【用途】用來決定要走哪個流程(B2、A1、C2、工時)
// 【Input】Account-帳號，Name-流程名稱
// 【Output】ProcessID-流程ID
func GetProcessID(ctx *gin.Context, account string, name string) string {
	client := gonita.New(account)
	body := client.GetProcessInstanceId()

	output := []model.GetProcessIDList{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, "bonita GetProcessID error"))

		return "error"
	}

	for _, value := range output {
		if value.Name == name {
			return value.ID
		}
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, "bonita GetProcessID 404"))
	return "error"
}

// GetUserExecutable
// 【用途】取得使用者可執行的單
// 【Input】Account-帳號，BonitaUserID-Bonita用戶ID
// 【Output】CaseID-單據ID，ID-任務ID，ParentCaseID-子流程ID，Name-任務名稱
func GetUserExecutable(ctx *gin.Context, account string, userID string) (output *[]model.GetCaseListOutput) {
	client := gonita.New(account)
	body := client.GetStateCaseList("9999", "ready", userID)

	err := json.Unmarshal(body, &output)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, "bonita GetStateCaseList error"))

		return nil
	}

	return output
}

// GetCasePendingTaskDetail
// 【用途】取得該單據待執行詳細資料
// 【Input】Account-帳號，CaseID-單據ID
// 【Output】 CaseID-單據ID，ID-任務ID，ParentCaseID-子流程ID，Name-任務名稱
func GetCasePendingTaskDetail(ctx *gin.Context, account string, caseID string) (output *[]model.GetCaseListOutput) {
	client := gonita.New(account)
	body := client.GetCasePendingTaskDetail(caseID)

	err := json.Unmarshal(body, &output)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, "bonita GetCasePendingTaskDetail error"))

		return nil
	}

	return output
}

// GetCaseID
// 【用途】啟單
// 【Input】ProcessID-流程ID，Account-Bonita登入帳號，Pm、Tm[]-B2起單，Assistant、Recipient-C2起單
// 【Output】 CaseID-單據ID，ID-任務ID，ParentCaseID-子流程ID，Name-任務名稱
func GetCaseID(ctx *gin.Context, processID string, input *model.CaseIDModelInput) (output *model.GetCaseIDOutput) {

	client := gonita.New(input.Account)
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return nil
	}

	body := client.CreateProcessCase(processID, string(marshal))
	//fmt.Println(body)

	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, "bonita CreateProcessCase error"))

		return nil
	}
	return output
}

// ReviewTask
// 【用途】審核任務
// 【Input】Account-帳號，TaskID-任務ID，GmApprovalStatus-B2_最高主管審核，
// 　　DmApprovalStatus-B2_部門主管審核，Tm[]-B2_突發任務，Status-A1&C2_審核，
// 　　Designee-A1_新增任務，Department、Pm-A1_PM人選確認並負責RD部門勾選，
// 　　Dstaff-A1_指派各部門人員(可能1人或多人)，Assistant、Recipient-C2起單
// 【Output】StatusCode-狀態碼
func ReviewTask(ctx *gin.Context, account string, taskID string, input *model.ReviewInput) (body int) {

	client := gonita.New(account)
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	body = client.ExecuteTask(taskID, string(marshal))

	return body
}

// AddDepartmentGetID
// 【用途】新增部門
// 【Input】Account-Bonita登入帳號，DID-部門表編號，Manager-部門主管，Name-中文名稱
// 　，EngName-英文名稱，Introduction-路徑，Fax，Tel，BonitaGroupID-Bonita部門
// 　　ID，BonitaParentGroupID-Bonita父部門ID，DisplayName-Bonita顯示的名稱，
// 　　ParentGroupID-Bonita父部門
// 【Output】ID-Bonita部門ID
func AddDepartmentGetID(ctx *gin.Context, input *department.Created) string {

	client := gonita.New(input.Account)

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return ""
	}

	body := client.AddGroup(string(marshal))

	output := model.GetDepartmentID{}

	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, "bonita AddGroup error"))

		return ""
	}

	if output.ID == "" {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, "bonita AddGroup error"))
		return ""
	}

	return output.ID
}

// UpdateDepartmentGetcode
// 【用途】更新部門
// 【Input】Account-Bonita登入帳號，DID-部門表編號，Manager-部門主管，Name-中文名稱
// 　　，EngName-英文名稱，Introduction-路徑，Fax，Tel，BonitaGroupID-Bonita部門
// 　　ID， BonitaParentGroupID-Bonita父部門ID，DisplayName-Bonita顯示的名稱，
// 　　ParentGroupID-Bonita父部門，GroupID-部門ID
// 【Output】StatusCode-狀態碼
func UpdateDepartmentGetcode(input *department.Updated, groupId string) int {

	client := gonita.New(input.Account)

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return -1
	}

	body := client.EditGroup(string(marshal), groupId)

	return body
}

// DeleteDepartmentGetcode
// 【用途】刪除部門
// 【Input】Account-Bonita登入帳號，DID-部門表編號，Manager-部門主管，Name-中文名稱，
// 　　EngName-英文名稱，Introduction-路徑，Fax，Tel，BonitaGroupID-Bonita部門
// 　　ID， BonitaParentGroupID-Bonita父部門ID，DisplayName-Bonita顯示的名稱，
// 　　ParentGroupID-Bonita父部門，GroupID-部門ID
// 【Output】StatusCode-狀態碼
func DeleteDepartmentGetcode(input *department.Updated, groupId string) int {

	client := gonita.New(input.Account)

	body := client.DeleteGroup(groupId)

	return body
}

// AddUserGetID
// 【用途】新增人員
// 【Input】CompanyID-公司ID，Account-帳號，Name-中文名稱，Password-密碼，
// 　　Dep-所屬部門，Phone-電話，Email-電子郵件，CreatedBy-創建者，BonitaUserID，
// 　　BonitaManagerID，UserName-bonita顯示的名稱，ManagerID-bonita父部門，
// 　　Enabled-bonita是否啟用使用者，CreatedAccount-創建者帳號
// 【Output】BonitaUserID，新增成功與否
func AddUserGetID(ctx *gin.Context, input *accounts.Created) int {

	client := gonita.New(input.CreatedAccount)
	bonita_input := *input
	bonita_input.Password = "12345"

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return 0
	}

	id, status := client.AddUser(string(marshal))

	if status != true {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, "bonita AddUser error"))
		return 0
	}

	return id
}

// UpdateUserGetcode
// 【用途】編輯人員
// 【Input】AccountID-編號，CompanyID-公司ID，Name-中文名稱，Password-密碼，
// 　　Dep-所屬部門， Phone-電話，Email-電子郵件，Status-帳號狀態，UpdateBy-更新者，
// 　　IsDeleted-是否刪除，BonitaUserID，BonitaManagerID，UserName-bonita顯示的名稱，
// 　　ManagerID-bonita父部門，Enabled-bonita是否啟用使用者，CreatedAccount-創建者帳號，　　　UserID-BonitaUserID
// 【Output】StatusCode-狀態碼
func UpdateUserGetcode(input *accounts.Updated, userId string) int {

	client := gonita.New(input.CreatedAccount)
	bonita_input := *input
	bonita_input.Password = "12345"

	marshal, err := json.Marshal(bonita_input)
	if err != nil {
		log.Error(err)

		return -1
	}

	body := client.EditUser(userId, string(marshal))

	return body
}

// DeleteUserGetcode
// 【用途】刪除人員
// 【Input】AccountID-編號，CompanyID-公司ID，Name-中文名稱，Password-密碼，
// 　　Dep-所屬部門， Phone-電話，Email-電子郵件，Status-帳號狀態，UpdateBy-更新者，
// 　　IsDeleted-是否刪除，BonitaUserID，BonitaManagerID，UserName-bonita顯示的名稱，
// 　　ManagerID-bonita父部門，Enabled-bonita是否啟用使用者，CreatedAccount-創建者帳號，
// 　　UserID-BonitaUserID
// 【Output】 StatusCode-狀態碼
func DeleteUserGetcode(input *accounts.Updated, userId string) int {

	client := gonita.New(input.CreatedAccount)

	body := client.DeleteUser(userId)

	return body
}

// AddUserMembership
// 【用途】設定人員隸屬部門
// 【Input】Account-bonita登入帳號，UserID-bonita隸屬部門用的userid，
// 　　GroupID-bonita隸屬部門用的group_id，RoleID-bonita隸屬部門用的role_id
// 【Output】成功與否
func AddUserMembership(input *model.GonitaMembership) bool {

	client := gonita.New(input.Account)

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return false
	}

	body := client.AddMembership(string(marshal))

	return body
}

// UpdateMembershipGetcode
// 【用途】編輯人員隸屬部門(for 單一角色)
// 【Input】Account-bonita登入帳號，UserID-bonita隸屬部門用的userid，
// 　　GroupID-bonita隸屬部門用的group_id，RoleID-bonita隸屬部門用的role_id，
// 　　GroupID-部門ID，RoleID-組織角色ID
// 【Output】成功與否
func UpdateMembershipGetcode(input *model.GonitaMembership, groupId string, roleId string) bool {

	client := gonita.New(input.Account)

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return false
	}

	body := client.EditUserMembership(groupId, roleId, string(marshal))

	return body
}

// DeleteMembershipGetcode
// 【用途】刪除人員隸屬部門
// 【Input】Account-bonita登入帳號，UserID-bonita隸屬部門用的userid，
// 　　GroupID-bonita隸屬部門用的group_id，RoleID-bonita隸屬部門用的role_id，
// 　　UserID-BonitaUserID，GroupID-部門ID，RoleID-組織角色ID
// 【Output】StatusCode-狀態碼
func DeleteMembershipGetcode(input *model.GonitaMembership, userId string, groupId string, roleId string) int {

	client := gonita.New(input.Account)

	body := client.DeleteMembership(userId, groupId, roleId)

	return body
}

// TransferTask
// 【用途】任務轉移
// 【Input】Account-bonita登入帳號，taskID-任務ID，userID-轉移至的人員bonita_user_id
// 【Output】StatusCode-狀態碼
func TransferTask(ctx *gin.Context, account string, taskID string, userID string) (body int) {

	client := gonita.New(account)

	body = client.UpdateAssignedId(userID, taskID)

	return body
}
