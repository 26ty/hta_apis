package company
//目前無使用
import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/companies"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)


// Created
// @Summary Companies.c 新增公司行號
// @description 新增公司行號
// @Tags Companies
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body companies.Created true "新增公司行號"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/companies [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	createBy := util.GenerateUUID()
	input := &companies.Created{}
	input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CompanyResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary Companies.1 條件搜尋公司行號
// @description 條件公司行號
// @Tags Companies
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param name query string false "公司名稱"
// @param uniform_number query string false "公司統一編號"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=companies.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/companies [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &companies.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CompanyResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Companies.2 取得單一公司行號
// @description 取得單一公司行號
// @Tags Companies
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param company_id path string true "公司編號"
// @success 200 object code.SuccessfulMessage{body=companies.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/companies/{companyID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	companyID := ctx.Param("companyID")
	input := &companies.Field{}
	input.CompanyID = companyID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CompanyResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Companies.d 刪除公司
// @description 刪除公司
// @Tags Companies
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param company_id path string true "公司編號"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/companies/{companyID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	companyID := ctx.Param("companyID")
	input := &companies.Updated{}
	input.CompanyID = companyID
	input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CompanyResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Companies.u 更新單一使用者
// @description 更新單一使用者
// @Tags Companies
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param companyID path string true "公司編號"
// @param * body companies.Updated true "更新公司"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/companies/{companyID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	companyID := ctx.Param("companyID")
	input := &companies.Updated{}
	input.CompanyID = companyID
	input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CompanyResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
