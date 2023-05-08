package task

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/task"
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

	output.TID = util.GenerateUUID()
	output.CreateTime = util.NowToUTC()
	//output.IsDeleted = false
	marshal, err = json.Marshal(output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	table := &model.Create_Table{}
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

func (s *service) GetByIDListTaskHour(input *model.Users) (quantity int64, output []*model.Task_Account_Labor_Hour, err error) {
	amount, fields, err := s.Entity.GetByIDListTaskHour(input)
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

func (s *service) GetByTaskListUser(input *model.Users) (quantity int64, output []*model.Task_Account, err error) {
	amount, field, err := s.Entity.GetByTaskListUser(input)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	return amount, output, nil
}

func (s *service) GetTaskListHourByUserID(input *model.Field) (quantity int64, output []*model.Task_Hour_User, err error) {
	amount, field, err := s.Entity.GetTaskListHourByUserID(input)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	return amount, output, nil
}

func (s *service) GetByTaskListHourDocumentsAndUserID(input *model.Field) (quantity int64, output []*model.Task_Hour_User, err error) {
	amount, field, err := s.Entity.GetByTaskListHourDocumentsAndUserID(input)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	return amount, output, nil
}

func (s *service) TaskListUser(input *model.Users) (quantity int64, output []*model.Task_Account, err error) {
	amount, field, err := s.Entity.TaskListUser(input)
	if err != nil {
		log.Error(err)

		return 0, output, err
	}

	marshal, err := json.Marshal(field)
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

func (s *service) GetByOriginIDAndUserID(input *model.Users) (quantity int64, output []*model.Task_OriginId, err error) {
	amount, field, err := s.Entity.GetByOriginIDAndUserID(input)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return 0, nil, err
	}

	return amount, output, nil
}

func (s *service) GetByTIDTaskListUser(input *model.Fields) (quantity int64, output []*model.Task_User_Account, err error) {
	amount, fields, err := s.Entity.GetByTIDTaskListUser(input)
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

func (s *service) GetByDocumentIDTaskListLast(input *model.Fields) (quantity int64, output []*model.Task_Template_Last, err error) {
	amount, fields, err := s.Entity.GetByDocumentIDTaskListLast(input)
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

func (s *service) GetByDocumentIDTaskList(input *model.Field) (quantity int64, output []*model.Task_Template, err error) {
	amount, fields, err := s.Entity.GetByDocumentIDTaskList(input)
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

func (s *service) GetByIDTaskBonitaUserList(input *model.Users) (output []*model.Bonita_ID_List, err error) {
	field, err := s.Entity.GetByIDTaskBonitaUserList(input)
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

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{TID: input.TID})
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
	field, err := s.Entity.GetByID(&model.Field{TID: input.TID})
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


