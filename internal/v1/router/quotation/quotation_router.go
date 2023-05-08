package quotation

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/quotation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("Quotation")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":qID", controller.GetByID)
		v10.DELETE(":qID", controller.Delete)
		v10.PATCH(":qID", controller.Updated)
		v10.PATCH("UpdatedCaseID/:qID", controller.UpdatedCaseID)
		v10.PATCH("ReviewTask/:account/:taskID", controller.ReviewTask)
		v10.GET("/GetByQIDDetail/:qID", controller.GetByQIDQuotationDetailListUser)
		v10.GET("/GetAllQuotationDetail", controller.QuotationDetailListUser)
	}

	return route
}
