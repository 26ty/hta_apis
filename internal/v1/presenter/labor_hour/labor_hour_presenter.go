package labor_hour

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"time"
	//"eirc.app/internal/pkg/util"
	"net/http"

	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/labor_hour"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary LaborHour.c 新增工時
// @description 新增工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body labor_hour.Created true "新增工時"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &labor_hour.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary LaborHour.1 條件搜尋工時
// @description 條件工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param title query string false "標題"
// @param content query string false "內容"
// @param nature query string false "類別"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=labor_hour.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &labor_hour.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.LaborHourResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary LaborHour.2 取得單一工時
// @description 取得單一工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HID path string true "工時ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/{HID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	hID := ctx.Param("HID")
	input := &labor_hour.Field{}
	input.HID = hID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIdLaborHourList
// @Summary LaborHour.6 篩選出該任務該使用者已提報工時列表
// @description 用t_id、user_id篩選出該任務該使用者已提報工時列表
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @param tID path string true "任務ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/GetByUserIdLaborHourList/{userID}/{tID} [get]
func (p *presenter) GetByUserIdLaborHourList(ctx *gin.Context) {
	input := &labor_hour.Field{}
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

	codeMessage := p.LaborHourResolver.GetByUserIdLaborHourList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByCuIdLaborHourList
// @Summary LaborHour.3 篩選出該會簽已提報工時列表
// @description 用cu_id篩選出該會簽已提報工時列表
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param cuID path string true "會簽ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/GetByCuIdLaborHourList/{cuID} [get]
func (p *presenter) GetByCuIdLaborHourList(ctx *gin.Context) {
	input := &labor_hour.Field{}
	// userID := ctx.Param("userID")
	// input.UserID = userID
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

	codeMessage := p.LaborHourResolver.GetByCuIdLaborHourList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIdCategoryList
// @Summary LaborHour.5 篩選取得該使用者的工時
// @description UserID&Category&DateForStart篩選取得該使用者的工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @param category path string true "類型ID"
// @param firstDate path string true "單月的第一天"
// @success 200 object code.SuccessfulMessage{body=[]labor_hour.GetUserCategoryLabor} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/GetByUserIdCategoryList/{userID}/{category}/{firstDate} [get]
func (p *presenter) GetByUserIdCategoryList(ctx *gin.Context) {
	input := &labor_hour.Field{}
	userID := ctx.Param("userID")
	input.UserID = userID
	category := ctx.Param("category")
	input.Category = category
	stringFirstDate := ctx.Param("firstDate")
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	timeFirstDate,err_date := time.Parse("2006-01-02", stringFirstDate)
	if err_date != nil {
		log.Error(err_date)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err_date.Error()))

		return
	}
	input.DateForStart = timeFirstDate
	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.LaborHourResolver.GetByUserIdCategoryList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIdCategory
// @Summary LaborHour.4 篩選取得該使用者的工時
// @description UserID&Category篩選取得該使用者的工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @param category path string true "類型ID"
// @success 200 object code.SuccessfulMessage{body=[]labor_hour.GetUserCategoryLabor} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/GetByUserIdCategory/{userID}/{category} [get]
func (p *presenter) GetByUserIdCategory(ctx *gin.Context) {
	input := &labor_hour.Field{}
	userID := ctx.Param("userID")
	input.UserID = userID
	category := ctx.Param("category")
	input.Category = category
	//stringFirstDate := ctx.Param("firstDate")
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	//timeFirstDate,err_date := time.Parse("2006-01-02", stringFirstDate)
	// if err_date != nil {
	// 	log.Error(err_date)
	// 	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err_date.Error()))

	// 	return
	// }
	//input.DateForStart = timeFirstDate
	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.LaborHourResolver.GetByUserIdCategory(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIdMonthList
// @Summary LaborHour.8 取得該使用者單月的工時表
// @description UserID篩選取得該使用者單月的工時表
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @param firstDate path string true "單月的第一天"
// @success 200 object code.SuccessfulMessage{body=labor_hour.GetUserAllLabors} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/GetByUserIdMonthList/{userID}/{firstDate} [get]
func (p *presenter) GetByUserIdMonthList(ctx *gin.Context) {
	input := &labor_hour.Field_Month{}
	userID := ctx.Param("userID")
	input.UserID = userID
	stringFirstDate := ctx.Param("firstDate")
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	timeFirstDate,err_date := time.Parse("2006-01-02", stringFirstDate)
	if err_date != nil {
		log.Error(err_date)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err_date.Error()))

		return
	}
	input.FirstDate = timeFirstDate
	// if input.Limit >= preset.DefaultLimit {
	// 	input.Limit = preset.DefaultLimit
	// }

	codeMessage := p.LaborHourResolver.GetByUserIdMonthList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByUserIdList
// @Summary LaborHour.7 取得該使用者的工時
// @description UserID篩選取得該使用者的工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param userID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=labor_hour.LaborHours} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/GetByUserIdList/{userID} [get]
func (p *presenter) GetByUserIdList(ctx *gin.Context) {
	input := &labor_hour.Field{}
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

	codeMessage := p.LaborHourResolver.GetByUserIdList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary LaborHour.d 刪除單一工時
// @description 刪除單一工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HID path string true "工時ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/{HID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	hID := ctx.Param("HID")
	input := &labor_hour.Updated{}
	input.HID = hID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary LaborHour.u 更新單一工時
// @description 更新單一工時
// @Tags LaborHour
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param HID path string true "工時ID"
// @param * body labor_hour.Updated true "更新工時"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/labor_hour/{HID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	hID := ctx.Param("HID")
	input := &labor_hour.Updated{}
	input.HID = hID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.LaborHourResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
