package project_template
//目前無使用
import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"

	//"eirc.app/internal/pkg/util"
	"net/http"

	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/project_template"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &project_template.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectTemplateResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) List(ctx *gin.Context) {
	input := &project_template.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.ProjectTemplateResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) GetByID(ctx *gin.Context) {
	ptID := ctx.Param("PtID")
	input := &project_template.Field{}
	input.PtID = ptID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectTemplateResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	ptID := ctx.Param("PtID")
	input := &project_template.Updated{}
	input.PtID = ptID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectTemplateResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	ptID := ctx.Param("PtID")
	input := &project_template.Updated{}
	input.PtID = ptID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.ProjectTemplateResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
