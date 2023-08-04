package service

type IUserService interface {
	CreateUser(name, email string)
	CheckUserExists(userId string) bool
}
