package gift_application_detail

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/gift_application_details"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增部品零件申請單明細
// @description 新增部品零件申請單明細
// @Tags GiftApplicationDetail
// @version 1.0
// @Accept json
// @produce json
// @param * body gift_application_details.Created true "新增部品零件申請單明細"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplicationDetail [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &gift_application_details.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationDetailResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部部品零件申請單明細
// @description 取得全部部品零件申請單明細
// @Tags GiftApplicationDetail
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=gift_application_details.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplicationDetail [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &gift_application_details.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.GiftApplicationDetailResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一部品零件申請單明細
// @description 取得單一部品零件申請單明細
// @Tags GiftApplicationDetail
// @version 1.0
// @Accept json
// @produce json
// @param gdID path string true "部品零件申請單明細ID"
// @success 200 object code.SuccessfulMessage{body=gift_application_details.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplicationDetail/{gdID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	gdID := ctx.Param("gdID") //跟router對應
	input := &gift_application_details.Field{}
	input.GdID = gdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationDetailResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一部品零件申請單明細
// @description 刪除單一部品零件申請單明細
// @Tags GiftApplicationDetail
// @version 1.0
// @Accept json
// @produce json
// @param gdID path string true "部品零件申請單明細ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplicationDetail/{gdID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	gdID := ctx.Param("gdID")
	input := &gift_application_details.Updated{}
	input.GdID = gdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationDetailResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一部品零件申請單明細
// @description 更新單一部品零件申請單明細
// @Tags GiftApplicationDetail
// @version 1.0
// @Accept json
// @produce json
// @param gdID path string true "部品零件申請單明細ID"
// @param * body gift_application_details.Updated true "更新部品零件申請單明細"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/GiftApplicationDetail/{gdID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	gdID := ctx.Param("gdID")
	input := &gift_application_details.Updated{}
	input.GdID = gdID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GiftApplicationDetailResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
