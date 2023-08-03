package impl

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service/impl/queue_service"
)

type QueueManager struct {
	factory map[dtos.TaskStatusType]service.IQueueService
}

func (q QueueManager) GetQueueByTaskType(statusType dtos.TaskStatusType) service.IQueueService {
	// TODO implement me
	return q.factory[statusType]
}

func NewQueueManager(taskType dtos.TaskStatusType) service.IQueueManagerService {

	return &QueueManager{
		factory: map[dtos.TaskStatusType]service.IQueueService{
			dtos.Completed: queue_service.NewCompletedQueueServiceImpl(),
			dtos.Pending:   queue_service.NewPendingQueueServiceImpl(),
			dtos.SpillOver: queue_service.NewSpillOverQueueServiceImpl(),
			dtos.Created:   queue_service.NewCreatedQueueServiceImpl(),
		},
	}
}
