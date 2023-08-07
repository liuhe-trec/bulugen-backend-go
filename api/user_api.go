package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "login success",
	})
}
