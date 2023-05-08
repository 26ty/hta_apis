package countersign_user

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/countersign_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary CountersignUser.c 新增會簽人員
// @description 新增會簽人員
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body countersign_user.Created true "新增會簽人員"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &countersign_user.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary CountersignUser.1 條件搜尋會簽人員
// @description 條件會簽人員
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=countersign_user.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &countersign_user.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CountersignUserResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByIDCountersignUserListUser
// @Summary CountersignUser.4 取得單一單據底下的所有會簽人員的詳細資料
// @description 取得單一單據底下的所有會簽人員的詳細資料，包含名稱、所屬部門、父部門
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據來源ID"
// @param countersignID query string false "會簽ID"
// @param cuID query string false "會簽人員ID"
// @success 200 object code.SuccessfulMessage{body=countersign_user.CountersignUser_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user/GetByIDCountersignUserListUser/{documentsID} [get]
func (p *presenter) GetByIDCountersignUserListUser(ctx *gin.Context) {
	documentsID := ctx.Param("DocumentsID")
	input := &countersign_user.Documents{}
	input.DocumentsID = documentsID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.GetByIDCountersignUserListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByIDCountersignUserListUser2
// @Summary CountersignUser.5 取得單一單據及單一會簽底下的所有會簽人員的詳細資料
// @description 取得單一單據及單一會簽底下的所有會簽人員的詳細資料，包含名稱、所屬部門、父部門
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據來源ID"
// @param countersignID path string true "會簽ID"
// @param cuID query string false "會簽人員ID"
// @success 200 object code.SuccessfulMessage{body=countersign_user.CountersignUser_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user/GetByIDCountersignUserListUser/{documentsID}/{countersignID} [get]
func (p *presenter) GetByIDCountersignUserListUser2(ctx *gin.Context) {
	documentsID := ctx.Param("DocumentsID")
	countersignID := ctx.Param("CountersignID")
	input := &countersign_user.Documents{}
	input.DocumentsID = documentsID
	input.CountersignId = countersignID

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.GetByIDCountersignUserListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByCuIDCountersignUserListUser
// @Summary CountersignUser.3 取得單一會簽人員的詳細資料
// @description 取得單一會簽人員的詳細資料，包含名稱、所屬部門、父部門
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID query string false "單據來源ID"
// @param countersignID query string false "會簽ID"
// @param cuID path string true "會簽人員ID"
// @success 200 object code.SuccessfulMessage{body=countersign_user.CountersignUser_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user/GetByCuIDCountersignUserListUser/{cuID} [get]
func (p *presenter) GetByCuIDCountersignUserListUser(ctx *gin.Context) {
	cuID := ctx.Param("CuID")
	input := &countersign_user.Documents{}
	input.CuID = cuID

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.GetByIDCountersignUserListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary CountersignUser.2 取得單一會簽人員
// @description 取得單一會簽人員
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cuID path string true "會簽人員ID"
// @success 200 object code.SuccessfulMessage{body=countersign_user.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user/{cuID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	cuID := ctx.Param("CuID")
	input := &countersign_user.Field{}
	input.CuID = cuID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary CountersignUser.d 刪除單一會簽人員
// @description 刪除單一會簽人員
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cuID path string true "會簽人員ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user/{cuID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	cuID := ctx.Param("CuID")
	input := &countersign_user.Updated{}
	input.CuID = cuID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary CountersignUser.u 更新單一會簽人員
// @description 更新單一會簽人員
// @Tags CountersignUser
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cuID path string true "會簽人員ID"
// @param * body countersign_user.Updated true "更新會簽人員"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/countersign_user/{cuID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	cuID := ctx.Param("CuID")
	input := &countersign_user.Updated{}
	input.CuID = cuID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignUserResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
