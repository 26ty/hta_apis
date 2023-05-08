package project_template

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/project_template"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("project_template")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":PtID", controller.GetByID)
		v10.DELETE(":PtID", controller.Delete)
		v10.PATCH(":PtID", controller.Updated)
	}

	return route
}
