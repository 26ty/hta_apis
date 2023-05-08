package personnel_affiliation

import (
	"encoding/json"
	"errors"

	bpm "eirc.app/internal/pkg/bpm"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	Model "eirc.app/internal/v1/structure"
	accountModel "eirc.app/internal/v1/structure/accounts"
	departmentModel "eirc.app/internal/v1/structure/department"
	jobtitleModel "eirc.app/internal/v1/structure/jobtitle"
	personnel_affiliationModel "eirc.app/internal/v1/structure/personnel_affiliation"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *personnel_affiliationModel.Created) interface{} {
	defer trx.Rollback()

	gonita_mebership := &Model.GonitaMembership{
		Account: input.Account,
		UserID:  input.BonitaUserID,
		GroupID: input.BonitaGroupID,
		RoleID:  input.BonitaRoleID,
	}

	status := bpm.AddUserMembership(gonita_mebership)

	if status == true {
		personnel_affiliation, err := r.PersonnelAffiliationService.WithTrx(trx).Created(input)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		trx.Commit()
		return code.GetCodeMessage(code.Successful, personnel_affiliation.PaID)
	} else {
		return code.GetCodeMessage(status, "bonita AddMembership error")
	}

}

func (r *resolver) List(input *personnel_affiliationModel.Fields) interface{} {
	output := &personnel_affiliationModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, personnel_affiliation, err := r.PersonnelAffiliationService.List(input)
	output.Total = quantity
	jobtitleByte, err := json.Marshal(personnel_affiliation)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(jobtitleByte, &output.PersonnelAffiliation)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *personnel_affiliationModel.Field) interface{} {
	personnel_affiliation, err := r.PersonnelAffiliationService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &personnel_affiliationModel.Single{}
	jobtitleByte, _ := json.Marshal(personnel_affiliation)
	err = json.Unmarshal(jobtitleByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserID(input *personnel_affiliationModel.Field) interface{} {
	personnel_affiliation, err := r.PersonnelAffiliationService.GetByUserID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &[]personnel_affiliationModel.Affiliation_Account{}
	jobtitleByte, _ := json.Marshal(personnel_affiliation)
	err = json.Unmarshal(jobtitleByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDepartmentID(input *personnel_affiliationModel.Field) interface{} {
	personnel_affiliation, err := r.PersonnelAffiliationService.GetByDepartmentID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &[]personnel_affiliationModel.Deparment_User{}
	jobtitleByte, _ := json.Marshal(personnel_affiliation)
	err = json.Unmarshal(jobtitleByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByParentDepartmentID(input *personnel_affiliationModel.Field) interface{} {
	department, err := r.DepartmentService.GetByID(&departmentModel.Field{DID: input.DepartmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	bonitaID := department.BonitaGroupID
	if department.BonitaParentGroupID != "" { //他不是孤兒(他是有父親的子部門)
		bonitaID = department.BonitaParentGroupID
	}
	personnel_affiliation, err := r.PersonnelAffiliationService.GetByParentDepartmentID(bonitaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &[]personnel_affiliationModel.Deparment_User{}
	jobtitleByte, _ := json.Marshal(personnel_affiliation)
	err = json.Unmarshal(jobtitleByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *personnel_affiliationModel.Updated) interface{} {
	personnel_affiliation, err := r.PersonnelAffiliationService.GetByID(&personnel_affiliationModel.Field{PaID: input.PaID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	account, err := r.AccountService.GetByID(&accountModel.Field{AccountID: personnel_affiliation.UserID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	department, err := r.DepartmentService.GetByID(&departmentModel.Field{DID: personnel_affiliation.DepartmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	jobtitle, err := r.JobtitleService.GetByID(&jobtitleModel.Field{JID: personnel_affiliation.JobtitleID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	gonita_mebership := &Model.GonitaMembership{
		Account: input.Account,
	}

	result := bpm.DeleteMembershipGetcode(gonita_mebership, account.BonitaUserID, department.BonitaGroupID, jobtitle.BonitaRoleID)

	if result != 200 {
		return code.GetCodeMessage(result, "bonita delete error")
	}

	err = r.PersonnelAffiliationService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *personnel_affiliationModel.Updated) interface{} {
	personnel_affiliation, err := r.PersonnelAffiliationService.GetByID(&personnel_affiliationModel.Field{PaID: input.PaID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	department, err := r.DepartmentService.GetByID(&departmentModel.Field{DID: personnel_affiliation.DepartmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	jobtitle, err := r.JobtitleService.GetByID(&jobtitleModel.Field{JID: personnel_affiliation.JobtitleID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PersonnelAffiliationService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	personnel_affiliation_new, err := r.PersonnelAffiliationService.GetByID(&personnel_affiliationModel.Field{PaID: input.PaID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	department_new, err := r.DepartmentService.GetByID(&departmentModel.Field{DID: personnel_affiliation_new.DepartmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	jobtitle_new, err := r.JobtitleService.GetByID(&jobtitleModel.Field{JID: personnel_affiliation_new.JobtitleID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	gonita_mebership := &Model.GonitaMembership{
		Account: input.Account,
		UserID:  input.BonitaUserID,
		GroupID: department_new.BonitaGroupID,
		RoleID:  jobtitle_new.BonitaRoleID,
	}

	result := bpm.UpdateMembershipGetcode(gonita_mebership, department.BonitaGroupID, jobtitle.BonitaRoleID)

	if result != true {
		return code.GetCodeMessage(result, "bonita update error")
	}

	return code.GetCodeMessage(code.Successful, personnel_affiliation.PaID)
}
