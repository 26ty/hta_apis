package plug_in

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/plug_in"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("plug_in")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/GetByPIDList/:ProjectID", controller.GetByPIDList)
		v10.GET(":PiID", controller.GetByID)
		v10.DELETE(":PiID", controller.Delete)
		v10.PATCH(":PiID", controller.Updated)
	}

	return route
}
