package factory_manufacturing

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/factory_manufacturing"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("FactoryManufacturing")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":fmID", controller.GetByID)
		v10.DELETE(":fmID", controller.Delete)
		v10.PATCH(":fmID", controller.Updated)
	}

	return route
}
