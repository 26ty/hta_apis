package factory

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/factorys"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增客戶廠別
// @description 新增客戶廠別
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param * body factorys.Created true "新增客戶廠別"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &factorys.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部客戶廠別
// @description 取得全部客戶廠別
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=factorys.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &factorys.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.FactoryResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// SearchFactory
// @Summary 取得客戶廠別搜尋列表
// @description 取得客戶廠別搜尋列表
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=factorys.SearchFactory} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory/SearchFactory [get]
func (p *presenter) SearchFactory(ctx *gin.Context) {
	input := &factorys.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.FactoryResolver.SearchFactory(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// FLMListUser
// @Summary 取得全部客戶廠別、聯絡人及生產列表
// @description 取得全部客戶廠別、聯絡人及生產列表
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=factorys.AllFLM} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory/GetAllFactory [get]
func (p *presenter) FLMListUser(ctx *gin.Context) {
	input := &factorys.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.FactoryResolver.FLMListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByFIDFLMListUser
// @Summary 取得單一客戶廠別、聯絡人及生產列表
// @description 取得單一客戶廠別、聯絡人及生產列表
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param fID path string true "客戶廠別ID"
// @success 200 object code.SuccessfulMessage{body=factorys.FLM} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory/GetByFIDFactory/{fID} [get]
func (p *presenter) GetByFIDFLMListUser(ctx *gin.Context) {
	fID := ctx.Param("fID")
	input := &factorys.Field{}
	input.FID = fID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryResolver.GetByFIDFLMListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一客戶廠別
// @description 取得單一客戶廠別
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param fID path string true "客戶廠別ID"
// @success 200 object code.SuccessfulMessage{body=factorys.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory/{fID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	fID := ctx.Param("fID") //跟router對應
	input := &factorys.Field{}
	input.FID = fID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一客戶廠別
// @description 刪除單一客戶廠別
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param fID path string true "客戶廠別ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory/{fID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	fID := ctx.Param("fID")
	input := &factorys.Updated{}
	input.FID = fID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一客戶廠別
// @description 更新單一客戶廠別
// @Tags Factory
// @version 1.0
// @Accept json
// @produce json
// @param fID path string true "客戶廠別ID"
// @param * body factorys.Updated true "更新客戶廠別"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/Factory/{fID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	fID := ctx.Param("fID")
	input := &factorys.Updated{}
	input.FID = fID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
