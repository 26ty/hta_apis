package labor_hour_modify

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/labor_hour_modify"
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

	output.HmID = util.GenerateUUID()
	output.CreateTime = util.NowToUTC()
	output.StatusTypeID = "7278b671-ad00-4ef2-b608-376a1f8c3967"
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

func (s *service) GetByUserIdList(input *model.Field) (quantity int64, output []*model.LaborHourModify, err error) {
	amount, fields, err := s.Entity.GetByUserIdList(input)
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

func (s *service) GetByUserIdLaborHourModifyList(input *model.Field) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.GetByUserIdLaborHourModifyList(input)
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

func (s *service) GetByCuIdLaborHourModifyList(input *model.Field) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.GetByCuIdLaborHourModifyList(input)
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

func (s *service) GetByCaseID(input string) (output *model.ReviewByDepartment, err error) {
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

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{HmID: input.HmID})
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Deleted(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{HmID: input.HmID})
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

func (s *service) UpdatedStatus(input *model.Updated_Review) (err error) {
	field, err := s.Entity.GetByID(&model.Field{HmID: input.HmID})
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
