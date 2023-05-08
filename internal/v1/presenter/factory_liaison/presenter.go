package factory_liaison

import (
	"eirc.app/internal/v1/resolver/factory_liaison"
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
	FactoryLiaisonResolver factory_liaison.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		FactoryLiaisonResolver: factory_liaison.New(db),
	}
}
