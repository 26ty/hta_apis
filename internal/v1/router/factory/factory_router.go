package factory

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/factory"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("Factory")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":fID", controller.GetByID)
		v10.DELETE(":fID", controller.Delete)
		v10.PATCH(":fID", controller.Updated)
		v10.GET("/GetByFIDFactory/:fID", controller.GetByFIDFLMListUser)
		v10.GET("/GetAllFactory", controller.FLMListUser)
		v10.GET("/SearchFactory", controller.SearchFactory)
	}

	return route
}
