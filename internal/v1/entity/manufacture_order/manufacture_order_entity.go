package manufacture_order

import (
	model "eirc.app/internal/v1/structure/manufacture_order"
	"strconv"
)

func (e *entity) Created(input *model.Create_Table) (err error) {
	err = e.db.Model(&model.Create_Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

// func (e *entity) ManufactureOrderCdListUser(input *model.Fields) (amount int64, output []*model.ManufactureOrder_Cd_Account, err error) {

// 	// err = e.db.Raw("select recipient_table.m_id,recipient_table.salesman_name,recipient_table.creater_name,recipient_table.recipient_name,manufacture_order.status,manufacture_order.date_for_open,manufacture_order.date_for_close,manufacture_order.code,manufacture_order.order_name,manufacture_order.salesman_id,manufacture_order.recipient_id,manufacture_order.creater,manufacture_order.project_id,customer_demand.code as cd_code from(select creater_table.m_id,creater_table.salesman_name,creater_table.creater_name,accounts.name as recipient_name from(select salesman_table.m_id,salesman_table.salesman_name,salesman_table.recipient_id,accounts.name as creater_name from(select manufacture_order .m_id,manufacture_order.creater,manufacture_order.recipient_id,accounts.name as salesman_name from manufacture_order left join accounts on accounts.account_id = manufacture_order.salesman_id) as salesman_table left join accounts on accounts.account_id = salesman_table.creater) as creater_table left join accounts on accounts.account_id = creater_table.recipient_id) as recipient_table left join manufacture_order on manufacture_order.m_id = recipient_table.m_id left join customer_demand on customer_demand.cd_id = manufacture_order.project_id").
// 	// 	Order("create_time desc").Find(&output).Error

// 	// e.db.Model(&model.ManufactureOrder_Cd_Account{}).Count(&amount)
// 	db := e.db.Model(&model.Table{})

// 	err = db.Select("manufacture_order.m_id,manufacture_order.code,manufacture_order.project_id,manufacture_order.order_name,manufacture_order.date_for_open,manufacture_order.date_for_close,manufacture_order.status,manufacture_order.recipient_id,manufacture_order.salesman_id,manufacture_order.creater,s_accounts.name as salesman_name,r_accounts.name as recipient_name,accounts.name as creater_name,customer_demand.code as cd_code,manufacture_order.bonita_case_id").
// 			Joins("left join accounts as s_accounts on s_accounts.account_id = manufacture_order.salesman_id").
// 			Joins("left join accounts as r_accounts on r_accounts.account_id = manufacture_order.recipient_id").
// 			Joins("left join customer_demand on customer_demand.cd_id = manufacture_order.project_id").
// 			Joins("left join accounts on accounts.account_id = manufacture_order.creater").
// 			Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
// 			Limit(int(input.Limit)).Order("manufacture_order.create_time desc").Find(&output).Error
// 	return amount, output, err
// }

func (e *entity) ManufactureOrderProjectListUser(input *model.Fields) (amount int64, output []*model.ManufactureOrder_Project_Account, err error) {

	// err = e.db.Raw("select recipient_table.m_id,recipient_table.salesman_name,recipient_table.creater_name,recipient_table.recipient_name,manufacture_order.status,manufacture_order.date_for_open,manufacture_order.date_for_close,manufacture_order.code,manufacture_order.order_name,manufacture_order.salesman_id,manufacture_order.recipient_id,manufacture_order.creater,manufacture_order.project_id,project.code as p_code from(select creater_table.m_id,creater_table.salesman_name,creater_table.creater_name,accounts.name as recipient_name from(select salesman_table.m_id,salesman_table.salesman_name,salesman_table.recipient_id,accounts.name as creater_name from(select manufacture_order .m_id,manufacture_order.creater,manufacture_order.recipient_id,accounts.name as salesman_name from manufacture_order left join accounts on accounts.account_id = manufacture_order.salesman_id) as salesman_table left join accounts on accounts.account_id = salesman_table.creater) as creater_table left join accounts on accounts.account_id = creater_table.recipient_id) as recipient_table left join manufacture_order on manufacture_order.m_id = recipient_table.m_id left join project on project.p_id = manufacture_order.project_id").
	// 	Order("create_time desc").Find(&output).Error

	// e.db.Model(&model.ManufactureOrder_Project_Account{}).Count(&amount)
	db := e.db.Model(&model.Table{})

	err = db.Select("manufacture_order.m_id,manufacture_order.code,manufacture_order.project_id,manufacture_order.order_name,manufacture_order.date_for_open,manufacture_order.date_for_close,manufacture_order.status,manufacture_order.recipient_id,manufacture_order.salesman_id,manufacture_order.creater,s_accounts.name as salesman_name,r_accounts.name as recipient_name,accounts.name as creater_name,customer_demand.code as cd_code,manufacture_order.project_detail,project.code as p_code,manufacture_order.bonita_case_id").
			Joins("left join accounts as s_accounts on s_accounts.account_id = manufacture_order.salesman_id").
			Joins("left join accounts as r_accounts on r_accounts.account_id = manufacture_order.recipient_id").
			Joins("left join project on project.p_id = manufacture_order.project_id").
			Joins("left join customer_demand on customer_demand.cd_id = manufacture_order.project_id").
			Joins("left join accounts on accounts.account_id = manufacture_order.creater").
			Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
			Limit(int(input.Limit)).Order("manufacture_order.create_time desc").Find(&output).Error
	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("m_id = ?", input.MID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByIDOne(input *model.Field) (output *model.One, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Select("manufacture_order.m_id,manufacture_order.project_id,customer_demand.code as customer_demand_code,manufacture_order.project_detail,project.code as project_code,manufacture_order.code,manufacture_order.order_name,manufacture_order.amount,manufacture_order.customer_id,manufacture_order.shipment_location,manufacture_order.date_for_open,manufacture_order.date_for_close,manufacture_order.date_for_estimated_shipment,manufacture_order.inner_id,manufacture_order.other_document_code,manufacture_order.remark,manufacture_order.creater,accounts.name as creater_name,manufacture_order.sales_assistant_id,manufacture_order.recipient_id,manufacture_order.salesman_id,manufacture_order.status,manufacture_order.copy_file,manufacture_order.bonita_case_id,manufacture_order.create_time ").
			Joins("left join project on project.p_id = manufacture_order.project_id").
			Joins("left join customer_demand on customer_demand.cd_id = manufacture_order.project_id").
			Joins("left join accounts on accounts.account_id = manufacture_order.creater").
			Where("m_id = ?", input.MID).First(&output).Error

	return output, err
}

func (e *entity) GetByPIDList(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("project_id = ?", input.ProjectID)

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("create_time desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByCaseID(input string) (output *model.Review, err error) {
	value, _ := strconv.Atoi(input)
	db := e.db.Model(&model.Table{})

	err = db.Select("manufacture_order.m_id,manufacture_order.project_id,customer_demand.code as customer_demand_code,project.code as project_code,manufacture_order.project_detail,manufacture_order.code,manufacture_order.order_name,manufacture_order.amount,manufacture_order.customer_id,c_accounts.name as customer_name,manufacture_order.shipment_location,manufacture_order.date_for_open,manufacture_order.date_for_close,manufacture_order.date_for_estimated_shipment,manufacture_order.inner_id,manufacture_order.other_document_code,manufacture_order.remark,manufacture_order.creater,accounts.name as creater_name,manufacture_order.sales_assistant_id,sa_accounts.name as sales_assistant_name,manufacture_order.recipient_id,r_accounts.name as recipient_name,manufacture_order.salesman_id,s_accounts.name as salesman_name,manufacture_order.status ").
		Joins("left join project on project.p_id = manufacture_order.project_id").
		Joins("left join customer_demand on customer_demand.cd_id = manufacture_order.project_id").
		Joins("left join accounts on accounts.account_id = manufacture_order.creater").
		Joins("left join accounts as c_accounts on c_accounts.account_id = manufacture_order.customer_id").
		Joins("left join accounts as sa_accounts on sa_accounts.account_id = manufacture_order.sales_assistant_id").
		Joins("left join accounts as r_accounts on r_accounts.account_id = manufacture_order.recipient_id").
		Joins("left join accounts as s_accounts on s_accounts.account_id = manufacture_order.salesman_id").
		Where("manufacture_order.bonita_case_id = ?",value).First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("m_id = ?", input.MID).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("m_id = ?", input.MID).Save(&input).Error

	return err
}

