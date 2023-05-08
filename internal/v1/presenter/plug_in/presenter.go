package plug_in

import (
	"eirc.app/internal/v1/resolver/plug_in"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByPIDList(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	PlugInResolver plug_in.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		PlugInResolver: plug_in.New(db),
	}
}
