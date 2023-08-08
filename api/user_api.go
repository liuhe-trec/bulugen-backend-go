package api

import (
	"bulugen-backend-go/service/dto"
	"bulugen-backend-go/utils"
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Tags User management
// @Summary User Login
// @Description User Login des
// @Param name formData string true "User name"
// @Param password formData string true "password"
// @Success 200 {string} string "login success"
// @Failure 401 {string} string true "login failed"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(ctx *gin.Context) {
	var iUserLogindto dto.UserLoginDTO
	errs := ctx.ShouldBind(&iUserLogindto)
	if errs != nil {
		Fail(ctx, ResponseJson{
			Msg: parseValidateErrors(errs.(validator.ValidationErrors), &iUserLogindto).Error(),
		})
		return
	}
	OK(ctx, ResponseJson{
		Data: iUserLogindto,
	})
}

func parseValidateErrors(errs validator.ValidationErrors, target any) error {
	var errResult error
	// 通过反射获取指针指向的元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprint("%S:%S Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}
