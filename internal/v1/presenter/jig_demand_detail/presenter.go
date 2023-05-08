package jig_demand_detail

import (
	"eirc.app/internal/v1/resolver/jig_demand_detail"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
	UpdatedByJigID(ctx *gin.Context)
}

type presenter struct {
	JigDemandDetailResolver jig_demand_detail.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		JigDemandDetailResolver: jig_demand_detail.New(db),
	}
}
