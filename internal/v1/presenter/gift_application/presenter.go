package gift_application

import (
	"eirc.app/internal/v1/resolver/gift_application"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
	DUpdatedCaseID(ctx *gin.Context)
	DReviewTask(ctx *gin.Context)
	GetByGIDGiftDetailListUser(ctx *gin.Context)
	GiftDetailListUser(ctx *gin.Context)
	GetDBonitaCaseListRestart(ctx *gin.Context)
	GetDBonitaCaseListDepartment(ctx *gin.Context)
	GetDBonitaCaseListViceTop(ctx *gin.Context)
	GetDBonitaCaseListTop(ctx *gin.Context)
	GetDBonitaCaseListAttm(ctx *gin.Context)
	GetDBonitaTaskId(ctx *gin.Context)
}

type presenter struct {
	GiftApplicationResolver gift_application.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		GiftApplicationResolver: gift_application.New(db),
	}
}
