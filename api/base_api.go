package api

import (
	"bulugen-backend-go/global"
	"bulugen-backend-go/utils"
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

func (b *BaseApi) AddError(newErr error) {
	b.Errors = utils.AppendError(b.Errors, newErr)
}

func (b *BaseApi) GetError() error {
	return b.Errors
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (b *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	// 绑定请求上下文
	b.Ctx = option.Ctx
	// 绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = b.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = b.Ctx.ShouldBind(option.DTO)
		}
		if errResult != nil {
			errResult = b.parseValidateErrors(errResult, option.DTO)
			b.AddError(errResult)
			b.Fail(ResponseJson{
				Msg: b.GetError().Error(),
			})
		}
	}
	return b
}

func (b *BaseApi) parseValidateErrors(errs error, target any) error {
	var errResult error
	// 错误类型断言
	errValidationErrs, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}
	// 通过反射获取指针指向的元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidationErrs {
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

func (b *BaseApi) OK(resp ResponseJson) {
	OK(b.Ctx, resp)
}

func (b *BaseApi) Fail(resp ResponseJson) {
	Fail(b.Ctx, resp)
}

func (b *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(b.Ctx, resp)
}
