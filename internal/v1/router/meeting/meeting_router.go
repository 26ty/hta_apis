package meeting

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/meeting"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("meeting")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET("/GetByUserIDMeetingListUser/:UserID", controller.GetByUserIDMeetingListUser)
		v10.GET("/GetByMIDMeetingListUser/:MID", controller.GetByMIDMeetingListUser)
		v10.GET("/GetByDIDMeetingListUser/:DocumentsID", controller.GetByDIDMeetingListUser)
		v10.GET("/GetByMIDMeetingUser/:MID", controller.GetByMIDMeetingUser)
		v10.GET("/MeetingUser", controller.MeetingUser)
		v10.GET(":MID", controller.GetByID)
		v10.DELETE(":MID", controller.Delete)
		v10.PATCH(":MID", controller.Updated)
	}

	return route
}
