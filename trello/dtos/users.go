package dtos

import uuid "github.com/satori/go.uuid"

type Users struct {
	userId string
	name   string
	email  string
}

func NewUser(name, email string) *Users {

	uuidValue, _ := uuid.NewV4()
	return &Users{
		userId: uuidValue.String(),
		name:   name,
		email:  email,
	}
}

func (user *Users) GetUserId() string {

	return user.userId
}
