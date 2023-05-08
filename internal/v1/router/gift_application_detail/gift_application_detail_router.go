package gift_application_detail

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/gift_application_detail"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("GiftApplicationDetail")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":gdID", controller.GetByID)
		v10.DELETE(":gdID", controller.Delete)
		v10.PATCH(":gdID", controller.Updated)
	}

	return route
}
