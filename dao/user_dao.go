package dao

import (
	"bulugen-backend-go/model"
	"bulugen-backend-go/service/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{BaseDao: NewBaseDao()}
	}
	return userDao
}

func (u *UserDao) GetUserByNameAndPassword(userName, password string) model.User {
	var iUser model.User
	u.Orm.Model(&iUser).Where("name=? and password=?", userName, password).Find(&iUser)
	return iUser
}

func (u *UserDao) CheckUserNameExist(userName string) bool {
	var totalNumber int64
	u.Orm.Model(&model.User{}).Where("name = ?", userName).Count(&totalNumber)
	return totalNumber > 0
}

func (u *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)
	err := u.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}
