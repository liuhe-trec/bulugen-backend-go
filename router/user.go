package router

import (
	"bulugen-backend-go/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InituserRouter() {
	RegistRouter(func(rgPublic, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
					},
				})
			})
			rgAuthUser.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   "1",
					"name": "zs",
				})
			})
		}
	})
}
