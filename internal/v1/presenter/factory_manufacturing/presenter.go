package factory_manufacturing

import (
	"eirc.app/internal/v1/resolver/factory_manufacturing"
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
	FactoryManufacturingResolver factory_manufacturing.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		FactoryManufacturingResolver: factory_manufacturing.New(db),
	}
}
