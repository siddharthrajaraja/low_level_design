package dtos

import uuid "github.com/satori/go.uuid"

type Users struct {
	userId string
	Name   string
	Email  string
}

func NewUser(name, email string) *Users {

	uuidValue, _ := uuid.NewV4()
	return &Users{
		userId: uuidValue.String(),
		Name:   name,
		Email:  email,
	}
}

func (user *Users) GetUserId() string {

	return user.userId
}
