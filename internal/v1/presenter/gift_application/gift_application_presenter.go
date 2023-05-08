package gift_application

import (
	"net/http"

	"eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/gift_applications"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增部品零件贈送申請單
// @description 新增部品零件贈送申請單
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param * body gift_applications.Created true "新增部品零件贈送申請單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &gift_applications.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部部品零件贈送申請單
// @description 取得全部部品零件贈送申請單
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=gift_applications.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &gift_applications.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.GiftApplicationResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GiftDetailListUser
// @Summary 取得全部部品零件贈送申請單及明細
// @description 取得全部部品零件贈送申請單及明細
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=gift_applications.AllGiftDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetAllGiftApplication [get]
func (p *presenter) GiftDetailListUser(ctx *gin.Context) {
	input := &gift_applications.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.GiftApplicationResolver.GiftDetailListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByGIDGiftDetailListUser
// @Summary 取得單一部品零件贈送申請單及明細
// @description 取得單一部品零件贈送申請單及明細
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param fID path string true "部品零件贈送申請單ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.GiftDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetByGIDDetail/{gID} [get]
func (p *presenter) GetByGIDGiftDetailListUser(ctx *gin.Context) {
	gID := ctx.Param("gID")
	input := &gift_applications.Field{}
	input.GID = gID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationResolver.GetByGIDGiftDetailListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetDBonitaTaskId
// @Summary Bonita D-T獲取單據待執行的任務ID
// @description Bonita D-T獲取單據待執行的任務ID
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param caseID path string true "Bonita 單據ID"
// @success 200 object code.SuccessfulMessage{body=model.GetCaseListOutput} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetDBonitaTaskId/{account}/{caseID} [get]
func (p *presenter) GetDBonitaTaskId(ctx *gin.Context) {
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
	output = bpm.GetCasePendingTaskDetail(ctx, input.Account, input.CaseID)

	if output != nil {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, output))
		return
	}

}

// GetDBonitaCaseListRestart
// @Summary Bonita D-R獲取使用者可執行的單 (部品零件贈送申請單重新送審)
// @description Bonita D-R獲取使用者可執行的單 (部品零件贈送申請單重新送審)
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param caseID path string true "Bonita 單據ID"
// @success 200 object code.SuccessfulMessage{body=model.GetCaseListOutput.ID} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetDBonitaCaseListRestart/{account}/{caseID} [get]
func (p *presenter) GetDBonitaCaseListRestart(ctx *gin.Context) {
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
	output = bpm.GetCasePendingTaskDetail(ctx, input.Account, input.CaseID)

	if output != nil {
		for _, value := range *output {
			if value.Name == "部品零件贈送申請單重新送審" {
				ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, value))
				return
			}
		}
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, nil))

}

// GetDBonitaCaseListDepartment
// @Summary Bonita D-1獲取使用者可執行的單 (單位主管審核)
// @description Bonita D-1獲取使用者可執行的單 (單位主管審核)
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param userID path string true "Bonita 用戶ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetDBonitaCaseListDepartment/{account}/{userID} [get]
func (p *presenter) GetDBonitaCaseListDepartment(ctx *gin.Context) {
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

	if output != nil {
		codeMessage := p.GiftApplicationResolver.GetByCaseIDtoDepartment(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetDBonitaCaseListViceTop
// @Summary Bonita D-2獲取使用者可執行的單 (副總審核)
// @description Bonita D-2獲取使用者可執行的單 (副總審核)
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param userID path string true "Bonita 用戶ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetDBonitaCaseListViceTop/{account}/{userID} [get]
func (p *presenter) GetDBonitaCaseListViceTop(ctx *gin.Context) {
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

	if output != nil {
		codeMessage := p.GiftApplicationResolver.GetByCaseIDtoViceTop(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetDBonitaCaseListTop
// @Summary Bonita D-3獲取使用者可執行的單 (總經理審核)
// @description Bonita D-3獲取使用者可執行的單 (總經理審核)
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param userID path string true "Bonita 用戶ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetDBonitaCaseListTop/{account}/{userID} [get]
func (p *presenter) GetDBonitaCaseListTop(ctx *gin.Context) {
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

	if output != nil {
		codeMessage := p.GiftApplicationResolver.GetByCaseIDtoTop(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetDBonitaCaseListAttm
// @Summary Bonita D-4獲取使用者可執行的單 (經辦結案確認)
// @description Bonita D-4獲取使用者可執行的單 (經辦結案確認)
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param userID path string true "Bonita 用戶ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/GetDBonitaCaseListAttm/{account}/{userID} [get]
func (p *presenter) GetDBonitaCaseListAttm(ctx *gin.Context) {
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

	if output != nil {
		codeMessage := p.GiftApplicationResolver.GetByCaseIDtoAttm(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetByID
// @Summary 取得單一部品零件贈送申請單
// @description 取得單一部品零件贈送申請單
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param gID path string true "部品零件贈送申請單ID"
// @success 200 object code.SuccessfulMessage{body=gift_applications.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/{gID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	gID := ctx.Param("gID") //跟router對應
	input := &gift_applications.Field{}
	input.GID = gID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一部品零件贈送申請單
// @description 刪除單一部品零件贈送申請單
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param gID path string true "部品零件贈送申請單ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/{gID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	gID := ctx.Param("gID")
	input := &gift_applications.Updated{}
	input.GID = gID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一部品零件贈送申請單
// @description 更新單一部品零件贈送申請單
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param gID path string true "部品零件贈送申請單ID"
// @param * body gift_applications.Updated true "更新部品零件贈送申請單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/{gID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	gID := ctx.Param("gID")
	input := &gift_applications.Updated{}
	input.GID = gID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// DUpdatedCaseID
// @Summary Bonita D-S啟單
// @description Bonita D-S啟單
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param gID path string true "部品零件贈送申請單ID"
// @param * body model.CaseIDModelInput true "啟單 input"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/UpdatedCaseID/{gID} [patch]
func (p *presenter) DUpdatedCaseID(ctx *gin.Context) {

	input := &model.CaseIDModelInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	D := bpm.GetProcessID(ctx, input.Account, "部品零件贈送申請單")
	if D == "error" {
		return
	}
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx, D, input)

	if output != nil {
		gID := ctx.Param("gID")
		update := &gift_applications.Updated_Bonita{}
		update.GID = gID
		update.BonitaCaseID = output.CaseID
		codeMessage := p.GiftApplicationResolver.Updated_Bonita(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}
}

// DReviewTask
// @Summary Bonita D-Review審核任務
// @description Bonita D-Review審核任務
// @Tags GiftApplication
// @version 1.0
// @Accept json
// @produce json
// @param account path string true "使用者帳號"
// @param taskID path string true "Bonita 任務ID"
// @param * body model.ReviewInput true "審核任務 input"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplication/DReviewTask/{account}/{taskID} [patch]
func (p *presenter) DReviewTask(ctx *gin.Context) {

	account := ctx.Param("account")
	taskID := ctx.Param("taskID")

	input := &model.ReviewInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	body := bpm.ReviewTask(ctx, account, taskID, input)

	if body != 204 {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(body, "bonita review error"))

		return
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, body))

}
