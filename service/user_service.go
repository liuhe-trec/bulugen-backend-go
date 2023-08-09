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

func (u *UserService) Login(iUserDto dto.UserLoginDTO) (model.User, error) {
	var errResult error
	iUser := u.Dao.GetUserByNameAndPassword(iUserDto.Name, iUserDto.Password)
	if iUser.ID == 0 {
		errResult = errors.New("invalid username or password")
	}
	return iUser, errResult
}

func (u *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if u.Dao.CheckUserNameExist(iUserAddDTO.Name) {
		return errors.New("user name exist")
	}
	return u.Dao.AddUser(iUserAddDTO)
}

func (u *UserService) GetUserByID(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return u.Dao.GetUserByID(iCommonIDDTO.ID)
}

func (u *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return u.Dao.GetUserList(iUserListDTO)
}

func (u *UserService) UpdateUser(iUpdateUserDTO *dto.UpdateUserDTO) error {
	if iUpdateUserDTO.ID == 0 {
		return errors.New("invalid user id")
	}
	//根据不同业务场景有追加不同的业务逻辑判断
	return u.Dao.UpdateUser(iUpdateUserDTO)
}

func (u *UserService) DeleteUserById(iCommonIDDTO *dto.CommonIDDTO) error {
	return u.Dao.DeleteUserById(iCommonIDDTO.ID)
}
