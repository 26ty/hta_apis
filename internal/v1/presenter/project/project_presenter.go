package project

import (
	"net/http"
	//"fmt"
	//"reflect"
	//"encoding/json"
	// "strconv"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	bpm "eirc.app/internal/pkg/bpm"
	preset "eirc.app/internal/v1/presenter"
	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/project"
	//gonita "eirc.app/internal/pkg/gonita/gonita"
	// gonita "gonita"
	//bpm "github.com/Rudy1021/go-bonita-client/bpm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Project.c 新增專案
// @description 新增專案
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body project.Created true "新增專案"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//createBy := util.GenerateUUID()
	input := &project.Created{}
	//input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Project.1 條件搜尋專案
// @description 條件專案
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param code query string false "專案代號"
// @param type query string false "類型"
// @param p_name query string false "專案名稱"
// @param customer_id query string false "客戶ID"
// @param salesman_id query string false "業務負責人ID"
// @param serviceman_id query string false "客服負責人ID"
// @param projectman_id query string false "專案負責人ID"
// @param status query string false "狀態"
// @param inner_id query string false "內部編號"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=project.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &project.Fields{}

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ProjectResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetB2BonitaCaseListStop
// @Summary Project.b2-end 獲取專案終止的任務ID
// @description B2獲取專案終止的任務ID
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param caseID path string true "Bonita專案ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListStop/{account}/{caseID} [get]
func (p *presenter) GetB2BonitaCaseListStop(ctx *gin.Context) {
	input := &model.GetDetailListInput{}
	caseID := ctx.Param("caseID")
	input.CaseID = caseID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetCasePendingTaskDetail(ctx,input.Account,input.CaseID)

	if output != nil{
		for _, value := range *output {
		if value.Name == "專案終止"{
			ctx.JSON(http.StatusOK, value.ID)
			}
		}	
	}
	
}

// GetB2BonitaCaseListTask
// @Summary Project.b2-ex-s 獲取新增任務的任務ID
// @description B2獲取新增任務的任務ID
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param caseID path string true "Bonita專案ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListTask/{account}/{caseID} [get]
func (p *presenter) GetB2BonitaCaseListTask(ctx *gin.Context) {
	input := &model.GetDetailListInput{}
	caseID := ctx.Param("caseID")
	input.CaseID = caseID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetCasePendingTaskDetail(ctx,input.Account,input.CaseID)

	if output != nil{
		for _, value := range *output {
		if value.Name == "新增任務"{
			ctx.JSON(http.StatusOK, value.ID)
			}
		}	
	}
	
}

// GetB2BonitaCaseListPM
// @Summary Project.b2-1 獲取使用者可執行的單(PM待回報專案)
// @description B2-1獲取使用者可執行的單(PM待回報專案)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]project.Project_Account} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListPM/{account}/{userID} [get]
func (p *presenter) GetB2BonitaCaseListPM(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)
	// client := gonita.New(input.Account)
	// body := client.GetStateCaseList("9999","ready",input.UserID)

	// output := &[]model.GetCaseListOutput{}

	// err := json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// List := &project.Field{}
	// update.PID = pID
	// update.BonitaCaseID = output.CaseID
	if output != nil {
		codeMessage := p.ProjectResolver.GetByCaseIDtoPM(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetB2BonitaCaseListTop
// @Summary Project.b2-2 獲取使用者可執行的單(專案完工送審)
// @description B2-2獲取使用者可執行的單(專案完工送審)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]project.Project_Account} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListTop/{account}/{userID} [get]
func (p *presenter) GetB2BonitaCaseListTop(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)
	// client := gonita.New(input.Account)
	// body := client.GetStateCaseList("9999","ready",input.UserID)

	// output := &[]model.GetCaseListOutput{}

	// err := json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// List := &project.Field{}
	// update.PID = pID
	// update.BonitaCaseID = output.CaseID
	if output != nil {
		codeMessage := p.ProjectResolver.GetBonitaCaseListPMCompleted(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetB2BonitaCaseListTM
// @Summary Project.b2-01 獲取使用者可執行的單(專案任務工作回報)
// @description B2-01獲取使用者可執行的單(專案任務工作回報)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]project.Tm_Return} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListTM/{account}/{userID} [get]
func (p *presenter) GetB2BonitaCaseListTM(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)
	// client := gonita.New(input.Account)
	// body := client.GetStateCaseList("9999","ready",input.UserID)

	// output := &[]model.GetCaseListOutput{}

	// err := json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// List := &project.Field{}
	// update.PID = pID
	// update.BonitaCaseID = output.CaseID
	if output != nil {
		codeMessage := p.ProjectResolver.GetB2BonitaCaseListTM(*output, userID)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetB2BonitaCaseListCountersign
// @Summary Project.b2-02 獲取使用者可執行的單(會簽)
// @description B2-02獲取使用者可執行的單(會簽)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]project.Tm_Return} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListCountersign/{account}/{userID} [get]
func (p *presenter) GetB2BonitaCaseListCountersign(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)
	// client := gonita.New(input.Account)
	// body := client.GetStateCaseList("9999","ready",input.UserID)

	// output := &[]model.GetCaseListOutput{}

	// err := json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// List := &project.Field{}
	// update.PID = pID
	// update.BonitaCaseID = output.CaseID
	if output != nil {
		codeMessage := p.ProjectResolver.GetB2BonitaCaseListCountersign(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetB2BonitaCaseListConfirm
// @Summary Project.b2-04獲取使用者可執行的單(確認會簽內容)
// @description b2-04獲取使用者可執行的單(確認會簽內容)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]project.Tm_Return} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListConfirm/{account}/{userID} [get]
func (p *presenter) GetB2BonitaCaseListConfirm(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)
	// client := gonita.New(input.Account)
	// body := client.GetStateCaseList("9999","ready",input.UserID)

	// output := &[]model.GetCaseListOutput{}

	// err := json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// List := &project.Field{}
	// update.PID = pID
	// update.BonitaCaseID = output.CaseID
	if output != nil {
		codeMessage := p.ProjectResolver.GetB2BonitaCaseListConfirm(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetB2BonitaCaseListDepartment
// @Summary Project.b2-03 獲取使用者可執行的單(專案任務完工送審)
// @description B2-03獲取使用者可執行的單(專案任務完工送審)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]project.Tm_Return} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetB2BonitaCaseListDepartment/{account}/{userID} [get]
func (p *presenter) GetB2BonitaCaseListDepartment(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)
	// client := gonita.New(input.Account)
	// body := client.GetStateCaseList("9999","ready",input.UserID)

	// output := &[]model.GetCaseListOutput{}

	// err := json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// List := &project.Field{}
	// update.PID = pID
	// update.BonitaCaseID = output.CaseID
	if output != nil {
		codeMessage := p.ProjectResolver.GetB2BonitaCaseListDepartment(*output, userID)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetByProjectBonitaUserList
// @Summary Project.3 取得指定專案的專案負責人資料
// @description 取得指定專案的專案負責人資料
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string true "專案ID"
// @success 200 object code.SuccessfulMessage{body=project.Bonita_ID_Lists} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetByProjectBonitaUserList/{pID} [get]
func (p *presenter) GetByProjectBonitaUserList(ctx *gin.Context) {
	input := &project.Users{}
	pID := ctx.Param("pID")
	input.PID = pID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectResolver.GetByProjectBonitaUserList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// ProjectListUser
// @Summary Project.4 將id轉成人名的專案列表
// @description 將id轉成人名的專案列表
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=project.Project_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/ProjectListUser [get]
func (p *presenter) ProjectListUser(ctx *gin.Context) {
	input := &project.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ProjectResolver.ProjectListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// ProduceSalesListUser
// @Summary Project.5 將id轉成人名的專案列表(產銷)
// @description 將id轉成人名的專案列表(產銷)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=project.Project_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/ProduceSalesListUser [get]
func (p *presenter) ProduceSalesListUser(ctx *gin.Context) {
	input := &project.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ProjectResolver.ProduceSalesListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// ProjectTemplateListUser
// @Summary Project.6 將id轉成人名的專案列表(專案範例)
// @description 將id轉成人名的專案列表(專案範例)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=project.Project_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/ProjectTemplateListUser [get]
func (p *presenter) ProjectTemplateListUser(ctx *gin.Context) {
	input := &project.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ProjectResolver.ProjectTemplateListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// ProjectAuthorizationListUser
// @Summary Project.7 將id轉成人名的專案列表(專案授權書)
// @description 將id轉成人名的專案列表(專案授權書)
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=project.Project_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/ProjectAuthorizationListUser [get]
func (p *presenter) ProjectAuthorizationListUser(ctx *gin.Context) {
	input := &project.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ProjectResolver.ProjectAuthorizationListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByProjectListUser
// @Summary Project.8 選取單一筆將id轉成人名的專案資料
// @description 選取單一筆將id轉成人名的專案資料
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string  true "專案ID專案ID"
// @success 200 object code.SuccessfulMessage{body=project.Project_Account} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/GetByProjectListUser/{pID} [get]
func (p *presenter) GetByProjectListUser(ctx *gin.Context) {
	pID := ctx.Param("pID")
	input := &project.Field{}
	input.PID = pID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectResolver.GetByProjectListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Project.2 取得單一專案
// @description 取得單一專案
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string true "專案ID"
// @success 200 object code.SuccessfulMessage{body=project.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/{pID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	pID := ctx.Param("pID")
	input := &project.Field{}
	input.PID = pID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Project.d 刪除單一專案
// @description 刪除單一專案
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string true "專案ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/{pID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	pID := ctx.Param("pID")
	input := &project.Updated{}
	input.PID = pID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// B2UpdatedCaseID
// @Summary Project.b2-s 起單
// @description B2-S起單
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string true "專案ID"
// @param * body project.Updated true "審核內容"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/B2UpdatedCaseID/{pID} [patch]
func (p *presenter) B2UpdatedCaseID(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	//client.Login("isabelle_wu")
	input := &model.CaseIDModelInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	B2 := bpm.GetProcessID(ctx,input.Account,"任務回報頁面")
	if B2 == "error"{
		return
	}
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx, B2, input)
	//fmt.Println(input.Tm)
	// client := gonita.New(input.Account)
	// marshal, err := json.Marshal(input)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }

	// body := client.CreateProcessCase(bpm.B2, string(marshal))
	// //fmt.Println(body)
	// output := &model.GetCaseIDOutput{}

	// err = json.Unmarshal(body, &output)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }
	// detailbody := client.GetCasePendingTaskDetail(strconv.Itoa(int(output.CaseID)))
	// detailoutput := &[]project.GetCasekDetailOutput{}
	// err2 := json.Unmarshal(detailbody, &detailoutput)
	// if err2 != nil {
	// 	log.Error(err2)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err2.Error()))

	// 	return
	// }
	if output != nil {
		pID := ctx.Param("pID")
		update := &project.Updated_Bonita{}
		update.PID = pID
		update.BonitaCaseID = output.CaseID

		// for _, value := range *detailoutput {
		// 	taskID, _ := strconv.ParseFloat(value.ID,32)
		// 	update.BonitaTaskID = float32(taskID)
		// 	update.Status = value.DisplayName
		// }
		//fmt.Println(client)

		codeMessage := p.ProjectResolver.Updated_Bonita(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// B2ReviewTask
// @Summary Project.b2-review 審核
// @description B2審核
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param taskID path string true "bonita任務ID"
// @param * body project.Updated true "審核內容"
// @success 200 object code.SuccessfulMessage{body=int} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/B2ReviewTask/{account}/{taskID} [patch]
func (p *presenter) B2ReviewTask(ctx *gin.Context) {

	account := ctx.Param("account")
	taskID := ctx.Param("taskID")

	input := &model.ReviewInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	body := bpm.ReviewTask(ctx, account, taskID, input)
	// client := gonita.New(account)
	// marshal, err := json.Marshal(input)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }
	// body := client.ExecuteTask(taskID, string(marshal))

	if body != 204 {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(body, "bonita review error"))

		return
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, body))

}


// Updated
// @Summary Project.u 更新單一專案
// @description 更新單一專案
// @Tags Project
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string true "專案ID"
// @param * body project.Updated true "更新專案"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/project/{pID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	pID := ctx.Param("pID")
	input := &project.Updated{}
	input.PID = pID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
