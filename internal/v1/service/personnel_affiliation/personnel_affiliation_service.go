package personnel_affiliation

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/personnel_affiliation"
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

	output.PaID = util.GenerateUUID()
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

func (s *service) GetByUserID(input *model.Field) (output []*model.Affiliation_Account, err error) {
	field, err := s.Entity.GetByUserID(input)
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

func (s *service) GetByDepartmentID(input *model.Field) (output []*model.Deparment_User, err error) {
	field, err := s.Entity.GetByDepartmentID(input)
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

func (s *service) GetByParentDepartmentID(bonita_group_id string) (output []*model.Deparment_User, err error) {
	field, err := s.Entity.GetByParentDepartmentID(bonita_group_id)
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
	field, err := s.Entity.GetByID(&model.Field{PaID: input.PaID})
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Deleted(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{PaID: input.PaID})
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
