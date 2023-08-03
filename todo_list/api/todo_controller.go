package api

import (
	"github.com/siddharthrajaraja/low_level_design/todo_list/dtos"
	"github.com/siddharthrajaraja/low_level_design/todo_list/service"
	"time"
)

type TodoController struct {
	todoService service.ITodoServiceController
	taskService service.ITaskService
}

func (t TodoController) AddTask(task dtos.Tasks, timeStamp *time.Time) {
	// TODO implement me
	t.taskService.AddTask(task, timeStamp, "")
}

func (t TodoController) GetTask(s string) dtos.Tasks {
	// TODO implement me
	panic("implement me")
}

func (t TodoController) ModifyTask(tasks dtos.Tasks) error {
	// TODO implement me
	panic("implement me")
}

func (t TodoController) RemoveTask(s string) {
	// TODO implement me
	panic("implement me")
}

func (t TodoController) ListTasks() []*dtos.Tasks {
	// TODO implement me
	panic("implement me")
}

func NewTodoController(todoService service.ITodoServiceController, taskService service.ITaskService) *TodoController {

	return &TodoController{
		todoService: todoService,
		taskService: taskService,
	}
}
