package factory_manufacturing

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/factory_manufacturings"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增廠別生產機型
// @description 新增廠別生產機型
// @Tags FactoryManufacturing
// @version 1.0
// @Accept json
// @produce json
// @param * body factory_manufacturings.Created_List true "新增廠別生產機型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryManufacturing [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &factory_manufacturings.Created_List{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryManufacturingResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部廠別生產機型
// @description 取得全部廠別生產機型
// @Tags FactoryManufacturing
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=factory_manufacturings.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryManufacturing [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &factory_manufacturings.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.FactoryManufacturingResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一廠別生產機型
// @description 取得單一廠別生產機型
// @Tags FactoryManufacturing
// @version 1.0
// @Accept json
// @produce json
// @param fmID path string true "廠別生產機型ID"
// @success 200 object code.SuccessfulMessage{body=factory_manufacturings.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryManufacturing/{fmID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	fmID := ctx.Param("fmID") //跟router對應
	input := &factory_manufacturings.Field{}
	input.FmID = fmID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryManufacturingResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一廠別生產機型
// @description 刪除單一廠別生產機型
// @Tags FactoryManufacturing
// @version 1.0
// @Accept json
// @produce json
// @param fmID path string true "廠別生產機型ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryManufacturing/{fmID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	fmID := ctx.Param("fmID")
	input := &factory_manufacturings.Updated{}
	input.FmID = fmID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryManufacturingResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一廠別生產機型
// @description 更新單一廠別生產機型
// @Tags FactoryManufacturing
// @version 1.0
// @Accept json
// @produce json
// @param fmID path string true "廠別生產機型ID"
// @param * body factory_manufacturings.Updated true "更新廠別生產機型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryManufacturing/{fmID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	fmID := ctx.Param("fmID")
	input := &factory_manufacturings.Updated{}
	input.FmID = fmID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryManufacturingResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
