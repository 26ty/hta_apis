package transaction_record

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/transaction_record"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Transaction_record.c 新增異動紀錄資料
// @description 新增異動紀錄資料
// @Tags Transaction_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body transaction_record.Created true "新增異動紀錄資料"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/transaction_record [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &transaction_record.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TransactionRecordResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Transaction_record.1 搜尋異動紀錄資料
// @description 搜尋異動紀錄資料
// @Tags Transaction_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=transaction_record.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/transaction_record [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &transaction_record.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.TransactionRecordResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


// GetByDocumentIDUserList
// @Summary Transaction_record.3 選取單一單據來源ID之異動紀錄資料
// @description 選取單一單據來源ID之異動紀錄資料
// @Tags Transaction_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param documentID path string true "單據來源ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=transaction_record.Record_user_lists} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/transaction_record/GetByDocumentIDUserList/{documentID} [get]
func (p *presenter) GetByDocumentIDUserList(ctx *gin.Context) {
	documentID := ctx.Param("DocumentID")
	input := &transaction_record.Fields{}
	input.DocumentID = documentID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.TransactionRecordResolver.GetByDocumentIDUserList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Transaction_record.2 取得單一異動紀錄
// @description 取得單一異動紀錄
// @Tags Transaction_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param trID path string true "取得單一異動紀錄"
// @success 200 object code.SuccessfulMessage{body=transaction_record.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/transaction_record/{trID}} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	trID := ctx.Param("TrID")
	input := &transaction_record.Field{}
	input.TrID = trID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TransactionRecordResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Transaction_record.d 刪除單一異動紀錄
// @description 刪除單一異動紀錄
// @Tags Transaction_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param trID path string true "刪除單一異動紀錄"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/transaction_record/{trID}} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trID := ctx.Param("TrID")
	input := &transaction_record.Updated{}
	input.TrID = trID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TransactionRecordResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Transaction_record.u 更新單一異動紀錄
// @description 更新單一異動紀錄
// @Tags Transaction_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param trID path string true "更新單一異動紀錄"
// @param * body transaction_record.Updated true "更新單一異動紀錄"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @Router /authority/v1.0/transaction_record/{trID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trID := ctx.Param("TrID")
	input := &transaction_record.Updated{}
	input.TrID = trID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.TransactionRecordResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
