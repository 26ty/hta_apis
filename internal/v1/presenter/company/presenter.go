package company

import (
	"eirc.app/internal/v1/resolver/company"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type presenter struct {
	CompanyResolver company.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		CompanyResolver: company.New(db),
	}
}
