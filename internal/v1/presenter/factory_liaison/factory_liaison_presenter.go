package factory_liaison

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/factory_liaisons"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增廠別聯絡人
// @description 新增廠別聯絡人
// @Tags FactoryLiaison
// @version 1.0
// @Accept json
// @produce json
// @param * body factory_liaisons.Created_List true "新增廠別聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryLiaison [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &factory_liaisons.Created_List{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryLiaisonResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部廠別聯絡人
// @description 取得全廠別聯絡人
// @Tags FactoryLiaison
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=factory_liaisons.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryLiaison [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &factory_liaisons.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.FactoryLiaisonResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一廠別聯絡人
// @description 取得單一廠別聯絡人
// @Tags FactoryLiaison
// @version 1.0
// @Accept json
// @produce json
// @param flID path string true "廠別聯絡人ID"
// @success 200 object code.SuccessfulMessage{body=factory_liaisons.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryLiaison{flID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	flID := ctx.Param("flID") //跟router對應
	input := &factory_liaisons.Field{}
	input.FlID = flID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryLiaisonResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一廠別聯絡人
// @description 刪除單一廠別聯絡人
// @Tags FactoryLiaison
// @version 1.0
// @Accept json
// @produce json
// @param flID path string true "廠別聯絡人ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryLiaison{flID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	flID := ctx.Param("flID")
	input := &factory_liaisons.Updated{}
	input.FlID = flID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryLiaisonResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一廠別聯絡人
// @description 更新單一廠別聯絡人
// @Tags FactoryLiaison
// @version 1.0
// @Accept json
// @produce json
// @param flID path string true "廠別聯絡人ID"
// @param * body factory_liaisons.Updated true "更新廠別聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/FactoryLiaison{flID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	flID := ctx.Param("flID")
	input := &factory_liaisons.Updated{}
	input.FlID = flID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FactoryLiaisonResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
