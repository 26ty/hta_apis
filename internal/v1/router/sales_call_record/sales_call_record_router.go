package sales_call_record

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/sales_call_record"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("SalesCallRecord")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":sID", controller.GetByID)
		v10.DELETE(":sID", controller.Delete)
		v10.PATCH(":sID", controller.Updated)
	}

	return route
}
