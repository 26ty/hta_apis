package labor_hour_modify

import (
	"eirc.app/internal/v1/resolver/labor_hour_modify"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetLaborBonitaCaseListStart(ctx *gin.Context)
	GetLaborBonitaCaseListDepartment(ctx *gin.Context)
	GetByUserIdLaborHourModifyList(ctx *gin.Context)
	GetByCuIdLaborHourModifyList(ctx *gin.Context) 
	GetByUserIdList(ctx *gin.Context) 
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
	UpdatedStatus(ctx *gin.Context) 
	LaborReviewTask(ctx *gin.Context)
}

type presenter struct {
	LaborHourModifyResolver labor_hour_modify.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		LaborHourModifyResolver: labor_hour_modify.New(db),
	}
}
