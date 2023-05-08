package antivirus_software

import (
	"eirc.app/internal/v1/resolver/antivirus_software"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByCaseIDtoTop(ctx *gin.Context)
	GetByPIDList(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
	AsReviewTask(ctx *gin.Context)
}

type presenter struct {
	AntivirusSoftwareResolver antivirus_software.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		AntivirusSoftwareResolver: antivirus_software.New(db),
	}
}
