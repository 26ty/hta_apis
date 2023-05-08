package task_user

import (
	"eirc.app/internal/v1/resolver/task_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByDocumnetIDListHour(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetName(ctx *gin.Context)
	Delete(ctx *gin.Context) 
	DeleteList(ctx *gin.Context)
	Updated(ctx *gin.Context)
	UpdatedStatus(ctx *gin.Context)
}

type presenter struct {
	TaskUserResolver task_user.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		TaskUserResolver: task_user.New(db),
	}
}
