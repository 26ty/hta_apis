package project_template

import (
	"eirc.app/internal/v1/resolver/project_template"
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
	ProjectTemplateResolver project_template.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ProjectTemplateResolver: project_template.New(db),
	}
}
