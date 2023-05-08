package machine_combined

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/machine_combined"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("machine_combined")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/MachineCombinedListLast", controller.MachineCombinedListLast)
		v10.GET("/GetByPIDMachineCombinedListLast/:PID", controller.GetByPIDMachineCombinedListLast)
		v10.GET(":McID", controller.GetByID)
		v10.DELETE(":McID", controller.Delete)
		v10.PATCH(":McID", controller.Updated)
	}

	return route
}
