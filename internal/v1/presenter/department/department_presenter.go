package department

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	bpm "eirc.app/internal/pkg/bpm"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/department"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary Department.c 新增部門
// @description 新增部門
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body department.Created true "新增部門"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &department.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	bonita_group_id := bpm.AddDepartmentGetID(ctx, input)
	if bonita_group_id != "" {

		input.BonitaGroupID = bonita_group_id

		codeMessage := p.DepartmentResolver.Created(trx, input)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// List
// @Summary Department.1 條件搜尋部門
// @description 條件部門
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param manager query string false "主管ID"
// @param name query string false "部門名稱"
// @param eng_name query string false "部門英文名稱"
// @param fax query string false "fax電話"
// @param tel query string false "電話"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=department.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &department.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.DepartmentResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// A1Department
// @Summary Department.3 A1會簽用部門
// @description A1會簽用部門(先進技術部、Vision、Motion、機械研發部、電控)
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=department.Deparment_Users} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department/A1Department [get]
func (p *presenter) A1Department(ctx *gin.Context) {
	input := &department.Field{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}


	codeMessage := p.DepartmentResolver.A1Department(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// AllDepartment
// @Summary Department.4 全部部門與底下員工
// @description 全部部門與底下員工
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=department.Deparment_Users} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department/AllDepartment [get]
func (p *presenter) AllDepartment(ctx *gin.Context) {
	input := &department.Field{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}


	codeMessage := p.DepartmentResolver.AllDepartment(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// DepartmentAccountList
// @Summary Department.5 全部部門並顯示主管名稱
// @description 全部部門並顯示主管名稱
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=department.Deparment_Accounts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department/DepartmentAccountList [get]
func (p *presenter) DepartmentAccountList(ctx *gin.Context) {
	input := &department.Users{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.DepartmentResolver.DepartmentAccountList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary Department.2 取得單一部門
// @description 取得單一部門
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param DID path string true "部門ID"
// @success 200 object code.SuccessfulMessage{body=department.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department/{DID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	dID := ctx.Param("DID")
	input := &department.Field{}
	input.DID = dID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.DepartmentResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary Department.d 刪除單一部門
// @description 刪除單一部門
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param DID path string true "部門ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department/{DID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	dID := ctx.Param("DID")
	account := ctx.Param("Account")
	input := &department.Updated{}
	input.DID = dID
	input.Account = account
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.DepartmentResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary Department.u 更新單一部門
// @description 更新單一部門
// @Tags Department
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param DID path string true "部門ID"
// @param * body department.Updated true "更新部門"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/department/{DID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	dID := ctx.Param("DID")
	input := &department.Updated{}
	input.DID = dID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	codeMessage := p.DepartmentResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
