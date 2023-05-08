package task_user

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	taskUserModel "eirc.app/internal/v1/structure/task_user"

	//accountModel "eirc.app/internal/v1/structure/accounts"
	//companyModel "eirc.app/internal/v1/structure/companies"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *taskUserModel.Created_List) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱
	// _, err := r.CompanyService.WithTrx(trx).GetByID(&companyModel.Field{CompanyID: input.CompanyID,
	// 	IsDeleted: util.PointerBool(false)})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	arr := []string{}
	for _, value := range input.Task_user {
		value.StatusTypeID = "98c543c6-944e-4860-b323-166ed5f3920e"
		task_user, err := r.TaskUserService.WithTrx(trx).Created(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		arr = append(arr, task_user.TuID)
	}

	// task_user, err := r.TaskUserService.WithTrx(trx).Created(input)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	trx.Commit()
	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) List(input *taskUserModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskUserModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, task_user, err := r.TaskUserService.List(input)
	output.Total = quantity
	taskUserByte, err := json.Marshal(task_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(taskUserByte, &output.Task_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDocumnetIDListHour(input *taskUserModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskUserModel.Task_user_Labor_Hours{}
	quantity, task_user, err := r.TaskUserService.GetByDocumnetIDListHour(input)
	output.Total = quantity
	taskUserByte, err := json.Marshal(task_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(taskUserByte, &output.Task_user)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *taskUserModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	task_user, err := r.TaskUserService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &taskUserModel.Single{}
	taskUserByte, _ := json.Marshal(task_user)
	err = json.Unmarshal(taskUserByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetName(input *taskUserModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	task_user, err := r.TaskUserService.GetName(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &taskUserModel.Task_user_Account{}
	taskUserByte, _ := json.Marshal(task_user)
	err = json.Unmarshal(taskUserByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) DeleteList(input *taskUserModel.Updated_List) interface{} {

	arr := []string{}
	for _, value := range input.Task_user {
		task_user, err := r.TaskUserService.GetByID(&taskUserModel.Field{TuID: value.TuID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}
	
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	
		err = r.TaskUserService.Deleted(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
		arr = append(arr, task_user.TuID +" Delete ok!")
	}

	return code.GetCodeMessage(code.Successful,arr )
}

func (r *resolver) Delete(input *taskUserModel.Updated) interface{} {

	_, err := r.TaskUserService.GetByID(&taskUserModel.Field{TuID: input.TuID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}
	
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	
		err = r.TaskUserService.Deleted(input)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

	return code.GetCodeMessage(code.Successful,"Delete ok!" )
}

func (r *resolver) Updated(input *taskUserModel.Updated_List) interface{} {

	arr := []string{}
	for _, value := range input.Task_user {
		task_user, err := r.TaskUserService.GetByID(&taskUserModel.Field{TuID: value.TuID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		err = r.TaskUserService.Updated(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
		arr = append(arr, task_user.TuID)
	}

	// task_user, err := r.TaskUserService.GetByID(&taskUserModel.Field{TuID: input.TuID})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// err = r.TaskUserService.Updated(input)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) UpdatedStatus(input *taskUserModel.Updated_Review) interface{} {

	task_user, err := r.TaskUserService.GetByID(&taskUserModel.Field{TuID: input.TuID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.TaskUserService.UpdatedStatus(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, task_user.TuID)
}
