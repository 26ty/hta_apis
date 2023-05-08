package quotation

import (
	"net/http"

	"eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/quotations"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增報價申請單
// @description 新增報價申請單
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param * body quotations.Created true "新增報價申請單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &quotations.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.QuotationResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部報價申請單
// @description 取得全部報價申請單
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=quotations.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &quotations.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.QuotationResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// QuotationDetailListUser
// @Summary 取得全部報價申請單及明細
// @description 取得全部報價申請單及明細
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=quotations.AllQuotationDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation/GetAllQuotationDetail [get]
func (p *presenter) QuotationDetailListUser(ctx *gin.Context) {
	input := &quotations.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.QuotationResolver.QuotationDetailListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByQIDQuotationDetailListUser
// @Summary 取得單一報價申請單及明細
// @description 取得單一報價申請單及明細
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param qID path string true "報價申請單ID"
// @success 200 object code.SuccessfulMessage{body=quotations.QuotationDetail} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation/GetByQIDDetail/{qID} [get]
func (p *presenter) GetByQIDQuotationDetailListUser(ctx *gin.Context) {
	qID := ctx.Param("qID")
	input := &quotations.Field{}
	input.QID = qID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.QuotationResolver.GetByQIDQuotationDetailListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一報價申請單
// @description 取得單一報價申請單
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param qID path string true "報價申請單ID"
// @success 200 object code.SuccessfulMessage{body=quotations.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation/{qID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	qID := ctx.Param("qID") //跟router對應
	input := &quotations.Field{}
	input.QID = qID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.QuotationResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一報價申請單
// @description 刪除單一報價申請單
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param qID path string true "報價申請單ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation/{qID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	qID := ctx.Param("qID")
	input := &quotations.Updated{}
	input.QID = qID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.QuotationResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一報價申請單
// @description 更新單一報價申請單
// @Tags Quotation
// @version 1.0
// @Accept json
// @produce json
// @param qID path string true "報價申請單ID"
// @param * body quotations.Updated true "更新報價申請單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Quotation/{qID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	qID := ctx.Param("qID")
	input := &quotations.Updated{}
	input.QID = qID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.QuotationResolver.Updated(input)
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
	ProccessID := bpm.GetProcessID(ctx, input.Account, "報價申請單管理")
	if ProccessID == "error" {
		return
	}
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx, ProccessID, input)

	if output != nil {
		qID := ctx.Param("qID")
		update := &quotations.Updated_Bonita{}
		update.QID = qID
		update.BonitaCaseID = output.CaseID
		codeMessage := p.QuotationResolver.Updated_Bonita(update)
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
