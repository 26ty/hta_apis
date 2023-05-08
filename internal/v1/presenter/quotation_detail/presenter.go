package quotation_detail

import (
	"eirc.app/internal/v1/resolver/quotation_detail"
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
	QuotationDetailResolver quotation_detail.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		QuotationDetailResolver: quotation_detail.New(db),
	}
}
