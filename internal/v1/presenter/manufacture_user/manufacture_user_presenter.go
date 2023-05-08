package manufacture_user

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/manufacture_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary ManufactureUser.c 新增製令副本參與人
// @description 新增製令副本參與人
// @Tags ManufactureUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body manufacture_user.Created true "新增製令副本參與人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_user [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &manufacture_user.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureUserResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary ManufactureUser.1 條件搜尋製令副本參與人
// @description 條件製令副本參與人
// @Tags ManufactureUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=manufacture_user.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_user [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &manufacture_user.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ManufactureUserResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary ManufactureUser.2 取得單一製令副本參與人
// @description 取得單一製令副本參與人
// @Tags ManufactureUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MuID path string true "製令副本參與人ID"
// @success 200 object code.SuccessfulMessage{body=manufacture_user.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_user/{MuID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	muID := ctx.Param("MuID")
	input := &manufacture_user.Field{}
	input.MuID = muID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureUserResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByManufactureID
// @Summary ManufactureUser.3 ManufactureID篩選顯示製令副本參與人列表
// @description ManufactureID篩選顯示製令副本參與人列表
// @Tags ManufactureUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param ManufactureID path string true "製令ID"
// @success 200 object code.SuccessfulMessage{body=manufacture_user.ManufactureAccounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_user/GetByManufactureID/{ManufactureID} [get]
func (p *presenter) GetByManufactureID(ctx *gin.Context) {
	manufactureID := ctx.Param("ManufactureID")
	input := &manufacture_user.Field{}
	input.ManufactureID = manufactureID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureUserResolver.GetByManufactureID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary ManufactureUser.d 刪除單一製令副本參與人
// @description 刪除單一製令副本參與人
// @Tags ManufactureUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MuID path string true "製令副本參與人ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_user/{MuID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	muID := ctx.Param("MuID")
	input := &manufacture_user.Updated{}
	input.MuID = muID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureUserResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary ManufactureUser.u 更新單一製令副本參與人
// @description 更新單一製令副本參與人
// @Tags ManufactureUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MuID path string true "製令副本參與人ID"
// @param * body manufacture_user.Updated true "更新製令副本參與人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_user/{MuID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	muID := ctx.Param("MuID")
	input := &manufacture_user.Updated{}
	input.MuID = muID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureUserResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
