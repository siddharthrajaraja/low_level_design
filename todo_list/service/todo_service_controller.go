package service

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"time"
)

type ITodoServiceController interface {
	AddTask(tasks dtos.Tasks, time time.Time)
	GetTask(string) dtos.Tasks
	ModifyTask(tasks dtos.Tasks) error
	RemoveTask(string)
	ListTasks() []*dtos.Tasks
}
