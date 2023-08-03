package service

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"time"
)

type IQueueManagerService interface {
	GetQueueByTaskType(dtos.TaskStatusType) IQueueService
}

type IQueueService interface {
	AddTask(tasks dtos.Tasks, time *time.Time)
	GetTask(tasks dtos.Tasks, startTime *time.Time, endTime *time.Time)
}
