package sales_call_record

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/sales_call_records"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增顧客拜訪紀錄
// @description 新增顧客拜訪紀錄
// @Tags SalesCallRecord
// @version 1.0
// @Accept json
// @produce json
// @param * body sales_call_records.Created true "新增顧客拜訪紀錄"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/SalesCallRecord [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &sales_call_records.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.SalesCallRecordResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部顧客拜訪紀錄
// @description 取得全部顧客拜訪紀錄
// @Tags SalesCallRecord
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=sales_call_records.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/SalesCallRecord [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &sales_call_records.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.SalesCallRecordResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一顧客拜訪紀錄
// @description 取得單一顧客拜訪紀錄
// @Tags SalesCallRecord
// @version 1.0
// @Accept json
// @produce json
// @param sID path string true "顧客拜訪紀錄ID"
// @success 200 object code.SuccessfulMessage{body=sales_call_records.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/SalesCallRecord/{sID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	sID := ctx.Param("sID") //跟router對應
	input := &sales_call_records.Field{}
	input.SID = sID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.SalesCallRecordResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一顧客拜訪紀錄
// @description 刪除單一顧客拜訪紀錄
// @Tags SalesCallRecord
// @version 1.0
// @Accept json
// @produce json
// @param sID path string true "顧客拜訪紀錄ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/SalesCallRecord/{sID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	sID := ctx.Param("sID")
	input := &sales_call_records.Updated{}
	input.SID = sID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.SalesCallRecordResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一顧客拜訪紀錄
// @description 更新單一顧客拜訪紀錄
// @Tags SalesCallRecord
// @version 1.0
// @Accept json
// @produce json
// @param sID path string true "顧客拜訪紀錄ID"
// @param * body sales_call_records.Updated true "更新顧客拜訪紀錄"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/SalesCallRecord/{sID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	sID := ctx.Param("sID")
	input := &sales_call_records.Updated{}
	input.SID = sID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.SalesCallRecordResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
