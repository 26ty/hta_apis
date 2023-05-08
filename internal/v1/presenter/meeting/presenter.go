package meeting

import (
	"eirc.app/internal/v1/resolver/meeting"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByUserIDMeetingListUser(ctx *gin.Context) 
	GetByMIDMeetingListUser(ctx *gin.Context)
	GetByDIDMeetingListUser(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByMIDMeetingUser(ctx *gin.Context)
	MeetingUser(ctx *gin.Context) 
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	MeetingResolver meeting.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		MeetingResolver: meeting.New(db),
	}
}
