package api

import (
	"bulugen-backend-go/service"
	"bulugen-backend-go/service/dto"
	"bulugen-backend-go/utils"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
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
	iUser, err := u.Service.Login(iUserLogindto)
	if err != nil {
		u.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	// 给前台返回
	u.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

func (u UserApi) AddUser(ctx *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserAddDTO}).GetError() != nil {
		return
	}
	err := u.Service.AddUser(&iUserAddDTO)
	if err != nil {
		u.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	u.OK(ResponseJson{
		Data: iUserAddDTO,
	})

}

func (u UserApi) GetUserByID(ctx *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iCommonIDDTO, BindParamsFromUri: true}).GetError() != nil {
		return
	}
	iUser, err := u.Service.GetUserByID(&iCommonIDDTO)
	if err != nil {
		u.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})
		return
	}
	u.OK(ResponseJson{
		Data: iUser,
	})
}
