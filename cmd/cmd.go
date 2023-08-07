package cmd

import (
	"bulugen-backend-go/conf"
	"bulugen-backend-go/global"
	"bulugen-backend-go/router"
	"fmt"
)

func Start() {
	// 初始化系统配置文件
	conf.InitConfig()
	// 初始化日志组件,因为要用到配置文件,所以必须在之后
	global.Logger = conf.InitLogger()
	// 初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("==============Clean==============")
}
