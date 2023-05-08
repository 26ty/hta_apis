package login

import (
	"eirc.app/internal/v1/presenter/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := login.New(db)
	route.Group("authority").Group("v1.0").Group("login").
		POST("web", controller.Web)
	route.Group("authority").Group("v1.0").Group("refresh").
		POST("", controller.Refresh)
	route.Group("authority").Group("v1.0").
		GET("GetBonitaCaseCount/:account/:userID", controller.GetBonitaCaseCount)
	route.Group("authority").Group("v1.0").
		GET("GetBonitaCaseDetail/:account/:userID", controller.GetBonitaCaseDetail)
	route.Group("authority").Group("v1.0").
		PATCH("BonitaTransferTask/:account", controller.BonitaTransferTask)
	route.Group("authority").Group("v1.0").
		POST("SendEmail", controller.SendEmail)
	return route
}
