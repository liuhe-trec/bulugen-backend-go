package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "bulugen-backend-go/docs"
	"bulugen-backend-go/global"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
*
可插拔路由,根据配置文件中的配置决定某个模块是否启用;
1.在基础路由中,只负责路由的初始化处理
*/
type IFnRegistRouter = func(rgPublic, rgAuth *gin.RouterGroup)

var (
	ginFnRouters []IFnRegistRouter
)

func RegistRouter(fn IFnRegistRouter) {
	if fn == nil {
		return
	}
	ginFnRouters = append(ginFnRouters, fn)
}

func InitRouter() {
	//优雅的关闭
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()
	// 定义两个组,public不需要鉴权另一个需要
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")
	// 初始化平台路由
	initBasePlatformRouter()

	// 注册自定义验证器
	registCustValidator()

	// 注册系统各个模块对应的路由信息
	for _, router := range ginFnRouters {
		router(rgPublic, rgAuth)
	}
	// 继承swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	serverPort := viper.GetString("server.port")
	if serverPort == "" {
		serverPort = "8999"
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: r,
	}
	// 启动一个goroutine来开启web服务,避免主线程信号监听被阻塞
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listening: %s", serverPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s\n", err.Error()))
			return
		}
	}()

	<-ctx.Done()
	cancelCtx()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("stop server error: %s\n", err.Error()))
		return
	}
	global.Logger.Info("Stop Server Success")
}

func initBasePlatformRouter() {
	InituserRouter()
}

// 注册自定义验证器
func registCustValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && strings.Index(value, "a") == 0 {
					return true
				}

			}
			return false
		})
	}
}
