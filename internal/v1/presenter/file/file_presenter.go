package file

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/file"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// Created
// @Summary File.c 新增檔案
// @description 新增檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body file.Created true "新增檔案"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file [post]
func (p *presenter) Created(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	defer trx.Rollback()
	input := &file.Created{}

	// input.CreatedBy = input.AccountID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	// filepath := `http://dindins3.s3-website-ap-southeast-1.amazonaws.com`

	//勝彰寫的dindins3-ap-southeast-1
	s3BucketName := "dindins3"
	input.FilePath = "files/" + util.GenerateUUID() + "." + input.FileExtension

	url := util.FileToS3(input, s3BucketName)

	if len(url) == 0 {
		ctx.JSON(http.StatusNotExtended, "上傳失敗")
	} else {
		input.DownloadUrl = url
		codeS3 := p.FileResolver.Created(trx, input)
		ctx.JSON(http.StatusOK, codeS3)
	}
}

// List
// @Summary File.1 條件搜尋檔案
// @description 條件檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=file.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &file.Fields{}
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	input.Limit, _ = strconv.ParseInt(limit, 10, 64)
	input.Page, _ = strconv.ParseInt(page, 10, 64)

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}
	codeS3 := p.FileResolver.List(input)
	ctx.JSON(http.StatusOK, codeS3)
}

// GetByDocumentID
// @Summary File.3 篩選出專案或客需單裡的全部檔案
// @description 篩選出專案或客需單裡的全部檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據ID"
// @success 200 object code.SuccessfulMessage{body=file.FilebydocumentIds} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/GetByDocumentID/{documentsID} [get]
func (p *presenter) GetByDocumentID(ctx *gin.Context) {
	DocumentsID := ctx.Param("documentsID")

	input := &file.Field{}
	input.DocumentsID = DocumentsID

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FileResolver.GetByDocumentID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByDocumentIDUserID
// @Summary File.4 篩選出該任務(TaskID)或會簽(CsID)及該使用者(UserID)上傳的所有附件列表
// @description 篩選出該任務(TaskID)或會簽(CsID)及該使用者(UserID)上傳的所有附件列表
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentsID path string true "單據ID"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=file.FilebydocumentIds} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/GetByDocumentIDUserID/{documentsID}/{userID} [get]
func (p *presenter) GetByDocumentIDUserID(ctx *gin.Context) {
	DocumentsID := ctx.Param("documentsID")
	UserID := ctx.Param("userID")


	input := &file.Users{}
	input.DocumentsID = DocumentsID
	input.UserID = UserID

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FileResolver.GetByDocumentIDUserID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary File.2 取得單一檔案
// @description 取得單一檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param fID path string true "檔案ID"
// @success 200 object code.SuccessfulMessage{body=file.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/{fID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	FID := ctx.Param("fID")
	input := &file.Field{}
	input.FID = FID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FileResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary File.d 刪除單一檔案
// @description 刪除單一檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param fID path string true "檔案ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/{fID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	FID := ctx.Param("fID")

	input := &file.Updated{}
	input.FID = FID

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FileResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary File.u 更新單一檔案
// @description 更新單一檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param fID path string true "檔案ID"
// @param * body file.Updated true "更新檔案"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/{fID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	//Todo 將UUID改成登入的使用者
	FID := ctx.Param("fID")
	input := &file.Updated{}
	input.FID = FID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	codeMessage := p.FileResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
