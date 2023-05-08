package labor_hour

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/labor_hour"
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

	output.HID = util.GenerateUUID()
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

func (s *service) GetByUserIdCategoryList(input *model.Field) (quantity int64, output []*model.GetUserCategoryLabor, err error) {
	amount, fields, err := s.Entity.GetByUserIdCategoryList(input)
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

func (s *service) GetByUserIdCategory(input *model.Field) (quantity int64, output []*model.GetUserCategoryLabor, err error) {
	amount, fields, err := s.Entity.GetByUserIdCategory(input)
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

func (s *service) GetByUserIdMonthList(input *model.Field_Month) (quantity int64, output []*model.GetUserAllLabor, err error) {
	amount, fields, err := s.Entity.GetByUserIdMonthList(input)
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

func (s *service) GetByUserIdOneSumLaborhour(input *model.Field_Month) (output *model.GetUserOneSumLabor, err error) {
	fields, err := s.Entity.GetByUserIdOneSumLaborhour(input)
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

func (s *service) GetByUserIdOneLaborhour(input *model.Field_Month) (output []*model.GetUserOneLabor, err error) {
	fields, err := s.Entity.GetByUserIdOneLaborhour(input)
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

func (s *service) GetByUserIdMonthSumList(input *model.Field_Month) (output []*model.GetUserAllSumLabor, err error) {
	fields, err := s.Entity.GetByUserIdMonthSumList(input)
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

func (s *service) GetByUserIdList(input *model.Field) (quantity int64, output []*model.LaborHour, err error) {
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

func (s *service) GetByUserIdLaborHourList(input *model.Field) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.GetByUserIdLaborHourList(input)
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

func (s *service) GetByCuIdLaborHourList(input *model.Field) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.GetByCuIdLaborHourList(input)
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

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{HID: input.HID})
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Deleted(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{HID: input.HID})
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
