package labor_hour_modify

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/labor_hour_modify"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("labor_hour_modify")
	{
		v10.POST(":account", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":HmID", controller.GetByID)
		v10.GET("/GetLaborBonitaCaseListStart/:account/:caseID", controller.GetLaborBonitaCaseListStart)
		v10.GET("/GetLaborBonitaCaseListDepartment/:account/:userID", controller.GetLaborBonitaCaseListDepartment)
		v10.GET("/GetByUserIdLaborHourModifyList/:userID/:tID", controller.GetByUserIdLaborHourModifyList)
		v10.GET("/GetByCuIdLaborHourModifyList/:cuID", controller.GetByCuIdLaborHourModifyList)
		v10.GET("/GetByUserIdList/:userID", controller.GetByUserIdList)
		v10.DELETE(":HmID", controller.Delete)
		v10.PATCH(":HmID", controller.Updated)
		v10.PATCH("/UpdatedStatus/:HmID", controller.UpdatedStatus)
		v10.PATCH("LaborReviewTask/:account/:taskID", controller.LaborReviewTask)
	}

	return route
}
