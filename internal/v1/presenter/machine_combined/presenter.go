package machine_combined

import (
	"eirc.app/internal/v1/resolver/machine_combined"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	MachineCombinedListLast(ctx *gin.Context)
	GetByPIDMachineCombinedListLast(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	MachineCombinedResolver machine_combined.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		MachineCombinedResolver: machine_combined.New(db),
	}
}
