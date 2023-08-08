package api

import (
	"bulugen-backend-go/service/dto"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
	}
}

// @Tags User management
// @Summary User Login
// @Description User Login des
// @Param name formData string true "User name"
// @Param password formData string true "password"
// @Success 200 {string} string "login success"
// @Failure 401 {string} string true "login failed"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	// 参数校验
	var iUserLogindto dto.UserLoginDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserLogindto}).GetError(); err != nil {
		return
	}
	// 给前台返回
	u.OK(ResponseJson{
		Data: iUserLogindto,
	})
}
