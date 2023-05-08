package personnel_affiliation

import (
	"eirc.app/internal/v1/resolver/personnel_affiliation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
	GetByDepartmentID(ctx *gin.Context) 
	GetByParentDepartmentID(ctx *gin.Context) 
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	PersonnelAffiliationResolver personnel_affiliation.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		PersonnelAffiliationResolver: personnel_affiliation.New(db),
	}
}
