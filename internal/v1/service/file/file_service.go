package file

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/file"
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

	output.FID = util.GenerateUUID()
	output.CreateTime = util.NowToUTC()
	output.IsDeleted = false

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

func (s *service) GetByDocumentID(input *model.Field) (quantity int64, output []*model.FilebydocumentId, err error) {
	amount, fields, err := s.Entity.GetByDocumentID(input)
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

func (s *service) GetByDocumentIDUserID(input *model.Users) (quantity int64, output []*model.FilebydocumentId, err error) {
	amount, fields, err := s.Entity.GetByDocumentIDUserID(input)
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
	field, err := s.Entity.GetByID(&model.Field{FID: input.FID})
	if err != nil {
		log.Error(err)
		return err
	}

	// field.UpdatedBy = input.UpdatedBy
	// field.UpdatedAt = util.PointerTime(util.NowToUTC())
	field.IsDeleted = true
	err = s.Entity.Deleted(&model.Field{FID: field.FID})

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{FID: input.FID})
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

	// field.CreatedBy = input.AccountID

	err = s.Entity.Updated(field)
	return err
}
