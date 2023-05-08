package file

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/file"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("file")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":fID", controller.GetByID)
		v10.GET("GetByDocumentID/:documentsID", controller.GetByDocumentID)
		v10.GET("GetByDocumentIDUserID/:documentsID/:userID", controller.GetByDocumentIDUserID)
		v10.DELETE(":fID", controller.Delete)
		v10.PATCH(":fID", controller.Updated)
	}

	return route
}
