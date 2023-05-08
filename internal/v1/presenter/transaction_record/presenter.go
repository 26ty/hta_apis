package transaction_record

import (
	"eirc.app/internal/v1/resolver/transaction_record"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByDocumentIDUserList(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	TransactionRecordResolver transaction_record.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		TransactionRecordResolver: transaction_record.New(db),
	}
}
