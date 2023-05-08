package sales_call_record

import (
	"eirc.app/internal/v1/resolver/sales_call_record"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	SalesCallRecordResolver sales_call_record.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		SalesCallRecordResolver: sales_call_record.New(db),
	}
}
