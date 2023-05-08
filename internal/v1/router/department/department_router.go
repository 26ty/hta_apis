package department

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/department"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("department")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/departmentUserList", controller.DepartmentAccountList)
		v10.GET("/A1Department", controller.A1Department)
		v10.GET("/AllDepartment", controller.AllDepartment)
		v10.GET(":DID", controller.GetByID)
		v10.DELETE(":Account/:DID", controller.Delete)
		v10.PATCH(":DID", controller.Updated)
	}

	return route
}
