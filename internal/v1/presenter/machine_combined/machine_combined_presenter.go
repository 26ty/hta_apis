package machine_combined

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/machine_combined"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary MachineCombined.c 新增組合機器
// @description 新增組合機器
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body machine_combined.Created true "新增組合機器"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &machine_combined.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MachineCombinedResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary MachineCombined.1 條件搜尋組合機器
// @description 條件組合機器
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=machine_combined.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &machine_combined.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MachineCombinedResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// MachineCombinedListLast
// @Summary MachineCombined.4 列出last_mc的mc_code
// @description 列出last_mc的mc_code
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=machine_combined.Machine_Combined_Lasts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined/MachineCombinedListLast [get]
func (p *presenter) MachineCombinedListLast(ctx *gin.Context) {
	input := &machine_combined.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MachineCombinedResolver.MachineCombinedListLast(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByPIDMachineCombinedListLast
// @Summary MachineCombined.3 PID篩選顯示組合機器詳細資料
// @description PID篩選顯示組合機器詳細資料
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PID path string true "專案ID"
// @success 200 object code.SuccessfulMessage{body=machine_combined.Machine_Combined_Lasts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined/GetByPIDMachineCombinedListLast/{PID} [get]
func (p *presenter) GetByPIDMachineCombinedListLast(ctx *gin.Context) {
	pID := ctx.Param("PID")
	input := &machine_combined.Fields{}
	input.ProjectID = pID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MachineCombinedResolver.GetByPIDMachineCombinedListLast(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary MachineCombined.2 取得單一組合機器
// @description 取得單一組合機器
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param McID path string true "組合機器ID"
// @success 200 object code.SuccessfulMessage{body=machine_combined.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined/{McID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	mcID := ctx.Param("McID")
	input := &machine_combined.Field{}
	input.McID = mcID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MachineCombinedResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary MachineCombined.d 刪除單一組合機器
// @description 刪除單一組合機器
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param McID path string true "組合機器ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined/{McID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mcID := ctx.Param("McID")
	input := &machine_combined.Updated{}
	input.McID = mcID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MachineCombinedResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary MachineCombined.u 更新單一組合機器
// @description 更新單一組合機器
// @Tags MachineCombined
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param McID path string true "組合機器ID"
// @param * body machine_combined.Updated true "更新組合機器"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/machine_combined/{McID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mcID := ctx.Param("McID")
	input := &machine_combined.Updated{}
	input.McID = mcID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MachineCombinedResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
