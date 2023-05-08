package task

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/task"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("task")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/GetByDocumentsTaskListUser/:documentsID", controller.TaskListUser)
		v10.GET("/GetByIDListTaskHour/:documentsID/:userID", controller.GetByIDListTaskHour)
		v10.GET("/GetByTaskListUser/:tID", controller.GetByTaskListUser)
		v10.GET(":tID", controller.GetByID)
		v10.GET("/GetByIDTaskBonitaUserList/:documentsID", controller.GetByIDTaskBonitaUserList)
		v10.GET("/GetByOriginIDAndUserID/:originID/:userID", controller.GetByOriginIDAndUserID)
		v10.GET("/GetByTaskListHourDocumentsID/:documentsID", controller.GetTaskListHourByUserID)
		v10.GET("/GetByTaskListHourDocumentsAndUserID/:documentsID/:accountID", controller.GetByTaskListHourDocumentsAndUserID)
		v10.GET("/GetByDocumentIDTaskList/:documentsID", controller.GetByDocumentIDTaskList)
		v10.DELETE("/:tID", controller.Delete)
		v10.DELETE("", controller.DeleteList)
		v10.PATCH("", controller.Updated)
	}

	return route
}
