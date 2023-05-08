package factory

import (
	"eirc.app/internal/v1/resolver/factory"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
	GetByFIDFLMListUser(ctx *gin.Context)
	FLMListUser(ctx *gin.Context)
	SearchFactory(ctx *gin.Context)
}

type presenter struct {
	FactoryResolver factory.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		FactoryResolver: factory.New(db),
	}
}
