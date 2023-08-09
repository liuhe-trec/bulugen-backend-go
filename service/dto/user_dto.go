package dto

import "bulugen-backend-go/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required"`
}

// 添加用户相关DTO

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"name" binding:"required" message:"密码不能为空"`
}

func (u *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = u.Name
	iUser.RealName = u.RealName
	iUser.Mobile = u.Mobile
	iUser.Email = u.Email
	iUser.Password = u.Password
}
