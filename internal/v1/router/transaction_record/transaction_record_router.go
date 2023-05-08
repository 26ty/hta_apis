package transaction_record

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/transaction_record"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("transaction_record")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/GetByDocumentIDUserList/:DocumentID", controller.GetByDocumentIDUserList)
		v10.GET(":TrID", controller.GetByID)
		v10.DELETE(":TrID", controller.Delete)
		v10.PATCH(":TrID", controller.Updated)
	}

	return route
}
