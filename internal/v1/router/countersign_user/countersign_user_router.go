package countersign_user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/countersign_user"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("countersign_user")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("GetByIDCountersignUserListUser/:DocumentsID", controller.GetByIDCountersignUserListUser)
		v10.GET("GetByIDCountersignUserListUser/:DocumentsID/:CountersignID", controller.GetByIDCountersignUserListUser2)
		v10.GET("GetByCuIDCountersignUserListUser/:CuID", controller.GetByCuIDCountersignUserListUser)
		v10.GET(":CuID", controller.GetByID)
		v10.DELETE(":CuID", controller.Delete)
		v10.PATCH(":CuID", controller.Updated)
	}

	return route
}
