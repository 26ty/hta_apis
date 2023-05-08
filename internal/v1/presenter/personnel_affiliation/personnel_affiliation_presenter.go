package personnel_affiliation

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	//bpm "eirc.app/internal/pkg/bpm"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/personnel_affiliation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary PersonnelAffiliation.c 新增人員隸屬
// @description 新增人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param create_account path string true "新增者的帳號"
// @param * body personnel_affiliation.Created true "新增人員隸屬"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/{create_account} [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &personnel_affiliation.Created{}
	create_account := ctx.Param("create_account")
	input.Account = create_account

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PersonnelAffiliationResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)

}

// List
// @Summary PersonnelAffiliation.1 條件搜尋人員隸屬
// @description 條件人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=personnel_affiliation.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &personnel_affiliation.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.PersonnelAffiliationResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary PersonnelAffiliation.2 取得單一人員隸屬
// @description 取得單一人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PaID path string true "人員隸屬ID"
// @success 200 object code.SuccessfulMessage{body=personnel_affiliation.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/{PaID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	paID := ctx.Param("PaID")
	input := &personnel_affiliation.Field{}
	input.PaID = paID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PersonnelAffiliationResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserID
// @Summary PersonnelAffiliation.3 取得單一使用者的人員隸屬
// @description 取得單一使用者的人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param UserID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]personnel_affiliation.Affiliation_Account} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/GetByUserID/{UserID} [get]
func (p *presenter) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("UserID")
	input := &personnel_affiliation.Field{}
	input.UserID = userID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PersonnelAffiliationResolver.GetByUserID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByDepartmentID
// @Summary PersonnelAffiliation.4 取得單一部門的人員隸屬
// @description 取得單一部門的人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param DepartmentID path string true "部門ID"
// @success 200 object code.SuccessfulMessage{body=[]personnel_affiliation.Deparment_User} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/GetByDepartmentID/{DepartmentID} [get]
func (p *presenter) GetByDepartmentID(ctx *gin.Context) {
	departmentID := ctx.Param("DepartmentID")
	input := &personnel_affiliation.Field{}
	input.DepartmentID = departmentID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PersonnelAffiliationResolver.GetByDepartmentID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByParentDepartmentID
// @Summary PersonnelAffiliation.5 取得部門的人員隸屬(包含同個父部門底下的所有子部門)
// @description 取得部門的人員隸屬(包含同個父部門底下的所有子部門)
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param DepartmentID path string true "部門ID"
// @success 200 object code.SuccessfulMessage{body=[]personnel_affiliation.Deparment_User} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/GetByParentDepartmentID/{DepartmentID} [get]
func (p *presenter) GetByParentDepartmentID(ctx *gin.Context) {
	departmentID := ctx.Param("DepartmentID")
	input := &personnel_affiliation.Field{}
	input.DepartmentID = departmentID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PersonnelAffiliationResolver.GetByParentDepartmentID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary PersonnelAffiliation.d 刪除單一人員隸屬
// @description 刪除單一人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PaID path string true "人員隸屬ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/{PaID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	paID := ctx.Param("PaID")
	created_account := ctx.Param("created_account")

	input := &personnel_affiliation.Updated{}
	input.PaID = paID
	input.Account = created_account

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PersonnelAffiliationResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary PersonnelAffiliation.u 更新單一人員隸屬
// @description 更新單一人員隸屬
// @Tags PersonnelAffiliation
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param PaID path string true "人員隸屬ID"
// @param * body personnel_affiliation.Updated true "更新人員隸屬"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/personnel_affiliation/{PaID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	paID := ctx.Param("PaID")
	created_account := ctx.Param("created_account")

	input := &personnel_affiliation.Updated{}
	input.PaID = paID
	input.Account = created_account

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	codeMessage := p.PersonnelAffiliationResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
