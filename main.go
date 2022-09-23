package main

import (
	"fmt"
	"gen-resume/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	err := router.Run(":2027")
	if err != nil {
		// 错误处理。端口被占用或其他错误
		fmt.Println(err.Error())
	}
}
