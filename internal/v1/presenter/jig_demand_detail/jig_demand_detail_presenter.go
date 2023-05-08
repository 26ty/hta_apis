package jig_demand_detail

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/jig_demand_details"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增治具需求單明細
// @description 新增治具需求單明細
// @Tags JigDemandDetail
// @version 1.0
// @Accept json
// @produce json
// @param * body jig_demand_details.Created true "新增治具需求單明細"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemandDetail [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &jig_demand_details.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandDetailResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 取得全部治具需求單明細
// @description 取得全部治具需求單明細
// @Tags JigDemandDetail
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=jig_demand_details.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemandDetail [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &jig_demand_details.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.JigDemandDetailResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一治具需求單明細
// @description 取得單一治具需求單明細
// @Tags JigDemandDetail
// @version 1.0
// @Accept json
// @produce json
// @param jdID path string true "治具需求單明細ID"
// @success 200 object code.SuccessfulMessage{body=jig_demand_details.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemandDetail/{jdID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	jdID := ctx.Param("jdID") //跟router對應
	input := &jig_demand_details.Field{}
	input.JdID = jdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandDetailResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一治具需求單明細
// @description 刪除單一治具需求單明細
// @Tags JigDemandDetail
// @version 1.0
// @Accept json
// @produce json
// @param jdID path string true "治具需求單明細ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemandDetail/{jdID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jdID := ctx.Param("jdID")
	input := &jig_demand_details.Updated{}
	input.JdID = jdID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandDetailResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一治具需求單明細
// @description 更新單一治具需求單明細
// @Tags JigDemandDetail
// @version 1.0
// @Accept json
// @produce json
// @param jdID path string true "治具需求單明細ID"
// @param * body jig_demand_details.Updated true "更新治具需求單明細"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemandDetail/{jdID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jdID := ctx.Param("jdID")
	input := &jig_demand_details.Updated{}
	input.JdID = jdID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandDetailResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// UpdatedByJigID
// @Summary 透過JigID更新單一治具需求單明細
// @description 透過JigID更新單一治具需求單明細
// @Tags JigDemandDetail
// @version 1.0
// @Accept json
// @produce json
// @param jigID path string true "治具需求單ID"
// @param * body jig_demand_details.Updated true "更新治具需求單明細"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /hta_crm/v1.0/JigDemandDetail/{jdID} [patch]
func (p *presenter) UpdatedByJigID(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	jigID := ctx.Param("jigID")
	input := &jig_demand_details.Updated{}
	input.JigID = jigID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.JigDemandDetailResolver.UpdatedByJigID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
