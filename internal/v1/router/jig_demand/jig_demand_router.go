package jig_demand

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/jig_demand"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("JigDemand")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":jID", controller.GetByID)
		v10.DELETE(":jID", controller.Delete)
		v10.PATCH(":jID", controller.Updated)
		v10.PATCH("UpdatedCaseID/:jID", controller.UpdatedCaseID)
		v10.PATCH("ReviewTask/:account/:taskID", controller.ReviewTask)
		v10.GET("/GetByJIDDetail/:jID", controller.GetByJIDJigDetailListUser)
		v10.GET("/GetAllJigDetail", controller.JigDetailListUser)
		v10.GET("/SearchJigDemand", controller.SearchJigDemand)
		v10.GET("/GetByUserIDListJD/:userID", controller.GetByUserIDListJD)
	}

	return route
}
