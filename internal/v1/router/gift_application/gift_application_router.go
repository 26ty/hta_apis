package gift_application

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/gift_application"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("GiftApplication")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":gID", controller.GetByID)
		v10.DELETE(":gID", controller.Delete)
		v10.PATCH(":gID", controller.Updated)
		v10.PATCH("UpdatedCaseID/:gID", controller.DUpdatedCaseID)
		v10.PATCH("ReviewTask/:account/:taskID", controller.DReviewTask)
		v10.GET("/GetByGIDDetail/:gID", controller.GetByGIDGiftDetailListUser)
		v10.GET("/GetAllGiftApplication", controller.GiftDetailListUser)
		v10.GET("/GetDBonitaCaseListRestart/:account/:caseID", controller.GetDBonitaCaseListRestart)
		v10.GET("/GetDBonitaCaseListDepartment/:account/:userID", controller.GetDBonitaCaseListDepartment)
		v10.GET("/GetDBonitaCaseListViceTop/:account/:userID", controller.GetDBonitaCaseListViceTop)
		v10.GET("/GetDBonitaCaseListTop/:account/:userID", controller.GetDBonitaCaseListTop)
		v10.GET("/GetDBonitaCaseListAttm/:account/:userID", controller.GetDBonitaCaseListAttm)
		v10.GET("/GetDBonitaTaskId/:account/:caseID", controller.GetDBonitaTaskId)
	}

	return route
}
