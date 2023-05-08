package project

import (
	"eirc.app/internal/v1/resolver/project"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetB2BonitaCaseListStop(ctx *gin.Context) 
	GetB2BonitaCaseListTask(ctx *gin.Context) 
	GetB2BonitaCaseListPM(ctx *gin.Context)
	GetB2BonitaCaseListTop(ctx *gin.Context)
	GetB2BonitaCaseListTM(ctx *gin.Context)
	GetB2BonitaCaseListCountersign(ctx *gin.Context)
	GetB2BonitaCaseListConfirm(ctx *gin.Context)
	GetB2BonitaCaseListDepartment(ctx *gin.Context)
	GetByProjectBonitaUserList(ctx *gin.Context)
	ProjectListUser(ctx *gin.Context)
	ProduceSalesListUser(ctx *gin.Context) 
	ProjectTemplateListUser(ctx *gin.Context) 
	ProjectAuthorizationListUser(ctx *gin.Context) 
	GetByProjectListUser(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	B2UpdatedCaseID(ctx *gin.Context)
	B2ReviewTask(ctx *gin.Context)
	// B2TopReviewTask(ctx *gin.Context)
	// B2DepartmentReviewTask(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	ProjectResolver project.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ProjectResolver: project.New(db),
	}
}
