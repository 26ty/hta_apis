package customer_demand

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/customer_demand"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("customer_demand")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/CustomerDemandListUser", controller.CustomerDemandListUser)
		v10.GET("/GetByCuIDCustomerDemandListUser/:CdId", controller.GetByCuIDCustomerDemandListUser)
		v10.GET("/GetByUserIDListCR/:UserId", controller.GetByUserIDListCR)
		v10.GET("/GetByUserIDListHCR/:UserId", controller.GetByUserIDListHCR)
		v10.GET(":CdId", controller.GetByID)
		v10.GET("/GetA1BonitaCaseListStart/:account/:caseID", controller.GetA1BonitaCaseListStart)
		v10.GET("/GetA1BonitaCaseListTask/:account/:userID/:caseID", controller.GetA1BonitaCaseListTask)
		v10.GET("/GetA1BonitaCaseListDepartment/:account/:userID", controller.GetA1BonitaCaseListDepartment)
		v10.GET("/GetA1BonitaCaseListDirector/:account/:userID", controller.GetA1BonitaCaseListDirector)
		v10.GET("/GetA1BonitaCaseListTop/:account/:userID", controller.GetA1BonitaCaseListTop)
		v10.GET("/GetA1BonitaCaseListDispatch/:account/:userID", controller.GetA1BonitaCaseListDispatch)
		v10.GET("/GetA1BonitaCaseListEvaluation/:account/:userID", controller.GetA1BonitaCaseListEvaluation)
		v10.GET("/GetA1BonitaCaseListCountersign/:account/:userID", controller.GetA1BonitaCaseListCountersign)
		v10.GET("/GetA1BonitaCaseListPMEvaluation/:account/:userID", controller.GetA1BonitaCaseListPMEvaluation)
		v10.GET("/GetA1BonitaCaseListBusiness/:account/:userID", controller.GetA1BonitaCaseListBusiness)
		v10.GET("/GetA1BonitaCaseListBusinessManager/:account/:userID", controller.GetA1BonitaCaseListBusinessManager)
		v10.GET("/GetA1BonitaCaseListBusinessDirector/:account/:userID", controller.GetA1BonitaCaseListBusinessDirector)
		v10.GET("/GetA1BonitaCaseListTaskFinish/:account/:userID", controller.GetA1BonitaCaseListTaskFinish)
		v10.GET("/GetA1BonitaCaseListTaskFinishManager/:account/:userID", controller.GetA1BonitaCaseListTaskFinishManager)
		v10.GET("/GetA1BonitaCaseListBusinessClose/:account/:userID", controller.GetA1BonitaCaseListBusinessClose)
		v10.GET("/GetA1BonitaCaseListDepartmentClose/:account/:userID", controller.GetA1BonitaCaseListDepartmentClose)
		v10.GET("/GetA1BonitaCaseListDirectorClose/:account/:userID", controller.GetA1BonitaCaseListDirectorClose)
		v10.GET("/GetA1BonitaCaseListTopClose/:account/:userID", controller.GetA1BonitaCaseListTopClose)
		v10.GET("/GetA1BonitaCaseListProductionClose/:account/:userID", controller.GetA1BonitaCaseListProductionClose)
		v10.GET("/GetA1BonitaCaseListCountersignClose/:account/:userID", controller.GetA1BonitaCaseListCountersignClose)
		v10.GET("/GetA1BonitaCaseListPMClose/:account/:userID", controller.GetA1BonitaCaseListPMClose)
		v10.DELETE(":CdId", controller.Delete)
		v10.PATCH("A1UpdatedCaseID/:CdId", controller.A1UpdatedCaseID)
		v10.PATCH("A1ReviewTask/:account/:taskID", controller.A1ReviewTask)
		v10.PATCH(":CdId", controller.Updated)
	}

	return route
}
