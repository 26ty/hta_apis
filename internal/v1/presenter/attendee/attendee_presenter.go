package attendee

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/attendee"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Attendee.c 新增會議參與人員
// @description 新增會議參與人員
// @Tags Attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body attendee.Created true "新增會議參與人員"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/attendee [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//有問題
	//AID := util.GenerateUUID()
	input := &attendee.Created{}
	//input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AttendeeResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Attendee.1 條件搜尋會議參與人員
// @description 條件會議參與人員
// @Tags Attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param meetID query string false "會議ID"
// @param userID query string false "人員ID"
// @param chairman query bool false "是否為主席"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=attendee.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/attendee [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &attendee.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.AttendeeResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Attendee.2 取得單一會議參與人員
// @description 取得單一會議參與人員
// @Tags Attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param aID path string true "會議參與人員ID"
// @success 200 object code.SuccessfulMessage{body=attendee.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/attendee/{aID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	aID := ctx.Param("aID")
	input := &attendee.Field{}
	input.AID = aID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AttendeeResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Attendee.d 刪除單一會議參與人員
// @description 刪除單一會議參與人員
// @Tags Attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param aID path string true "會議參與人員ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/attendee/{aID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	aID := ctx.Param("aID")
	input := &attendee.Updated{}
	input.AID = aID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AttendeeResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Attendee.u 更新單一或多筆會議參與人員
// @description 更新單一或多筆會議參與人員
// @Tags Attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body attendee.Updated_List true "更新會議參與人員"
// @success 200 object code.SuccessfulMessage{body=[]string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/attendee [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	//aID := ctx.Param("aID")
	input := &attendee.Updated_List{}
	//input.AID = aID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AttendeeResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

