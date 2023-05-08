package labor_hour

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/labor_hour"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("labor_hour")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":HID", controller.GetByID)
		v10.GET("/GetByUserIdLaborHourList/:userID/:tID", controller.GetByUserIdLaborHourList)
		v10.GET("/GetByCuIdLaborHourList/:cuID", controller.GetByCuIdLaborHourList)
		v10.GET("/GetByUserIdList/:userID", controller.GetByUserIdList)
		v10.GET("/GetByUserIdCategoryList/:userID/:category/:firstDate", controller.GetByUserIdCategoryList)
		v10.GET("/GetByUserIdCategory/:userID/:category", controller.GetByUserIdCategory)
		v10.GET("/GetByUserIdMonthList/:userID/:firstDate", controller.GetByUserIdMonthList)
		v10.DELETE(":HID", controller.Delete)
		v10.PATCH(":HID", controller.Updated)
	}

	return route
}
