package jig_demand

import (
	"eirc.app/internal/v1/resolver/jig_demand"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
	UpdatedCaseID(ctx *gin.Context)
	ReviewTask(ctx *gin.Context)
	GetByJIDJigDetailListUser(ctx *gin.Context)
	JigDetailListUser(ctx *gin.Context)
	SearchJigDemand(ctx *gin.Context)
	GetByUserIDListJD(ctx *gin.Context)
}

type presenter struct {
	JigDemandResolver jig_demand.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		JigDemandResolver: jig_demand.New(db),
	}
}
