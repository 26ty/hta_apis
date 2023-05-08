package login

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/jwe"
	"eirc.app/internal/v1/service/customer_demand"
	"eirc.app/internal/v1/service/project"
	"eirc.app/internal/v1/service/manufacture_order"
	"eirc.app/internal/v1/service/labor_hour_modify"

	jweModel "eirc.app/internal/v1/structure/jwe"
	loginsModel "eirc.app/internal/v1/structure/logins"
	Model "eirc.app/internal/v1/structure"
	"gorm.io/gorm"
)

type Resolver interface {
	Web(input *loginsModel.Login) interface{}
	Refresh(input *jweModel.Refresh) interface{}
	GetBonitaCaseCount(input []Model.GetCaseListOutput) interface{}
	GetBonitaCaseDetail(input []Model.GetCaseListOutput) interface{}
}

type resolver struct {
	Account account.Service
	JWE     jwe.Service
	CustomerDemandService customer_demand.Service
	ProjectService project.Service
	ManufactureOrderService manufacture_order.Service
	LaborHourModifyService labor_hour_modify.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		Account: account.New(db),
		JWE:     jwe.New(),
		CustomerDemandService: customer_demand.New(db),
		ProjectService: project.New(db),
		ManufactureOrderService: manufacture_order.New(db),
		LaborHourModifyService: labor_hour_modify.New(db),
	}
}
