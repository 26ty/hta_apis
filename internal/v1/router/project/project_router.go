package project

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/project"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("project")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/GetB2BonitaCaseListStop/:account/:caseID", controller.GetB2BonitaCaseListStop)
		v10.GET("/GetB2BonitaCaseListTask/:account/:caseID", controller.GetB2BonitaCaseListTask)
		v10.GET("/GetB2BonitaCaseListPM/:account/:userID", controller.GetB2BonitaCaseListPM)
		v10.GET("/GetB2BonitaCaseListTop/:account/:userID", controller.GetB2BonitaCaseListTop)
		v10.GET("/GetB2BonitaCaseListTM/:account/:userID", controller.GetB2BonitaCaseListTM)
		v10.GET("/GetB2BonitaCaseListCountersign/:account/:userID", controller.GetB2BonitaCaseListCountersign)
		v10.GET("/GetB2BonitaCaseListConfirm/:account/:userID", controller.GetB2BonitaCaseListConfirm)
		v10.GET("/GetB2BonitaCaseListDepartment/:account/:userID", controller.GetB2BonitaCaseListDepartment)
		v10.GET("/GetByProjectListUser/:pID", controller.GetByProjectListUser)
		v10.GET("/GetByProjectBonitaUserList/:pID", controller.GetByProjectBonitaUserList)
		v10.GET("/ProjectListUser", controller.ProjectListUser)
		v10.GET("/ProduceSalesListUser", controller.ProduceSalesListUser)
		v10.GET("/ProjectTemplateListUser", controller.ProjectTemplateListUser)
		v10.GET("/ProjectAuthorizationListUser", controller.ProjectAuthorizationListUser)
		v10.GET(":pID", controller.GetByID)
		v10.DELETE(":pID", controller.Delete)
		v10.PATCH("B2UpdatedCaseID/:pID", controller.B2UpdatedCaseID)
		v10.PATCH("B2ReviewTask/:account/:taskID", controller.B2ReviewTask)
		// v10.PATCH("B2TopReviewTask/:account/:taskID", controller.B2TopReviewTask)
		// v10.PATCH("B2DepartmentReviewTask/:account/:taskID", controller.B2DepartmentReviewTask)
		v10.PATCH(":pID", controller.Updated)
	}

	return route
}
