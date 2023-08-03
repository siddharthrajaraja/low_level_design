package dtos

type TaskStatusType string

const (
	Created   TaskStatusType = "Created"
	Pending   TaskStatusType = "Pending"
	SpillOver TaskStatusType = "SpillOver"
	Completed TaskStatusType = "Completed"
)
