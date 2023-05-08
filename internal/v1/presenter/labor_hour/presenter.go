package labor_hour

import (
	"eirc.app/internal/v1/resolver/labor_hour"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByUserIdLaborHourList(ctx *gin.Context)
	GetByCuIdLaborHourList(ctx *gin.Context) 
	GetByUserIdList(ctx *gin.Context) 
	GetByUserIdCategoryList(ctx *gin.Context)
	GetByUserIdCategory(ctx *gin.Context)
	GetByUserIdMonthList(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	LaborHourResolver labor_hour.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		LaborHourResolver: labor_hour.New(db),
	}
}
