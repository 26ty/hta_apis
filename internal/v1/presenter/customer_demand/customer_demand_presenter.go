package customer_demand

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"
	model "eirc.app/internal/v1/structure"
	bpm "eirc.app/internal/pkg/bpm"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/customer_demand"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary CustomerDemand.c 新增客需單
// @description 新增客需單
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body customer_demand.Created true "新增客需單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &customer_demand.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary CustomerDemand.1 條件搜尋客需單
// @description 條件客需單
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=customer_demand.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &customer_demand.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CustomerDemandResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetA1BonitaCaseListStart
// @Summary CustomerDemand.a1-r 獲取重啟單的任務ID
// @description 獲取重啟單的任務ID
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaCaseID path string  true "Bonita單據ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListStart/{account}/{bonitaCaseID} [get]
func (p *presenter) GetA1BonitaCaseListStart(ctx *gin.Context) {
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
			if value.Name == "建立客需單"{
				ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, value.ID))
				return
			}
		}	
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, nil))
	
}

// GetA1BonitaCaseListTask
// @Summary CustomerDemand.a1-ex-s 獲取客需單新增任務的任務ID
// @description 獲取客需單新增任務的任務ID
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @param bonitaCaseID path string  true "Bonita單據ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListTask/{account}/{bonitaUserID}/{bonitaCaseID} [get]
func (p *presenter) GetA1BonitaCaseListTask(ctx *gin.Context) {
	input := &model.GetDetailListInput{}
	caseID := ctx.Param("caseID")
	input.CaseID = caseID
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		for _, value := range *output {
			if value.Name == "新增任務" && value.CaseID == input.CaseID{
				ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, value.ID))
				return
			}
		}	
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, nil))
}

// GetA1BonitaCaseListDepartment
// @Summary CustomerDemand.a1-1 獲取使用者可執行的單 (業務經理審核)
// @description 獲取使用者可執行的單 (業務經理審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListDepartment/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListDepartment(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoDepartment(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListDirector
// @Summary CustomerDemand.a1-2 獲取使用者可執行的單 (處長回覆簽核意見並指定專案經理)
// @description 獲取使用者可執行的單 (處長回覆簽核意見並指定專案經理)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListDirector/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListDirector(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoDirector(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListTop
// @Summary CustomerDemand.a1-3 獲取使用者可執行的單 (總經理PM人選確認並負責RD部門勾選)
// @description 獲取使用者可執行的單 (總經理PM人選確認並負責RD部門勾選)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListTop/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListTop(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoTop(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListDispatch
// @Summary CustomerDemand.a1-4.1 獲取使用者可執行的單 (會簽主管指派各部門人員(可能1人或多人))
// @description 獲取使用者可執行的單 (會簽主管指派各部門人員(可能1人或多人))
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListDispatch/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListDispatch(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoDispatch(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListEvaluation
// @Summary CustomerDemand.a1-4.1.1 獲取使用者可執行的單 (會簽人員送交評估報告)
// @description 獲取使用者可執行的單 (會簽人員送交評估報告)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListEvaluation/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListEvaluation(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoEvaluation(*output,userID)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListCountersign
// @Summary CustomerDemand.a1-4.1.2 獲取使用者可執行的單 (會簽主管審核)
// @description 獲取使用者可執行的單 (會簽主管審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListCountersign/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListCountersign(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoCountersign(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListPMEvaluation
// @Summary CustomerDemand.a1-4.1.3 獲取使用者可執行的單 (PM提交評估報告)
// @description 獲取使用者可執行的單 (PM提交評估報告)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review2} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListPMEvaluation/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListPMEvaluation(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoPMEvaluation(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListBusiness
// @Summary CustomerDemand.a1-4.1.4 獲取使用者可執行的單 (業務簽核)
// @description 獲取使用者可執行的單 (業務簽核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListBusiness/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListBusiness(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoBusiness(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListBusinessManager
// @Summary CustomerDemand.a1-4.1.5 獲取使用者可執行的單 (業務經理簽核)
// @description 獲取使用者可執行的單 (業務經理簽核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListBusinessManager/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListBusinessManager(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoBusinessManager(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListBusinessDirector
// @Summary CustomerDemand.a1-4.1.6 獲取使用者可執行的單 (業務處長簽核)
// @description 獲取使用者可執行的單 (業務處長簽核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListBusinessDirector/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListBusinessDirector(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoBusinessDirector(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListTaskFinish
// @Summary CustomerDemand.a1-ex-1 獲取使用者可執行的單 (任務完工送審)
// @description 獲取使用者可執行的單 (任務完工送審)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review_Task} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListTaskFinish/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListTaskFinish(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoTaskFinish(*output,userID)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListTaskFinishManager
// @Summary CustomerDemand.a1-ex-2 獲取使用者可執行的單 (任務直屬主管審核)
// @description 獲取使用者可執行的單 (任務直屬主管審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review_Task} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListTaskFinishManager/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListTaskFinishManager(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoTaskFinishManager(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListBusinessClose
// @Summary CustomerDemand.a1-5 獲取使用者可執行的單 (業務結案審核)
// @description 獲取使用者可執行的單 (業務結案審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListBusinessClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListBusinessClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoBusinessClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListDepartmentClose
// @Summary CustomerDemand.a1-6 獲取使用者可執行的單 (業務經理結案審核)
// @description 獲取使用者可執行的單 (業務經理結案審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListDepartmentClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListDepartmentClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoDepartmentClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListDirectorClose
// @Summary CustomerDemand.a1-7 獲取使用者可執行的單 (業務處長結案審核)
// @description 獲取使用者可執行的單 (業務處長結案審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review2} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListDirectorClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListDirectorClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoDirectorClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListTopClose
// @Summary CustomerDemand.a1-8 獲取使用者可執行的單 (總經理結案審核)
// @description 獲取使用者可執行的單 (總經理結案審核)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListTopClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListTopClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoTopClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListProductionClose
// @Summary CustomerDemand.a1-9 獲取使用者可執行的單 (製造部主管通知)
// @description 獲取使用者可執行的單 (製造部主管通知)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListProductionClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListProductionClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoProductionClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListCountersignClose
// @Summary CustomerDemand.a1-10 獲取使用者可執行的單 (回報製令完工)
// @description 獲取使用者可執行的單 (回報製令完工)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListCountersignClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListCountersignClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoCountersignClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetA1BonitaCaseListPMClose
// @Summary CustomerDemand.a1-11 獲取使用者可執行的單 (製令結案)
// @description 獲取使用者可執行的單 (製令結案)
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string  true "使用者帳號"
// @param bonitaUserID path string  true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.Customer_Review} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetA1BonitaCaseListPMClose/{account}/{bonitaUserID} [get]
func (p *presenter) GetA1BonitaCaseListPMClose(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.CustomerDemandResolver.GetByCaseIDtoPMClose(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// CustomerDemandListUser
// @Summary CustomerDemand.3 取得所有客需單的詳細資料
// @description 取得所有客需單的詳細資料，將ID轉成人名
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=customer_demand.Customer_Demand_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/CustomerDemandListUser [get]
func (p *presenter) CustomerDemandListUser(ctx *gin.Context) {
	input := &customer_demand.Field{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.CustomerDemandListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByCuIDCustomerDemandListUser
// @Summary CustomerDemand.4 取得單一客需單的詳細資料
// @description 取得單一客需單的詳細資料，將ID轉成人名
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cdID path string true "客需單ID"
// @success 200 object code.SuccessfulMessage{body=customer_demand.Customer_Demand_Account} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetByCuIDCustomerDemandListUser/{cdID} [get]
func (p *presenter) GetByCuIDCustomerDemandListUser(ctx *gin.Context) {
	cdId := ctx.Param("CdId")
	input := &customer_demand.Field{}
	input.CdId = cdId
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.GetByCuIDCustomerDemandListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIDListCR
// @Summary CustomerDemand.5 取得單一使用者的客需單任務
// @description 取得單一使用者的客需單任務
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=customer_demand.CRs} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetByUserIDListCR/{userID} [get]
func (p *presenter) GetByUserIDListCR(ctx *gin.Context) {
	userId := ctx.Param("UserId")
	input := &customer_demand.Users{}
	input.UserID = userId
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.GetByUserIDListCR(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIDListHCR
// @Summary CustomerDemand.6 取得單一使用者的客需單會簽
// @description 取得單一使用者的客需單會簽
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]customer_demand.H_CR} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/GetByUserIDListHCR/{userID} [get]
func (p *presenter) GetByUserIDListHCR(ctx *gin.Context) {
	userId := ctx.Param("UserId")
	input := &customer_demand.Users{}
	input.UserID = userId
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.GetByUserIDListHCR(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary CustomerDemand.2 取得單一客需單
// @description 取得單一客需單
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cdID path string true "客需單ID"
// @success 200 object code.SuccessfulMessage{body=customer_demand.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/{cdID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	cdId := ctx.Param("CdId")
	input := &customer_demand.Field{}
	input.CdId = cdId
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary CustomerDemand.d 刪除單一客需單
// @description 刪除單一客需單
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cdID path string true "客需單ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/{cdID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	cdId := ctx.Param("CdId")
	input := &customer_demand.Updated{}
	input.CdId = cdId
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// A1UpdatedCaseID
// @Summary CustomerDemand.a1-s 客需單起單
// @description 客需單起單，將bonita的單據ID導入對應的客需單
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cdID path string true "客需單ID"
// @param * body model.CaseIDModelInput true "客需單起單"
// @success 204 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/A1UpdatedCaseID/{cdID} [patch]
func (p *presenter) A1UpdatedCaseID(ctx *gin.Context) {

	input := &model.CaseIDModelInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	A1 := bpm.GetProcessID(ctx,input.Account,"客需單管理")
	if A1 == "error"{
		return
	}
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx,A1,input)

	if output != nil{
		cdId := ctx.Param("CdId")
		update := &customer_demand.Updated_Bonita{}
		update.CdId = cdId
		update.BonitaCaseID = output.CaseID
		codeMessage := p.CustomerDemandResolver.Updated_Bonita(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}
}

// A1ReviewTask
// @Summary CustomerDemand.a1-review 更新客需單任務審核狀態
// @description 更新客需單任務審核狀態
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "使用者帳號"
// @param taskID path string true "任務ID"
// @param * body model.ReviewInput true "更新客需單任務審核狀態"
// @success 204 object code.SuccessfulMessage{body=int} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/A1ReviewTask/{account}/{taskID} [patch]
func (p *presenter) A1ReviewTask(ctx *gin.Context) {

	account := ctx.Param("account")
	taskID := ctx.Param("taskID")

	input := &model.ReviewInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	body := bpm.ReviewTask(ctx,account,taskID,input)
	// client := gonita.New(account)
	// marshal, err := json.Marshal(input)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

	// 	return
	// }
	// body := client.ExecuteTask(taskID, string(marshal))
	
	if body != 204{
		ctx.JSON(http.StatusOK, code.GetCodeMessage(body, "bonita review error"))

		return
	}

	ctx.JSON(http.StatusOK,code.GetCodeMessage(code.Successful, body))

}

// Updated
// @Summary CustomerDemand.u 更新單一客需單
// @description 更新單一客需單
// @Tags CustomerDemand
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cdID path string true "客需單ID"
// @param * body accounts.Updated true "更新客需單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/customer_demand/{cdID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	cdId := ctx.Param("CdId")
	input := &customer_demand.Updated{}
	input.CdId = cdId
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CustomerDemandResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
