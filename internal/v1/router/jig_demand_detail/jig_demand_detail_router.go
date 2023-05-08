package jig_demand_detail

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/jig_demand_detail"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("JigDemandDetail")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":jdID", controller.GetByID)
		v10.DELETE(":jdID", controller.Delete)
		v10.PATCH(":jdID", controller.Updated)
		v10.PATCH("UpdatedByJigID/:jigID", controller.UpdatedByJigID)
	}

	return route
}
