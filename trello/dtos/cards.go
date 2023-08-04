package dtos

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cards struct {
	id           string
	name         string
	description  string
	assignedUser string
	priority     TaskPriority
	createdAt    int
	finishedAt   int
	eta          int
	status       ListType
}

func NewCards(name, description string) *Cards {

	uuidValue, _ := uuid.NewV4()
	return &Cards{
		id:          uuidValue.String(),
		name:        name,
		description: description,
		priority:    P0,
		createdAt:   time.Now().Nanosecond(),
		finishedAt:  -1,
		eta:         -1,
		status:      New,
	}
}

func (c *Cards) GetId() string {

	return c.id
}

func (c *Cards) GetName() string {

	return c.name
}

func (c *Cards) SetName(name string) {

	c.name = name
}

func (c *Cards) GetDescription() string {

	return c.description
}

func (c *Cards) SetDescription(description string) {

	c.description = description
}

func (c *Cards) GetAssignedUser() string {

	return c.assignedUser
}

func (c *Cards) SetAssignedUser(userId string) {

	c.assignedUser = userId
}

func (c *Cards) GetPriority() TaskPriority {
	return c.priority
}

func (c *Cards) SetPriority(taskPriority TaskPriority) {
	c.priority = taskPriority
}

func (c *Cards) GetCreatedAt() int {

	return c.createdAt
}

func (c *Cards) SetCreatedAt(timeStamp int) {

	c.createdAt = timeStamp
}

func (c *Cards) GetFinishedAt() int {

	return c.finishedAt
}

func (c *Cards) SetFinishedAt(timeStamp int) {

	c.finishedAt = timeStamp
}

func (c *Cards) GetEta() int {

	return c.eta
}

func (c *Cards) SetEta(timeStamp int) {

	c.eta = timeStamp
}

func (c *Cards) GetStatus() ListType {

	return c.status
}

func (c *Cards) SetStatus(status ListType) {

	c.status = status
}
