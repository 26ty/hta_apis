package countersign

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/countersign"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("countersign")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":CsID", controller.GetByID)
		v10.DELETE(":CsID", controller.Delete)
		v10.PATCH(":CsID", controller.Updated)
	}

	return route
}
