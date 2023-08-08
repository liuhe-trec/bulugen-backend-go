package service

import (
	"bulugen-backend-go/dao"
	"bulugen-backend-go/model"
	"bulugen-backend-go/service/dto"
	"errors"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (u UserService) Login(iUserDto dto.UserLoginDTO) (model.User, error) {
	var errResult error
	iUser := u.Dao.GetUserByNameAndPassword(iUserDto.Name, iUserDto.Password)
	if iUser.ID == 0 {
		errResult = errors.New("invalid username or password")
	}
	return iUser, errResult
}
