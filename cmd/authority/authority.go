package main

// authority lambda
import (
	"eirc.app/internal/pkg/dao/gorm"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/v1/router"
	"eirc.app/internal/v1/router/account"
	"eirc.app/internal/v1/router/antivirus_software"
	"eirc.app/internal/v1/router/attendee"
	"eirc.app/internal/v1/router/company"
	"eirc.app/internal/v1/router/countersign"
	"eirc.app/internal/v1/router/countersign_type"
	"eirc.app/internal/v1/router/countersign_user"
	"eirc.app/internal/v1/router/customer"
	"eirc.app/internal/v1/router/customer_demand"
	"eirc.app/internal/v1/router/department"
	"eirc.app/internal/v1/router/factory"
	"eirc.app/internal/v1/router/factory_liaison"
	"eirc.app/internal/v1/router/factory_manufacturing"
	"eirc.app/internal/v1/router/file"
	"eirc.app/internal/v1/router/gateway_data"
	"eirc.app/internal/v1/router/gift_application"
	"eirc.app/internal/v1/router/gift_application_detail"
	"eirc.app/internal/v1/router/jig_demand"
	"eirc.app/internal/v1/router/jig_demand_detail"
	"eirc.app/internal/v1/router/jobtitle"
	"eirc.app/internal/v1/router/labor_hour"
	"eirc.app/internal/v1/router/labor_hour_modify"
	"eirc.app/internal/v1/router/login"
	"eirc.app/internal/v1/router/machine_combined"
	"eirc.app/internal/v1/router/manufacture_order"
	"eirc.app/internal/v1/router/manufacture_type"
	"eirc.app/internal/v1/router/manufacture_user"
	"eirc.app/internal/v1/router/meeting"
	"eirc.app/internal/v1/router/personnel_affiliation"
	"eirc.app/internal/v1/router/plug_in"
	"eirc.app/internal/v1/router/project"
	"eirc.app/internal/v1/router/project_template"
	"eirc.app/internal/v1/router/quotation"
	"eirc.app/internal/v1/router/quotation_detail"
	"eirc.app/internal/v1/router/sales_call_record"
	"eirc.app/internal/v1/router/task"
	"eirc.app/internal/v1/router/task_user"
	"eirc.app/internal/v1/router/todo_type"
	"eirc.app/internal/v1/router/transaction_record"

	"github.com/apex/gateway"
)

func main() {
	db, err := gorm.New()
	if err != nil {
		log.Error(err)
		return
	}

	route := router.Default()
	route = file.GetRoute(route, db)
	route = meeting.GetRoute(route, db)
	route = account.GetRoute(route, db)
	route = company.GetRoute(route, db)
	route = labor_hour.GetRoute(route, db)
	route = labor_hour_modify.GetRoute(route, db)
	route = login.GetRoute(route, db)
	route = attendee.GetRoute(route, db)
	route = project.GetRoute(route, db)
	route = plug_in.GetRoute(route, db)
	route = machine_combined.GetRoute(route, db)
	route = antivirus_software.GetRoute(route, db)
	route = countersign_type.GetRoute(route, db)
	route = countersign_user.GetRoute(route, db)
	route = countersign.GetRoute(route, db)
	route = task.GetRoute(route, db)
	route = task_user.GetRoute(route, db)
	route = customer_demand.GetRoute(route, db)
	route = project_template.GetRoute(route, db)
	route = manufacture_order.GetRoute(route, db)
	route = manufacture_type.GetRoute(route, db)
	route = manufacture_user.GetRoute(route, db)
	route = transaction_record.GetRoute(route, db)
	route = department.GetRoute(route, db)
	route = jobtitle.GetRoute(route, db)
	route = personnel_affiliation.GetRoute(route, db)
	route = todo_type.GetRoute(route, db)
	route = gateway_data.GetRoute(route, db)
	route = gift_application.GetRoute(route, db)
	route = gift_application_detail.GetRoute(route, db)
	route = quotation.GetRoute(route, db)
	route = quotation_detail.GetRoute(route, db)
	route = jig_demand.GetRoute(route, db)
	route = jig_demand_detail.GetRoute(route, db)
	route = sales_call_record.GetRoute(route, db)
	route = factory.GetRoute(route, db)
	route = factory_manufacturing.GetRoute(route, db)
	route = factory_liaison.GetRoute(route, db)
	route = customer.GetRoute(route, db)
	log.Fatal(gateway.ListenAndServe(":8080", route))
}
