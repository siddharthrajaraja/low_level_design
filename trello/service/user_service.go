package service

type IUserService interface {
	CreateUser(name, email string) string
	CheckUserExists(userId string) bool
}
