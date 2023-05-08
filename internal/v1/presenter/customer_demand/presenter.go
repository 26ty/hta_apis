package customer_demand

import (
	"eirc.app/internal/v1/resolver/customer_demand"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetA1BonitaCaseListStart(ctx *gin.Context)
	GetA1BonitaCaseListTask(ctx *gin.Context) 
	GetA1BonitaCaseListDepartment(ctx *gin.Context)
	GetA1BonitaCaseListDirector(ctx *gin.Context)
	GetA1BonitaCaseListTop(ctx *gin.Context)
	GetA1BonitaCaseListDispatch(ctx *gin.Context)
	GetA1BonitaCaseListEvaluation(ctx *gin.Context)
	GetA1BonitaCaseListCountersign(ctx *gin.Context)
	GetA1BonitaCaseListPMEvaluation(ctx *gin.Context) 
	GetA1BonitaCaseListBusiness(ctx *gin.Context)
	GetA1BonitaCaseListBusinessManager(ctx *gin.Context)
	GetA1BonitaCaseListBusinessDirector(ctx *gin.Context) 
	GetA1BonitaCaseListTaskFinish(ctx *gin.Context) 
	GetA1BonitaCaseListTaskFinishManager(ctx *gin.Context) 
	GetA1BonitaCaseListBusinessClose(ctx *gin.Context) 
	GetA1BonitaCaseListDepartmentClose(ctx *gin.Context) 
	GetA1BonitaCaseListDirectorClose(ctx *gin.Context)
	GetA1BonitaCaseListTopClose(ctx *gin.Context) 
	GetA1BonitaCaseListProductionClose(ctx *gin.Context)
	GetA1BonitaCaseListCountersignClose(ctx *gin.Context)
	GetA1BonitaCaseListPMClose(ctx *gin.Context)
	CustomerDemandListUser(ctx *gin.Context)
	GetByCuIDCustomerDemandListUser(ctx *gin.Context)
	GetByUserIDListCR(ctx *gin.Context)
	GetByUserIDListHCR(ctx *gin.Context) 
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	A1UpdatedCaseID(ctx *gin.Context) 
	A1ReviewTask(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	CustomerDemandResolver customer_demand.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		CustomerDemandResolver: customer_demand.New(db),
	}
}
