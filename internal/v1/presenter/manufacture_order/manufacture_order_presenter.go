package manufacture_order

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	bpm "eirc.app/internal/pkg/bpm"
	model "eirc.app/internal/v1/structure"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/manufacture_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary ManufactureOrder.c 新增製令
// @description 新增製令
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body manufacture_order.Created true "新增製令"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &manufacture_order.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureOrderResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary ManufactureOrder.1 條件搜尋製令
// @description 條件製令
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=manufacture_order.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &manufacture_order.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ManufactureOrderResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// func (p *presenter) ManufactureOrderCdListUser(ctx *gin.Context) {
// 	input := &manufacture_order.Fields{}
// 	if err := ctx.ShouldBindQuery(input); err != nil {
// 		log.Error(err)
// 		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

// 		return
// 	}

// 	if input.Limit >= preset.DefaultLimit {
// 		input.Limit = preset.DefaultLimit
// 	}

// 	codeMessage := p.ManufactureOrderResolver.ManufactureOrderCdListUser(input)
// 	ctx.JSON(http.StatusOK, codeMessage)
// }

// ManufactureOrderProjectListUser
// @Summary ManufactureOrder.3 顯示Project&Customer_demand的code
// @description 顯示Project&Customer_demand的code
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=manufacture_order.ManufactureOrder_Project_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/ManufactureOrderProjectListUser [get]
func (p *presenter) ManufactureOrderProjectListUser(ctx *gin.Context) {
	input := &manufacture_order.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ManufactureOrderResolver.ManufactureOrderProjectListUser(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary ManufactureOrder.2 取得單一製令
// @description 取得單一製令
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "製令ID"
// @success 200 object code.SuccessfulMessage{body=manufacture_order.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/{MID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	mID := ctx.Param("MID")
	input := &manufacture_order.Field{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureOrderResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByIDOne
// @Summary ManufactureOrder.4 取得單一Project&Customer_demand的code
// @description 取得單一Project&Customer_demand的code
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "製令ID"
// @success 200 object code.SuccessfulMessage{body=manufacture_order.Ones} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetByIDOne/{MID} [get]
func (p *presenter) GetByIDOne(ctx *gin.Context) {
	mID := ctx.Param("MID")
	input := &manufacture_order.Field{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureOrderResolver.GetByIDOne(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByPIDList
// @Summary ManufactureOrder.5 條件搜尋專案內的所有組合機器
// @description 條件搜尋專案內的所有組合機器
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PID path string true "專案ID"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=manufacture_order.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetByPIDList/{PID} [get]
func (p *presenter) GetByPIDList(ctx *gin.Context) {
	pID := ctx.Param("PID")
	input := &manufacture_order.Fields{}
	input.ProjectID = pID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ManufactureOrderResolver.GetByPIDList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetC2BonitaCaseListStart
// @Summary ManufactureOrder.c1-r 獲取使用者可執行的單(重啟單)
// @description C1-R獲取使用者可執行的單(重啟單)
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param caseID path string true "bonita單據ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetC2BonitaCaseListStart/{account}/{caseID} [get]
func (p *presenter) GetC2BonitaCaseListStart(ctx *gin.Context) {
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
			if value.Name == "填寫製令資料"{
				ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, value.ID))
				return
			}
		}	
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, nil))
	
}

// GetC2BonitaCaseListDepartment
// @Summary  ManufactureOrder.c1-1 獲取使用者可執行的單(單位主管審核)
// @description  C1-1獲取使用者可執行的單(單位主管審核)
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]manufacture_order.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetC2BonitaCaseListDepartment/{account}/{userID} [get]
func (p *presenter) GetC2BonitaCaseListDepartment(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)

	if output != nil {
		codeMessage := p.ManufactureOrderResolver.GetByCaseIDtoDepartment(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetC2BonitaCaseListManufacture
// @Summary  ManufactureOrder.c1-2 獲取使用者可執行的單(生管-製造-審核)
// @description  C1-2獲取使用者可執行的單(生管-製造-審核)
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]manufacture_order.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetC2BonitaCaseListManufacture/{account}/{userID} [get]
func (p *presenter) GetC2BonitaCaseListManufacture(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)

	if output != nil {
		codeMessage := p.ManufactureOrderResolver.GetByCaseIDtoManufacture(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetC2BonitaCaseListTop
// @Summary  ManufactureOrder.c1-3 獲取使用者可執行的單(總經理審核)
// @description  C1-3獲取使用者可執行的單(總經理審核)
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]manufacture_order.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetC2BonitaCaseListTop/{account}/{userID} [get]
func (p *presenter) GetC2BonitaCaseListTop(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)

	if output != nil {
		codeMessage := p.ManufactureOrderResolver.GetByCaseIDtoTop(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetC2BonitaCaseListConfirm
// @Summary  ManufactureOrder.c1-4 獲取使用者可執行的單(確認單號開啟)
// @description  C1-4獲取使用者可執行的單(確認單號開啟)
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]manufacture_order.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetC2BonitaCaseListConfirm/{account}/{userID} [get]
func (p *presenter) GetC2BonitaCaseListConfirm(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)

	if output != nil {
		codeMessage := p.ManufactureOrderResolver.GetByCaseIDtoConfirm(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetC2BonitaCaseListSave
// @Summary  ManufactureOrder.c1-5 獲取使用者可執行的單(儲存製令單號)
// @description  C1-5獲取使用者可執行的單(儲存製令單號)
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]manufacture_order.Review} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/GetC2BonitaCaseListSave/{account}/{userID} [get]
func (p *presenter) GetC2BonitaCaseListSave(ctx *gin.Context) {
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
	output = bpm.GetUserExecutable(ctx, input.Account, input.UserID)

	if output != nil {
		codeMessage := p.ManufactureOrderResolver.GetByCaseIDtoSave(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// Delete
// @Summary ManufactureOrder.d 刪除單一製令
// @description 刪除單一製令
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "製令ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/{MID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mID := ctx.Param("MID")
	input := &manufacture_order.Updated{}
	input.MID = mID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureOrderResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// C2UpdatedCaseID
// @Summary ManufactureOrder.c1-s起單
// @description C1起單
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "製令ID"
// @param * body model.CaseIDModelInput true "起單內容"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/C2UpdatedCaseID/{MID} [patch]
func (p *presenter) C2UpdatedCaseID(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	//client.Login("isabelle_wu")
	input := &model.CaseIDModelInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	C2 := bpm.GetProcessID(ctx,input.Account,"製令開啟通知作業")
	if C2 == "error"{
		return
	}
	output := &model.GetCaseIDOutput{}
	output = bpm.GetCaseID(ctx, C2, input)
	
	if output != nil {
		mID := ctx.Param("mID")
		update := &manufacture_order.Updated_Bonita{}
		update.MID = mID
		update.BonitaCaseID = output.CaseID

		codeMessage := p.ManufactureOrderResolver.Updated_Bonita(update)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// C2ReviewTask
// @Summary ManufactureOrder.c1-review 審核
// @description C1審核
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param taskID path string true "bonita任務ID"
// @param * body model.ReviewInput true "審核狀態"
// @success 200 object code.SuccessfulMessage{body=int} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/C2ReviewTask/{account}/{taskID} [patch]
func (p *presenter) C2ReviewTask(ctx *gin.Context) {

	account := ctx.Param("account")
	taskID := ctx.Param("taskID")

	input := &model.ReviewInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	body := bpm.ReviewTask(ctx, account, taskID, input)
	

	if body != 204 {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(body, "bonita review error"))

		return
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, body))

}

// Updated
// @Summary ManufactureOrder.u 更新單一製令
// @description 更新單一製令
// @Tags ManufactureOrder
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param MID path string true "製令ID"
// @param * body manufacture_order.Updated true "更新製令"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/manufacture_order/{MID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	mID := ctx.Param("MID")
	input := &manufacture_order.Updated{}
	input.MID = mID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ManufactureOrderResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
