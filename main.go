package main

import (
	"flag"
	"fmt"
	"gen-resume/bootstrap"
	bootstrapConfig "gen-resume/config"
	"gen-resume/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	bootstrapConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化 Gin 实例
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理。端口被占用或其他错误
		fmt.Println(err.Error())
	}
}
