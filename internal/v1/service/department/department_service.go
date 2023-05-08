package department

import (
	"encoding/json"

	_ "errors"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/department"
)

func (s *service) Created(input *model.Created) (output *model.Base, err error) {
	// fields := &model.Fields{}
	// // fields.Limit = 1
	// // fields.Page = 1
	// fields.Manager = util.PointerString(input.Manager)
	// fields.Name = util.PointerString(input.Name)
	// //fields.EngName = util.PointerString(input.EngName)
	// amount, _, err := s.Entity.List(fields)
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }

	// if amount > 0 {
	// 	log.Info("department already exists. Department: ", input.Name, ",ManagerID:", input.Manager)
	// 	return nil, errors.New("department already exists")
	// }

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

	output.DID = util.GenerateUUID()
	output.CreateTime = util.NowToUTC()
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

func (s *service) A1Department(input *model.Field) (output []*model.Table, err error) {
	fields, err := s.Entity.A1Department(input)
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

func (s *service) AllDepartment(input *model.Field) (output []*model.Table, err error) {
	fields, err := s.Entity.AllDepartment(input)
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

func (s *service) DepartmentAccountList(input *model.Users) (quantity int64, output []*model.Deparment_Account, err error) {
	amount, fields, err := s.Entity.DepartmentAccountList(input)
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
	field, err := s.Entity.GetByID(&model.Field{DID: input.DID})
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Deleted(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{DID: input.DID})
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
