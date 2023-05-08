package todo_type

import (
	"net/http"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/todo_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Todo_type.c 新增代辦事項分類
// @description 新增代辦事項分類
// @Tags Todo_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body todo_type.Created true "新增代辦事項分類"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/todo_type [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &todo_type.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
		codeMessage := p.Todo_typeResolver.Created(trx, input)
		ctx.JSON(http.StatusOK, codeMessage)
	
}

// List
// @Summary Todo_type.1 搜尋代辦事項分類
// @description 搜尋代辦事項分類
// @Tags Todo_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=todo_type.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/todo_type [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &todo_type.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.Todo_typeResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Todo_type.2 取得單一代辦事項分類
// @description 取得單一代辦事項分類
// @Tags Todo_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param ttID path string true "取得單一代辦事項分類"
// @success 200 object code.SuccessfulMessage{body=todo_type.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/todo_type/{ttID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	ttID := ctx.Param("ttID")
	input := &todo_type.Field{}
	input.TtID = ttID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Todo_typeResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserID
// @Summary Todo_type.3 條件搜尋單一代辦事項分類
// @description 條件搜尋單一代辦事項分類
// @Tags Todo_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "條件搜尋單一代辦事項分類"
// @success 200 object code.SuccessfulMessage{body=todo_type.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/todo_type/GetByUserID/{userID} [get]
func (p *presenter) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	input := &todo_type.Field{}
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Todo_typeResolver.GetByUserID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Todo_type.d 刪除單一代辦事項分類
// @description 刪除單一代辦事項分類
// @Tags Todo_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param ttID path string true "刪除單一代辦事項分類"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/todo_type/{ttID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {

	ttID := ctx.Param("ttID")
	input := &todo_type.Updated{}
	input.TtID = ttID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Todo_typeResolver.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Todo_type.u 更新單一代辦事項分類
// @description 更新單一代辦事項分類
// @Tags Todo_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param ttID path string true "更新單一代辦事項分類"
// @param * body todo_type.Updated true "更新單一代辦事項分類"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/todo_type/{ttID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	ttID := ctx.Param("ttID")
	input := &todo_type.Updated{}
	input.TtID = ttID

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Todo_typeResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
