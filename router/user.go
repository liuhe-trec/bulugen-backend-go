package router

import (
	"bulugen-backend-go/api"

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
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserByID)
			rgAuthUser.POST("/list", userApi.GetUserList)
		}
	})
}
