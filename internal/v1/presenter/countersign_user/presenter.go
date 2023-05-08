package countersign_user

import (
	"eirc.app/internal/v1/resolver/countersign_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByIDCountersignUserListUser(ctx *gin.Context)
	GetByIDCountersignUserListUser2(ctx *gin.Context)
	GetByCuIDCountersignUserListUser(ctx *gin.Context) 
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	CountersignUserResolver countersign_user.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		CountersignUserResolver: countersign_user.New(db),
	}
}
