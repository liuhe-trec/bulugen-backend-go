package middleware

import (
	"bulugen-backend-go/api"
	"bulugen-backend-go/global"
	"bulugen-backend-go/global/constants"
	"bulugen-backend-go/model"
	"bulugen-backend-go/service"
	"bulugen-backend-go/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_INVALID_TOKEN     = 10401  // Token无效
	ERR_CODE_TOKEN_PARSE       = 100402 // 解析Token失败
	ERR_CODE_TOKEN_NOT_MATCHED = 100403 // Token登录时Token不一致
	ERR_CODE_TOKEN_EXPIRED     = 100404 // Token 过期
	ERR_CODE_TOKEN_RENEW       = 100405 // TOken 续期失败
	TOKEN_NAME                 = "Authorization"
	TOKEN_PREFIX               = "Bearer: "
	RENEW_TOKEN_DURATION       = 10 * 60 * time.Second
)

func tokenErr(ctx *gin.Context, code int) {
	api.Fail(ctx, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid Token",
	})
}

func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 从请求头中拿到token
		token := ctx.GetHeader(TOKEN_NAME)
		// token不存在,直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(ctx, ERR_CODE_INVALID_TOKEN)
			return
		}
		// Token无法解析
		token = token[len(TOKEN_PREFIX):]
		iJwtCustClaims, err := utils.ParseToken(token)
		iUserID := iJwtCustClaims.ID
		if err != nil || iUserID == 0 {
			tokenErr(ctx, ERR_CODE_TOKEN_PARSE)
			return
		}

		userIDStr := strconv.Itoa(int(iUserID))
		redisUserIDKey := strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", userIDStr, -1)
		// Token与访问者登录对应的Token不一致,直接返回
		redisToken, err := global.RedisClient.Get(redisUserIDKey)
		if err != nil || token != redisToken {
			tokenErr(ctx, ERR_CODE_TOKEN_NOT_MATCHED)
			return
		}
		// Token过期,直接返回
		tokenExpireDuration, err := global.RedisClient.GetExpireDuration(redisUserIDKey)
		if err != nil || tokenExpireDuration <= 0 {
			tokenErr(ctx, ERR_CODE_TOKEN_EXPIRED)
			return
		}
		// Token续期
		if tokenExpireDuration.Seconds() < float64(RENEW_TOKEN_DURATION) {
			newToken, err := service.GenerateAndCacheLoginUserToken(iUserID, iJwtCustClaims.Name)
			if err != nil {
				tokenErr(ctx, ERR_CODE_TOKEN_RENEW)
				return
			}
			ctx.Header("token", newToken)
		}
		// 因为jwt的token是用id和name生成的,所以这里不需要查库
		// userInfo, err := dao.NewUserDao().GetUserByID(iUserID)
		// if err != nil {
		// 	tokenErr(ctx)
		// 	return
		// }
		// ctx.Set(constants.LOGIN_USER, userInfo)
		ctx.Set(constants.LOGIN_USER, model.LoginUser{
			ID:   iUserID,
			Name: iJwtCustClaims.Name,
		})
		ctx.Next()
	}
}
