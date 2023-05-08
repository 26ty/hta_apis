package factory_liaison

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/factory_liaison"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("FactoryLiaison")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":flID", controller.GetByID)
		v10.DELETE(":flID", controller.Delete)
		v10.PATCH(":flID", controller.Updated)
	}

	return route
}
