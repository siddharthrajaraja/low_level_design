package service

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"time"
)

type ITaskService interface {
	AddTask(tasks dtos.Tasks, timesStamp *time.Time, taskType dtos.TaskStatusType)
	GetTask(string) dtos.Tasks
	ModifyTask(tasks dtos.Tasks) error
	RemoveTask(string)
	ListTasks() []*dtos.Tasks
}
