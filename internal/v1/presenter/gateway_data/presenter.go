package gateway_data

import (
	"eirc.app/internal/v1/resolver/gateway_data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByClassificationTitle(ctx *gin.Context)
	GetByDataDemand(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	GatewayDataResolver gateway_data.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		GatewayDataResolver: gateway_data.New(db),
	}
}
