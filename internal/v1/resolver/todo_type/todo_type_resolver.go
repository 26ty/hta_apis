package todo_type

import (
	"encoding/json"
	"errors"
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	todo_typeModel "eirc.app/internal/v1/structure/todo_type"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *todo_typeModel.Created) interface{} {
	defer trx.Rollback()
	
	todo_type, err := r.Todo_typeService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	
	trx.Commit()
	return code.GetCodeMessage(code.Successful, todo_type.TtID)
	
}

func (r *resolver) List(input *todo_typeModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &todo_typeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, todo_type, err := r.Todo_typeService.List(input)
	output.Total = quantity
	todo_typeByte, err := json.Marshal(todo_type)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(todo_typeByte, &output.Todo_type)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *todo_typeModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	todo_type, err := r.Todo_typeService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &todo_typeModel.Single{}
	todo_typeByte, _ := json.Marshal(todo_type)
	err = json.Unmarshal(todo_typeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByUserID(input *todo_typeModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	todo_type, err := r.Todo_typeService.GetByUserID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &[]todo_typeModel.Single{}
	todo_typeByte, _ := json.Marshal(todo_type)
	err = json.Unmarshal(todo_typeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Delete(input *todo_typeModel.Updated) interface{} {
	_, err := r.Todo_typeService.GetByID(&todo_typeModel.Field{TtID: input.TtID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}
	
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	
		err = r.Todo_typeService.Deleted(input)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}


func (r *resolver) Updated(input *todo_typeModel.Updated) interface{} {
	
	todo_type, err := r.Todo_typeService.GetByID(&todo_typeModel.Field{TtID: input.TtID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.Todo_typeService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, todo_type.TtID)
}
