package manufacture_user

import (
	"eirc.app/internal/v1/resolver/manufacture_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByManufactureID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	ManufactureUserResolver manufacture_user.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ManufactureUserResolver: manufacture_user.New(db),
	}
}
