package impl

import (
	"github.com/siddharthrajaraja/low_level_design/trello/dtos"
	"github.com/siddharthrajaraja/low_level_design/trello/service"
)

type UserServiceImpl struct {
	UserMap map[string]*dtos.Users
}

func (u UserServiceImpl) CreateUser(name, email string) string {

	user := dtos.NewUser(name, email)
	u.UserMap[user.GetUserId()] = user
	return user.GetUserId()
}

func (u UserServiceImpl) CheckUserExists(id string) bool {

	_, isPresent := u.UserMap[id]
	return isPresent
}

func NewUserServiceImpl() service.IUserService {

	userMap := make(map[string]*dtos.Users)
	return &UserServiceImpl{
		UserMap: userMap,
	}
}
