package queue_service

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service"
	"time"
)

type CreatedQueueServiceImpl struct {
	queue map[time.Time][]string
}

func (q CreatedQueueServiceImpl) AddTask(tasks dtos.Tasks, time *time.Time) {
	// TODO implement me
	panic("implement me")
}

func (q CreatedQueueServiceImpl) GetTask(tasks dtos.Tasks, startTime *time.Time, endTime *time.Time) {
	// TODO implement me
	panic("implement me")
}

func NewCreatedQueueServiceImpl() service.IQueueService {

	queue := make(map[time.Time][]string, 0)
	return &CreatedQueueServiceImpl{
		queue: queue,
	}
}
