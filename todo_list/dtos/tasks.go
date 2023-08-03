package dtos

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Tasks struct {
	taskID         string
	title          string
	taskStatusType TaskStatusType

	status    TaskStatusType
	createdAt time.Time
	updatedAt time.Time
}

func NewTask(title string, status TaskStatusType) *Tasks {

	uuidValue := uuid.NewV4()
	return &Tasks{
		taskID:    uuidValue.String(),
		title:     title,
		status:    status,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
