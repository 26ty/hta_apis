package meeting

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/meeting"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Meeting.c 新增會議
// @description 新增會議
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body meeting.Created true "新增會議"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &meeting.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MeetingResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Meeting.1 條件搜尋會議
// @description 條件會議
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param name query string false "會議名稱"
// @param room query string false "會議室"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=meeting.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &meeting.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MeetingResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// MeetingUser
// @Summary Meeting.7 條件搜尋將id轉成人名的會議詳細資料
// @description 條件搜尋將id轉成人名的會議詳細資料
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=meeting.AllMeetingListUserParticipant} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/MeetingUser [get]
func (p *presenter) MeetingUser(ctx *gin.Context) {
	input := &meeting.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MeetingResolver.MeetingUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIDMeetingListUser
// @Summary Meeting.6 取得該使用者的所有會議
// @description 取得該使用者的所有會議
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=meeting.MeetingListUsers} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/GetByUserIDMeetingListUser/{userID} [get]
func (p *presenter) GetByUserIDMeetingListUser(ctx *gin.Context) {
	userID := ctx.Param("UserID")
	input := &meeting.Users{}
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MeetingResolver.GetByUserIDMeetingListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByMIDMeetingListUser
// @Summary Meeting.4 列出該會議下的參與者account_id,name
// @description 列出該會議下的參與者account_id,name
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "會議ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=meeting.MeetingListUsers} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/GetByMIDMeetingListUser/{MID} [get]
func (p *presenter) GetByMIDMeetingListUser(ctx *gin.Context) {
	mID := ctx.Param("MID")
	input := &meeting.Fields{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MeetingResolver.GetByMIDMeetingListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByDIDMeetingListUser
// @Summary Meeting.3 列出屬於該documentsID下的會議列表(將id轉成人名的)
// @description 列出屬於該documentsID下的會議列表(將id轉成人名的)
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param DocumentsID path string true "所屬單據ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=meeting.MeetingListUsers} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/GetByDIDMeetingListUser/{DocumentsID} [get]
func (p *presenter) GetByDIDMeetingListUser(ctx *gin.Context) {
	documentsID := ctx.Param("DocumentsID")
	input := &meeting.Fields{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.MeetingResolver.GetByDIDMeetingListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByMIDMeetingUser
// @Summary Meeting.5 選取單一筆將id轉成人名的會議資料
// @description 選取單一筆將id轉成人名的會議資料
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "會議ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=meeting.MeetingListUserParticipant} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/GetByMIDMeetingUser/{MID} [get]
func (p *presenter) GetByMIDMeetingUser(ctx *gin.Context) {
	mID := ctx.Param("MID")
	input := &meeting.Field{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MeetingResolver.GetByMIDMeetingUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Meeting.2 取得單一會議
// @description 取得單一會議
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "會議ID"
// @success 200 object code.SuccessfulMessage{body=meeting.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/{MID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	mID := ctx.Param("MID")
	input := &meeting.Field{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MeetingResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Meeting.d 刪除單一會議
// @description 刪除單一會議
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "會議ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/{MID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mID := ctx.Param("MID")
	input := &meeting.Updated{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MeetingResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Meeting.u 更新單一會議
// @description 更新單一會議
// @Tags Meeting
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "會議ID"
// @param * body meeting.Updated true "更新會議"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/meeting/{MID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mID := ctx.Param("MID")
	input := &meeting.Updated{}
	input.MID = mID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.MeetingResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
