package customer_demand

import (
	model "eirc.app/internal/v1/structure/customer_demand"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Create_Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	CustomerDemandListUser(input *model.Field) (amount int64, output []*model.Customer_Demand_Account, err error)
	GetByCuIDCustomerDemandListUser(input *model.Field) (output *model.Customer_Demand_Account, err error)
	GetByUserIDListCR(input *model.Users) (amount int64,output []*model.CR, err error) 
	GetByUserIDListHCR(input *model.Users) (amount int64,output []*model.H_CR, err error) 
	GetByID(input *model.Field) (output *model.Table, err error)
	GetByCaseID(input string) (output *model.Customer_Review, err error) 
	GetByCaseID2(input string) (output *model.Customer_Review2, err error)
	GetByCaseIDCountersignUserID(input string,userID string) (output *model.Customer_Review, err error) 
	GetByCaseIDCountersignParentcaseID(input string,bonita_parentcase_id string) (output *model.Customer_Review, err error) 
	GetByCaseIDTaskUser(caseId string,userID string,status_type_id string) (output []*model.Customer_Review_Task, err error) 
	GetByCaseIDTaskUserStatus(caseId string,bonita_parentcase_id string) (output *model.Customer_Review_Task, err error) 
	Deleted(input *model.Table) (err error)
	Updated(input *model.Table) (err error)
}

type entity struct {
	db *gorm.DB
}

func New(db *gorm.DB) Entity {
	return &entity{
		db: db,
	}
}

func (e *entity) WithTrx(tx *gorm.DB) Entity {
	return &entity{
		db: tx,
	}
}
