package todo_type

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/todo_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("todo_type")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":ttID", controller.GetByID)
		v10.GET("GetByUserID/:userID", controller.GetByUserID)
		v10.DELETE(":ttID", controller.Delete)
		v10.PATCH(":ttID", controller.Updated)
	}

	return route
}
