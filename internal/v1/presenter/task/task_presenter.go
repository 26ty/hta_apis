package task

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/task"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Task.c 新增任務
// @description 新增任務
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body task.Created true "新增任務"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//createBy := util.GenerateUUID()
	input := &task.Created_List{}
	//input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Task.1 條件搜尋任務
// @description 條件任務
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param t_name query string false "任務名字"
// @param remark query string false "備註"
// @param landmark query bool false "里程碑"
// @param rank query int false "排序"
// @param last_task query string false "上一任務"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=task.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &task.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.TaskResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByIDListTaskHour
// @Summary Task.3 工時-任務負責人工時列表(有無提報都全列)
// @description 工時-任務負責人工時列表(有無提報都全列)
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據ID"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_Account_Labor_Hours} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByIDListTaskHour/{documentsID}/{userID} [get]
func (p *presenter) GetByIDListTaskHour(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	userID := ctx.Param("userID")
	input := &task.Users{}
	input.DocumentsID = documentsID
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByIDListTaskHour(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByTaskListUser
// @Summary Task.4 選取單一筆把id轉換成人名的任務資料
// @description 選取單一筆把id轉換成人名的任務資料
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tID path string true "任務ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByTaskListUser/{tID} [get]
func (p *presenter) GetByTaskListUser(ctx *gin.Context) {
	tID := ctx.Param("tID")
	input := &task.Users{}
	input.TID = tID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByTaskListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetTaskListHourByUserID
// @Summary Task.5 資源-任務負責人已報工時列表(只列出已提報者列表)
// @description 資源-任務負責人已報工時列表(只列出已提報者列表)
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_Hour_Users} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByTaskListHourDocumentsID/{documentsID} [get]
func (p *presenter) GetTaskListHourByUserID(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	input := &task.Field{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetTaskListHourByUserID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByTaskListHourDocumentsAndUserID
// @Summary Task.6 資源-任務負責人個人累計提報工時
// @description 資源-任務負責人個人累計提報工時
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據ID"
// @param accountID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_Hour_Users} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByTaskListHourDocumentsAndUserID/{documentsID}/{accountID} [get]
func (p *presenter) GetByTaskListHourDocumentsAndUserID(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	input := &task.Field{}
	input.DocumentsID = documentsID
	accountID := ctx.Param("accountID")
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByTaskListHourDocumentsAndUserID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// TaskListUser
// @Summary Task.7 列出屬於該documentsID下的任務列表，且id轉換為人名
// @description 列出屬於該documentsID下的任務列表，且id轉換為人名
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByDocumentsTaskListUser/{documentsID} [get]
func (p *presenter) TaskListUser(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	input := &task.Users{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.TaskResolver.TaskListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByOriginIDAndUserID
// @Summary Task.8 篩選出該使用者的分類列表
// @description 篩選出該使用者的分類列表
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param originID path string true "來源ID"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_OriginIds} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByOriginIDAndUserID/{originID}/{userID} [get]
func (p *presenter) GetByOriginIDAndUserID(ctx *gin.Context) {
	originID := ctx.Param("originID")
	userID := ctx.Param("userID")
	input := &task.Users{}
	input.OriginID = originID
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByOriginIDAndUserID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByTIDTaskListUser
// @Summary Task.9 選取同一task_id，把id轉換成人名的任務資料列表
// @description 選取同一task_id，把id轉換成人名的任務資料列表
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tID path string true "任務ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=task.Task_User_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByTIDTaskListUser/{tID} [get]
func (p *presenter) GetByTIDTaskListUser(ctx *gin.Context) {
	tID := ctx.Param("tID")
	input := &task.Fields{}
	input.TID = tID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.TaskResolver.GetByTIDTaskListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByDocumentIDTaskListLast
// @Summary Task.10 專案範本-選取該專案範本代號的任務列表(舊版，已無法使用)
// @description 專案範本-選取該專案範本代號的任務列表(舊版，已無法使用)
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param ptID path string true "專案範本ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=task.Task_User_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByDocumentIDTaskListLast/{ptID} [get]
func (p *presenter) GetByDocumentIDTaskListLast(ctx *gin.Context) {
	ptID := ctx.Param("ptID")
	input := &task.Fields{}
	input.PtID = ptID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.TaskResolver.GetByDocumentIDTaskListLast(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByDocumentIDTaskList
// @Summary Task.11 專案範本-選取該專案範本的任務列表
// @description 專案範本-選取該專案範本的任務列表
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "專案範本ID"
// @success 200 object code.SuccessfulMessage{body=task.Task_Templates} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByDocumentIDTaskList/{documentsID} [get]
func (p *presenter) GetByDocumentIDTaskList(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	input := &task.Field{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByDocumentIDTaskList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByIDTaskBonitaUserList
// @Summary Task.12 單據篩選任務與bonita的資料
// @description 單據篩選任務與bonita的資料
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "專案範本ID"
// @success 200 object code.SuccessfulMessage{body=task.Bonita_ID_Lists} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/GetByIDTaskBonitaUserList/{documentsID} [get]
func (p *presenter) GetByIDTaskBonitaUserList(ctx *gin.Context) {
	documentsID := ctx.Param("documentsID")
	input := &task.Users{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByIDTaskBonitaUserList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Task.2 取得單一任務
// @description 取得單一任務
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tID path string true "任務ID"
// @success 200 object code.SuccessfulMessage{body=task.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/{tID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	tID := ctx.Param("tID")
	input := &task.Field{}
	input.TID = tID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// DeleteList
// @Summary Task.d2 刪除多筆任務
// @description 刪除多筆任務
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body task.Updated_List true "刪除任務"
// @success 200 object code.SuccessfulMessage{body=[]string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task [delete]
func (p *presenter) DeleteList(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	// tID := ctx.Param("tID")
	// input := &task.Updated{}
	// input.TID = tID
	input := &task.Updated_List{}
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.DeleteList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Task.d1 刪除單一任務
// @description 刪除單一任務
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tID path string true "任務ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/{tID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	tID := ctx.Param("tID")
	input := &task.Updated{}
	input.TID = tID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Task.u 更新單一任務
// @description 更新單一任務
// @Tags Task
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param tID path string true "任務ID"
// @param * body task.Updated true "更新任務"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/task/{tID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	//tID := ctx.Param("tID")
	input := &task.Updated_List{}
	//input.TID = tID
	//input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TaskResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
