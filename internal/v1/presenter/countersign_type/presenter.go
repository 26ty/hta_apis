package countersign_type

import (
	"eirc.app/internal/v1/resolver/countersign_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	CountersignTypeResolver countersign_type.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		CountersignTypeResolver: countersign_type.New(db),
	}
}
