package countersign_type
//目前無使用
import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	//"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/countersign_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)


func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	//FID := util.GenerateUUID()
	input := &countersign_type.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) List(ctx *gin.Context) {
	input := &countersign_type.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CountersignTypeResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) GetByID(ctx *gin.Context) {
	ctID := ctx.Param("CtID")
	input := &countersign_type.Field{}
	input.CtID = ctID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	ctID := ctx.Param("CtID")
	input := &countersign_type.Updated{}
	input.CtID = ctID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}


func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	ctID := ctx.Param("CtID")
	input := &countersign_type.Updated{}
	input.CtID = ctID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.CountersignTypeResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
