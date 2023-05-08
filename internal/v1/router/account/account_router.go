package account

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/account"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("account")
	{
		v10.POST(":create_account",middleware.Verify(), middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/GetAccountNameList", controller.AccountNameList)
		v10.GET("/AccountNameDepartmentList", controller.AccountNameDepartmentList)
		v10.GET(":accountID", controller.GetByID)
		v10.DELETE(":accountID/:created_account", controller.Delete)
		v10.PATCH(":accountID/:created_account", controller.Updated)
		//為了導入EMAIL用的(暫時寫死)
		//v10.PATCH("UpdatedCsv/:filename/:created_account", controller.UpdatedCsv)
	}

	return route
}
