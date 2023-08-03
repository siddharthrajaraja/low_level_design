package impl

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service"
	"time"
)

type TaskServiceImpl struct {
	queueManagerService service.IQueueManagerService
}

func (t TaskServiceImpl) AddTask(tasks dtos.Tasks, timesStamp *time.Time, taskType dtos.TaskStatusType) {
	// TODO implement me
	t.queueManagerService.GetQueueByTaskType(taskType).AddTask(tasks, timesStamp)
}

func (t TaskServiceImpl) GetTask(s string) dtos.Tasks {
	// TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) ModifyTask(tasks dtos.Tasks) error {
	// TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) RemoveTask(s string) {
	// TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) ListTasks() []*dtos.Tasks {
	// TODO implement me
	panic("implement me")
}

func NewTaskService(queueManegerService service.IQueueManagerService) service.ITaskService {

	return &TaskServiceImpl{
		queueManagerService: queueManegerService,
	}
}
