package quotation

import (
	"eirc.app/internal/v1/resolver/quotation"
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
	GetByQIDQuotationDetailListUser(ctx *gin.Context)
	QuotationDetailListUser(ctx *gin.Context)
}

type presenter struct {
	QuotationResolver quotation.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		QuotationResolver: quotation.New(db),
	}
}
