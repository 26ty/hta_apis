package gateway_data

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/gateway_data"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("gateway_data")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":GdID", controller.GetByID)
		v10.GET("/GetByDataDemand/:account/:userID/:gdID", controller.GetByDataDemand)
		v10.POST("/GetByClassificationTitle", controller.GetByClassificationTitle)
		v10.DELETE(":GdID", controller.Delete)
		v10.PATCH(":GdID", controller.Updated)
	}

	return route
}
