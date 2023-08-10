package service

import (
	"bulugen-backend-go/dao"
	"bulugen-backend-go/global"
	"bulugen-backend-go/global/constants"
	"bulugen-backend-go/model"
	"bulugen-backend-go/service/dto"
	"bulugen-backend-go/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
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

func GenerateAndCacheLoginUserToken(uid uint, username string) (string, error) {
	token, err := utils.GenerateToken(uid, username)
	if err == nil {
		err = global.RedisClient.Set(strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(uid)), -1), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
	}
	return token, err
}

func (u *UserService) Login(iUserDto dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token = ""
	iUser, err := u.Dao.GetUserByName(iUserDto.Name)
	// 用户名或密码不正确
	if err != nil || !utils.CompareHashAndPassword(iUser.Password, iUserDto.Password) {
		errResult = errors.New("invalid username or password")
	} else { // 登录成功,生成token
		token, err = GenerateAndCacheLoginUserToken(iUser.ID, iUser.Name)
		if err != nil {
			errResult = fmt.Errorf("generate token error: %w", err)
		}
	}
	return iUser, token, errResult
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
