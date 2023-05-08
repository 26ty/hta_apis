package personnel_affiliation

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/personnel_affiliation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("personnel_affiliation")
	{
		v10.POST(":create_account", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":PaID", controller.GetByID)
		v10.GET("GetByDepartmentID/:DepartmentID", controller.GetByDepartmentID)
		v10.GET("GetByParentDepartmentID/:DepartmentID", controller.GetByParentDepartmentID)
		v10.GET("GetByUserID/:UserID", controller.GetByUserID)
		v10.DELETE(":PaID/:created_account", controller.Delete)
		v10.PATCH(":PaID/:created_account", controller.Updated)
	}

	return route
}
