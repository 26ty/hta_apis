package task

import (
	"eirc.app/internal/v1/resolver/task"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	TaskListUser(ctx *gin.Context)
	GetByIDListTaskHour(ctx *gin.Context)
	GetByTaskListUser(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByIDTaskBonitaUserList(ctx *gin.Context)
	GetByOriginIDAndUserID(ctx *gin.Context)
	GetTaskListHourByUserID(ctx *gin.Context)
	GetByTaskListHourDocumentsAndUserID(ctx *gin.Context)
	GetByTIDTaskListUser(ctx *gin.Context)
	GetByDocumentIDTaskListLast(ctx *gin.Context)
	GetByDocumentIDTaskList(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
	Delete(ctx *gin.Context) 
	Updated(ctx *gin.Context)
}

type presenter struct {
	TaskResolver task.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		TaskResolver: task.New(db),
	}
}
