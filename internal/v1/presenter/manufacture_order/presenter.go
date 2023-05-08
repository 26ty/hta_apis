package manufacture_order

import (
	"eirc.app/internal/v1/resolver/manufacture_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	// ManufactureOrderCdListUser(ctx *gin.Context)
	ManufactureOrderProjectListUser(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByIDOne(ctx *gin.Context)
	GetByPIDList(ctx *gin.Context)
	GetC2BonitaCaseListStart(ctx *gin.Context) 
	GetC2BonitaCaseListDepartment(ctx *gin.Context) 
	GetC2BonitaCaseListManufacture(ctx *gin.Context)
	GetC2BonitaCaseListTop(ctx *gin.Context)
	GetC2BonitaCaseListConfirm(ctx *gin.Context) 
	GetC2BonitaCaseListSave(ctx *gin.Context) 
	Delete(ctx *gin.Context)
	C2UpdatedCaseID(ctx *gin.Context)
	C2ReviewTask(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	ManufactureOrderResolver manufacture_order.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ManufactureOrderResolver: manufacture_order.New(db),
	}
}
