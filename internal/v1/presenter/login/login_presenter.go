package login

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	"net/http"

	"eirc.app/internal/v1/resolver/login"
	"eirc.app/internal/v1/structure/jwe"
	"eirc.app/internal/v1/structure/logins"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bpm "eirc.app/internal/pkg/bpm"
	shared "eirc.app/internal/pkg/shared"
	model "eirc.app/internal/v1/structure"
)

type Presenter interface {
	Web(ctx *gin.Context)
	Refresh(ctx *gin.Context)
	GetBonitaCaseCount(ctx *gin.Context)
	GetBonitaCaseDetail(ctx *gin.Context)
	BonitaTransferTask(ctx *gin.Context)
	SendEmail(ctx *gin.Context)
}

type presenter struct {
	Login login.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		Login: login.New(db),
	}
}

// Web
// @Summary Login.2 使用者登入
// @description 使用者登入
// @Tags Login
// @version 1.0
// @Accept json
// @produce json
// @param * body logins.Login true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/login/web [post]
func (p *presenter) Web(ctx *gin.Context) {
	input := &logins.Login{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := p.Login.Web(input)
	// client := bpm.Bc
	// client.Login(input.Account)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Refresh
// @Summary Login.1 換新的令牌
// @description 換新的令牌
// @Tags Login
// @version 1.0
// @Accept json
// @produce json
// @param * body jwe.Refresh true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/login/refresh [post]
func (p *presenter) Refresh(ctx *gin.Context) {
	input := &jwe.Refresh{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := p.Login.Refresh(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetBonitaCaseCount
// @Summary Other.1 計算待審核單據
// @description 計算待審核單據
// @Tags Other
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]model.GonitaListCount} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/GetBonitaCaseCount/{account}/{userID} [get]
func (p *presenter) GetBonitaCaseCount(ctx *gin.Context) {
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
		codeMessage := p.Login.GetBonitaCaseCount(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// GetBonitaCaseDetail
// @Summary Other.2 待審核單據清單(任務轉移)
// @description 計算待審核單據清單(任務轉移)
// @Tags Other
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param userID path string true "bonita使用者ID"
// @success 200 object code.SuccessfulMessage{body=[]model.GetCaseDetailListOutput} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/GetBonitaCaseDetail/{account}/{userID} [get]
func (p *presenter) GetBonitaCaseDetail(ctx *gin.Context) {
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
		codeMessage := p.Login.GetBonitaCaseDetail(*output)
		ctx.JSON(http.StatusOK, codeMessage)
	}

}

// BonitaTransferTask
// @Summary Other.4 任務轉移
// @description 任務轉移
// @Tags Other
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param account path string true "帳號"
// @param * body model.TransferTaskInput true "任務轉移內容"
// @success 200 object code.SuccessfulMessage{body=int} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/BonitaTransferTask/{account} [patch]
func (p *presenter) BonitaTransferTask(ctx *gin.Context) {

	account := ctx.Param("account")

	input := &model.TransferTaskInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}
	for _, value := range input.TransferTask {
		body := bpm.TransferTask(ctx, account, value.BonitaTaskID, value.BonitaUserID)

		if body != 200 {
			ctx.JSON(http.StatusOK, code.GetCodeMessage(body, "bonita updateAssignedId error"))

			return
		}
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, 200))

}

// SendEmail
// @Summary Other.3 寄送EMAIL
// @description 寄送EMAIL
// @Tags Other
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body model.SendEmailInput true "寄送EMAIL內容"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/SendEmail [post]
func (p *presenter) SendEmail(ctx *gin.Context) {
	input := &model.SendEmailInput{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := shared.SendEmail(input.Host, input.Port, input.Name, input.Username, input.Password, input.To, input.Subject, input.Body)
	// client := bpm.Bc
	// client.Login(input.Account)
	if codeMessage != nil {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, codeMessage))
		return
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, "發送成功!"))
}
