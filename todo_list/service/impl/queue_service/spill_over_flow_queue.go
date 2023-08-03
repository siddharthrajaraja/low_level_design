package queue_service

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service"
	"time"
)

type SpillOverQueueServiceImpl struct {
	queue map[time.Time][]string
}

func (q SpillOverQueueServiceImpl) AddTask(tasks dtos.Tasks, time *time.Time) {
	// TODO implement me
	panic("implement me")
}

func (q SpillOverQueueServiceImpl) GetTask(tasks dtos.Tasks, startTime *time.Time, endTime *time.Time) {
	// TODO implement me
	panic("implement me")
}

func NewSpillOverQueueServiceImpl() service.IQueueService {

	queue := make(map[time.Time][]string, 0)
	return &SpillOverQueueServiceImpl{
		queue: queue,
	}
}
