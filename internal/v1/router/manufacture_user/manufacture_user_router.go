package manufacture_user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/manufacture_user"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("manufacture_user")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":MuID", controller.GetByID)
		v10.GET("GetByManufactureID/:ManufactureID", controller.GetByManufactureID)
		v10.DELETE(":MuID", controller.Delete)
		v10.PATCH(":MuID", controller.Updated)
	}

	return route
}
