package dao

import "bulugen-backend-go/model"

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
