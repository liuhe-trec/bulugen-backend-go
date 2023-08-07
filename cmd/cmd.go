package cmd

import (
	"bulugen-backend-go/conf"
	"bulugen-backend-go/global"
	"bulugen-backend-go/router"
	"bulugen-backend-go/utils"
	"fmt"
)

func Start() {
	// 错误链
	var initErr error
	// 初始化系统配置文件
	conf.InitConfig()
	// 初始化日志组件,因为要用到配置文件,所以必须在之后
	global.Logger = conf.InitLogger()
	// 初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// 初始化Redis连接
	redisClient, err := conf.InitRedis()
	global.RedisClient = redisClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// 初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("==============Clean==============")
}
