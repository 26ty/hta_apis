package gateway_data

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/gateway_data"
	model "eirc.app/internal/v1/structure"
	_ "eirc.app/internal/v1/structure/gg_data_demand"
	bpm "eirc.app/internal/pkg/bpm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary GatewayData.c 新增gateway
// @description 新增gateway
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body gateway_data.Created true "新增gateway"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &gateway_data.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GatewayDataResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary GatewayData.1 條件搜尋gateway
// @description 條件gateway
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=gateway_data.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &gateway_data.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.GatewayDataResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary GatewayData.2 取得單一gateway
// @description 取得單一gateway
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param gdID path string true "gateway_dataID"
// @success 200 object code.SuccessfulMessage{body=gateway_data.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data/{gdID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	gdID := ctx.Param("GdID")
	input := &gateway_data.Field{}
	input.GdID = gdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GatewayDataResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByClassificationTitle
// @Summary GatewayData.3 取得指定的gateway
// @description 用類別跟標題取得指定的gateway
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body gateway_data.Field true "取得gateway"
// @success 200 object code.SuccessfulMessage{body=gateway_data.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data/GetByClassificationTitle [post]
func (p *presenter) GetByClassificationTitle(ctx *gin.Context) {
	input := &gateway_data.Field{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GatewayDataResolver.GetByClassificationTitle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


// GetByDataDemand
// @Summary  GatewayData.4 獲取使用者可執行的單
// @description  獲取使用者可執行的單
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @param gdID path string true "gateway_dataID"
// @success 200 object code.SuccessfulMessage{body=[]gg_data_demand.Reviews} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data/GetByDataDemand/{account}/{userID}/{gdID} [get]
func (p *presenter) GetByDataDemand(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	gdID := ctx.Param("gdID")
	input.GdID = gdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)

	if output != nil {
		codeMessage := p.GatewayDataResolver.GetByDataDemand(*output,input.GdID,input.UserID)
		ctx.JSON(http.StatusOK, codeMessage)
	}
}

// Delete
// @Summary GatewayData.d 刪除單一gateway
// @description 刪除單一gateway
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param gdID path string true "gateway_dataID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data/{gdID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	gdID := ctx.Param("GdID")
	input := &gateway_data.Updated{}
	input.GdID = gdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GatewayDataResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary GatewayData.u 更新單一gateway
// @description 更新單一gateway
// @Tags GatewayData
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param gdID path string true "gatewayID"
// @param * body gateway_data.Updated true "更新gateway"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/gateway_data/{gdID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	gdID := ctx.Param("GdID")
	input := &gateway_data.Updated{}
	input.GdID = gdID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.GatewayDataResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
