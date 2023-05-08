package manufacture_type

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/manufacture_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary ManufactureType.c 新增製令類型
// @description 新增製令類型
// @Tags ManufactureType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body manufacture_type.Created true "新增製令類型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_type [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &manufacture_type.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureTypeResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary ManufactureType.1 條件搜尋製令類型
// @description 條件製令類型
// @Tags ManufactureType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=manufacture_type.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_type [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &manufacture_type.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ManufactureTypeResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary ManufactureType.2 取得單一製令類型
// @description 取得單一製令類型
// @Tags ManufactureType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MtID path string true "製令類型ID"
// @success 200 object code.SuccessfulMessage{body=manufacture_type.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_type/{MtID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	mtID := ctx.Param("MtID")
	input := &manufacture_type.Field{}
	input.MtID = mtID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureTypeResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary ManufactureType.d 刪除單一製令類型
// @description 刪除單一製令類型
// @Tags ManufactureType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MtID path string true "製令類型ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_type/{MtID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mtID := ctx.Param("MtID")
	input := &manufacture_type.Updated{}
	input.MtID = mtID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureTypeResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary ManufactureType.u 更新單一製令類型
// @description 更新單一製令類型
// @Tags ManufactureType
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MtID path string true "製令類型ID"
// @param * body manufacture_type.Updated true "更新製令類型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_type/{MtID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mtID := ctx.Param("MtID")
	input := &manufacture_type.Updated{}
	input.MtID = mtID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureTypeResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
