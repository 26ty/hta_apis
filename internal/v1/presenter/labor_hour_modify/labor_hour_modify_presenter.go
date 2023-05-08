package labor_hour_modify

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"
	model "eirc.app/internal/v1/structure"
	bpm "eirc.app/internal/pkg/bpm"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/labor_hour_modify"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary LaborHourModify.c 新增工時異動
// @description 新增工時異動
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body labor_hour_modify.Created true "新增工時異動"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &labor_hour_modify.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//登入Bonita
	account := ctx.Param("account")
	bonita_input := &model.CaseIDModelInput{}
	bonita_input.Account = account

	//取得工時異動的ProccessID
	Labor := bpm.GetProcessID(ctx,bonita_input.Account,"工時異動")
	if Labor == "error"{
		return
	}
	//起單並取得CaseID
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx,Labor,bonita_input)

	if output != nil{
		input.BonitaCaseID = output.CaseID
		codeMessage := p.LaborHourModifyResolver.Created(trx, input)
		ctx.JSON(http.StatusOK, codeMessage)
	}	
}

// List
// @Summary LaborHourModify.1 條件搜尋工時異動
// @description 條件工時異動
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param category query string false "種類ID"
// @param title query string false "標題"
// @param content query string false "內容"
// @param nature query string false "類型"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=labor_hour_modify.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &labor_hour_modify.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.LaborHourModifyResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary LaborHourModify.2 取得單一工時異動
// @description 取得單一工時異動
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HmID path string true "工時異動ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour_modify.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/{HmID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	hmID := ctx.Param("HmID")
	input := &labor_hour_modify.Field{}
	input.HmID = hmID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourModifyResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetLaborBonitaCaseListStart
// @Summary LaborHourModify.lh-r 工時異動獲取使用者可執行的單(異動記錄重新送審)
// @description 工時異動獲取使用者可執行的單(異動記錄重新送審)
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param caseID path string true "Bonita單據ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/GetLaborBonitaCaseListStart/{account}/{caseID} [get]
func (p *presenter) GetLaborBonitaCaseListStart(ctx *gin.Context) {
	input := &model.GetDetailListInput{}
	caseID := ctx.Param("caseID")
	input.CaseID = caseID
	account := ctx.Param("account")
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetCasePendingTaskDetail(ctx,input.Account,input.CaseID)

	if output != nil{
		for _, value := range *output {
			if value.Name == "異動記錄重新送審"{
				ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, value.ID))
				return
			}
		}	
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, nil))
	
}

// GetLaborBonitaCaseListDepartment
// @Summary LaborHourModify.lh-1 工時異動獲取使用者可執行的單(主管審核)
// @description 工時異動獲取使用者可執行的單(主管審核)
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "Bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]labor_hour_modify.ReviewByDepartment} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/GetLaborBonitaCaseListDepartment/{account}/{userID} [get]
func (p *presenter) GetLaborBonitaCaseListDepartment(ctx *gin.Context) {
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
	output := &[]model.GetCaseListOutput{}
	output = bpm.GetUserExecutable(ctx,input.Account,input.UserID)

	if output != nil{
		codeMessage := p.LaborHourModifyResolver.GetByCaseIDtoDepartment(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// GetByUserIdLaborHourModifyList
// @Summary LaborHourModify.4 篩選出該任務該使用者已提報工時列表
// @description 用t_id、user_id篩選出該任務該使用者已提報工時列表
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @param tID path string true "任務ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour_modify.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/GetByUserIdLaborHourModifyList/{userID}/{tID} [get]
func (p *presenter) GetByUserIdLaborHourModifyList(ctx *gin.Context) {
	input := &labor_hour_modify.Field{}
	userID := ctx.Param("userID")
	input.UserID = userID
	tID := ctx.Param("tID")
	input.TID = tID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.LaborHourModifyResolver.GetByUserIdLaborHourModifyList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByCuIdLaborHourModifyList
// @Summary LaborHourModify.3 篩選出該會簽已提報工時列表
// @description 用cu_id篩選出該會簽已提報工時列表
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cuID path string true "會簽ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour_modify.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/GetByCuIdLaborHourModifyList/{cuID} [get]
func (p *presenter) GetByCuIdLaborHourModifyList(ctx *gin.Context) {
	input := &labor_hour_modify.Field{}
	userID := ctx.Param("userID")
	input.UserID = userID
	cuID := ctx.Param("cuID")
	input.CuID = cuID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.LaborHourModifyResolver.GetByCuIdLaborHourModifyList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIdList
// @Summary LaborHourModify.5 取得該使用者的工時
// @description UserID篩選取得該使用者的工時
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour_modify.LaborHourModifys} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/GetByUserIdList/{userID} [get]
func (p *presenter) GetByUserIdList(ctx *gin.Context) {
	input := &labor_hour_modify.Field{}
	userID := ctx.Param("userID")
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.LaborHourModifyResolver.GetByUserIdList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary LaborHourModify.d 刪除單一工時異動
// @description 刪除單一工時異動
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HmID path string true "工時異動ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/{HmID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	hmID := ctx.Param("HmID")
	input := &labor_hour_modify.Updated{}
	input.HmID = hmID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourModifyResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary LaborHourModify.u 更新單一工時異動
// @description 更新單一工時異動
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HmID path string true "工時異動ID"
// @param * body labor_hour_modify.Updated true "更新工時異動"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/{HmID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	hmID := ctx.Param("HmID")
	input := &labor_hour_modify.Updated{}
	input.HmID = hmID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourModifyResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// UpdatedStatus
// @Summary LaborHourModify.u-2 審核狀態變更
// @description status_type_id審核狀態變更
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HmID path string true "工時異動ID"
// @param * body labor_hour_modify.Updated_Review true "更新工時異動審核狀態"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/UpdatedStatus/{HmID} [patch]
func (p *presenter) UpdatedStatus(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	hmID := ctx.Param("HmID")
	input := &labor_hour_modify.Updated_Review{}
	input.HmID = hmID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourModifyResolver.UpdatedStatus(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// LaborReviewTask
// @Summary LaborHourModify.lh-review 工時異動主管審核通過
// @description 工時異動主管審核通過
// @Tags LaborHourModify
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param taskID path string true "Bonita任務ID"
// @param * body model.ReviewInput true "更新工時異動bonita審核狀態"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour_modify/LaborReviewTask/{account}/{taskID} [patch]
func (p *presenter) LaborReviewTask(ctx *gin.Context) {
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

	if input.HmID == ""{
		ctx.JSON(http.StatusOK,code.GetCodeMessage(code.Successful, body))
	}else{
		update := &labor_hour_modify.Field{}
		update.HmID = input.HmID
		codeMessage := p.LaborHourModifyResolver.Replace(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}


}
