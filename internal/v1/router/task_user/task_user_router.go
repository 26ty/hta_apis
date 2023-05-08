package task_user

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/task_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("task_user")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":tuID", controller.GetByID)
		v10.GET("GetName/:tuID", controller.GetName)
		v10.GET("GetByDocumnetIDListHour/:documentsID", controller.GetByDocumnetIDListHour)
		v10.DELETE("/:tuID", controller.Delete)
		v10.DELETE("", controller.DeleteList)
		v10.PATCH("", controller.Updated)
		v10.PATCH("UpdatedStatus/:tuID", controller.UpdatedStatus)
	}

	return route
}
