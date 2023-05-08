package customer

import (
	"eirc.app/internal/v1/resolver/customer"
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
	CustomerResolver customer.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		CustomerResolver: customer.New(db),
	}
}
