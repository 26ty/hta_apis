package department

import (
	bpm "eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	// accountModel "eirc.app/internal/v1/structure/accounts"
	// companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	departmentModel "eirc.app/internal/v1/structure/department"
	personnel_affiliationModel "eirc.app/internal/v1/structure/personnel_affiliation"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *departmentModel.Created) interface{} {
	defer trx.Rollback()

	department, err := r.DepartmentService.WithTrx(trx).Created(input)
	if err != nil {
		// if err.Error() == "department already exists" {
		// 	return code.GetCodeMessage(code.BadRequest, err.Error())
		// }

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, department.DID)
}

func (r *resolver) List(input *departmentModel.Fields) interface{} {
	output := &departmentModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, department, err := r.DepartmentService.List(input)
	output.Total = quantity
	departmentByte, err := json.Marshal(department)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(departmentByte, &output.Department)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) A1Department(input *departmentModel.Field) interface{} {
	department, err := r.DepartmentService.A1Department(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &departmentModel.Deparment_Users{}
	departmentByte, _ := json.Marshal(department)
	err = json.Unmarshal(departmentByte, &output.Department)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	for _, value := range output.Department {
		if value.DID == "3fb4e9b0-544d-407f-a24e-15ead22706f2" || value.DID == "dd8bd64c-050f-4bfa-993c-a5aeb9f91535" {
			personnel_affiliation, err := r.PersonnelAffiliationService.GetByParentDepartmentID(value.BonitaParentGroupID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return code.GetCodeMessage(code.DoesNotExist, err)
				}

				log.Error(err)
				return code.GetCodeMessage(code.InternalServerError, err)
			}

			jobtitleByte, _ := json.Marshal(personnel_affiliation)
			err = json.Unmarshal(jobtitleByte, &value.Users)
			if err != nil {
				log.Error(err)
				return code.GetCodeMessage(code.InternalServerError, err)
			}
		} else {
			personnel_affiliation, err := r.PersonnelAffiliationService.GetByDepartmentID(&personnel_affiliationModel.Field{DepartmentID: value.DID})
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return code.GetCodeMessage(code.DoesNotExist, err)
				}

				log.Error(err)
				return code.GetCodeMessage(code.InternalServerError, err)
			}

			jobtitleByte, _ := json.Marshal(personnel_affiliation)
			err = json.Unmarshal(jobtitleByte, &value.Users)
			if err != nil {
				log.Error(err)
				return code.GetCodeMessage(code.InternalServerError, err)
			}
		}

	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) AllDepartment(input *departmentModel.Field) interface{} {
	department, err := r.DepartmentService.AllDepartment(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &departmentModel.Deparment_Users{}
	departmentByte, _ := json.Marshal(department)
	err = json.Unmarshal(departmentByte, &output.Department)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	for _, value := range output.Department {
		personnel_affiliation, err := r.PersonnelAffiliationService.GetByDepartmentID(&personnel_affiliationModel.Field{DepartmentID: value.DID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		personnel_affiliationByte, _ := json.Marshal(personnel_affiliation)
		err = json.Unmarshal(personnel_affiliationByte, &value.Users)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) DepartmentAccountList(input *departmentModel.Users) interface{} {
	output := &departmentModel.Deparment_Accounts{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, department, err := r.DepartmentService.DepartmentAccountList(input)
	output.Total = quantity
	departmentByte, err := json.Marshal(department)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(departmentByte, &output.Department)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *departmentModel.Field) interface{} {
	department, err := r.DepartmentService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &departmentModel.Single{}
	departmentByte, _ := json.Marshal(department)
	err = json.Unmarshal(departmentByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *departmentModel.Updated) interface{} {
	department, err := r.DepartmentService.GetByID(&departmentModel.Field{DID: input.DID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	result := bpm.DeleteDepartmentGetcode(input, department.BonitaGroupID)

	if result != 200 {
		return code.GetCodeMessage(result, "bonita delete error")
	}

	err = r.DepartmentService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *departmentModel.Updated) interface{} {
	department, err := r.DepartmentService.GetByID(&departmentModel.Field{DID: input.DID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	input.ParentGroupId = input.BonitaParentGroupID

	if input.BonitaParentGroupID == "" {
		input.BonitaParentGroupID = " "
	}
	result := bpm.UpdateDepartmentGetcode(input, department.BonitaGroupID)

	if result != 200 {
		return code.GetCodeMessage(result, "bonita update error")
	}

	err = r.DepartmentService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, department.DID)
}
