package countersign

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/countersign"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary Countersign.c 新增會簽
// @description 新增會簽
// @Tags Countersign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body countersign.Created true "新增會簽"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &countersign.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Countersign.1 條件搜尋會簽
// @description 條件會簽
// @Tags Countersign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=countersign.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &countersign.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CountersignResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Countersign.2 取得單一會簽
// @description 取得單一會簽
// @Tags Countersign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param csID path string true "會簽ID"
// @success 200 object code.SuccessfulMessage{body=countersign.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign/{csID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	csID := ctx.Param("CsID")
	input := &countersign.Field{}
	input.CsID = csID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Countersign.d 刪除單一會簽
// @description 刪除單一會簽
// @Tags Countersign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param csID path string true "會簽ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign/{csID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	csID := ctx.Param("CsID")
	input := &countersign.Updated{}
	input.CsID = csID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Countersign.u 更新單一會簽
// @description 更新單一會簽
// @Tags Countersign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param csID path string true "會簽ID"
// @param * body countersign.Updated true "更新會簽"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign/{csID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	csID := ctx.Param("CsID")
	input := &countersign.Updated{}
	input.CsID = csID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
