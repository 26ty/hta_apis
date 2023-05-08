package account

import (
	"net/http"
	"strconv"
	"encoding/json"

	csv "eirc.app/internal/pkg/shared"
	bpm "eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	"eirc.app/internal/v1/structure/accounts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Account.c 新增使用者
// @description 新增使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body accounts.Created true "新增使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &accounts.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	input.CompanyID = ctx.MustGet("company_id").(string)
	input.CreatedBy = ctx.MustGet("account_id").(string) //從Token去得到AccountId

	create_account := ctx.Param("create_account")
	input.CreatedAccount = create_account

	input.UserName = input.Account
	input.ManagerID = input.BonitaManagerID
	input.Enabled = strconv.FormatBool(true)


	bonita_user_id := bpm.AddUserGetID(ctx, input)
	if bonita_user_id != 0 {

		input.BonitaUserID = strconv.Itoa(bonita_user_id) 

		codeMessage := p.AccountResolver.Created(trx, input)
		ctx.JSON(http.StatusOK, codeMessage)
	}
	
}

// List
// @Summary Account.1 條件搜尋使用者
// @description 條件使用者
// @Tags Account
// @version 1.0
// @Accept json	
// @produce json
// @param Authorization header string  true "JWE Token"
// @param companyID query string false "組織ID"
// @param account query string false "帳號"
// @param name query string false "名稱"
// @success 200 object code.SuccessfulMessage{body=accounts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &accounts.Users{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.AccountResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// AccountNameList
// @Summary Account.4 查看全部使用者
// @description 查看全部使用者的ID跟部門名稱
// @Tags Account
// @version 1.0
// @Accept json	
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=accounts.Account_Names} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account/GetAccountNameList [get]
func (p *presenter) AccountNameList(ctx *gin.Context) {
	input := &accounts.Users{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.AccountResolver.AccountNameList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// AccountNameDepartmentList
// @Summary Account.3 所有員工的所屬部門及職稱
// @description 所有員工的所屬部門及職稱(包含所屬部門的父部門名稱)
// @Tags Account
// @version 1.0
// @Accept json	
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=accounts.Account_Names} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account/AccountNameDepartmentList [get]
func (p *presenter) AccountNameDepartmentList(ctx *gin.Context) {
	input := &accounts.Users{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.AccountResolver.AccountNameDepartmentList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Account.2 取得單一使用者
// @description 取得單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=accounts.Single} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account/{accountID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	input := &accounts.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AccountResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Account.d 刪除單一使用者
// @description 刪除單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "使用者ID"
// @param account path string true "使用者帳號"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account/{accountID}/{account} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	accountID := ctx.Param("accountID")
	created_account := ctx.Param("created_account")

	input := &accounts.Updated{}
	input.AccountID = accountID
	input.CreatedAccount = created_account
	input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AccountResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Account.u. 更新單一使用者
// @description 更新單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "使用者ID"
// @param account path string true "使用者帳號"
// @param * body accounts.Updated true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 404 object code.ErrorMessage{detailed=string} "Not Found"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/account/{accountID}/{account} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	accountID := ctx.Param("accountID")
	created_account := ctx.Param("created_account")

	input := &accounts.Updated{}
	input.AccountID = accountID
	input.CreatedAccount = created_account

	input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AccountResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

//為了導入EMAIL用的(暫時寫死)
func (p *presenter) UpdatedCsv(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := util.GenerateUUID()
	filename := ctx.Param("filename")
	created_account := ctx.Param("created_account")
	input := &accounts.UpdatedCsvList{}
	account ,csv_err := csv.CsvToUser(filename,created_account)

	if csv_err != nil {
		log.Error(csv_err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, csv_err.Error()))

		return
	}

	marshal, err := json.Marshal(account)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	err = json.Unmarshal(marshal, &input.Account)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.AccountResolver.UpdatedCsv(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
