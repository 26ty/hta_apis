package jig_demand

import (
	"net/http"

	bpm "eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/jig_demands"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增治具需求單
// @description 新增治具需求單
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param * body jig_demands.Created true "新增治具需求單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &jig_demands.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部治具需求單
// @description 取得全部治具需求單
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=jig_demands.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &jig_demands.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.JigDemandResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// SearchJigDemand
// @Summary 取得治具需求單搜尋列表
// @description 取得治具需求單搜尋列表
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=jig_demands.SearchJigDemand} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/SearchJigDemand [get]
func (p *presenter) SearchJigDemand(ctx *gin.Context) {
	input := &jig_demands.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.JigDemandResolver.SearchJigDemand(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// JigDetailListUser
// @Summary 取得全部治具需求單及明細
// @description 取得全部治具需求單及明細
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=jig_demands.AllJigDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/GetAllJigDetail [get]
func (p *presenter) JigDetailListUser(ctx *gin.Context) {
	input := &jig_demands.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.JigDemandResolver.JigDetailListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByJIDJigDetailListUser
// @Summary 取得單一治具需求單及明細
// @description 取得單一治具需求單及明細
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param jID path string true "治具需求單ID"
// @success 200 object code.SuccessfulMessage{body=jig_demands.JigDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/GetByJIDDetail/{jID} [get]
func (p *presenter) GetByJIDJigDetailListUser(ctx *gin.Context) {
	jID := ctx.Param("jID")
	input := &jig_demands.Field{}
	input.JID = jID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandResolver.GetByJIDJigDetailListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIDListJD
// @Summary 取得單一使用者的治具需求單任務
// @description 取得單一使用者的治具需求單任務
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=jig_demands.JDs} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/GetByUserIDListJD/{userID} [get]
func (p *presenter) GetByUserIDListJD(ctx *gin.Context) {
	userID := ctx.Param("userID")
	input := &jig_demands.Users{}
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandResolver.GetByUserIDListJD(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一治具需求單
// @description 取得單一治具需求單
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param jID path string true "治具需求單ID"
// @success 200 object code.SuccessfulMessage{body=jig_demands.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/{jID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	jID := ctx.Param("jID") //跟router對應
	input := &jig_demands.Field{}
	input.JID = jID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一治具需求單
// @description 刪除單一治具需求單
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param jID path string true "治具需求單ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/{jID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jID := ctx.Param("jID")
	input := &jig_demands.Updated{}
	input.JID = jID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一治具需求單
// @description 更新單一治具需求單
// @Tags JigDemand
// @version 1.0
// @Accept json
// @produce json
// @param jID path string true "治具需求單ID"
// @param * body jig_demands.Updated true "更新治具需求單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemand/{jID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jID := ctx.Param("jID")
	input := &jig_demands.Updated{}
	input.JID = jID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// 啟動單據
func (p *presenter) UpdatedCaseID(ctx *gin.Context) {

	input := &model.CaseIDModelInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	ProccessID := bpm.GetProcessID(ctx, input.Account, "治具需求單")
	if ProccessID == "error" {
		return
	}
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx, ProccessID, input)

	if output != nil {
		jID := ctx.Param("jID")
		update := &jig_demands.Updated_Bonita{}
		update.JID = jID
		update.BonitaCaseID = output.CaseID
		codeMessage := p.JigDemandResolver.Updated_Bonita(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}
}

// 審核任務
func (p *presenter) ReviewTask(ctx *gin.Context) {

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
