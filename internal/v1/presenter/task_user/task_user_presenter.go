package task_user

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/task_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Task_user.c 新增任務負責人
// @description 新增任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body task_user.Created_List true "新增任務負責人"
// @success 200 object code.SuccessfulMessage{body=[]string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task_user [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//有問題
	//AID := util.GenerateUUID()
	input := &task_user.Created_List{}
	//input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Task_user.1 搜尋任務負責人
// @description 搜尋任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID query string false "負責人ID"
// @param taskID query string false "任務編號"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=task_user.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task_user [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &task_user.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.TaskUserResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByDocumnetIDListHour
// @Summary Task_user.3 條件搜尋任務負責人
// @description 條件搜尋任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "條件搜尋任務負責人"
// @success 200 object code.SuccessfulMessage{body=task_user.Task_user_Labor_Hours} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task_user/GetByDocumnetIDListHour/{documentsID} [get]
func (p *presenter) GetByDocumnetIDListHour(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	input := &task_user.Field{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.GetByDocumnetIDListHour(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Task_user.2 取得單一任務負責人
// @description 取得單一任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tuID path string true "取得單一任務負責人"
// @success 200 object code.SuccessfulMessage{body=task_user.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/task_user/{tuID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	tuID := ctx.Param("tuID")
	input := &task_user.Field{}
	input.TuID = tuID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetName
// @Summary Task_user.4 取得任務負責人名稱
// @description 取得任務負責人名稱
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tuID path string true "取得任務負責人名稱"
// @success 200 object code.SuccessfulMessage{body=task_user.Task_user_Account} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/task_user/GetName/{tuID} [get]
func (p *presenter) GetName(ctx *gin.Context) {
	tuID := ctx.Param("tuID")
	input := &task_user.Field{}
	input.TuID = tuID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.GetName(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// DeleteList
// @Summary Task_user.d2 刪除多筆任務負責人
// @description 刪除多筆任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/task_user [delete]
func (p *presenter) DeleteList(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	// tuID := ctx.Param("tuID")
	// input := &task_user.Updated{}
	// input.TuID = tuID
	input := &task_user.Updated_List{}
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.DeleteList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Task_user.d1 刪除單一任務負責人
// @description 刪除單一任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tuID path string true "刪除單一任務負責人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/task_user/{tuID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	tuID := ctx.Param("tuID")
	input := &task_user.Updated{}
	input.TuID = tuID

	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Task_user.u1 更新單一或多筆任務負責人
// @description 更新單一或多筆任務負責人
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body task_user.Updated_List true "更新任務負責人"
// @success 200 object code.SuccessfulMessage{body=[]string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/task_user [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	//tuID := ctx.Param("tuID")
	input := &task_user.Updated_List{}
	//input.TuID = tuID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// UpdatedStatus
// @Summary Task_user.u2 更新任務審核狀態
// @description 更新任務審核狀態
// @Tags Task_user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tuID path string true "更新任務審核狀態"
// @param * body task_user.Updated_Review true "更新任務審核狀態"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/task_user/UpdatedStatus/{tuID} [patch]
func (p *presenter) UpdatedStatus(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	tuID := ctx.Param("tuID")
	input := &task_user.Updated_Review{}
	input.TuID = tuID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskUserResolver.UpdatedStatus(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
