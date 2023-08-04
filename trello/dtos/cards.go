package dtos

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cards struct {
	id           string
	Name         string
	Description  string
	AssignedUser string
	Priority     TaskPriority
	CreatedAt    int
	FinishedAt   int
	Eta          int
	Status       ListType
}

func NewCards(name, description string) *Cards {

	uuidValue, _ := uuid.NewV4()
	return &Cards{
		id:          uuidValue.String(),
		Name:        name,
		Description: description,
		Priority:    P0,
		CreatedAt:   time.Now().Nanosecond(),
		FinishedAt:  -1,
		Eta:         -1,
		Status:      New,
	}
}

func (c *Cards) GetId() string {

	return c.id
}

func (c *Cards) GetName() string {

	return c.Name
}

func (c *Cards) SetName(name string) {

	c.Name = name
}

func (c *Cards) GetDescription() string {

	return c.Description
}

func (c *Cards) SetDescription(description string) {

	c.Description = description
}

func (c *Cards) GetAssignedUser() string {

	return c.AssignedUser
}

func (c *Cards) SetAssignedUser(userId string) {

	c.AssignedUser = userId
}

func (c *Cards) GetPriority() TaskPriority {
	return c.Priority
}

func (c *Cards) SetPriority(taskPriority TaskPriority) {
	c.Priority = taskPriority
}

func (c *Cards) GetCreatedAt() int {

	return c.CreatedAt
}

func (c *Cards) SetCreatedAt(timeStamp int) {

	c.CreatedAt = timeStamp
}

func (c *Cards) GetFinishedAt() int {

	return c.FinishedAt
}

func (c *Cards) SetFinishedAt(timeStamp int) {

	c.FinishedAt = timeStamp
}

func (c *Cards) GetEta() int {

	return c.Eta
}

func (c *Cards) SetEta(timeStamp int) {

	c.Eta = timeStamp
}

func (c *Cards) GetStatus() ListType {

	return c.Status
}

func (c *Cards) SetStatus(status ListType) {

	c.Status = status
}
