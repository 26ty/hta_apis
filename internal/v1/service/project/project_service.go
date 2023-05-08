package project

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/project"
)

func (s *service) Created(input *model.Created) (output *model.Base, err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	output.PID = util.GenerateUUID()
	output.CreateTime = util.NowToUTC()
	output.IsTemplate = false
	//output.IsDeleted = false
	marshal, err = json.Marshal(output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	table := &model.Table{}
	err = json.Unmarshal(marshal, &table)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = s.Entity.Created(table)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) List(input *model.Fields) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.List(input)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	return amount, output, err
}

func (s *service) GetByProjectBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error) {
	fields, err := s.Entity.GetByProjectBonitaUserList(input)
	if err != nil {
		log.Error(err)

		return output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)

		return output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return output, err
	}

	return output, err
}

func (s *service) ProjectListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) {
	amount, fields, err := s.Entity.ProjectListUser(input)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	return amount, output, err
}

func (s *service) ProduceSalesListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) {
	amount, fields, err := s.Entity.ProduceSalesListUser(input)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	return amount, output, err
}

func (s *service) ProjectTemplateListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) {
	amount, fields, err := s.Entity.ProjectTemplateListUser(input)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	return amount, output, err
}

func (s *service) ProjectAuthorizationListUser(input *model.Fields) (quantity int64, output []*model.Project_Account, err error) {
	amount, fields, err := s.Entity.ProjectAuthorizationListUser(input)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	return amount, output, err
}

func (s *service) GetByProjectListUser(input *model.Field) (output *model.Project_Account, err error) {
	field, err := s.Entity.GetByProjectListUser(input)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByID(input *model.Field) (output *model.Base, err error) {
	field, err := s.Entity.GetByID(input)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByCaseID(input string) (output *model.Project_Account, err error) {
	field, err := s.Entity.GetByCaseID(input)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByCaseIDBonitaUserID(caseId string,userId string) (output []*model.Tm_Return, err error) {
	field, err := s.Entity.GetByCaseIDBonitaUserID(caseId,userId)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByCaseIDTaskUserParentcaseID(caseId string,bonita_parentcase_id string) (output *model.Tm_Return, err error) {
	field, err := s.Entity.GetByCaseIDTaskUserParentcaseID(caseId,bonita_parentcase_id)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByCaseIDTaskUserStatus2(caseId string,status_type_id string,status_type_id2 string) (output []*model.Tm_Return, err error) {
	field, err := s.Entity.GetByCaseIDTaskUserStatus2(caseId,status_type_id,status_type_id2)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByCaseIDTaskUserDepartment(caseId string,status_type_id string,dep string) (output []*model.Tm_Return, err error) {
	field, err := s.Entity.GetByCaseIDTaskUserDepartment(caseId,status_type_id,dep)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{PID: input.PID})
	if err != nil {
		log.Error(err)

		return err
	}

	//field.UpdatedBy = input.UpdatedBy
	//field.UpdatedAt = util.PointerTime(util.NowToUTC())
	//field.IsDeleted = true
	err = s.Entity.Deleted(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{PID: input.PID})
	if err != nil {
		log.Error(err)

		return err
	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Updated(field)

	return err
}

func (s *service) Updated_Bonita(input *model.Updated_Bonita) (err error) {
	field, err := s.Entity.GetByID(&model.Field{PID: input.PID})
	if err != nil {
		log.Error(err)

		return err
	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Updated(field)

	return err
}
