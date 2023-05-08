package manufacture_order

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/manufacture_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("manufacture_order")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		// v10.GET("/ManufactureOrderCdListUser", controller.ManufactureOrderCdListUser)
		v10.GET("/ManufactureOrderProjectListUser", controller.ManufactureOrderProjectListUser)
		v10.GET(":MID", controller.GetByID)
		v10.GET("/GetByIDOne/:MID", controller.GetByIDOne)
		v10.GET("/GetByPIDList/:PID", controller.GetByPIDList)
		v10.GET("/GetC2BonitaCaseListStart/:account/:caseID", controller.GetC2BonitaCaseListStart)
		v10.GET("/GetC2BonitaCaseListDepartment/:account/:userID", controller.GetC2BonitaCaseListDepartment)
		v10.GET("/GetC2BonitaCaseListManufacture/:account/:userID", controller.GetC2BonitaCaseListManufacture)
		v10.GET("/GetC2BonitaCaseListTop/:account/:userID", controller.GetC2BonitaCaseListTop)
		v10.GET("/GetC2BonitaCaseListConfirm/:account/:userID", controller.GetC2BonitaCaseListConfirm)
		v10.GET("/GetC2BonitaCaseListSave/:account/:userID", controller.GetC2BonitaCaseListSave)
		v10.DELETE(":MID", controller.Delete)
		v10.PATCH("C2UpdatedCaseID/:mID", controller.C2UpdatedCaseID)
		v10.PATCH("C2ReviewTask/:account/:taskID", controller.C2ReviewTask)
		v10.PATCH(":MID", controller.Updated)
	}

	return route
}
