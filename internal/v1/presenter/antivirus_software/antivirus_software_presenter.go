package antivirus_software

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/antivirus_software"
	model "eirc.app/internal/v1/structure"
	bpm "eirc.app/internal/pkg/bpm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Created
// @Summary AntivirusSoftware.c 新增防毒軟件
// @description 新增防毒軟件
// @Tags AntivirusSoftware
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body antivirus_software.Created true "新增防毒軟件"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/antivirus_software [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &antivirus_software.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//起單要給的資料
	input_bonita := &model.CaseIDModelInput{}
	input_bonita.Account = input.Account
	//找到新增軟體審核單的流程ID，
	AS := bpm.GetProcessID(ctx,input_bonita.Account,"新增軟體審核單")
	if AS == "error"{
		return
	}
	//將起單要給的資料與流程ID一同傳給bonita進行起單
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx,AS,input_bonita)

	if output != nil{
		input.BonitaCaseID = output.CaseID
		codeMessage := p.AntivirusSoftwareResolver.Created(trx, input)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// List
// @Summary AntivirusSoftware.1 條件搜尋防毒軟件
// @description 條件防毒軟件
// @Tags AntivirusSoftware
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=antivirus_software.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/antivirus_software [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &antivirus_software.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.AntivirusSoftwareResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary AntivirusSoftware.2 取得單一防毒軟件
// @description 取得單一防毒軟件
// @Tags AntivirusSoftware
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param asID path string true "防毒軟件ID"
// @success 200 object code.SuccessfulMessage{body=antivirus_software.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/antivirus_software/{asID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	asID := ctx.Param("AsID")
	input := &antivirus_software.Field{}
	input.AsID = asID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AntivirusSoftwareResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) GetByCaseIDtoTop(ctx *gin.Context) {
	input := &model.GetCaseListInput{}
	userID := ctx.Param("userID")
	input.UserID = userID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	//使用者可執行的單據
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.AntivirusSoftwareResolver.GetByCaseIDtoTop(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetByPIDList
// @Summary AntivirusSoftware.3 取得單一Project底下的所有防毒軟件
// @description 取得單一Project底下的所有防毒軟件
// @Tags AntivirusSoftware
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param pID path string true "專案ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=antivirus_software.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/antivirus_software/GetByPIDList/{pID} [get]
func (p *presenter) GetByPIDList(ctx *gin.Context) {
	pID := ctx.Param("PID")
	input := &antivirus_software.Fields{}
	input.ProjectID = pID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.AntivirusSoftwareResolver.GetByPIDList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary AntivirusSoftware.d 刪除單一防毒軟件
// @description 刪除單一防毒軟件
// @Tags AntivirusSoftware
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param asID path string true "防毒軟件ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/antivirus_software/{asID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	asID := ctx.Param("AsID")
	input := &antivirus_software.Updated{}
	input.AsID = asID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AntivirusSoftwareResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary AntivirusSoftware.u 更新單一防毒軟件
// @description 更新單一防毒軟件
// @Tags AntivirusSoftware
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param asID path string true "防毒軟件ID"
// @param * body antivirus_software.Updated true "更新防毒軟件"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/antivirus_software/{asID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	asID := ctx.Param("AsID")
	input := &antivirus_software.Updated{}
	input.AsID = asID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AntivirusSoftwareResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

func (p *presenter) AsReviewTask(ctx *gin.Context) {

	account := ctx.Param("account")
	taskID := ctx.Param("taskID")

	input := &model.ReviewInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	body := bpm.ReviewTask(ctx,account,taskID,input)

	if body != 204{
		ctx.JSON(http.StatusOK, code.GetCodeMessage(body, "bonita review error"))

		return
	}

	if *input.Status == true{
		update := &antivirus_software.Updated{}
		update.AsID = input.AsID
		update.Status = "已通過"
		codeMessage := p.AntivirusSoftwareResolver.Updated(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}else{
		update := &antivirus_software.Updated{}
		update.AsID = input.AsID
		update.Status = "已退回"
		codeMessage := p.AntivirusSoftwareResolver.Updated(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}