package queue_service

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service"
	"time"
)

type PendingQueueServiceImpl struct {
	queue map[time.Time][]string
}

func (q PendingQueueServiceImpl) AddTask(tasks dtos.Tasks, time *time.Time) {
	// TODO implement me
	panic("implement me")
}

func (q PendingQueueServiceImpl) GetTask(tasks dtos.Tasks, startTime *time.Time, endTime *time.Time) {
	// TODO implement me
	panic("implement me")
}

func NewPendingQueueServiceImpl() service.IQueueService {

	queue := make(map[time.Time][]string, 0)
	return &PendingQueueServiceImpl{
		queue: queue,
	}
}
