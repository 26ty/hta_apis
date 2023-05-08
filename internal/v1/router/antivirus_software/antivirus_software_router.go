package antivirus_software

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/antivirus_software"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("antivirus_software")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":AsID", controller.GetByID)
		v10.GET("/GetByCaseIDtoTop/:account/:userID", controller.GetByCaseIDtoTop)
		v10.GET("/GetByPIDList/:PID", controller.GetByPIDList)
		v10.DELETE(":AsID", controller.Delete)
		v10.PATCH(":AsID", controller.Updated)
		v10.PATCH("AsReviewTask/:account/:taskID", controller.AsReviewTask)

	}

	return route
}
