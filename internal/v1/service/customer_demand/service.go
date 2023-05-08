package customer_demand

import (
	"eirc.app/internal/v1/entity/customer_demand"
	model "eirc.app/internal/v1/structure/customer_demand"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	CustomerDemandListUser(input *model.Field) (quantity int64, output []*model.Customer_Demand_Account, err error)
	GetByCuIDCustomerDemandListUser(input *model.Field) (output *model.Customer_Demand_Account, err error)
	GetByUserIDListCR(input *model.Users) (quantity int64, output []*model.CR, err error)
	GetByUserIDListHCR(input *model.Users) (quantity int64, output []*model.H_CR, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetByCaseID(input string) (output *model.Customer_Review, err error)
	GetByCaseID2(input string) (output *model.Customer_Review2, err error)
	GetByCaseIDCountersignUserID(input string,userID string) (output *model.Customer_Review, err error) 
	GetByCaseIDCountersignParentcaseID(input string,bonita_parentcase_id string) (output *model.Customer_Review, err error) 
	GetByCaseIDTaskUser(caseId string,userID string,status_type_id string) (output []*model.Customer_Review_Task, err error) 
	GetByCaseIDTaskUserStatus(caseId string,bonita_parentcase_id string) (output *model.Customer_Review_Task, err error) 
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	Updated_Bonita(input *model.Updated_Bonita) (err error)

}

type service struct {
	Entity customer_demand.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: customer_demand.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
