package jobtitle

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	//bpm "eirc.app/internal/pkg/bpm"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/jobtitle"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Jobtitle.c 新增職稱
// @description 新增職稱
// @Tags Jobtitle
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body jobtitle.Created true "新增職稱"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/jobtitle [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &jobtitle.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JobtitleResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)

}

// List
// @Summary Jobtitle.1 條件搜尋職稱
// @description 條件職稱
// @Tags Jobtitle
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param name query string false "職位名稱"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=jobtitle.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/jobtitle [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &jobtitle.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.JobtitleResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Jobtitle.2 取得單一職稱
// @description 取得單一職稱
// @Tags Jobtitle
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param JID path string true "職稱ID"
// @success 200 object code.SuccessfulMessage{body=jobtitle.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/jobtitle/{JID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	jID := ctx.Param("JID")
	input := &jobtitle.Field{}
	input.JID = jID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JobtitleResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Jobtitle.d 刪除單一職稱
// @description 刪除單一職稱
// @Tags Jobtitle
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param JID path string true "職稱ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/jobtitle/{JID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jID := ctx.Param("JID")
	input := &jobtitle.Updated{}
	input.JID = jID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JobtitleResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Jobtitle.u 更新單一職稱
// @description 更新單一職稱
// @Tags Jobtitle
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param JID path string true "職稱ID"
// @param * body jobtitle.Updated true "更新職稱"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/jobtitle/{JID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jID := ctx.Param("JID")
	input := &jobtitle.Updated{}
	input.JID = jID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	codeMessage := p.JobtitleResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
