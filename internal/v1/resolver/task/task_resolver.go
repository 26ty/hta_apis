package task

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	taskModel "eirc.app/internal/v1/structure/task"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *taskModel.Created_List) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱
	// _, err := r.TaskService.WithTrx(trx).GetByID(&taskModel.Field{PID: input.PID})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	arr := []string{}
	for _, value := range input.Task {
		task, err := r.TaskService.WithTrx(trx).Created(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		arr = append(arr, task.TID)
	}

	// task, err := r.TaskService.WithTrx(trx).Created(input)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err.Error())
	// }

	trx.Commit()
	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) List(input *taskModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, task, err := r.TaskService.List(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByIDListTaskHour(input *taskModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Task_Account_Labor_Hours{}
	quantity, task, err := r.TaskService.GetByIDListTaskHour(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByTaskListUser(input *taskModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Task_Accounts{}
	quantity, task, err := r.TaskService.GetByTaskListUser(input)
	output.Total = quantity
	taskByte, _ := json.Marshal(task)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetTaskListHourByUserID(input *taskModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)

	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	output := &taskModel.Task_Hour_Users{}
	quantity, task, err := r.TaskService.GetTaskListHourByUserID(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByTaskListHourDocumentsAndUserID(input *taskModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)

	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	output := &taskModel.Task_Hour_Users{}
	quantity, task, err := r.TaskService.GetByTaskListHourDocumentsAndUserID(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) TaskListUser(input *taskModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Task_Accounts{}
	// output.Limit = input.Limit
	// output.Page = input.Page
	quantity, task, err := r.TaskService.TaskListUser(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByOriginIDAndUserID(input *taskModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)

	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	output := &taskModel.Task_OriginIds{}
	quantity, task, err := r.TaskService.GetByOriginIDAndUserID(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByTIDTaskListUser(input *taskModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Task_User_Accounts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, task, err := r.TaskService.GetByTIDTaskListUser(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDocumentIDTaskListLast(input *taskModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Task_Template_Lasts{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, task, err := r.TaskService.GetByDocumentIDTaskListLast(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByDocumentIDTaskList(input *taskModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Task_Templates{}
	quantity, task, err := r.TaskService.GetByDocumentIDTaskList(input)
	output.Total = quantity
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByIDTaskBonitaUserList(input *taskModel.Users) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &taskModel.Bonita_ID_Lists{}
	task, err := r.TaskService.GetByIDTaskBonitaUserList(input)
	taskByte, err := json.Marshal(task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(taskByte, &output.Task)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *taskModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	task, err := r.TaskService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &taskModel.Single{}
	taskByte, _ := json.Marshal(task)
	err = json.Unmarshal(taskByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) DeleteList(input *taskModel.Updated_List) interface{} {
	arr := []string{}
	for _, value := range input.Task {
		task, err := r.TaskService.GetByID(&taskModel.Field{TID: value.TID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}
	
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	
		err = r.TaskService.Deleted(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
		arr = append(arr, task.TID+" Delete ok!")
	}

	return code.GetCodeMessage(code.Successful, arr)
}

func (r *resolver) Delete(input *taskModel.Updated) interface{} {
	_, err := r.TaskService.GetByID(&taskModel.Field{TID: input.TID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}
	
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
	
		err = r.TaskService.Deleted(input)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *taskModel.Updated_List) interface{} {
	arr := []string{}
	for _, value := range input.Task {
		task, err := r.TaskService.GetByID(&taskModel.Field{TID: value.TID})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.GetCodeMessage(code.DoesNotExist, err)
			}

			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}

		err = r.TaskService.Updated(value)
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
		arr = append(arr, task.TID)
	}

	// task, err := r.TaskService.GetByID(&taskModel.Field{TID: input.TID})
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return code.GetCodeMessage(code.DoesNotExist, err)
	// 	}

	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	// err = r.TaskService.Updated(input)
	// if err != nil {
	// 	log.Error(err)
	// 	return code.GetCodeMessage(code.InternalServerError, err)
	// }

	return code.GetCodeMessage(code.Successful, arr)
}

