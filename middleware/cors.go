package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	conf := cors.Config{
		// 因为有 AllowCredentials: true, 所以这个配置会失效
		// AllowAllOrigins:  true,
		// 上面的替代方案
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(conf)
}
