package plug_in

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/plug_in"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary PlugIn.c 新增外掛部品
// @description 新增外掛部品
// @Tags PlugIn
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body plug_in.Created true "新增外掛部品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/plug_in [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &plug_in.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PlugInResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary PlugIn.1 條件搜尋外掛部品
// @description 條件外掛部品
// @Tags PlugIn
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=plug_in.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/plug_in [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &plug_in.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.PlugInResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByPIDList
// @Summary PlugIn.3 PID篩選顯示plug_in列表
// @description PID篩選顯示plug_in列表
// @Tags PlugIn
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param ProjectID path string true "專案ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=plug_in.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/plug_in/GetByPIDList/{ProjectID} [get]
func (p *presenter) GetByPIDList(ctx *gin.Context) {
	projectID := ctx.Param("ProjectID")
	input := &plug_in.Fields{}
	input.ProjectID = projectID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.PlugInResolver.GetByPIDList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary PlugIn.2 取得單一外掛部品
// @description 取得單一外掛部品
// @Tags PlugIn
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PiID path string true "外掛部品ID"
// @success 200 object code.SuccessfulMessage{body=plug_in.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/plug_in/{PiID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	piID := ctx.Param("PiID")
	input := &plug_in.Field{}
	input.PiID = piID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PlugInResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary PlugIn.d 刪除單一外掛部品
// @description 刪除單一外掛部品
// @Tags PlugIn
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PiID path string true "外掛部品ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/plug_in/{PiID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	piID := ctx.Param("PiID")
	input := &plug_in.Updated{}
	input.PiID = piID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PlugInResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary PlugIn.u 更新單一外掛部品
// @description 更新單一外掛部品
// @Tags PlugIn
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PiID path string true "外掛部品ID"
// @param * body plug_in.Updated true "更新外掛部品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/plug_in/{PiID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	piID := ctx.Param("PiID")
	input := &plug_in.Updated{}
	input.PiID = piID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PlugInResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
